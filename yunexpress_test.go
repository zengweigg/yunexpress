package yunexpress

import (
	"fmt"
	"github.com/zengweigg/yunexpress/config"
	"testing"
)

func Test001(m *testing.T) {
	// 初始化
	cfg := config.LoadConfig()
	yunClient := NewYunService(*cfg)
	// 测试1
	resp, err := yunClient.Services.Base.GetCountryList()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("测试1响应", resp.Code, resp.Message)

	// // 测试2
	// fmt.Println("--------------------------------------------------")
	// resp2, err := yunClient.Services.Base.GetShippingList("AD")
	// if err != nil {
	// 	fmt.Printf("Error: %v\n", err)
	// 	return
	// }
	// fmt.Println("测试2响应", resp2.Code, resp2.Message)
}
