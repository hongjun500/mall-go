// @author hongjun500
// @date 2023/6/14 15:09
// @tool ThinkPadX1隐士
// Created with GoLand 2022.2
// Description: 封装适用于 mall 项目的 API 返回码

package gin_common

// ResultCode API 返回码
type ResultCode struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type IErrorCode interface {
	GetCode() int
	GetMessage() string
	SetMessage(string)
}

const (
	// SUCCESS 操作成功
	SUCCESS = 200
	// FAILED 操作失败
	FAILED = 500
	// VALIDATE_FAILED 参数检验失败
	VALIDATE_FAILED = 404
	// UNAUTHORIZED 未登录或token过期
	UNAUTHORIZED = 401
	// FORBIDDEN 没有相关权限
	FORBIDDEN = 403
)

var IErrorCodeConst = map[int]IErrorCode{
	SUCCESS:         &ResultCode{Code: SUCCESS, Message: "操作成功"},
	FAILED:          &ResultCode{Code: FAILED, Message: "操作失败"},
	VALIDATE_FAILED: &ResultCode{Code: VALIDATE_FAILED, Message: "参数检验失败"},
	UNAUTHORIZED:    &ResultCode{Code: UNAUTHORIZED, Message: "未登录或token过期"},
	FORBIDDEN:       &ResultCode{Code: FORBIDDEN, Message: "没有相关权限"},
}

func (r *ResultCode) GetCode() int {
	return r.Code
}

func (r *ResultCode) GetMessage() string {
	return r.Message
}

func (r *ResultCode) SetMessage(s string) {
	r.Message = s
}
