package example

// messages get msg from const comment
var messages = map[int]string{

	Test1: "测试1",
	Test2: "测试2",
	Test3: "测试3",
	Test4: "测试4",
}

// GetErrMsg get error msg
func GetErrMsg(code int) string {
	if msg, ok := messages[code]; ok {
		return msg
	}
	return ""
}
