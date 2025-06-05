package yunexpress

import (
	"errors"
	"github.com/bytedance/sonic"
	"github.com/zengweigg/yunexpress/model"
)

type trackService service

// 2.5 查询跟踪号
func (s trackService) GetTrackingNumber(req string) (model.YunTrackingNumberResp, error) {
	respData := model.YunTrackingNumberResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetQueryParam("CustomerOrderNumber", req).
		Get("Waybill/GetTrackingNumber")
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

// 2.15-1 查询物流轨迹信息
func (s trackService) GetTrackInfo(req string) (model.YunTrackInfoResp, error) {
	respData := model.YunTrackInfoResp{}
	if req == "" {
		return respData, errors.New("请输入物流系统运单号，客户订单或跟踪号")
	}
	// 请求数据
	resp, err := s.httpClient.R().
		SetQueryParam("OrderNumber", req).
		Get("Tracking/GetTrackInfo")
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

// 2.15-2 查询物流轨迹信息 查询全程
func (s trackService) GetTrackAllInfo(req string) (model.YunTrackInfoResp, error) {
	respData := model.YunTrackInfoResp{}
	if req == "" {
		return respData, errors.New("请输入物流系统运单号，客户订单或跟踪号")
	}
	// 请求数据
	resp, err := s.httpClient.R().
		SetQueryParam("OrderNumber", req).
		Get("Tracking/GetTrackAllInfo")
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

// 2.18 按单号订阅轨迹
func (s trackService) CreatedOrderSubscribe(reqData model.YunCreatedOrderSubscribePost) (model.YunCreatedOrderSubscribeResp, error) {
	respData := model.YunCreatedOrderSubscribeResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(reqData).
		Post("Tracking/GetTrackAllInfo")
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

// 2.19 按单号取消轨迹订阅
func (s trackService) CancelOrderSubscribe(reqData model.YunCreatedOrderSubscribePost) (model.YunCreatedOrderSubscribeResp, error) {
	respData := model.YunCreatedOrderSubscribeResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(reqData).
		Post("Tracking/CancelOrderSubscribe")
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

// 2.20 获取按订单号订阅轨迹的数据
func (s trackService) GetOrderSubscribe(reqData model.YunOrderSubscribePost) (model.YunOrderSubscribeResp, error) {
	respData := model.YunOrderSubscribeResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(reqData).
		Post("tracking/GetOrderSubscribe")
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

// 2.21 按运输方式订阅轨迹
func (s trackService) CreatedProductSubscribe(reqData model.YunCreatedProductSubscribePost) (model.YunCreatedProductSubscribeResp, error) {
	respData := model.YunCreatedProductSubscribeResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(reqData).
		Post("tracking/GetOrderSubscribe")
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

// 2.22 取消按运输方式订阅
func (s trackService) CancelProductSubscribe(reqData model.YunCancelProductSubscribePost) (model.YunCancelProductSubscribeResp, error) {
	respData := model.YunCancelProductSubscribeResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(reqData).
		Post("tracking/CancelProductSubscribe")
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

// 2.23 获取按运输方式订阅的数据 tracking/GetProductSubscribe
func (s trackService) GetProductSubscribe(reqData model.YunGetProductSubscribePost) (model.YunGetProductSubscribeResp, error) {
	respData := model.YunGetProductSubscribeResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(reqData).
		Post("tracking/GetProductSubscribe")
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
