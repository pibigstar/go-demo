package weigo

//通过id获取mid
func (api *APIClient) get_statuses_querymid(params map[string]interface{}, result interface{}) error {
	return api.GET("emotions", params, result)
}
