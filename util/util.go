package utils

import (
	"reflect"
	"strings"
)

type IStruct interface {
	GetStructData() interface{}
}

// struct转map
// 使用反射实现，完美地兼容了json标签的处理
// https://www.jb51.net/article/273626.htm
func StructToMap(st IStruct) map[string]interface{} {
	m := make(map[string]interface{})
	in := st.GetStructData()
	val := reflect.ValueOf(in)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return m
	}

	relType := val.Type()
	for i := 0; i < relType.NumField(); i++ {
		name := relType.Field(i).Name
		tag := relType.Field(i).Tag.Get("json")
		if tag != "" {
			index := strings.Index(tag, ",")
			if index == -1 {
				name = tag
			} else {
				name = tag[:index]
			}
		}
		m[name] = val.Field(i).Interface()
	}
	return m
}

// 数组去重
// https://www.jianshu.com/p/2d9f1b3b2f3c
func RemoveDuplicateElement(arr []string) []string {
	result := make([]string, 0, len(arr))
	temp := map[string]struct{}{}
	for _, item := range arr {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
