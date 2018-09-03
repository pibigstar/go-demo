/**
 * Author: vuiletgo(china.violetgo@gmail.com)
 * Date: 13-12-15
 * Version: 0.02
 */
package weigo

//返回最近一小时内的热门话题
func (api *APIClient) GET_trends_hourly(params map[string]interface{}, result interface{}) error {
	return api.GET("trends/hourly", params, result)
}

//返回最近一天内的热门话题
func (api *APIClient) GET_trends_daily(params map[string]interface{}, result interface{}) error {
	return api.GET("trends/daily", params, result)
}

//返回最近一周内的热门话题
func (api *APIClient) GET_trends_weekly(params map[string]interface{}, result interface{}) error {
	return api.GET("searchtrends/weekly", params, result)
}
