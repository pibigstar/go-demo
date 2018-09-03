/**
 * Author: vuiletgo(china.violetgo@gmail.com)
 * Date: 13-12-08
 * Version: 0.02
 */
package weigo

//通过地址编码获取地址名称
func (api *APIClient) GET_common_code_to_location(params map[string]interface{}, result interface{}) error {
	return api.GET("common/code_to_location", params, result)
}

//获取城市列表
func (api *APIClient) GET_common_get_city(params map[string]interface{}, result interface{}) error {
	return api.GET("common/get_city", params, result)
}

//获取省份列表
func (api *APIClient) GET_common_get_province(params map[string]interface{}, result interface{}) error {
	return api.GET("common/get_province", params, result)
}

//获取国家列表
func (api *APIClient) GET_common_get_country(params map[string]interface{}, result interface{}) error {
	return api.GET("common/get_country", params, result)
}

//获取时区配置表
func (api *APIClient) GET_common_get_timezone(params map[string]interface{}, result interface{}) error {
	return api.GET("common/get_timezone", params, result)
}
