package model

import (
	"gopkg.in/guregu/null.v4"
	"time"
)

type ChildTrackingInfo struct {
	ChildOrderNumber    string      `json:"ChildOrderNumber"`    // 子单号
	ChildTrackingNumber null.String `json:"ChildTrackingNumber"` // 子单跟踪号
}

type OrderTracking struct {
	CustomerOrderNumber string              `json:"CustomerOrderNumber"` // 订单号
	TrackingNumber      string              `json:"TrackingNumber"`      // 跟踪号
	WayBillNumber       string              `json:"WayBillNumber"`       // 运单号
	Status              int                 `json:"Status"`              // 1-跟踪号生成成功;2-跟踪号生成中;3- 跟踪号无需生成;4-跟踪号生成失败
	Remark              string              `json:"Remark"`              // 详情
	ChildTrackings      []ChildTrackingInfo `json:"ChildTrackings"`      // 子单跟踪号信息
}

type YunTrackingNumberResp struct {
	YunApiResponse
	Items []OrderTracking `json:"Items"`
}

type OrderTrackingInfo struct {
	CountryCode          string                `json:"CountryCode"`          // 目的地国家简码
	WaybillNumber        string                `json:"WaybillNumber"`        // 运单号
	TrackingNumber       string                `json:"TrackingNumber"`       // 跟踪号
	ProviderName         string                `json:"ProviderName"`         // 末端服务商名称
	ProviderTelephone    string                `json:"ProviderTelephone"`    // 末端服务商联系方式
	ProviderSite         string                `json:"ProviderSite"`         // 末端服务商官网
	POD                  string                `json:"POD"`                  // 单个POD链接(URL地址)
	PODs                 []string              `json:"PODs"`                 // 多个POD链接(URL地址)
	CreatedBy            string                `json:"CreatedBy"`            // 创建人
	PackageState         int                   `json:"PackageState"`         // 包裹状态(0-未知,1-已提交,2-运输中,3-已签收,4-已收货,5-订单取消,6-投递失败,7-已退回)
	IntervalDays         int                   `json:"IntervalDays"`         // 包裹签收天数
	OrderTrackingDetails []OrderTrackingDetail `json:"OrderTrackingDetails"` // 订单跟踪详情
}

type OrderTrackingDetail struct {
	ProcessDate          time.Time        `json:"ProcessDate"`          // 包裹请求日期
	ProcessContent       string           `json:"ProcessContent"`       // 包裹请求内容
	ProcessLocation      string           `json:"ProcessLocation"`      // 包裹请求地址
	TrackNodeCode        string           `json:"TrackNodeCode"`        // 轨迹节点代码
	TrackCodeDescription string           `json:"TrackCodeDescription"` // 轨迹节点英文描述
	ProcessCountry       string           `json:"ProcessCountry"`       // 轨迹发生地所在国家
	ProcessProvince      string           `json:"ProcessProvince"`      // 轨迹发生地所在省州
	ProcessCity          string           `json:"ProcessCity"`          // 轨迹发生地所在城市
	AbnormalReasons      []AbnormalReason `json:"AbnormalReasons"`      // 轨迹异常原因
}

type AbnormalReason struct {
	AbnormalReasonCode string `json:"AbnormalReasonCode"` // 轨迹异常原因代码
	AbnormalReason     string `json:"AbnormalReason"`     // 轨迹异常原因描述
}

type YunTrackInfoResp struct {
	YunApiResponse
	Item OrderTrackingInfo `json:"Item"`
}

type YunCreatedOrderSubscribePost struct {
	DisplayMode     int      `json:"DisplayMode" validate:"required"`      // 轨迹内容:0-全程轨迹;1-头程轨迹;2-末端轨迹;3-隐藏轨迹;4-电子预报信息
	QueryMode       int      `json:"QueryMode" validate:"required"`        // 查询单号：1-运单号;2-订单号;4-跟踪号
	ShipperHawbcode []string `json:"shipper_hawbcode" validate:"required"` // 运单号数组
}

type CreatedOrderSubscribe struct {
	Success   bool          `json:"Success"`   // 状态：true-成功；false-失败
	ErrorCode int           `json:"ErrorCode"` // 0成功1000参数错误
	Message   string        `json:"Message"`   // 返回的错误消息
	Data      SubscribeData `json:"Data"`      // 返回数据内容
}

type SubscribeData struct {
	NoCount     int      `json:"NoCount"`     // 订阅失败的数量
	OkCount     int      `json:"OkCount"`     // 订阅成功数量
	FailNumbers []string `json:"FailNumbers"` // 订阅失败的业务ID
}

type YunCreatedOrderSubscribeResp struct {
	YunApiResponse
	Item CreatedOrderSubscribe `json:"Item"`
}

type YunOrderSubscribePost struct {
	ShipperHawbcode []string `json:"shipper_hawbcode,omitempty"`      // 运单号集合(非必填)
	CurrentPage     int      `json:"currentpage" validate:"required"` // 当前页(必填)
	PageSize        int      `json:"pagesize" validate:"required"`    // 页数(必填)
	StartDate       string   `json:"StartDate" validate:"required"`   // 订阅时间起始(必填)
	EndDate         string   `json:"EndDate" validate:"required"`     // 订阅时间结束(必填)
}

