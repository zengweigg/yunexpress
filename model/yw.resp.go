package model

type YWBlank struct {
}

type YWApiResponse struct {
	Msg     string `json:"message"` // 消息内容
	Code    string `json:"code"`    // 消息编码 0:成功 >0:失败 <0:系统异常
	Success bool   `json:"success"` // 是否成功 true:成功 false:失败
}
