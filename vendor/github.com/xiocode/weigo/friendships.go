/**
 * Author: Tony.Shao(xiocode@gmail.com)
 * Date: 13-03-06
 * Version: 0.02
 */
package weigo

/////////////////////////////////////////////// 读取接口 /////////////////////////////////////////////////

//获取用户信息
func (api *APIClient) GET_friendships_friends(params map[string]interface{}, result *Friendships) error {
	return api.GET("friendships/friends", params, result)
}

//批量获取当前登录用户的关注人的备注信息 ** 高级接口 **
func (api *APIClient) GET_friendships_friends_remark_batch(params map[string]interface{}, result interface{}) error {
	return api.GET("friendships/friends/remark_batch", params, result)
}

//获取共同关注人列表
func (api *APIClient) GET_friendships_friends_in_common(params map[string]interface{}, result *Friendships) error {
	return api.GET("friendships/friends/in_common", params, result)
}

//获取双向关注列表
func (api *APIClient) GET_friendships_friends_bilateral(params map[string]interface{}, result *Friendships) error {
	return api.GET("friendships/friends/bilateral", params, result)
}

//获取双向关注UID列表
func (api *APIClient) GET_friendships_friends_bilateral_ids(params map[string]interface{}, result *FriendsIDS) error {
	return api.GET("friendships/friends/bilateral/ids", params, result)
}

//获取用户关注对象UID列表
func (api *APIClient) GET_friendships_friends_ids(params map[string]interface{}, result *FriendsIDS) error {
	return api.GET("friendships/friends/ids", params, result)
}

//获取用户粉丝列表
func (api *APIClient) GET_friendships_followers(params map[string]interface{}, result *Friendships) error {
	return api.GET("friendships/followers", params, result)
}

//获取用户粉丝UID列表
func (api *APIClient) GET_friendships_followers_ids(params map[string]interface{}, result *FriendsIDS) error {
	return api.GET("friendships/followers/ids", params, result)
}

//获取用户优质粉丝列表
func (api *APIClient) GET_friendships_followers_active(params map[string]interface{}, result *[]User) error {
	return api.GET("friendships/followers/active", params, result)
}

//获取我的关注人中关注了指定用户的人
func (api *APIClient) GET_friendships_friends_chain_followers(params map[string]interface{}, result *Friendships) error {
	return api.GET("friendships/friends_chain/followers", params, result)
}

//TODO result type
//获取两个用户关系的详细情况
func (api *APIClient) GET_friendships_show(params map[string]interface{}, result interface{}) error {
	return api.GET("friendships/show", params, result)
}

/////////////////////////////////////////////// 写入接口 /////////////////////////////////////////////////

//关注某用户
func (api *APIClient) POST_friendships_create(params map[string]interface{}, result *User) error {
	return api.POST("friendships/create", params, result)
}

//取消关注某用户
func (api *APIClient) POST_friendships_destroy(params map[string]interface{}, result *User) error {
	return api.POST("friendships/destroy", params, result)
}

//移除当前登录用户的粉丝 ** 高级接口 **
func (api *APIClient) POST_friendships_followers_destroy(params map[string]interface{}, result *User) error {
	return api.POST("friendships/followers/destroy", params, result)
}

//更新关注人备注 ** 高级接口 **
func (api *APIClient) POST_friendships_remark_update(params map[string]interface{}, result *User) error {
	return api.POST("friendships/remark/update", params, result)
}
