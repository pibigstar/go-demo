Weigo
===========

Weigo是Go语言版本的SinaWeibo的SDK.以下是api接口功能说明

###未完成的功能有

    提醒
    收藏
    帐号
    搜索
    短链
    标签
    推荐
    话题

###微博
<table>
    <tr><th width="230">接口名</th><th width="250">方法</th><th width="260">参数</th><th>对应API</th></tr>
    <tr><td>获取最新的公共微博</td><td>GET_statuses_public_timeline</td><td>map[string]interface{}, *Statuses</td><td>statuses/public_timeline</td></tr>
    <tr><td>获取用户发布的微博</td><td>GET_statuses_user_timeline</td><td>map[string]interface{}, *Statuses</td><td>statuses/user_timeline</td></tr>
    <tr><td>获取当前登录用户<br />及其所关注用户的最新微博</td><td>GET_statuses_friends_timeline</td><td>map[string]interface{}, *Statuses</td><td>statuses/friends_timeline</td></tr>
    <tr><td>获取当前登录用户<br />及其所关注用户的最新微博</td><td>GET_statuses_home_timeline</td><td>map[string]interface{}, *Statuses</td><td>statuses/home_timeline</td></tr>
    <tr><td>返回一条原创微博的最新转发微博</td><td>GET_statuses_repost_timeline</td><td>map[string]interface{}, *Reposts</td><td>statusesstatuses/repost_timeline</td></tr>
    <tr><td>批量获取指定的一批用户的微博列表</td><td>GET_statuses_timeline_batch</td><td>map[string]interface{}, *Statuses</td><td>statuses/timeline_batch</td></tr>
    <tr><td>获取用户发布的微博的ID</td><td>GET_statuses_user_timeline_ids</td><td>map[string]interface{}, *TimelineIDs</td><td>statuses/repost_timeline/ids</td></tr>
    <tr><td>获取一条原创微博的最新转发微博的ID</td><td>GET_statuses_repost_timeline_ids</td><td>map[string]interface{}, *TimelineIDs</td><td>statuses/repost_timeline/ids</td></tr>
    <tr><td>获取当前登录用户及其所关注用户的最新微博的ID</td><td>GET_statuses_friends_timeline_ids</td><td>map[string]interface{}, *TimelineIDs</td><td>statuses/friends_timeline/ids</td></tr>
    <tr><td>返回用户转发的最新微博</td><td>GET_statuses_repost_by_me</td><td>map[string]interface{}, *Statuses</td><td>statuses/repost_by_me</td></tr>
    <tr><td>获取@当前用户的最新微博</td><td>GET_statuses_mentions</td><td>map[string]interface{}, *Statuses</td><td>statusesstatuses/mentions</td></tr>
    <tr><td>获取@当前用户的最新微博的ID</td><td>GET_statuses_mentions_ids</td><td>map[string]interface{}, *TimelineIDs</td><td>statuses/mentions/ids</td></tr>
    <tr><td>获取双向关注用户的最新微博</td><td>GET_statuses_bilateral_timeline</td><td>map[string]interface{}, *Statuses</td><td>statuses/bilateral_timeline</td></tr>
    <tr><td>根据ID获取单条微博信息</td><td>GET_statuses_show</td><td>map[string]interface{}, *Status</td><td>statuses/show</td></tr>
    <tr><td>根据微博ID批量获取微博信息</td><td>GET_statuses_show_batch</td><td>map[string]interface{}, *Statuses</td><td>statuses/show_batch</td></tr>
    <tr><td>通过id获取mid</td><td>GET_statuses_querymid</td><td>map[string]interface{}, *string</td><td>statuses/querymid</td></tr>
    <tr><td>通过mid获取id</td><td>GET_statuses_queryid</td><td>map[string]interface{}, *string</td><td>statuses/queryid</td></tr>
    <tr><td>批量获取指定微博的转发数评论数</td><td>GET_statuses_count</td><td>map[string]interface{}, *[]StatusCount</td><td>statuses/count</td></tr>
    <tr><td>获取当前登录用户关注的人<br />发给其的定向微博</td><td>GET_statuses_to_me</td><td>map[string]interface{}, *Statuses</td><td>statuses/to_me</td></tr>
    <tr><td>获取当前登录用户关注的人<br />发给其的定向微博ID列表</td><td>GET_statuses_to_me_ids</td><td>map[string]interface{}, *TimelineIDs</td><td>statuses/to_me/ids</td></tr>
    <tr><td>转发一条微博信息</td><td>POST_statuses_repost</td><td>map[string]interface{}, *Status</td><td>statuses/repost</td></tr>
    <tr><td>删除微博信息</td><td>POST_statuses_destroy</td><td>map[string]interface{}, *Status</td><td>statuses/destroy</td></tr>
    <tr><td>发布一条微博信息</td><td>POST_statuses_update</td><td>map[string]interface{}, *Status</td><td>statuses/update</td></tr>
    <tr><td>上传图片并发布一条微博</td><td>POST_statuses_upload</td><td>map[string]interface{}, *Status</td><td>statusesstatuses/upload</td></tr>
    <tr><td>发布一条微博<br/>同时指定上传的图片或图片url</td><td>POST_statuses_upload_url_text</td><td>map[string]interface{}, *Status</td><td>statuses/upload_url_text</td></tr>
    <tr><td>屏蔽某条微博</td><td>POST_statuses_filter_create</td><td>map[string]interface{}, *Status</td><td>statuses/filter/create</td></tr>
    <tr><td>屏蔽某个@我的微博<br />及后续由其转发引起的@提及</td><td>POST_statuses_mentions_shield</td><td>map[string]interface{}, interface{}</td><td>statuses/mentions/shield</td></tr>
