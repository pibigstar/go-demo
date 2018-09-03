/**
 * Author: vuiletgo(china.violetgo@gmail.com)
 * Date: 13-12-15
 * Version: 0.02
 */
package weigo

//获取系统推荐用户
func (api *APIClient) GET_suggestions_users_hot(params map[string]interface{}, result *[]User) error {
	return api.GET("suggestions/users/hot", params, result)
}

//获取用户可能感兴趣的人
func (api *APIClient) GET_suggestions_users_may_interested(params map[string]interface{}, result interface{}) error {
	return api.GET("suggestions/users/may_interested", params, result)
}

//获取系统推荐用户
func (api *APIClient) GET_suggestions_users_by_status(params map[string]interface{}, result *SuggestionsUser) error {
	return api.GET("suggestions/users/by_status", params, result)
}

//主Feed微博按兴趣推荐排序
func (api *APIClient) GET_suggestions_statuses_reorder(params map[string]interface{}, result *Topic) error {
	return api.GET("suggestions/statuses/reorder", params, result)
}

//主Feed微博按兴趣推荐排序的微博ID
func (api *APIClient) GET_suggestions_statuses_reorder_ids(params map[string]interface{}, result interface{}) error {
	return api.GET("suggestions/statuses/reorder/ids", params, result)
}

//热门收藏
func (api *APIClient) GET_suggestions_favorites_hot(params map[string]interface{}, result *[]Status) error {
	return api.GET("suggestions/favorites/hot", params, result)
}

//写入接口

//不感兴趣的人
func (api *APIClient) POST_suggestions_users_not_interested(params map[string]interface{}, result *User) error {
	return api.POST("suggestions/users/not_interested", params, result)
}
