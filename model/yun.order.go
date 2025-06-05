package model

import "gopkg.in/guregu/null.v4"

type Receiver struct {
	CountryCode     string `json:"CountryCode"`               // 收件人国家(2位简码,50)
	FirstName       string `json:"FirstName"`                 // 收件人姓(50)
	LastName        string `json:"LastName,omitempty"`        // 收件人名(50)
	Company         string `json:"Company,omitempty"`         // 收件人公司(200)
	Street          string `json:"Street"`                    // 收件人详细地址(200)
	StreetAddress1  string `json:"StreetAddress1,omitempty"`  // 收件人详细地址1(200)
	StreetAddress2  string `json:"StreetAddress2,omitempty"`  // 收件人详细地址2(200)
	City            string `json:"City"`                      // 收件人城市(100)
	State           string `json:"State,omitempty"`           // 收件人省/州(100)
	Zip             string `json:"Zip,omitempty"`             // 收件人邮编(50)
	Phone           string `json:"Phone,omitempty"`           // 收件人电话(200)
	HouseNumber     string `json:"HouseNumber,omitempty"`     // 街道门牌号(50)
	Email           string `json:"Email,omitempty"`           // 收件人邮箱(100)
	MobileNumber    string `json:"MobileNumber,omitempty"`    // 收件人手机号(50)
	CertificateCode string `json:"CertificateCode,omitempty"` // 收件人ID(50)
}

type Sender struct {
	CountryCode string `json:"CountryCode,omitempty"` // 发件人国家(2位简码,10)
	FirstName   string `json:"FirstName,omitempty"`   // 发件人姓(50)
	LastName    string `json:"LastName,omitempty"`    // 发件人名(50)
	Company     string `json:"Company,omitempty"`     // 发件人公司(500)
	Street      string `json:"Street,omitempty"`      // 发件人地址(200)
	City        string `json:"City,omitempty"`        // 发件人城市(100)
	State       string `json:"State,omitempty"`       // 发件人省/州(100)
	Zip         string `json:"Zip,omitempty"`         // 发件人邮编(50)
	Phone       string `json:"Phone,omitempty"`       // 发件人电话(20)
	Email       string `json:"Email,omitempty"`       // 发件人邮箱(20)
	UsciCode    string `json:"UsciCode,omitempty"`    // 统一社会信用代码(50)
}

type Parcel struct {
	EName         string  `json:"EName"`                   // 申报名称(英文,200)
	CName         string  `json:"CName,omitempty"`         // 申报名称(中文,500)
	HSCode        string  `json:"HSCode,omitempty"`        // 海关编码(50)
	Quantity      int     `json:"Quantity"`                // 申报数量
	UnitPrice     float64 `json:"UnitPrice"`               // 申报价格(单价)
	UnitWeight    float64 `json:"UnitWeight"`              // 申报重量(单重,kg)
	Remark        string  `json:"Remark,omitempty"`        // 订单备注(500)
	ProductUrl    string  `json:"ProductUrl,omitempty"`    // 产品销售链接(200)
	SKU           string  `json:"SKU,omitempty"`           // 商品SKU(50)
	InvoiceRemark string  `json:"InvoiceRemark,omitempty"` // 配货信息(50)
	CurrencyCode  string  `json:"CurrencyCode"`            // 申报币种(50)
	Attachment    string  `json:"Attachment,omitempty"`    // SKU附件(255)
	InvoicePart   string  `json:"InvoicePart,omitempty"`   // 材质(255)
	InvoiceUsage  string  `json:"InvoiceUsage,omitempty"`  // 用途(255)
}

type OrderExtra struct {
	ExtraCode  string `json:"ExtraCode"`            // 额外服务代码(20)
	ExtraName  string `json:"ExtraName"`            // 额外服务名称(20)
	ExtraValue string `json:"ExtraValue,omitempty"` // 额外服务值(30)
	ExtraNote  string `json:"ExtraNote,omitempty"`  // 额外服务备注(200)
}

type ChildOrder struct {
	BoxNumber    string        `json:"BoxNumber,omitempty"`    // 箱子编号(200)
	ReferenceID  string        `json:"ReferenceID,omitempty"`  // FBA订单参考ID(100)
	Length       int           `json:"Length,omitempty"`       // 预估包裹单边长(cm)
	Width        int           `json:"Width,omitempty"`        // 预估包裹单边宽(cm)
	Height       int           `json:"Height,omitempty"`       // 预估包裹单边高(cm)
	BoxWeight    float64       `json:"BoxWeight,omitempty"`    // 预估包裹总重量(kg)
	ChildDetails []ChildDetail `json:"ChildDetails,omitempty"` // 单箱SKU信息
}