</table>
<br />

####相关数据结构如下

    type Statuses struct {
        Statuses        *[]Status     `json:"statuses"`
        Hasvisible      bool          `json:"hasvisible"`
        Previous_cursor int64         `json:"previous_cursor"`
        Next_cursor     int64         `json:"next_cursor"`
        Total_number    int64         `json:"total_number"`
        Marks           []interface{} `json:"marks,omitempty"`
    }
    
    type Reposts struct {
        Reposts         *[]Status     `json:"reposts"`
        Hasvisible      bool          `json:"hasvisible"`
        Previous_cursor int64         `json:"previous_cursor"`
        Next_cursor     int64         `json:"next_cursor"`
        Total_number    int64         `json:"total_number"`
        Marks           []interface{} `json:"marks,omitempty"`
    }

    type TimelineIDs struct {
        Statuses        []string      `json:"statuses"`
        Hasvisible      bool          `json:"hasvisible"`
        Previous_cursor int64         `json:"previous_cursor"`
        Next_cursor     int64         `json:"next_cursor"`
        Total_number    int64         `json:"total_number"`
        Marks           []interface{} `json:"marks,omitempty"`
    }

    type Status struct {
        Id                      int64       `json:"id"`
        Mid                     string      `json:"mid"`
        Idstr                   string      `json:"idstr"`
        Text                    string      `json:"text"`
        Source                  string      `json:"source"`
        Favorited               bool        `json:"favorited"`
        Truncated               bool        `json:"truncated"`
        In_reply_to_status_id   string      `json:"in_reply_to_status_id"`   //暂未支持
        In_reply_to_user_id     string      `json:"in_reply_to_user_id"`     //暂未支持
        In_reply_to_screen_name string      `json:"in_reply_to_screen_name"` //暂未支持
        Thumbnail_pic           string      `json:"thumbnail_pic"`
        Bmiddle_pic             string      `json:"bmiddle_pic"`
        Original_pic            string      `json:"original_pic"`
        Geo                     interface{} `json:"geo"` //{"type": "Point","coordinates": [21.231153,110.418708]}
        User                    *User       `json:"user,omitempty"`
        Retweeted_status        *Status     `json:"retweeted_status,omitempty"`
        Reposts_count           int64       `json:"reposts_count"`
        Comments_count          int64       `json:"comments_count"`
        Attitudes_count         int64       `json:"attitudes_count"`
        Mlevel                  int64       `json:"mlevel"`  //暂未支持
        Visible                 interface{} `json:"visible"` //{"type": 0,"list_id": 0}
        Created_at              string      `json:"created_at"`
    }
    
    
    type User struct {
        Id                 int64   `json:"id"`
        Idstr              string  `json:"idstr"`
        Screen_name        string  `json:"screen_name"`
        Name               string  `json:"name"`
        Province           string  `json:"province"`
        City               string  `json:"city"`
        Location           string  `json:"location"`
        Description        string  `json:"description"`
        Url                string  `json:"url"`
        Profile_image_url  string  `json:"profile_image_url"`
        Profile_url        string  `json:"profile_url"`
        Domain             string  `json:"domain"`
        Weihao             string  `json:"weihao"`
        Gender             string  `json:"gender"`
        Followers_count    int64   `json:"followers_count"`
        Friends_count      int64   `json:"friends_count"`
        Statuses_count     int64   `json:"statuses_count"`
        Favourites_count   int64   `json:"favourites_count"`
        Created_at         string  `json:"created_at"`
        Following          bool    `json:"following"`
        Allow_all_act_msg  bool    `json:"allow_all_act_msg"`
        Geo_enabled        bool    `json:"geo_enabled"`
        Verified           bool    `json:"verified"`
        Verified_type      int64   `json:"verified_type"`
        Remark             string  `json:"remark"`
        Status             *Status `json:"status,omitempty"`
        Allow_all_comment  bool    `json:"allow_all_comment"`
        Avatar_large       string  `json:"avatar_large"`
        Verified_reason    string  `json:"verified_reason"`
        Follow_me          bool    `json:"follow_me"`
        Online_status      int64   `json:"online_status"`
        Bi_followers_count int64   `json:"bi_followers_count"`
        Lang               string  `json:"lang"`
        Star               int64   `json:"star"`
        Mbtype             int64   `json:"mbtype"`
        Mbrank             int64   `json:"mbrank"`
        Block_word         int64   `json:"block_word"`
    }


