/**
对 gin 做一些封装
*/

package gin_common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type GinEngine struct {
	GinEngine *gin.Engine
}

// GinCommonResponse 通用返回信息
type GinCommonResponse struct {

	// success or fail
	Status string `json:"status"`
	// 返回的数据 任意类型
	// 如果有错误，则把错误信息也封装在此
	/*
		{
			"errCode": "fail",
			"errMsg": "用户名已存在"
		}
	*/
	Data any `json:"data"`
}

// GinCommonError 通用错误信息
type GinCommonError struct {
	ErrCode int    `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
}

var CommonErrorConst = map[int]string{
	UnknownError:             "未知错误",
	ParameterValidationError: "参数不合法",
	UsernameAlreadyExists:    "用户名已存在",
}

// 通用错误码
const (
	UnknownError             = 100000
	ParameterValidationError = 200000
	UsernameAlreadyExists    = 300000
	PasswordError            = 300000 + iota
	// ParameterMissingError
	// ParameterTypeError
)

// CreateAny 创建一个通用的返回信息,不取用 Http 状态码,而是自己定义 status 为 success 或 fail
func CreateAny(result any, status string, context *gin.Context) {
	context.JSON(http.StatusOK, GinCommonResponse{
		Status: status,
		Data:   result,
	})
}

// CreateSuccess 创建一个成功的返回信息
func CreateSuccess(result any, context *gin.Context) {
	CreateAny(result, "success", context)
}

// Create 创建一个成功的没有返回值的返回信息
func Create(context *gin.Context) {
	CreateSuccess(nil, context)
}

// CreateFail 创建一个失败的返回信息
func CreateFail(result any, context *gin.Context) {
	switch errCode := result.(type) {
	case int:
		// 失败时将错误信息封装到 Data 中
		commonError := GinCommonError{ErrCode: errCode, ErrMsg: CommonErrorConst[errCode]}
		CreateAny(commonError, "fail", context)
	}
}
