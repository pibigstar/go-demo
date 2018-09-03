/**
 * Author: vuiletgo(china.violetgo@gmail.com)
 * Date: 13-12-15
 * Version: 0.02
 */
package weigo

//长链转短链
func (api *APIClient) short_url_shorten(params map[string]interface{}, result interface{}) error {
	return api.GET("short_url/shorten", params, result)
}

//短链转长链
func (api *APIClient) GET_short_url_expand(params map[string]interface{}, result interface{}) error {
	return api.GET("short_url/expand", params, result)
}

//获取短链接在微博上的微博分享数
func (api *APIClient) GET_short_url_share_counts(params map[string]interface{}, result interface{}) error {
	return api.GET("short_url/share/counts", params, result)
}

//获取包含指定单个短链接的最新微博内容
func (api *APIClient) GET_short_url_share_statuses(params map[string]interface{}, result interface{}) error {
	return api.GET("short_url/share/statuses", params, result)
}

//获取短链接在微博上的微博评论数
func (api *APIClient) GET_short_url_comment_counts(params map[string]interface{}, result interface{}) error {
	return api.GET("short_url/comment/counts", params, result)
}

//获取包含指定单个短链接的最新微博评论
func (api *APIClient) GET_short_url_comment_comments(params map[string]interface{}, result interface{}) error {
	return api.GET("short_url/comment/comments", params, result)
}
