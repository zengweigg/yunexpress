package model

type YunBlank struct {
}

type YunApiResponse struct {
	Message string `json:"Message"` // 消息内容
	Code    string `json:"Code"`    // 消息编码 0000:成功
}