type ChildDetail struct {
	SKU      string `json:"SKU,omitempty"`      // 商品SKU(50)
	Quantity int    `json:"Quantity,omitempty"` // 申报数量
}

type Platform struct {
	PlatformName string `json:"PlatformName,omitempty"` // 平台名称(50)
	Province     string `json:"Province,omitempty"`     // 州省(50)
	Address      string `json:"Address,omitempty"`      // 地址(200)
	PostalCode   string `json:"PostalCode,omitempty"`   // 邮编(20)
	PhoneNumber  string `json:"PhoneNumber,omitempty"`  // 联系电话(50)
	Email        string `json:"Email,omitempty"`        // 邮箱(100)
}

type YunOrderPost struct {
	CustomerOrderNumber  string       `json:"CustomerOrderNumber"`            // 客户订单号,不能重复(50)
	ShippingMethodCode   string       `json:"ShippingMethodCode"`             // 运输方式代码(50)
	TrackingNumber       string       `json:"TrackingNumber,omitempty"`       // 包裹跟踪号(50)
	TransactionNumber    string       `json:"TransactionNumber,omitempty"`    // 平台交易号(50)
	TaxNumber            string       `json:"TaxNumber,omitempty"`            // 增值税号(50)
	EoriNumber           string       `json:"EoriNumber,omitempty"`           // 欧盟税号(50)
	TaxCountryCode       string       `json:"TaxCountryCode,omitempty"`       // 税号所属国家(2)
	ImporterName         string       `json:"ImporterName,omitempty"`         // 进口商名称(255)
	ImporterAddress      string       `json:"ImporterAddress,omitempty"`      // 进口商地址(255)
	TaxProve             string       `json:"TaxProve,omitempty"`             // 税号证明(255)
	TaxRemark            string       `json:"TaxRemark,omitempty"`            // 税号备注(255)
	WarehouseAddressCode string       `json:"WarehouseAddressCode,omitempty"` // 仓库代码(50)
	SizeUnits            string       `json:"SizeUnits,omitempty"`            // 长度单位(cm/inch)
	Length               int          `json:"Length,omitempty"`               // 预估包裹单边长
	Width                int          `json:"Width,omitempty"`                // 预估包裹单边宽
	Height               int          `json:"Height,omitempty"`               // 预估包裹单边高
	PackageCount         int          `json:"PackageCount"`                   // 运单包裹的件数(>0)
	Weight               float64      `json:"Weight"`                         // 预估包裹总重量(kg,3位小数)
	Receiver             Receiver     `json:"Receiver"`                       // 收件人信息
	Sender               Sender       `json:"Sender,omitempty"`               // 发件人信息
	ApplicationType      int          `json:"ApplicationType,omitempty"`      // 申报类型(1-Gift,2-Sameple,3-Documents,4-Others)
	ReturnOption         bool         `json:"ReturnOption,omitempty"`         // 是否退回(1-退回,0-不退回)
	TariffPrepay         bool         `json:"TariffPrepay,omitempty"`         // 关税预付(1-参加,0-不参加)
	InsuranceOption      int          `json:"InsuranceOption,omitempty"`      // 包裹投保类型(0-不参保,1-按件,2-按比例)
	Coverage             float64      `json:"Coverage,omitempty"`             // 保险的最高额度(RMB)
	SensitiveTypeID      int          `json:"SensitiveTypeID,omitempty"`      // 特殊货品类型
	Parcels              []Parcel     `json:"Parcels"`                        // 申报信息
	SourceCode           string       `json:"SourceCode,omitempty"`           // 订单来源代码(3)
	ChildOrders          []ChildOrder `json:"ChildOrders,omitempty"`          // 箱子明细信息(FBA必填)
	OrderExtra           []OrderExtra `json:"OrderExtra,omitempty"`           // 附加服务
	IossCode             string       `json:"IossCode,omitempty"`             // IOSS备案识别码(50)
	DangerousGoodsType   string       `json:"DangerousGoodsType,omitempty"`   // 危险品类型(30)
	Platform             Platform     `json:"Platform,omitempty"`             // 平台信息
}

type OrderResponse struct {
	CustomerOrderNumber  string        `json:"CustomerOrderNumber"`  // 客户的订单号
	Success              int           `json:"Success"`              // 运单请求状态，1-成功，0-失败
	TrackType            string        `json:"TrackType"`            // 1-已产生跟踪号，2-等待后续更新跟踪号,3-不需要跟踪号
	Remark               string        `json:"Remark"`               // 运单请求失败反馈失败原因
	RequireSenderAddress int           `json:"RequireSenderAddress"` // 0-不需要分配地址，1-需要分配地址
	AgentNumber          null.String   `json:"AgentNumber"`          // 代理单号
	WayBillNumber        string        `json:"WayBillNumber"`        // YT 运单号
	TrackingNumber       null.String   `json:"TrackingNumber"`       // 跟踪号
	ShipperBoxs          []ShipperBoxs `json:"ShipperBoxs"`          // 箱子信息
}

