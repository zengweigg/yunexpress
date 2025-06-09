package yunexpress

import (
	"github.com/bytedance/sonic"
	"github.com/zengweigg/yunexpress/model"
)

type baseService service

// 2.1 查询国家简码
func (s baseService) GetCountryList() (model.YunCountryResp, error) {
	respData := model.YunCountryResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		Get("Common/GetCountry")
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

// 2.2 查询运输方式
func (s baseService) GetShippingList(countryCode string) (model.YunShippingResp, error) {
	respData := model.YunShippingResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetQueryParam("CountryCode", countryCode).
		Get("Common/GetShippingMethods")
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

// 2.3 查询货品类型
func (s baseService) GetProductTypeList() (model.YunProductTypeResp, error) {
	respData := model.YunProductTypeResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		Get("Common/GetGoodsType")
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

// 2.4 查询价格
func (s baseService) GetPriceTrial(req model.YunPriceReq) (model.YunPriceResp, error) {
	respData := model.YunPriceResp{}
	// 请求数据
	params, err := StructToMap(req)
	if err != nil {
		return respData, err
	}
	resp, err := s.httpClient.R().
		SetQueryParams(params).
		Get("Freight/GetPriceTrial")
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

// 2.6 查询发件人信息
func (s baseService) GetSender(req string) (model.YunSenderResp, error) {
	respData := model.YunSenderResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetQueryParam("OrderNumber", req).
		Get("WayBill/GetSender")
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

// 2.14 用户注册
func (s baseService) Register(reqData model.YunRegisterReq) (model.YunRegisterResp, error) {
	respData := model.YunRegisterResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(reqData).
		Post("Common/Register")
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

// 2.16 查询末端派送商
func (s baseService) GetCarrier(reqData model.YunCarrierPost) (model.YunCarrierResp, error) {
	respData := model.YunCarrierResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(reqData).
		Post("Waybill/GetCarrier")
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

// 2.17 IOSS号备案
func (s baseService) RegisterIoss(reqData model.YunRegisterIossPost) (model.YunRegisterIossResp, error) {
	respData := model.YunRegisterIossResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(reqData).
		Post("WayBill/RegisterIoss")
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

// 2.24 IOSS号作废
func (s baseService) DisableIoss(reqData model.YunDisableIossPost) (model.YunDisableIossResp, error) {
	respData := model.YunDisableIossResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(reqData).
		Post("WayBill/DisableIoss")
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
