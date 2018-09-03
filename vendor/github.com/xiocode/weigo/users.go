/**
 * Author: Tony.Shao(xiocode@gmail.com)
 * Date: 13-03-06
 * Version: 0.02
 */
package weigo

/////////////////////////////////////////////// 读取接口 /////////////////////////////////////////////////

//获取用户信息
func (api *APIClient) GET_users_show(params map[string]interface{}, result *User) error {
	return api.GET("users/show", params, result)
}

//通过个性域名获取用户信息
func (api *APIClient) GET_users_domain_show(params map[string]interface{}, result *User) error {
	return api.GET("users/domain_show", params, result)
}

//批量获取用户的粉丝数、关注数、微博数
func (api *APIClient) GET_users_counts(params map[string]interface{}, result *[]UserCounts) error {
	return api.GET("users/counts", params, result)
}

//废弃?
/*
func (api *APIClient) GET_users_show_rank(params map[string]interface{}, result *UserRank) error {
	return api.GET("users/show_rank", params, result)
}
*/
