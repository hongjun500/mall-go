// @author hongjun500
// @date 2023/6/9 13:20
// @tool ThinkPadX1隐士
// Created with GoLand 2022.2
// Description: 针对 encoding/json 包的一些封装

package convert

import (
	"encoding/json"
	"log"
)

// AnyToBytes 转换任意类型为字节切片
func AnyToBytes(data any) []byte {
	bytes, err := json.Marshal(data)
	if err != nil {
		log.Printf("any to bytes error: %v", err.Error())
		return nil
	}
	return bytes
}

// AnyToJson 转换任意类型为 json 字符串
func AnyToJson(data any) string {
	bytes := AnyToBytes(data)
	if bytes == nil {
		log.Printf("any to json error: %v", "bytes is nil")
		return ""
	}
	return string(bytes)
}

// BytesToAny 将字节切片转换成任意类型, data 为指针类型
func BytesToAny(bytes []byte, data any) error {
	err := json.Unmarshal(bytes, data)
	if err != nil {
		return err
	}
	return nil
}

// JsonToAny 将 json 字符串转换成任意类型, data 为指针类型
func JsonToAny(jsonStr string, data any) error {
	bytes := []byte(jsonStr)
	err := BytesToAny(bytes, data)
	if err != nil {
		return err
	}
	return nil
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
