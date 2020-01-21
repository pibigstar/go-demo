package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const (
	// 授权中心的URL地址
	authServerURL = "http://localhost:9000"
	// 请求标识
	StateValue = "pi"
	StateKey   = "state"
)

var (
	config = oauth2.Config{
		// 客户端ID
		ClientID: "123456",
		// 客户端密钥
		ClientSecret: "pibigstar",
		// 访问范围
		Scopes: []string{"all"},
		// 授权成功之后的回调地址
		RedirectURL: "http://localhost:8000/oauth2",
		// 节点地址
		Endpoint: oauth2.Endpoint{
			// 授权处理
			AuthURL: authServerURL + "/authorize",
			// 获取token地址
			TokenURL: authServerURL + "/token",
		},
	}
	// 全局token
	globalToken *oauth2.Token
)

func main() {
	// 用户点击 第三方登录，将当前网页的标识设置进去
	http.HandleFunc("/pi", toAuthorizeHandler)
	// 授权成功之后的回调地址
	http.HandleFunc("/oauth2", authSuccessHandler)

	// 刷新token
	http.HandleFunc("/refresh", refreshTokenHandler)
	// 判断是否已有token
	http.HandleFunc("/try", tryHandler)

	// 通过用户名和密码授权，授权中心必须支持：PasswordCredentials 模式
	http.HandleFunc("/pwd", pwdTokenHandler)
	// 通过client进行授权，授权中心必须支持：ClientCredentials 模式
	http.HandleFunc("/client", clientHandler)

	log.Println("Client is running at 8000 port.")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// 跳转到授权中心
func toAuthorizeHandler(w http.ResponseWriter, r *http.Request) {
	// 获取config中的 AuthURL 并添加 State 标识
	u := config.AuthCodeURL(StateValue)
	// 访问授权中心的授权地址
	http.Redirect(w, r, u, http.StatusFound)
}

// 授权成功之后的回调地址
func authSuccessHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	state := r.Form.Get(StateKey)
	// 判断是否是自己的搜全回调
	if state != StateValue {
		http.Error(w, "State invalid", http.StatusBadRequest)
		return
	}
	// 获取返回Code
	code := r.Form.Get("code")
	if code == "" {
		http.Error(w, "Code not found", http.StatusBadRequest)
		return
	}

	// 通过code换取一个token
	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// 将token设置为全局的
	globalToken = token

	// 写回前端
	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	e.Encode(token)
}

func refreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	if globalToken == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	globalToken.Expiry = time.Now()
	token, err := config.TokenSource(context.Background(), globalToken).Token()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	globalToken = token
	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	e.Encode(token)
}

func tryHandler(w http.ResponseWriter, r *http.Request) {
	if globalToken == nil {
		http.Redirect(w, r, "/pi", http.StatusFound)
		return
	}

	resp, err := http.Get(fmt.Sprintf("%s/test?access_token=%s", authServerURL, globalToken.AccessToken))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer resp.Body.Close()

	io.Copy(w, resp.Body)
}

func pwdTokenHandler(w http.ResponseWriter, r *http.Request) {
	token, err := config.PasswordCredentialsToken(context.Background(), "admin", "admin")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	globalToken = token
	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	e.Encode(token)
}

func clientHandler(w http.ResponseWriter, r *http.Request) {
	cfg := clientcredentials.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		TokenURL:     config.Endpoint.TokenURL,
	}

	token, err := cfg.Token(context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	e.Encode(token)
}
