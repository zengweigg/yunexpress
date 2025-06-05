package yunexpress

import (
	"errors"
	"github.com/bytedance/sonic"
	"github.com/zengweigg/yunexpress/model"
)

type orderService service

// 2.7 运单申请 WayBill/CreateOrder
func (s orderService) CreateOrder(reqData model.YunOrderPost) (model.YunOrderResp, error) {
	respData := model.YunOrderResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(reqData).
		Post("WayBill/CreateOrder")
	if err != nil {
		return respData, err
	}
	// 解析数据
	err = sonic.Unmarshal(resp.Body(), &respData)
	if err != nil {
		return respData, nil
	}
	return respData, nil
}

// 2.8 查询运单
func (s orderService) GetOrder(req string) (model.YunOrderDetailResp, error) {
	respData := model.YunOrderDetailResp{}
	if req == "" {
		return respData, errors.New("请输入物流系统运单号，客户订单或跟踪号")
	}
	// 请求数据
	resp, err := s.httpClient.R().
		SetQueryParam("OrderNumber", req).
		Get("WayBill/GetOrder")
	if err != nil {
		return respData, err
	}
	// 解析数据
	err = sonic.Unmarshal(resp.Body(), &respData)
	if err != nil {
		return respData, nil
	}
	return respData, nil
}

// 2.9 修改订单预报重量
func (s orderService) UpdateOrder(reqData model.YunUpdateWeightPost) (model.YunUpdateWeightResp, error) {
	respData := model.YunUpdateWeightResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(reqData).
		Post("WayBill/UpdateWeight")
	if err != nil {
		return respData, err
	}
	// 解析数据
	err = sonic.Unmarshal(resp.Body(), &respData)
	if err != nil {
		return respData, nil
	}
	return respData, nil
}

// 2.10 订单删除
func (s orderService) DeleteOrder(reqData model.YunDelOrderPost) (model.YunDelOrderResp, error) {
	respData := model.YunDelOrderResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(reqData).
		Post("WayBill/Delete")
	if err != nil {
		return respData, err
	}
	// 解析数据
	err = sonic.Unmarshal(resp.Body(), &respData)
	if err != nil {
		return respData, nil
	}
	return respData, nil
}

// 2.11 订单拦截
func (s orderService) InterceptOrder(reqData model.YunInterceptOrderPost) (model.YunInterceptOrderResp, error) {
	respData := model.YunInterceptOrderResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(reqData).
		Post("WayBill/Intercept")
	if err != nil {
		return respData, err
	}
	// 解析数据
	err = sonic.Unmarshal(resp.Body(), &respData)
	if err != nil {
		return respData, nil
	}
	return respData, nil
}

// 2.12 标签打印
func (s orderService) PrintLabel(reqData model.YunPrintLabelPost) (model.YunPrintLabelResp, error) {
	respData := model.YunPrintLabelResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(reqData).
		Post("Label/Print")
	if err != nil {
		return respData, err
	}
	// 解析数据
	err = sonic.Unmarshal(resp.Body(), &respData)
	if err != nil {
		return respData, nil
	}
	return respData, nil
}

// 2.13 查询物流运费明细
func (s orderService) GetShippingFeeDetail(req string) (model.YunFeeDetailResp, error) {
	respData := model.YunFeeDetailResp{}
	if req == "" {
		return respData, errors.New("请输入运单号")
	}
	// 请求数据
	resp, err := s.httpClient.R().
		SetQueryParam("WayBillNumber", req).
		Get("Freight/GetShippingFeeDetail")
	if err != nil {
		return respData, err
	}
	// 解析数据
	err = sonic.Unmarshal(resp.Body(), &respData)
	if err != nil {
		return respData, nil
	}
	return respData, nil
}
