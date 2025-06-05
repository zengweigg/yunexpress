package yunexpress

import (
	"encoding/base64"
	"fmt"
)

// GetToken 获取token
func GetToken(apikey, apisecret string) (token string, err error) {
	// 获取用户认证 Token
	// 客户编号：客户新注册时由业务部门提供的客户身份唯一标识编号
	// ApiSecret：申请 API 接口服务时由业务部门提供的一个密钥
	// Token 值：由客户编号和 ApiSecret 构成中间用”&”分隔，并要 Base64 编码。
	// 测试客户编号为：ITC0893791
	// 测试 ApiSecret 为：axzc2utvPbfc9UbJDOh+7w== 由 ‘ITC0893791&axzc2utvPbfc9UbJDOh+7w==’ 字符(不含单引号)通
	// Base64 编码后产生的字符则为该用户的 Token 值
	// 对字符串进行 Base64 编码
	str := fmt.Sprintf("%s&%s", apikey, apisecret)
	if apikey == "" || apisecret == "" {
		return token, fmt.Errorf("apikey or apisecret is empty")
	}
	encodedString := base64.StdEncoding.EncodeToString([]byte(str))
	token = "Basic " + encodedString
	return token, nil
}
