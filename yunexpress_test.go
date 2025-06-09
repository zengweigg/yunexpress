package yunexpress

import (
	"fmt"
	"github.com/zengweigg/yunexpress/config"
	"github.com/zengweigg/yunexpress/model"
	"testing"
)

func Test001(m *testing.T) {
	// 初始化 111
	cfg := config.LoadConfig()
	yunClient := NewYunService(*cfg)
	// 测试1
	// resp, err := yunClient.Services.Base.GetCountryList()
	// if err != nil {
	// 	fmt.Printf("Error: %v\n", err)
	// 	return
	// }
	// fmt.Println("测试1响应", resp.Code, resp.Message)
	receiver := model.Receiver{
		CountryCode:  "DE",
		FirstName:    "Andreas Loertzer",
		Street:       "Forstwald Straße 674",
		City:         "Krefeld",
		State:        "Nordrhein-Westfalen",
		Zip:          "47804",
		Phone:        "015141965790",
		Email:        "qtnh1thg3gh5p05@marketplace.amazon.de",
		MobileNumber: "015141965790",
	}

	respdata := model.YunOrderPost{
		CustomerOrderNumber: "test_order_1",
		ShippingMethodCode:  "A01PH",
		PackageCount:        1,
		Weight:              2,
		Receiver:            receiver,
	}
	resp2, err := yunClient.Services.Order.CreateOrder(respdata)
	// // 测试2
	// fmt.Println("--------------------------------------------------")
	// resp2, err := yunClient.Services.Base.GetShippingList("AD")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("测试2响应", resp2.Code, resp2.Message)
}
