package yunexpress

import (
	"fmt"
	"github.com/zengweigg/yunexpress/config"
	"github.com/zengweigg/yunexpress/model"
	"gopkg.in/guregu/null.v4"
	"testing"
)

func Test001(m *testing.T) {
	// 初始化
	cfg := config.LoadConfig()
	ywClient := NewYunService(*cfg)
	// 构造测试请求数据
	// postData := model.YWBlank{}
	// resp, err := ywClient.Services.XiaoBao.QueryCountryList(postData)
	// if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	// }
	// fmt.Println("666", resp.YWApiResponse, resp.Code, resp.Data, resp.Msg, resp.Success)
	receiverInfo := model.ReceiverInfo{
		Name:        "glassware",
		Phone:       "18231730588",
		Email:       "529932298@qq.com",
		Company:     "yanwen",
		Country:     "115",
		State:       "he bei sheng",
		City:        "cang zhou shi",
		ZipCode:     "10110",
		HouseNumber: "#4501124",
		Address:     "he fang jie cang zhou kai fa qu",
		TaxNumber:   "qwer123",
	}
	senderInfo := model.SenderInfo{
		Name:        "glassware",
		Phone:       "18231730588",
		Email:       "529932298@qq.com",
		Company:     "yanwen",
		Country:     "115",
		State:       "he bei sheng",
		City:        "cang zhou shi",
		ZipCode:     "101110",
		HouseNumber: "#4501124",
		Address:     "he fang jie cang zhou kai fa qu",
		TaxNumber:   "qwer123",
	}
	productList := make([]model.Product, 1)
	productList[0] = model.Product{
		GoodsNameCh: "杯子",
		GoodsNameEn: "CUP",
		Price:       "12.5",
		Hscode:      "8400000001",
		URL:         "http://www.aliexpress.com/item//32681820727.html",
		Material:    "material",
		Quantity:    "10",
		Weight:      "1000",
		Imei:        "IMEI1001",
	}
	parcelInfo := model.ParcelInfo{
		ProductList:   productList,
		HasBattery:    1,
		Currency:      "USD",
		TotalPrice:    "50.01",
		TotalQuantity: "10",
		TotalWeight:   "243",
		Height:        "10",
		Width:         "10",
		Length:        "10",
		Ioss:          "123456",
	}
	csp := model.YWCsp{
		Csp: "19237423SDFGEG",
	}
	importCustomsInfo := model.ImportCustomsInfo{
		TaxPolicyExtends: csp,
	}
	pointID := model.YWPointID{
		PointID: "1001",
	}
	postData := model.YWOrderPost{
		ChannelID:         "481",
		OrderSource:       null.StringFrom("portal"),
		UserID:            "99000015",
		OrderNumber:       "KI1000000001A",
		DateOfReceipt:     "2025-02-17",
		Remark:            "拣货单信息",
		ReceiverInfo:      receiverInfo,
		SenderInfo:        senderInfo,
		ParcelInfo:        parcelInfo,
		PoPStation:        pointID,
		ImportCustomsInfo: importCustomsInfo,
	}
	resp, err := ywClient.Services.XiaoBao.CreateExpressOrder(postData)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("666", resp.Data, resp.Code, resp.Msg, resp.Success)
}
