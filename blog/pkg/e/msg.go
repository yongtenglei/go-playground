package e

var MsgMaps = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "请求参数错误",

	ERROR_EXIST_TAG:     "已存在该标签名称",
	ERROR_NOT_EXIST_TAG: "该标签不存在",
}

func GetMsg(code int) string {
	msg, ok := MsgMaps[code]
	if ok {
		return msg
	}
	return MsgMaps[ERROR]
}
