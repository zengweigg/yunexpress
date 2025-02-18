package yunexpress

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/go-resty/resty/v2"
	"github.com/zengweigg/yunexpress/config"
)

type service struct {
	config     *config.Config // Config
	logger     Logger         // Logger
	httpClient *resty.Client  // HTTP client
}

type services struct {
	XiaoBao xiaoBaoService
	//Gts  gtsService
	//Icms icmsService
	//Icsm icsmService
}

// GetSign MD5加密
func GetSign(text string) string {
	// 创建一个 md5 哈希对象
	hasher := md5.New()
	// 将输入字符串写入哈希对象
	hasher.Write([]byte(text))
	// 计算哈希值
	hashBytes := hasher.Sum(nil)
	// 将哈希值转换为十六进制字符串
	hashStr := hex.EncodeToString(hashBytes)
	return hashStr
}
