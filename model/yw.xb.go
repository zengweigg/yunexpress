package model

import (
	"gopkg.in/guregu/null.v4"
)

type YWCountryData struct {
	ID     string `json:"id"`     // 国家Id
	Code   string `json:"code"`   // 国家代码
	NameCh string `json:"nameCh"` // 国家中文名称
	NameEn string `json:"nameEn"` // 国家英文名称
}

type YWCountryResp struct {
	YWApiResponse
	Data []struct {
		CountryCode string `json:"countryCode"`
		CountryName string `json:"countryName"`
	} `json:"data"`
}

type YWPointID struct {
	PointID string `json:"pointId"`
}

type YWOrderPost struct {
	ChannelID         string            `json:"channelId"`
	OrderSource       null.String       `json:"orderSource"`
	UserID            string            `json:"userId"`
	OrderNumber       string            `json:"orderNumber"`
	DateOfReceipt     string            `json:"dateOfReceipt"`
	Remark            string            `json:"remark"`
	ReceiverInfo      ReceiverInfo      `json:"receiverInfo"`
	ParcelInfo        ParcelInfo        `json:"parcelInfo"`
	SenderInfo        SenderInfo        `json:"senderInfo"`
	PoPStation        YWPointID         `json:"poPStation"`
	ImportCustomsInfo ImportCustomsInfo `json:"importCustomsInfo"`
}

type ReceiverInfo struct {
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Company     string `json:"company"`
	Country     string `json:"country"`
	State       string `json:"state"`
	City        string `json:"city"`
	ZipCode     string `json:"zipCode"`
	HouseNumber string `json:"houseNumber"`
	Address     string `json:"address"`
	TaxNumber   string `json:"taxNumber"`
}

type ParcelInfo struct {
	ProductList   []Product `json:"productList"`
	HasBattery    int       `json:"hasBattery"`
	Currency      string    `json:"currency"`
	TotalPrice    string    `json:"totalPrice"`
	TotalQuantity string    `json:"totalQuantity"`
	TotalWeight   string    `json:"totalWeight"`
	Height        string    `json:"height"`
	Width         string    `json:"width"`
	Length        string    `json:"length"`
	Ioss          string    `json:"ioss"`
}

type Product struct {
	GoodsNameCh string `json:"goodsNameCh"`
	GoodsNameEn string `json:"goodsNameEn"`
	Price       string `json:"price"`
	Hscode      string `json:"hscode"`
	URL         string `json:"url"`
	Material    string `json:"material"`
	Quantity    string `json:"quantity"`
	Weight      string `json:"weight"`
	Imei        string `json:"imei"`
}

type SenderInfo struct {
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Company     string `json:"company"`
	Country     string `json:"country"`
	State       string `json:"state"`
	City        string `json:"city"`
	ZipCode     string `json:"zipCode"`
	HouseNumber string `json:"houseNumber"`
	Address     string `json:"address"`
	TaxNumber   string `json:"taxNumber"`
}

type YWCsp struct {
	Csp string `json:"csp"`
}

type ImportCustomsInfo struct {
	TaxPolicyExtends YWCsp `json:"taxPolicyExtends"`
}

type YWOrderResp struct {
	YWApiResponse
	Data struct {
		WaybillNumber   string `json:"waybill_number"`
		OrderNumber     string `json:"order_number"`
		ReferenceNumber string `json:"reference_number"`
	} `json:"data"`
}
