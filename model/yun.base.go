package model

// 通用ApiResponse
type YunApiResponse struct {
	Message string `json:"Message"` // 消息内容
	Code    string `json:"Code"`    // 消息编码 0000:成功
}

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

type YunProductType struct {
	Id    string `json:"Id"`    // 货品类型 ID
	CName string `json:"CName"` // 货品类型中文名称
	EName string `json:"EName"` // 货品类型英文名称
}

type YunProductTypeResp struct {
	YunApiResponse
	Items []YunProductType `json:"Items"`
}

type ExtraService struct {
	ExtraService string `json:"ExtraService"` // 保价服务代码，必填，最大长度50 BJFA：固定金额 BJFR：运费*费率 BJDR：申报价值*费率 BJFDR：(运费+申报价值)*费率
}

type YunPriceReq struct {
	CountryCode       string         `json:"CountryCode"`                 // 国家简码，必填，最大长度5
	Weight            float64        `json:"Weight"`                      // 包裹重量(kg)，必填，支持3位小数
	Length            int            `json:"Length,omitempty"`            // 包裹长度(cm)，非必填，默认1
	Width             int            `json:"Width,omitempty"`             // 包裹宽度(cm)，非必填，默认1
	Height            int            `json:"Height,omitempty"`            // 包裹高度(cm)，非必填，默认1
	PackageType       int            `json:"PackageType"`                 // 包裹类型，1-带电，0-普货，必填，默认1
	Origin            string         `json:"Origin,omitempty"`            // 计费地点，非必填，默认YT-SZ
	ExtraServicesList []ExtraService `json:"ExtraServicesList,omitempty"` // 保价服务代码列表，非必填
}

type Quotation struct {
	Code             string  `json:"Code"`        // 可服务的运输方式代码
	CName            string  `json:"CName"`       // 可服务的运输方式中文名称
	EName            string  `json:"EName"`       // 可服务的运输方式英文名称
	ShippingFee      float64 `json:"ShippingFee"` // 基础运费
	SignatureFee     float64 `json:"SignatureFee"`
	RegistrationFee  float64 `json:"RegistrationFee"` // 挂号费
	FuelFee          float64 `json:"FuelFee"`         // 燃油费
	SundryFee        float64 `json:"SundryFee"`       // 杂费
	TariffPrepayFee  float64 `json:"TariffPrepayFee"` // 关税预付服务费
	InsuredFee       float64 `json:"InsuredFee"`      // 保价费
	TotalFee         float64 `json:"TotalFee"`        // 总费用
	Weight           float64 `json:"Weight"`
	DeliveryDays     string  `json:"DeliveryDays"` // 预计时效
	Remark           string  `json:"Remark"`       // 说明(当 Remark 为空时不显示)
	Track            string  `json:"Track"`
	Currency         string  `json:"Currency"`
	ProductGroupname string  `json:"product_groupname"`
	GoodsType        string  `json:"GoodsType"`
}

type YunPriceResp struct {
	YunApiResponse
	Items []Quotation `json:"Items"`
}

type OrderSender struct {
	OrderNumber   string `json:"OrderNumber"`   // 客户输入的查询号
	WayBillNumber string `json:"WayBillNumber"` // 运单号
	CountryCode   string `json:"CountryCode"`   // 发件人国家简码
	FirstName     string `json:"FirstName"`     // 发件人姓
	LastName      string `json:"LastName"`      // 发件人名
	Company       string `json:"Company"`       // 发件人公司
	Address       string `json:"Address"`       // 发件人地址
	City          string `json:"City"`          // 发件人城市
	State         string `json:"State"`         // 发件人省州
	Zip           string `json:"Zip"`           // 发件人邮编
	Phone         string `json:"Phone"`         // 发件人电话
}

type YunSenderResp struct {
	YunApiResponse
	Items OrderSender `json:"Items"`
}

