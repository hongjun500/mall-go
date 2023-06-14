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

	// http 状态码
	Code int `json:"code"`

	// 返回的信息 例如: 操作成功或者操作失败
	Message string `json:"message"`

	// success or fail
	Status string `json:"status"`
	// 返回的数据是任意类型	如果有错误，则把错误信息也封装在此
	/*
		{
			"err_code": 300000,
			"err_msg": "用户名已存在"
		}
	*/
	Data any `json:"data"`
}

// GinCommonError 通用错误信息
type GinCommonError struct {
	ErrCode int    `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
}

// 通用错误码
const (
	UnknownError = 100000
	CustomError  = 100001

	ParameterValidationError = 200000

	DatabaseError = 200001

	UsernameAlreadyExists   = 300000
	UsernameOrPasswordError = 300001
	AccountForbidden        = 300401
	AccountLocked           = 300402
	Unauthorized            = 300403

	TokenGenFail = 300404
	TokenExpired = 300405
	TokenInvalid = 300406
)

var CommonErrorConst = map[int]string{
	UnknownError: "未知错误",
	CustomError:  "自定义错误",

	ParameterValidationError: "参数不合法",
	DatabaseError:            "数据库操作错误",

	UsernameAlreadyExists:   "用户名已存在",
	UsernameOrPasswordError: "用户名或密码错误",
	AccountForbidden:        "当前账号没有相关权限",
	AccountLocked:           "账号被锁定",
	Unauthorized:            "暂未登录或token已经过期",

	TokenGenFail: "token生成失败",
	TokenExpired: "token已过期",
	TokenInvalid: "token不合法",
}

// CreateAny 创建一个通用的返回信息,不取用 Http 状态码,而是自己定义 status 为 success 或 fail
// 2023/6/14 15:09 改动: 增加了 Code 和 Message 字段 适用于 mall 项目
func CreateAny(result any, Code int, Message, status string, context *gin.Context) {
	context.JSON(http.StatusOK, GinCommonResponse{
		Code:    Code,
		Message: Message,
		Status:  status,
		Data:    result,
	})
}

// CreateSuccess 创建一个成功的返回信息
// 2023/6/14 15:09 改动: 增加了 Code 和 Message 字段 适用于 mall 项目
func CreateSuccess(result any, context *gin.Context) {
	CreateAny(result, SUCCESS, IErrorCodeConst[SUCCESS].GetMessage(), "success", context)
}

// CreateSuccessWithMessage 创建一个成功的返回信息,并且自定义返回信息
func CreateSuccessWithMessage(result any, message string, context *gin.Context) {
	CreateAny(result, SUCCESS, message, "success", context)
}

// Create 创建一个成功但没有返回值的返回信息
func Create(context *gin.Context) {
	CreateSuccess(nil, context)
}

// CreateFail 创建一个失败的返回信息,并且将具体的错误信息封装到 Data 中
func CreateFail(result any, context *gin.Context) {
	switch errCodeMsg := result.(type) {
	case int:
		// 失败时将错误信息封装到 Data 中
		err := GinCommonError{ErrCode: errCodeMsg, ErrMsg: CommonErrorConst[errCodeMsg]}
		CreateFiledParam(context, IErrorCodeConst[FAILED], "", err)
	// 接收一个自定义错误信息
	case string:
		customErr := GinCommonError{ErrCode: CustomError, ErrMsg: errCodeMsg}
		CreateFiledParam(context, IErrorCodeConst[FAILED], "", customErr)
	default:
		CreateFiledParam(context, IErrorCodeConst[FAILED], "", GinCommonError{ErrCode: UnknownError, ErrMsg: CommonErrorConst[UnknownError]})
	}
}

func CreateFailed(context *gin.Context) {
	CreateFiledParam(context, IErrorCodeConst[FAILED], IErrorCodeConst[FAILED].GetMessage())
}

// CreateFiledParam 创建一个失败的返回信息,接收 http 状态码，并且将具体的错误信息封装到 Data 中
func CreateFiledParam(context *gin.Context, errs ...any) {
	errorCode := errs[0].(IErrorCode)
	message := errs[1].(string)
	customError := errs[2].(GinCommonError)
	if errorCode != nil && message == "" {
		CreateAny(customError, errorCode.GetCode(), errorCode.GetMessage(), "fail", context)
	} else if message != "" && errorCode == nil {
		CreateAny(customError, FAILED, message, "fail", context)
	} else if errorCode == nil && message == "" {
		CreateAny(customError, FAILED, IErrorCodeConst[FAILED].GetMessage(), "fail", context)
	} else {
		CreateAny(customError, errorCode.GetCode(), message, "fail", context)
	}
}

// CreateValidateFailed 创建一个参数验证失败的返回信息
func CreateValidateFailed(context *gin.Context, message string) {
	CreateFiledParam(context, VALIDATE_FAILED, message, GinCommonError{ErrCode: ParameterValidationError, ErrMsg: CommonErrorConst[ParameterValidationError]})
}

// CreateUnauthorized 创建一个未授权的返回信息
func CreateUnauthorized(context *gin.Context) {
	CreateFiledParam(context, UNAUTHORIZED, IErrorCodeConst[UNAUTHORIZED].GetMessage(), GinCommonError{ErrCode: Unauthorized, ErrMsg: CommonErrorConst[Unauthorized]})
}

// CreateForbidden 创建一个禁止访问的返回信息
func CreateForbidden(context *gin.Context) {
	CreateFiledParam(context, FORBIDDEN, IErrorCodeConst[FORBIDDEN].GetMessage(), GinCommonError{ErrCode: AccountForbidden, ErrMsg: CommonErrorConst[AccountForbidden]})
}
