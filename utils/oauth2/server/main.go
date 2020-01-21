package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"github.com/go-session/session"
	oredis "gopkg.in/go-oauth2/redis.v3"
	"gopkg.in/oauth2.v3"
	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/generates"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
)

const (
	UserToken = "user-token"
	UserId    = "pibigstar"
)

var (
	srv *server.Server
)

func main() {
	mgr := manage.NewDefaultManager()
	cfg := &manage.Config{
		AccessTokenExp:    time.Hour * 1,
		RefreshTokenExp:   time.Hour * 24 * 3,
		IsGenerateRefresh: false,
	}
	// 设置授权码模式令牌的配置参数
	mgr.SetAuthorizeCodeTokenCfg(cfg)

	// 设置令牌存储方式
	redisStore := oredis.NewRedisStore(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0,
	})
	mgr.MapTokenStorage(redisStore)
	// 设置令牌的生成方式
	mgr.MapAccessGenerate(generates.NewJWTAccessGenerate([]byte("pibigstar"), jwt.SigningMethodHS512))

	// 设置 Client端
	clientStore := store.NewClientStore()
	clientStore.Set("123456", &models.Client{
		// 分配给第三方的ID
		ID: "123456",
		// 分配给第三方的密钥
		Secret: "pibigstar",
		// 第三方的域名
		Domain: "http://localhost:8000",
	})
	mgr.MapClientStorage(clientStore)

	// 创建server实例
	srv = server.NewServer(server.NewConfig(), mgr)
	// 设置支持的授权类型
	srv.SetAllowedGrantType(oauth2.AuthorizationCode, oauth2.PasswordCredentials)
	// 根据请求获取用户标识
	srv.SetUserAuthorizationHandler(userAuthorHandler)
	// 根据请求的用户名和密码获取用户标识
	srv.SetPasswordAuthorizationHandler(passwordAuthorHandler)

	// 设置内部错误处理Handler
	srv.SetInternalErrorHandler(internalErrHandler)
	// 设置返回错误Response的Handler
	srv.SetResponseErrorHandler(responseErrHandler)

	// 启动 http server
	go RunServer()

	log.Println("Server is running at 9000 port.")
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func RunServer() {
	// GET:去登录页面，POST处理点击登录，跳转到 /auth
	http.HandleFunc("/login", loginHandler)
	// 跳转到授权页面
	http.HandleFunc("/auth", toAuthHandler)
	// 授权处理
	http.HandleFunc("/authorize", authorizeHandler)
	// 客户端通过code换取token
	http.HandleFunc("/token", tokenHandler)
	// client 请求 服务中心数据
	http.HandleFunc("/test", testHandler)
}

func userAuthorHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	store, err := session.Start(nil, w, r)
	if err != nil {
		return
	}
	// 判断用户是否已经登录
	id, ok := store.Get(UserToken)
	if !ok {
		if r.Form == nil {
			r.ParseForm()
		}

		store.Set("ReturnUri", r.Form)
		store.Save()
		// 用户如果还没有登录，则跳转到登录页面
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusFound)
		return
	}
	userID = id.(string)
	// 从store里面移除掉token
	store.Delete(UserToken)
	store.Save()
	return
}

func passwordAuthorHandler(username, password string) (string, error) {
	if username == "admin" && password == "admin" {
		return UserId, nil
	}
	return "", nil
}

func internalErrHandler(err error) (re *errors.Response) {
	log.Println("Internal Error:", err.Error())
	return
}

func responseErrHandler(re *errors.Response) {
	log.Println("Response Error:", re.Error.Error())
}

func authorizeHandler(w http.ResponseWriter, r *http.Request) {
	// 获取一个 session 的 store
	store, err := session.Start(nil, w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 获取授权成功之后的回调地址
	// 当授权成功之后会通知此地址，并加上 code
	if v, ok := store.Get("ReturnUri"); ok {
		r.Form = v.(url.Values)
	}

	store.Delete("ReturnUri")
	store.Save()

	// 处理授权请求
	err = srv.HandleAuthorizeRequest(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func tokenHandler(w http.ResponseWriter, r *http.Request) {
	err := srv.HandleTokenRequest(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

// 登录页面以及登录请求处理
func loginHandler(w http.ResponseWriter, r *http.Request) {
	store, err := session.Start(nil, w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == "POST" {
		// 如果是 POST请求，则认为用户直接登录了，跳转到 授权页面
		// TODO： 这里应该判断用户的账号和密码是否正确，正确的话，将用户信息放到 store里面
		if r.FormValue("username") == "admin" && r.FormValue("password") == "admin" {
			store.Set(UserToken, UserId)
			store.Save()

			w.Header().Set("Location", "/auth")
			w.WriteHeader(http.StatusFound)
			return
		}
		fmt.Fprint(w, "登录失败")
		return
	}

	// 如果是Get请求，则跳转到登录页面
	outputHTML(w, r, "utils/oauth2/server/static/login.html")
}

// 跳转到授权页面
func toAuthHandler(w http.ResponseWriter, r *http.Request) {
	store, err := session.Start(nil, w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 判断是否有 已登录用户ID，如果没有则跳转到登录页面
	if _, ok := store.Get(UserToken); !ok {
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusFound)
		return
	}

	outputHTML(w, r, "utils/oauth2/server/static/auth.html")
}

func outputHTML(w http.ResponseWriter, req *http.Request, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer file.Close()
	fi, _ := file.Stat()
	http.ServeContent(w, req, file.Name(), fi.ModTime(), file)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	// 校验token
	token, err := srv.ValidationBearerToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data := map[string]interface{}{
		"expires_in": int64(token.GetAccessCreateAt().Add(token.GetAccessExpiresIn()).Sub(time.Now()).Seconds()),
		"client_id":  token.GetClientID(),
		"user_id":    token.GetUserID(),
	}
	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	e.Encode(data)
}
