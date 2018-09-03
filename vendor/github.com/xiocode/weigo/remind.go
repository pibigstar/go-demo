/**
 * Author: vuiletgo(china.violetgo@gmail.com)
 * Date: 13-12-15
 * Version: 0.02
 */
package weigo

//获取某个用户的各种消息未读数
func (api *APIClient) GET_remind_unread_count(params map[string]interface{}, result interface{}) error {
	return api.GET("remind/unread_count", params, result)
}

//对当前登录用户某一种消息未读数进行清零
func (api *APIClient) POST_remind_set_count(params map[string]interface{}, result interface{}) error {
	return api.POST("remind/set_count", params, result)
}
