package yunexpress

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/zengweigg/yunexpress/config"
	"reflect"
)

type service struct {
	config     *config.Config // Config
	logger     Logger         // Logger
	httpClient *resty.Client  // HTTP client
}

type services struct {
	Base  baseService
	Order orderService
	Track trackService
}

// StructToMap 结构体转换为 map[string]string{}
func StructToMap(obj interface{}) (map[string]string, error) {
	result := make(map[string]string)
	// 获取对象的反射值
	val := reflect.ValueOf(obj)
	// 检查是否是指针类型，如果是则获取其指向的值
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	// 检查是否是结构体类型
	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("input is not a struct")
	}
	// 获取结构体的类型
	typ := val.Type()
	// 遍历结构体的字段
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fieldValue := val.Field(i)
		// 获取字段的名称
		fieldName := field.Name
		// 获取字段的json标签
		jsonTag := field.Tag.Get("json")
		if jsonTag != "" && jsonTag != "-" {
			fieldName = jsonTag
		}
		// 获取字段的值
		// result[fieldName] = fieldValue.Interface()
		if fieldValue.Interface() == nil {
			result[fieldName] = ""
		} else {
			// 将字段值转为字符串
			value := fieldValue.Interface()
			result[fieldName] = fmt.Sprintf("%v", value)
		}
	}
	return result, nil
}