type ShipperBoxs struct {
	BoxNumber       string `json:"BoxNumber"`       // 箱子号码
	ShipperHawbcode string `json:"ShipperHawbcode"` // 物流运单子单号
}

type YunOrderResp struct {
	YunApiResponse
	Item OrderResponse `json:"Item"`
}

type OrderDetail struct {
	WayBillNumber      string       `json:"WayBillNumber"`               // 物流系统运单号
	OrderNumber        string       `json:"OrderNumber"`                 // 客户订单号
	ShippingMethodCode string       `json:"ShippingMethodCode"`          // 发货的方式
	TrackingNumber     string       `json:"TrackingNumber,omitempty"`    // 包裹跟踪号
	TransactionNumber  string       `json:"TransactionNumber,omitempty"` // 平台交易号
	Length             int          `json:"Length"`                      // 预估包裹单边长(cm)
	Width              int          `json:"Width"`                       // 预估包裹单边宽(cm)
	Height             int          `json:"Height"`                      // 预估包裹单边高(cm)
	Status             int          `json:"Status"`                      // 订单状态
	PackageNumber      int          `json:"PackageNumber"`               // 运单的包裹件数
	Weight             float64      `json:"Weight"`                      // 预估包裹总重量(kg)
	Receiver           ReceiverResp `json:"Receiver"`                    // 收件人信息
	Sender             SenderResp   `json:"Sender,omitempty"`            // 发件人信息
	ApplicationType    int          `json:"ApplicationType"`             // 申报类型
	ReturnOption       bool         `json:"ReturnOption"`                // 是否退回
	InsuranceType      int          `json:"InsuranceType"`               // 包裹投保类型
	InsureAmount       float64      `json:"InsureAmount"`                // 保险的最高额度(RMB)
	SensitiveTypeID    int          `json:"SensitiveTypeID,omitempty"`   // 特殊货品类型
	Parcels            []ParcelResp `json:"Parcels"`                     // 申报信息
	ChargeWeight       float64      `json:"ChargeWeight"`                // 计费重量
}

type ReceiverResp struct {
	TaxId          string `json:"TaxId,omitempty"`          // 收件人企业税号
	CountryCode    string `json:"CountryCode"`              // 收件人国家简码
	FirstName      string `json:"FirstName"`                // 收件人姓
	LastName       string `json:"LastName"`                 // 收件人名字
	Company        string `json:"Company,omitempty"`        // 收件人公司名称
	Street         string `json:"Street"`                   // 收件人详情地址
	StreetAddress1 string `json:"StreetAddress1,omitempty"` // 收件人详细地址1
	StreetAddress2 string `json:"StreetAddress2,omitempty"` // 收件人详细地址2
	City           string `json:"City"`                     // 收件人所在城市
	State          string `json:"State"`                    // 收货人省/州
	Zip            string `json:"Zip"`                      // 收货人邮编
	Phone          string `json:"Phone"`                    // 收货人电话
	HouseNumber    string `json:"HouseNumber,omitempty"`    // 收件人地址门牌号
}

type SenderResp struct {
	CountryCode string `json:"CountryCode"`       // 发件人国家简码
	FirstName   string `json:"FirstName"`         // 发件人姓名
	LastName    string `json:"LastName"`          // 发件人名字
	Company     string `json:"Company,omitempty"` // 发件人公司名称
	Street      string `json:"Street"`            // 发件人详情地址
	City        string `json:"City"`              // 发件人所在城市
	State       string `json:"State"`             // 发货人省/州
	Zip         string `json:"Zip"`               // 发货人邮编
	Phone       string `json:"Phone"`             // 发货人电话
}

type ParcelResp struct {
	Sku           string  `json:"Sku,omitempty"`           // 商品SKU
	EName         string  `json:"EName"`                   // 申报名称(英文)
	CName         string  `json:"CName,omitempty"`         // 申报名称(中文)
	HSCode        string  `json:"HSCode,omitempty"`        // 申报编码
	Quantity      int     `json:"Quantity"`                // 申报数量
	UnitPrice     float64 `json:"UnitPrice"`               // 申报价格
	UnitWeight    float64 `json:"UnitWeight"`              // 申报重量
	Remark        string  `json:"Remark,omitempty"`        // 材质/用途备注
	ProductUrl    string  `json:"ProductUrl,omitempty"`    // 产品销售链接
	InvoiceRemark string  `json:"InvoiceRemark,omitempty"` // 配货信息
	CurrencyCode  string  `json:"CurrencyCode"`            // 申报币种
	ClassCode     string  `json:"ClassCode,omitempty"`     // 品类编码
	Brand         string  `json:"Brand,omitempty"`         // 品牌
	ModelType     string  `json:"ModelType,omitempty"`     // 型号规格
	Unit          string  `json:"Unit,omitempty"`          // 单位
}

