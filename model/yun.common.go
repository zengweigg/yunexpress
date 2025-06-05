package model

type YunCountry struct {
	CountryCode string `json:"CountryCode"` // 国家代码
	EName       string `json:"EName"`       // 英文名称
	CName       string `json:"CName"`       // 中文名称
}

type YunCountryResp struct {
	YunApiResponse
	Items []YunCountry `json:"Items"`
}

type YunShipping struct {
	Code              string `json:"Code"`              // 代码
	CName             string `json:"CName"`             // 中文名称
	EName             string `json:"EName"`             // 英文名称
	HasTrackingNumber bool   `json:"HasTrackingNumber"` // 是否有跟踪号
	DisplayName       string `json:"DisplayName"`       // 显示名称
}

type QueryShippingData struct {
	CountryCode string `json:"CountryCode,omitempty"` // 国家简码，未填写国家代表查询所有运输方式
}

type YunShippingResp struct {
	YunApiResponse
	Items []YunCountry `json:"Items"`
}
