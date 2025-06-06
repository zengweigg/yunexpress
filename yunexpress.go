package yunexpress

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/go-resty/resty/v2"
	"github.com/zengweigg/yunexpress/config"
	"net/http"
	"time"
)

const (
	Version   = "1.0.0"
	userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36"
)

type YunClient struct {
	config     *config.Config
	httpClient *resty.Client
	logger     Logger   // Logger
	Services   services // API Services
}

func NewYunService(cfg config.Config) *YunClient {
	YunClient := &YunClient{
		config: &cfg,
		logger: createLogger(),
	}

	token, err := GetToken(cfg.APIKey, cfg.APISecret)
	if err != nil {
		panic(err)
	}
	// OnBeforeRequest：设置请求发送前的钩子函数，允许在请求发送之前对请求进行修改或添加逻辑。
	// OnAfterResponse：设置响应接收后的钩子函数，允许在接收到响应后处理响应数据或执行其他逻辑。
	// SetRetryCount：设置请求失败时的最大重试次数。
	// SetRetryWaitTime：设置每次重试之间的等待时间（最小等待时间）。
	// SetRetryMaxWaitTime：设置每次重试之间的最大等待时间，实际等待时间会在最小和最大等待时间之间随机选取。
	// AddRetryCondition：添加自定义的重试条件，当满足该条件时触发重试机制。
	httpClient := resty.
		New().
		SetDebug(YunClient.config.Debug).
		SetHeaders(map[string]string{
			"Content-Type":  "application/json",
			"Accept":        "application/json",
			"User-Agent":    userAgent,
			"Authorization": token,
		})
	if cfg.Sandbox {
		// 测试
		httpClient.SetBaseURL("http://omsapi.uat.yunexpress.com/api/")
	} else {
		// 正式
		httpClient.SetBaseURL("http://oms.api.yunexpress.com/api/")
	}
	httpClient.
		SetTimeout(time.Duration(cfg.Timeout) * time.Second).
		SetRetryCount(2).
		SetRetryWaitTime(5 * time.Second).
		SetRetryMaxWaitTime(10 * time.Second).
		AddRetryCondition(func(r *resty.Response, err error) bool {
			text := r.Request.URL
			if r == nil {
				return false
			}
			if err != nil {
				text += fmt.Sprintf(", error: %s", err.Error())
				YunClient.logger.Debugf("Retry request: %s", text)
				return true // 如果有错误则重试
			}
			// 检查响应状态码是否不是200
			if r.StatusCode() != http.StatusOK {
				text += fmt.Sprintf(", error: %d", r.StatusCode())
				YunClient.logger.Debugf("Retry request: %s", text)
				return true
			}
			type ResponseBody struct {
				Code string `json:"Code"`
			}
			// 解析响应体JSON
			var responseBody ResponseBody
			if err := sonic.Unmarshal(r.Body(), &responseBody); err != nil {
				text += fmt.Sprintf(", error: %s", string(r.Body()))
				YunClient.logger.Debugf("Retry request: %s", text)
				return true // 如果解析错误则重试
			}
			// 检查apiResultCode字段是否不是0
			if responseBody.Code != "0000" {
				text += fmt.Sprintf(", error: %s", responseBody.Code)
				YunClient.logger.Debugf("Retry request: %s", text)
				return true
			}
			return false
		})
	YunClient.httpClient = httpClient
	xService := service{
		config:     &cfg,
		logger:     YunClient.logger,
		httpClient: YunClient.httpClient,
	}
	YunClient.Services = services{
		Base:  (baseService)(xService),  // 基础通用
		Order: (orderService)(xService), // 订单
		Track: (trackService)(xService), // 追踪
	}
	return YunClient
}
