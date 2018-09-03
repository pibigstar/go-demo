/**
 * Author: Tony.Shao(xiocode@gmail.com)
 * Date: 13-03-06
 * Version: 0.02
 */
package weigo

//返回指定用户的标签列表
func (api *APIClient) GET_tags(params map[string]interface{}, result interface{}) error {
	return api.GET("tags", params, result)
}

//批量获取用户标签
func (api *APIClient) GET_tags_tags_batch(params map[string]interface{}, result *[]Tags) error {
	return api.GET("tags/tags_batch", params, result)
}

//返回系统推荐的标签列表
func (api *APIClient) GET_tags_suggestions(params map[string]interface{}, result interface{}) error {
	return api.GET("tags/suggestions", params, result)
}

//写入接口

//添加用户标签
func (api *APIClient) POST_tags_create(params map[string]interface{}, result interface{}) error {
	return api.POST("tags/create", params, result)
}

//删除用户标签
func (api *APIClient) POST_tags_destroy(params map[string]interface{}, result interface{}) error {
	return api.POST("tags/destroy", params, result)
}

//批量删除用户标签
func (api *APIClient) POST_tags_destroy_batch(params map[string]interface{}, result interface{}) error {
	return api.POST("tags/destroy_batch", params, result)
}
