/**
 * Author: vuiletgo(china.violetgo@gmail.com)
 * Date: 13-12-11
 * Version: 0.02
 */
package weigo

//读取接口

//获取当前用户的收藏列表
func (api *APIClient) GET_favorites(params map[string]interface{}, result *Favorites) error {
	return api.GET("favorites", params, result)
}

//获取当前用户的收藏列表的ID
func (api *APIClient) GET_favorites_ids(params map[string]interface{}, result *FavoritesID) error {
	return api.GET("favorites/ids", params, result)
}

//获取单条收藏信息
func (api *APIClient) GET_favorites_show(params map[string]interface{}, result *Favorites) error {
	return api.GET("favorites/show", params, result)
}

//获取当前用户某个标签下的收藏列表
func (api *APIClient) GET_favorites_by_tags(params map[string]interface{}, result *Favorites) error {
	return api.GET("favorites/by_tags", params, result)
}

//当前登录用户的收藏标签列表
func (api *APIClient) GET_favorites_tags(params map[string]interface{}, result interface{}) error {
	return api.GET("favorites/tags", params, result)
}

//获取当前用户某个标签下的收藏列表的ID
func (api *APIClient) GET_favorites_by_tags_ids(params map[string]interface{}, result *FavoritesID) error {
	return api.GET("favorites/by_tags/ids", params, result)
}

//写入接口

//添加收藏
func (api *APIClient) POST_favorites_create(params map[string]interface{}, result *Favorites) error {
	return api.POST("favorites/create", params, result)
}

//删除收藏
func (api *APIClient) POST_favorites_destroy(params map[string]interface{}, result *Favorites) error {
	return api.POST("favorites/destroy", params, result)
}

//批量删除收藏
func (api *APIClient) POST_favorites_destroy_batch(params map[string]interface{}, result interface{}) error {
	return api.POST("favorites/destroy_batch", params, result)
}

//更新收藏标签
func (api *APIClient) POST_favorites_tags_update(params map[string]interface{}, result *Favorites) error {
	return api.POST("favorites/tags/update", params, result)
}

//更新当前用户所有收藏下的指定标签
func (api *APIClient) POST_favorites_tags_update_batch(params map[string]interface{}, result interface{}) error {
	return api.POST("favorites/tags/update_batch	", params, result)
}

//删除当前用户所有收藏下的指定标签
func (api *APIClient) POST_favorites_tags_destroy_batch(params map[string]interface{}, result interface{}) error {
	return api.POST("favorites/tags/destroy_batch", params, result)
}
