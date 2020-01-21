# oauth2授权
> 文档中心： https://go-oauth2.github.io/zh/

- server: 授权中心
- client: 第三方应用

## 授权流程

1. 用户在client点击通过 `server`登录
2. 如果用户没有在`server`登录，则用户去登录
3. 用户已登录，弹出用户授权页面
4. 授权成功，返回 `code`
5. 第三方应用通过`code`请求授权中心换取token