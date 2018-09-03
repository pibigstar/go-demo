package weigo

type TokenInfo struct {
	UID       int    `json:"uid"`
	AppKey    string `json:"appkey"`
	Scope     string `json:"scope"`
	CreatedAt int    `json:"create_at"`
	ExpireIn  int    `json:"expire_in"`
}

//获取一个token info
func (api *APIClient) GET_token_info(params map[string]interface{}, result *TokenInfo) error {
	return api.Auth("get_token_info", params, result)
}