type YunRegisterReq struct {
	UserName  string `json:"UserName"`  // 用户名(50)，必填
	PassWord  string `json:"PassWord"`  // 密码(50)，必填
	Contact   string `json:"Contact"`   // 联系人(200)，必填
	Mobile    string `json:"Mobile"`    // 联系人电话(200)，必填
	Telephone string `json:"Telephone"` // 联系人电话(200)，必填
	Name      string `json:"Name"`      // 客户名称/公司名称(50)，必填
	Email     string `json:"Email"`     // 邮箱(50)，必填
	Address   string `json:"Address"`   // 详细地址(50)，必填
	PlatForm  int    `json:"PlatForm"`  // 平台ID(1)，必填(通途平台--2)
}

type CustomerModel struct {
	CustomerCode string `json:"CustomerCode"` // 用户代码
	UserName     string `json:"UserName"`     // 用户登录账号
	SecretKey    string `json:"SecretKey"`    // API对接密钥
}

type YunRegisterResp struct {
	YunApiResponse
	Item []CustomerModel `json:"Item"`
}

type YunCarrierPost struct {
	OrderNumbers []string `json:"OrderNumbers"`
}

type CarrierInfo struct {
	OrderCode      string `json:"OrderCode"`      // 请求查询单号
	OrderCodeType  string `json:"OrderCodeType"`  // 查询单号类型(C-客户单号,Y-运单号,T-跟踪号)
	CarrierCode    string `json:"CarrierCode"`    // 末端派送商代码
	CarrierCName   string `json:"CarrierCName"`   // 末端派送商中文名
	CarrierEName   string `json:"CarrierEName"`   // 末端派送商英文名
	CarrierPhone   string `json:"CarrierPhone"`   // 联系电话
	CarrierWebsite string `json:"CarrierWebsite"` // 官网
}

type YunCarrierResp struct {
	YunApiResponse
	Item []CarrierInfo `json:"Item"`
}

type YunRegisterIossPost struct {
	IossType     string   `json:"IossType" validate:"required,len=1"`                                  // IOSS类型(0-个人,1-平台)，必填
	PlatformName string   `json:"PlatformName,omitempty" maxLength:"100"`                              // 平台名称(100)，类型为"1"时需提供
	IossNumber   string   `json:"IossNumber" validate:"required,len=12,regexp=^[a-zA-Z]{2}[0-9]{10}$"` // IOSS号(12)，必填，2位字母+10位数字
	Company      string   `json:"Company,omitempty" maxLength:"200"`                                   // IOSS号注册公司名称(200)
	Country      string   `json:"Country,omitempty" maxLength:"2"`                                     // 国家简码(2)
	Street       string   `json:"Street,omitempty" maxLength:"200"`                                    // 街道地址(200)
	City         string   `json:"City,omitempty" maxLength:"100"`                                      // 所在城市(100)
	Province     string   `json:"Province,omitempty" maxLength:"100"`                                  // 所在省/州(100)
	PostalCode   string   `json:"PostalCode,omitempty" maxLength:"50"`                                 // 邮编(50)
	MobilePhone  string   `json:"MobilePhone,omitempty" maxLength:"200"`                               // 手机号(200)
	Email        string   `json:"Email,omitempty" maxLength:"100"`                                     // 电子邮箱(100)
	FileUrl      []string `json:"FileUrl,omitempty"`                                                   // IOSS注册文件，IOSS号审核不通过、发起修改时必填
}

type YunRegisterIossResp struct {
	YunApiResponse
	IossCode   string `json:"IossCode"`   // 备案成功返回的IOSS号对应的备案识别码
	RequestId  string `json:"RequestId"`  // 请求唯一标识
	TimeStamp  string `json:"TimeStamp"`  // 时间戳
	IossStatus string `json:"IossStatus"` // IOSS状态(1-通过,3-待审核)
	IossNote   string `json:"IossNote"`   // IOSS备注
}

type YunDisableIossPost struct {
	Code string `json:"code" maxLength:"50"` // IOSS号备案返回的识别码
}

type YunDisableIossResp struct {
	YunApiResponse
	Item bool `json:"Item"`
}
