/**
 * Author: Tony.Shao(xiocode@gmail.com)
 * Date: 13-03-06
 * Version: 0.02
 */
package weigo

import (
	"github.com/going/toolkit/dig"
)

/////////////////////////////////////////////// 读取接口 /////////////////////////////////////////////////

//获取用户发布的微博
func (api *APIClient) GET_statuses_user_timeline(params map[string]interface{}, result *Statuses) error {
	return api.GET("statuses/user_timeline", params, result)
}

//获取最新的公共微博
func (api *APIClient) GET_statuses_public_timeline(params map[string]interface{}, result *Statuses) error {
	return api.GET("statuses/public_timeline", params, result)
}

//获取当前登录用户及其所关注用户的最新微博
func (api *APIClient) GET_statuses_friends_timeline(params map[string]interface{}, result *Statuses) error {
	return api.GET("statuses/friends_timeline", params, result)
}

//获取当前登录用户及其所关注用户的最新微博
func (api *APIClient) GET_statuses_home_timeline(params map[string]interface{}, result *Statuses) error {
	return api.GET("statuses/home_timeline", params, result)
}

//返回一条原创微博的最新转发微博
func (api *APIClient) GET_statuses_repost_timeline(params map[string]interface{}, result *Reposts) error {
	return api.GET("statuses/repost_timeline", params, result)
}

//批量获取指定的一批用户的微博列表 ** 高级接口 **
func (api *APIClient) GET_statuses_timeline_batch(params map[string]interface{}, result *Statuses) error {
	return api.GET("statuses/timeline_batch", params, result)
}

//获取用户发布的微博的ID
func (api *APIClient) GET_statuses_user_timeline_ids(params map[string]interface{}, result *TimelineIDs) error {
	return api.GET("statuses//user_timeline/ids", params, result)
}

//获取一条原创微博的最新转发微博的ID
func (api *APIClient) GET_statuses_repost_timeline_ids(params map[string]interface{}, result *TimelineIDs) error {
	return api.GET("statuses/repost_timeline/ids", params, result)
}

//获取当前登录用户及其所关注用户的最新微博的ID
func (api *APIClient) GET_statuses_friends_timeline_ids(params map[string]interface{}, result *TimelineIDs) error {
	return api.GET("statuses/friends_timeline/ids", params, result)
}

//返回用户转发的最新微博
func (api *APIClient) GET_statuses_repost_by_me(params map[string]interface{}, result *Statuses) error {
	return api.GET("statuses/repost_by_me", params, result)
}

//获取@当前用户的最新微博
func (api *APIClient) GET_statuses_mentions(params map[string]interface{}, result *Statuses) error {
	return api.GET("statuses/mentions", params, result)
}

//获取@当前用户的最新微博的ID
func (api *APIClient) GET_statuses_mentions_ids(params map[string]interface{}, result *TimelineIDs) error {
	return api.GET("statuses/mentions/ids", params, result)
}

//获取双向关注用户的最新微博
func (api *APIClient) GET_statuses_bilateral_timeline(params map[string]interface{}, result *Statuses) error {
	return api.GET("statuses/bilateral_timeline", params, result)
}

//根据ID获取单条微博信息
func (api *APIClient) GET_statuses_show(params map[string]interface{}, result *Status) error {
	return api.GET("statuses/show", params, result)
}

//根据微博ID批量获取微博信息
func (api *APIClient) GET_statuses_show_batch(params map[string]interface{}, result *Statuses) error {
	return api.GET("statuses/show_batch", params, result)
}

//通过id获取mid
func (api *APIClient) GET_statuses_querymid(params map[string]interface{}, mid *string) error {
	result := new(map[string]interface{})
	err := api.GET("statuses/querymid", params, result)
	dig.Get(result, mid, "mid")
	return err
}

//通过mid获取id
func (api *APIClient) GET_statuses_queryid(params map[string]interface{}, id *string) error {
	result := new(map[string]interface{})
	err := api.GET("statuses/queryid", params, result)
	dig.Get(result, id, "id")
	return err
}

//批量获取指定微博的转发数评论数
func (api *APIClient) GET_statuses_count(params map[string]interface{}, result *[]StatusCount) error {
	return api.GET("statuses/count", params, result)
}

//获取当前登录用户关注的人发给其的定向微博 ** 高级接口 **
func (api *APIClient) GET_statuses_to_me(params map[string]interface{}, result *Statuses) error {
	return api.GET("statuses/to_me", params, result)
}

//获取当前登录用户关注的人发给其的定向微博ID列表 ** 高级接口 **
func (api *APIClient) GET_statuses_to_me_ids(params map[string]interface{}, result *TimelineIDs) error {
	return api.GET("statuses/to_me/ids", params, result)
}

/////////////////////////////////////////////// 写入接口 /////////////////////////////////////////////////

//转发一条微博信息
func (api *APIClient) POST_statuses_repost(params map[string]interface{}, result *Status) error {
	return api.POST("statuses/repost", params, result)
}

//删除微博信息
func (api *APIClient) POST_statuses_destroy(params map[string]interface{}, result *Status) error {
	return api.POST("statuses/destroy", params, result)
}

//发布一条微博信息
func (api *APIClient) POST_statuses_update(params map[string]interface{}, result *Status) error {
	return api.POST("statuses/update", params, result)
}

//上传图片并发布一条微博
func (api *APIClient) POST_statuses_upload(params map[string]interface{}, result *Status) error {
	return api.UPLOAD("statuses/upload", params, result)
}

//发布一条微博同时指定上传的图片或图片url ** 高级接口 **
func (api *APIClient) POST_statuses_upload_url_text(params map[string]interface{}, result *Status) error {
	return api.POST("statuses/upload_url_text", params, result)
}

//屏蔽某条微博 ** 高级接口 **
func (api *APIClient) POST_statuses_filter_create(params map[string]interface{}, result *Status) error {
	return api.POST("statuses/filter/create", params, result)
}

//屏蔽某个@我的微博及后续由其转发引起的@提及 ** 高级接口 **
func (api *APIClient) POST_statuses_mentions_shield(params map[string]interface{}, result interface{}) error {
	return api.POST("statuses/mentions/shield", params, result)
}