###用户

<table>
    <tr><th width="230">接口名</th><th width="250">方法</th><th width="260">参数</th><th>对应API</th></tr>
    <tr><td>获取用户信息</td><td>GET_users_show</td><td>map[string]interface{}, *User</td><td>users/show</td></tr>
    <tr><td>通过个性域名获取用户信息</td><td>GET_users_domain_show</td><td>map[string]interface{}, *User</td><td>users/domain_show</td></tr>
    <tr><td>批量获取用户的粉丝数、关注数、微博数</td><td>GET_users_counts</td><td>map[string]interface{}, *UserCounts</td><td>users/counts</td></tr>
</table>
<br />

####相关数据结构如下
    type UserCounts struct {
        Id                    int64 `json:"id"`
        Followers_count       int64 `json:"followers_count"`
        Friends_count         int64 `json:"friends_count"`
        Statuses_count        int64 `json:"statuses_count"`
        Private_friends_count int64 `json:"private_friends_count,omitempty"`
    }


###评论

<table>
    <tr><th width="230">接口名</th><th width="250">方法</th><th width="260">参数</th><th>对应API</th></tr>
    <tr><td>获取某条微博的评论列表</td><td>GET_comments_show</td><td>map[string]interface{}, *Comments</td><td>comments/show</td></tr>
    <tr><td>我发出的评论列表</td><td>GET_comments_by_me</td><td>map[string]interface{}, *Comments</td><td>comments/by_me</td></tr>
    <tr><td>我收到的评论列表</td><td>GET_comments_to_me</td><td>map[string]interface{}, *Comments</td><td>comments/to_me</td></tr>
    <tr><td>获取用户发送及收到的评论列表</td><td>GET_comments_timeline</td><td>map[string]interface{}, *Comments</td><td>comments/timeline</td></tr>
    <tr><td>获取@到我的评论</td><td>GET_comments_mentions</td><td>map[string]interface{}, *Comments</td><td>comments/mentions</td></tr>
    <tr><td>批量获取评论内容</td><td>GET_comments_show_batch</td><td>map[string]interface{}, *Comments</td><td>comments/show_batch</td></tr>
    <tr><td>评论一条微博</td><td>POST_comments_create</td><td>map[string]interface{}, *Comments</td><td>comments/create</td></tr>
    <tr><td>删除一条评论</td><td>POST_comments_destroy</td><td>map[string]interface{}, *Comments</td><td>comments/destroy</td></tr>
    <tr><td>批量删除评论</td><td>POST_comments_destroy_batch</td><td>map[string]interface{}, *[]Comments</td><td>comments/destroy_batch</td></tr>
    <tr><td>回复一条评论</td><td>POST_comments_reply</td><td>map[string]interface{}, *Comments</td><td>comments/reply</td></tr>
</table>
<br />

