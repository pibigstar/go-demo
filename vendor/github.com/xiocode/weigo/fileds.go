/**
 * Author: Tony.Shao(xiocode@gmail.com)
 * Date: 13-03-06
 * Version: 0.02
 * modify by violetgo
 */
package weigo

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

//UserTimeline, HomeTimeline...
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

type Comments struct {
	Comments        *[]Comment    `json:"comments,omitempty"`
	Hasvisible      bool          `json:"hasvisible"`
	Previous_cursor int64         `json:"previous_cursor"`
	Next_cursor     int64         `json:"next_cursor"`
	Total_number    int64         `json:"total_number"`
	Marks           []interface{} `json:"marks,omitempty"`
}

type UserCounts struct {
	Id                    int64 `json:"id"`
	Followers_count       int64 `json:"followers_count"`
	Friends_count         int64 `json:"friends_count"`
	Statuses_count        int64 `json:"statuses_count"`
	Private_friends_count int64 `json:"private_friends_count,omitempty"`
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

type UserRank struct {
	Uid  int64 `json:"uid"`
	Rank int64 `json:"rank"`
}

//friends_timeline_ids, repost_timeline_ids
type TimelineIDs struct {
	Statuses        []string      `json:"statuses"`
	Hasvisible      bool          `json:"hasvisible"`
	Previous_cursor int64         `json:"previous_cursor"`
	Next_cursor     int64         `json:"next_cursor"`
	Total_number    int64         `json:"total_number"`
	Marks           []interface{} `json:"marks,omitempty"`
}

type StatusCount struct {
	Id        int64 `json:"id"`
	Comments  int64 `json:"comments"`
	Reposts   int64 `json:"reposts"`
	Attitudes int64 `json:"attitudes,omitempty"`
}

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

type Config struct {
	Comment  int64 `json:"comment"`
	Geo      int64 `json:"geo"`
	Message  int64 `json:"message"`
	Realname int64 `json:"realname"`
	Badge    int64 `json:"badge"`
	Mobile   int64 `json:"mobile"`
	Webim    int64 `json:"webim"`
}

type LimitStatus struct {
	Ip_limit              int64  `json:"ip_limit"`
	Limit_time_unit       string `json:"limit_time_unit"`
	Remaining_ip_hits     int64  `json:"remaining_ip_hits"`
	Remaining_user_hits   int64  `json:"remaining_user_hits"`
	Reset_time            string `json:"reset_time"`
	Reset_time_in_seconds int64  `json:"reset_time_in_seconds"`
	User_limit            int64  `json:"user_limit"`
}

type UserID struct {
	Uid int64 `json:"uid"`
}

type Email struct {
	Email string `json:"email"`
}

type School struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type Tags struct {
	Tags []map[string]interface{} `json:"tags"`
	Id   int64                    `json:"id"`
}

type Favorite struct {
	Status         *Status `json:"status"`
	Tags           *Tags   `json:"tags"`
	Favorited_time string  `json:"favorited_time"`
}

type Favorites struct {
	Favorites    *[]Favorite `json:"favorites"`
	Total_number int64       `json:"total_number"`
}

type FavoriteID struct {
	StatusID       int64  `json:"status"`
	Tags           *Tags  `json:"tags"`
	Favorited_time string `json:"favorited_time"`
}

type FavoritesID struct {
	Favorites    *[]FavoriteID `json:"favorites"`
	Total_number int64         `json:"total_number"`
}

type Topic struct {
	Statuses     *[]Status `json:"statuses"`
	Total_number int64     `json:"total_number"`
}

type SuggestionsUser struct {
	Users        *[]User `json:"users"`
	Total_number int64   `json:"total_number"`
}
