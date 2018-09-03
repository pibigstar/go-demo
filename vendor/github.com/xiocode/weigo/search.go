/**
 * Author: vuiletgo(china.violetgo@gmail.com)
 * Date: 13-12-15
 * Version: 0.02
 */
package weigo

//搜用户搜索建议
func (api *APIClient) GET_search_suggestions_users(params map[string]interface{}, result interface{}) error {
	return api.GET("search/suggestions/users", params, result)
}

//搜学校搜索建议
func (api *APIClient) GET_search_suggestions_schools(params map[string]interface{}, result interface{}) error {
	return api.GET("search/suggestions/schools", params, result)
}

//搜应用搜索建议
func (api *APIClient) GET_search_suggestions_apps(params map[string]interface{}, result interface{}) error {
	return api.GET("search/suggestions/apps", params, result)
}

//联想搜索
func (api *APIClient) GET_search_suggestions_at_users(params map[string]interface{}, result interface{}) error {
	return api.GET("search/suggestions/at_users", params, result)
}

//搜索某一话题下的微博
func (api *APIClient) GET_search_topics(params map[string]interface{}, result *Topic) error {
	return api.GET("search/topics", params, result)
}