type YunOrderDetailResp struct {
	YunApiResponse
	Item OrderDetail `json:"Item"`
}

type YunDelOrderPost struct {
	OrderNumber string `json:"OrderNumber"` // 单号
	OrderType   string `json:"OrderType"`   // 单号类型
}

type OrderDelete struct {
	Status      string `json:"Status"`      // 返回的结果码Success(成功)/Failure(失败)
	Type        int    `json:"Type"`        // 请求的单号类型
	OrderNumber string `json:"OrderNumber"` // 请求的单号
	Remark      string `json:"Remark"`      // 处理信息
}

type YunDelOrderResp struct {
	YunApiResponse
	Item OrderDelete `json:"Item"`
}

type YunInterceptOrderPost struct {
	OrderNumber string `json:"OrderNumber"` // 单号
	OrderType   int    `json:"OrderType"`   // 单号类型：1-云途单号,2-客户订单号,3-跟踪号
	Remark      string `json:"Remark"`      // 拦截原因
}

type OrderIntercept struct {
	Status      string `json:"Status"`      // 拦截结果码Success(成功)/Failure(失败)
	Type        int    `json:"Type"`        // 拦截的单号类型
	OrderNumber string `json:"OrderNumber"` // 拦截的单号
	Remark      string `json:"Remark"`      // 处理信息
}

type YunInterceptOrderResp struct {
	YunApiResponse
	Item OrderIntercept `json:"Item"`
}

type YunPrintLabelPost struct {
	OrderNumbers string `json:"OrderNumbers"` // 物流系统运单号，客户订单或跟踪号
}

type OrderLabelPrint struct {
	Url        string           `json:"Url"`        // 打印标签地址
	OrderInfos []LabelOrderInfo `json:"OrderInfos"` // 订单详细信息
}

type LabelOrderInfo struct {
	CustomerOrderNumber string `json:"CustomerOrderNumber"` // 客户订单号
	Error               string `json:"Error"`               // 错误信息
	Code                int    `json:"Code"`                // 错误代码(100-正确，200-不存在打印模板)
}

type YunPrintLabelResp struct {
	YunApiResponse
	Item []OrderLabelPrint `json:"Item"`
}

type FeeDetail struct {
	ShippingMethodName  string  `json:"ShippingMethodName"`  // 销售产品
	CountryCode         string  `json:"CountryCode"`         // 国家二字码
	CountryName         string  `json:"CountryName"`         // 国家中文名
	CustomerOrderNumber string  `json:"CustomerOrderNumber"` // 客户订单号
	WayBillNumber       string  `json:"WayBillNumber"`       // 运单号
	TrackingNumber      string  `json:"TrackingNumber"`      // 跟踪号
	GrossWeight         float64 `json:"GrossWeight"`         // 称重重量
	OccurrenceTime      string  `json:"OccurrenceTime"`      // 费用发生时间
	Freight             string  `json:"Freight"`             // 运费
	FuelSurcharge       string  `json:"FuelSurcharge"`       // 燃油费
	RegistrationFee     string  `json:"RegistrationFee"`     // 挂号费
	ProcessingFee       string  `json:"ProcessingFee"`       // 处理费
	OtherFee            string  `json:"OtherFee"`            // 其它杂费
	TotalFee            string  `json:"TotalFee"`            // 总费用
	ChargeWeight        float64 `json:"ChargeWeight"`        // 计费重
	StandardMoney       string  `json:"StandardMoney"`       // 本位币总费用
}

type YunFeeDetailResp struct {
	YunApiResponse
	Item FeeDetail `json:"Item"`
}

type YunUpdateWeightPost struct {
	OrderNumber string `json:"OrderNumber"` // 订单号
	Weight      string `json:"Weight"`      // 重量 2.88
}

type OrderWeight struct {
	OrderNumber string `json:"OrderNumber"` // 订单号
	Status      string `json:"Status"`      // 处理结果(success/failure)
	Remark      string `json:"Remark"`      // 错误信息
}

type YunUpdateWeightResp struct {
	YunApiResponse
	Item OrderWeight `json:"Item"`
}
