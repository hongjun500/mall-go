// @author hongjun500
// @date 2023/6/9 13:20
// @tool ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description: 针对 encoding/json 包的一些封装

package convert

import "encoding/json"

// StructToJson 将对象转换成 json 字符串
func StructToJson(obj any) string {
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}

// MapToJson 将 map 转换成 json 字符串
func MapToJson(obj map[string]any) string {
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}

// SliceToJson 将切片转换成 json 字符串
func SliceToJson(obj []any) string {
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}

// SliceMapToJson 将切片 map 转换成 json 字符串
func SliceMapToJson(obj []map[string]any) string {
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}

// SliceStructToJson 将切片结构体转换成 json 字符串
func SliceStructToJson(obj []any) string {
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}

// JsonToStruct 将 json 字符串转换成结构体 obj
func JsonToStruct(jsonStr string, obj any) error {
	err := json.Unmarshal([]byte(jsonStr), &obj)
	if err != nil {
		return err
	}
	return nil
}

// JsonToMap 将 json 字符串转换成 map
func JsonToMap(jsonStr string) (map[string]any, error) {
	var result map[string]any
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// JsonToSlice 将 json 字符串转换成切片
func JsonToSlice(jsonStr string) ([]any, error) {
	var result []any
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// JsonToSliceMap 将 json 字符串转换成切片 map
func JsonToSliceMap(jsonStr string) ([]map[string]any, error) {
	var result []map[string]any
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// JsonToSliceStruct 将 json 字符串转换成切片结构体
func JsonToSliceStruct(jsonStr string, obj any) (any, error) {
	err := json.Unmarshal([]byte(jsonStr), &obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// JsonToInterface 将 json 字符串转换成 interface
func JsonToInterface(jsonStr string) (any, error) {
	var result any
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// JsonFormat 格式化 json 字符串
func JsonFormat(jsonStr string) string {
	var result any
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		return ""
	}
	jsonBytes, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}
