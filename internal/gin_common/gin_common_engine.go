/**
对 gin 做一些封装
*/

package gin_common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
	Unauthorized            = 300401
	AccountLocked           = 300402
	AccountForbidden        = 300403

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
	Unauthorized:            "暂未登录或token已经过期",
	AccountLocked:           "账号被锁定",
	AccountForbidden:        "当前账号没有相关权限",

	TokenGenFail: "token生成失败",
	TokenExpired: "token已过期",
	TokenInvalid: "token不合法",
}

// CreateAny 创建一个通用的返回信息,不取用 Http 状态码,而是自己定义 status 为 success 或 fail
// 2023/6/14 15:09 改动: 增加了 Code 和 Message 字段 适用于 mall 项目
func CreateAny(context *gin.Context, Code int, Message, status string, result any) {
	context.JSON(http.StatusOK, GinCommonResponse{
		Code:    Code,
		Message: Message,
		Status:  status,
		Data:    result,
	})
}

// CreateSuccess 创建一个成功的返回信息
// 2023/6/14 15:09 改动: 增加了 Code 和 Message 字段 适用于 mall 项目
func CreateSuccess(context *gin.Context, result any) {
	CreateAny(context, SUCCESS, IErrorCodeConst[SUCCESS].GetMessage(), "success", result)
}

// CreateSuccessWithMessage 创建一个成功的返回信息,并且自定义返回信息
func CreateSuccessWithMessage(context *gin.Context, message string, result any) {
	CreateAny(context, SUCCESS, message, "success", result)
}

// Create 创建一个成功但没有返回值的返回信息
func Create(context *gin.Context) {
	CreateSuccess(context, nil)
}

// CreateFiledParam 创建一个失败的返回信息,接收 http 状态码，并且将具体的错误信息封装到 Data 中
func CreateFiledParam(context *gin.Context, errs ...any) {
	resultCode := errs[0].(*ResultCode)
	message := errs[1].(string)
	customError := errs[2].(GinCommonError)
	if resultCode != nil && message == "" {
		CreateAny(context, resultCode.GetCode(), resultCode.GetMessage(), "fail", customError)
	} else if message != "" && resultCode == nil {
		CreateAny(context, FAILED, message, "fail", customError)
	} else if resultCode == nil && message == "" {
		CreateAny(context, FAILED, IErrorCodeConst[FAILED].GetMessage(), "fail", customError)
	} else {
		CreateAny(context, resultCode.GetCode(), message, "fail", customError)
	}
}

// CreateFail 创建一个失败的返回信息,并且将具体的错误信息封装到 Data 中
func CreateFail(context *gin.Context, result any) {
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

// CreateValidateFailed 创建一个参数验证失败的返回信息
func CreateValidateFailed(context *gin.Context, message string) {
	CreateFiledParam(context, IErrorCodeConst[VALIDATE_FAILED], message, GinCommonError{ErrCode: ParameterValidationError, ErrMsg: CommonErrorConst[ParameterValidationError]})
}

// CreateUnauthorized 创建一个未授权的返回信息
func CreateUnauthorized(context *gin.Context) {
	CreateFiledParam(context, IErrorCodeConst[UNAUTHORIZED], IErrorCodeConst[UNAUTHORIZED].GetMessage(), GinCommonError{ErrCode: Unauthorized, ErrMsg: CommonErrorConst[Unauthorized]})
}

// CreateForbidden 创建一个禁止访问的返回信息
func CreateForbidden(context *gin.Context) {
	CreateFiledParam(context, IErrorCodeConst[FORBIDDEN], IErrorCodeConst[FORBIDDEN].GetMessage(), GinCommonError{ErrCode: AccountForbidden, ErrMsg: CommonErrorConst[AccountForbidden]})
}
