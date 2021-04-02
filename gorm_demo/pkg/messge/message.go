package messge

var msgMap = map[int]string{
	SUCCESS:        "请求成功",
	FAIL:           "请求失败",
	PARAMSFAIL:     "参数异常",
	AUTHTIME:       "接口加密时间错误",
	AUTHSESSIONKEY: "sessionkey非法",
	AUTHRESERR:     "加密结果验证异常",
}

// 根据错误id获取错误的详细信息
func GetMessage(code int) string {
	msg, ok := msgMap[code]
	if ok {
		return msg
	}

	return msgMap[FAIL]
}
