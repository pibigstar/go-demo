/**
 * Author: vuiletgo(china.violetgo@gmail.com)
 * Date: 13-12-08
 * Version: 0.02
 */
package weigo

//获取隐私设置信息
func (api *APIClient) GET_account_privacy(params map[string]interface{}, result *Config) error {
	return api.GET("account/get_privacy", params, result)
}

//获取所有学校列表
func (api *APIClient) GET_account_profile_school_list(params map[string]interface{}, result *[]School) error {
	return api.GET("account/profile/school_list", params, result)
}

//获取当前用户API访问频率限制
func (api *APIClient) GET_account_rate_limit_status(params map[string]interface{}, result *LimitStatus) error {
	return api.GET("account/rate_limit_status", params, result)
}

//获取用户的联系邮箱 **高级**
func (api *APIClient) GET_account_get_uid(params map[string]interface{}, uid *UserID) error {
	return api.GET("account/get_uid", params, uid)
}

//OAuth授权之后获取用户UID
func (api *APIClient) GET_account_get_email(params map[string]interface{}, email *Email) error {
	return api.GET("account/profile/email", params, email)
}