type YunOrderSubscribe struct {
	Success bool                 `json:"Success"` // 状态：true-成功；false-失败
	Data    SubscriptionResponse `json:"Data"`    // 单号订阅信息
}

type SubscriptionResponse struct {
	CountNum int                     `json:"CountNum"` // 查询总数
	Item     []YunOrderSubscribeItem `json:"Item"`     // 订阅信息
}

type YunOrderSubscribeItem struct {
	Id              int64       `json:"Id"`              // Id
	CustomerOrderId string      `json:"CustomerOrderId"` // 客户订单号
	WaybillNumber   string      `json:"WaybillNumber"`   // 运单号
	TrackingNumber  null.String `json:"TrackingNumber"`  // 跟踪号
	SubscribeType   int         `json:"SubscribeType"`   // 订阅类型(轨迹内容):0-全程轨迹;1-头程轨迹;2-末端轨迹;3-隐藏轨迹;4-电子预报信息+末端轨迹
	QueryNumber     string      `json:"QueryNumber"`     // 查询单号：1-运单号;2-订单号;4-跟踪号;3-运单号、订单号;5-运单号、跟踪号;6-订单号、跟踪号;7-运单号、订单号、跟踪号
	SubscribeTime   time.Time   `json:"SubscribeTime"`   // 订阅时间(DateTime)
	Subscriber      string      `json:"Subscriber"`      // 订阅人
}

type YunOrderSubscribeResp struct {
	YunApiResponse
	Item YunOrderSubscribe `json:"Item"`
}

type YunCreatedProductSubscribePost struct {
	DisplayMode  int                `json:"display_mode"`            // 轨迹内容: 0: 全程轨迹;1: 头程轨迹(国内轨迹);2: 末端轨迹(国外派送轨迹);3: 隐藏轨迹(不显示任何轨迹);4: 电子预报信息
	ProductInfo  []ProductSubscribe `json:"product_info"`            // 产品信息
	QueryMode    int                `json:"query_mode,omitempty"`    // 查询单号： 1:运单号;2:订单号;4:跟踪号;3:运单号、订单号;5:运单号、跟踪号;6:订单号、跟踪号;7:运单号、订单号、跟踪号;
	CountryCodes []string           `json:"country_codes,omitempty"` // 国家:为空即所有国家
}

type ProductSubscribe struct {
	ProductName string `json:"product_name"` // 产品名称
	ProductCode string `json:"product_code"` // 产品代码
}

type ProductSubscribeResp struct {
	Success   bool          `json:"Success"`   // 状态：true-成功；false-失败
	ErrorCode int           `json:"ErrorCode"` // 0成功1000参数错误
	Message   string        `json:"Message"`   // 返回的错误消息
	Data      SubscribeData `json:"Data"`      // 返回数据内容
}

type YunCreatedProductSubscribeResp struct {
	YunApiResponse
	Item ProductSubscribeResp `json:"Item"`
}

type YunCancelProductSubscribePost struct {
	Ids []string `json:"ids"`
}

type YunCancelProductSubscribeResp struct {
	YunApiResponse
	Item ProductSubscribeResp `json:"Item"`
}

type YunGetProductSubscribePost struct {
	DisplayMode string `json:"display_mode,omitempty"`        // 订阅类型(0:全程轨迹;1:头程轨迹;2:末端轨迹;3:隐藏轨迹;4:电子预报信息)
	ProductCode string `json:"product_code,omitempty"`        // 产品代码
	PageCount   int    `json:"pageCount" validate:"required"` // 页数(必填)
	PageIndex   int    `json:"pageIndex" validate:"required"` // 当前页(必填)
}

type ProductSubs struct {
	Success bool           `json:"Success"` // 状态：true-成功；false-失败
	Data    ProductSubData `json:"Data"`    // 单号订阅信息
}

type ProductSubData struct {
	CountNum int                  `json:"CountNum"` // 查询总数
	Item     []ProductSubDataItem `json:"Item"`     // 单号订阅信息
}

type ProductSubDataItem struct {
	Id            int64     `json:"Id"`            // 订阅Id
	UserCode      string    `json:"UserCode"`      // 客户编码
	ProductCode   string    `json:"ProductCode"`   // 销售产品代码
	ProductName   string    `json:"ProductName"`   // 销售产品名称
	Status        int       `json:"Status"`        // 状态(0-未订阅,1-已订阅)
	CountryCode   string    `json:"CountryCode"`   // 国家二字简码
	CountryName   string    `json:"CountryName"`   // 国家名称
	SubscribeType int       `json:"SubscribeType"` // 订阅类型:0-全程轨迹;1-头程轨迹;2-末端轨迹;3-隐藏轨迹;4-电子预报信息+末端轨迹
	QueryNumber   string    `json:"QueryNumber"`   // 查询单号:1-运单号;2-订单号;4-跟踪号;3-运单号+订单号;5-运单号+跟踪号;6-订单号+跟踪号;7-运单号+订单号+跟踪号
	SubscribeTime time.Time `json:"SubscribeTime"` // 订阅时间(DateTime)
	Operator      string    `json:"Operator"`      // 订阅人
}

type YunGetProductSubscribeResp struct {
	YunApiResponse
	Item ProductSubs `json:"Item"`
}