####相关数据结构如下
    type Comments struct {
            Comments        *[]Comment    `json:"comments,omitempty"`
            Hasvisible      bool          `json:"hasvisible"`
            Previous_cursor int64         `json:"previous_cursor"`
            Next_cursor     int64         `json:"next_cursor"`
            Total_number    int64         `json:"total_number"`
            Marks           []interface{} `json:"marks,omitempty"`
    }
    
    type Comment struct {
        Created_at    string   `json:"created_at"`
        Id            int64    `json:"id"`
        Text          string   `json:"text"`
        Source        string   `json:"source"`
        User          *User    `json:"user,omitempty"`
        Mid           string   `json:"mid"`
        Idstr         string   `json:"idstr"`
        Status        *Status  `json:"status,omitempty"`
        Reply_comment *Comment `json:"reply_comment,omitempty"`
}


###关系

<table>
    <tr><th width="230">接口名</th><th width="250">方法</th><th width="260">参数</th><th>对应API</th></tr>
    <tr><td>获取用户的关注列表</td><td>GET_friendships_friends</td><td>map[string]interface{}, *Friendships</td><td>friendships/friends</td></tr>
    <tr><td>批量获取当前登录用户的关注人的备注信息</td><td>GET_friendships_friends_remark_batch</td><td>map[string]interface{}, interface{}</td><td>friendships/friends/remark_batch</td></tr>
    <tr><td>获取共同关注人列表</td><td>GET_friendships_friends_in_common</td><td>map[string]interface{}, *Friendships</td><td>friendships/friends/in_common</td></tr>
    <tr><td>获取双向关注列表</td><td>GET_friendships_friends_bilateral</td><td>map[string]interface{}, *Friendships</td><td>friendships/friends/bilateral</td></tr>
    <tr><td>获取双向关注UID列表</td><td>GET_friendships_friends_bilateral_ids</td><td>map[string]interface{}, *FriendsIDS</td><td>friendships/friends/bilateral/ids</td></tr>
    <tr><td>获取用户关注对象UID列表</td><td>GET_friendships_friends_ids</td><td>map[string]interface{}, *FriendsIDS</td><td>friendships/friends/ids</td></tr>
    <tr><td>获取用户粉丝列表</td><td>GET_friendships_followers</td><td>map[string]interface{}, *Friendships</td><td>friendships/followers</td></tr>
    <tr><td>获取用户粉丝UID列表</td><td>GET_friendships_followers_ids</td><td>map[string]interface{}, *FriendsIDS</td><td>friendships/followers/ids</td></tr>
    <tr><td>获取用户优质粉丝列表</td><td>GET_friendships_followers_active</td><td>map[string]interface{}, *[]User</td><td>friendships/followers/active</td></tr>
    <tr><td>获取我的关注人中关注了指定用户的人</td><td>GET_friendships_friends_chain_followers</td><td>map[string]interface{}, *Friendships</td><td>friendships/friends_chain/followers</td></tr>
    <tr><td>获取两个用户关系的详细情况</td><td>GET_friendships_show</td><td>map[string]interface{}, interface{}</td><td>friendships/show</td></tr>
    <tr><td>关注某用户</td><td>POST_friendships_create</td><td>map[string]interface{}, *User</td><td>friendships/create</td></tr>
    <tr><td>取消关注某用户</td><td>POST_friendships_destroy</td><td>map[string]interface{}, *User</td><td>friendships/destroy</td></tr>
    <tr><td>移除当前登录用户的粉丝</td><td>POST_friendships_followers_destroy</td><td>map[string]interface{}, *User</td><td>friendships/followers/destroy</td></tr>
    <tr><td>更新关注人备注</td><td>POST_friendships_remark_update</td><td>map[string]interface{}, *User</td><td>friendships/remark/update</td></tr>
</table>
<br />

    type Friendships struct {
        Users           *[]User `json:"users"`
        Previous_cursor int64   `json:"previous_cursor,omitempty"`
        Next_cursor     int64   `json:"next_cursor,omitempty"`
        Total_number    int64   `json:"total_number,omitempty"`
    }
    
    type FriendsIDS struct {
        Ids             *[]int64 `json:"ids"`
        Previous_cursor int64    `json:"previous_cursor,omitempty"`
        Next_cursor     int64    `json:"next_cursor,omitempty"`
        Total_number    int64    `json:"total_number,omitempty"`
    }
    

