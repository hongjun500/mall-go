/**
 * @author hongjun500
 * @date 2023/6/4
 * @tool ThinkPadX1隐士
 * Created with GoLand 2022.2
 * Description: 使用 golang-security 实现 JWT 认证,并将其封装成 gin 的中间件
 */

package mid

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/conf"
	"github.com/hongjun500/mall-go/internal/gin_common"
	"github.com/hongjun500/mall-go/pkg/security"
)

// GinJWTMiddleware 自定义使用 golang-security 实现 JWT 认证的 gin 中间件
func GinJWTMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		status := http.StatusOK
		header := context.GetHeader(conf.GlobalJwtConfigProperties.TokenHeader)
		if header != "" && strings.HasPrefix(header, conf.GlobalJwtConfigProperties.TokenHead) {
			authToken := header[len(conf.GlobalJwtConfigProperties.TokenHead):]
			username, userId, err := security.GetUsernameAndUserIdFromToken(authToken)
			if err != nil {
				log.Printf("get username from token[%v] is fail %d\n", authToken, err)
				status = gin_common.TokenInvalid
			}
			log.Printf("checking username: %v", username)
			if username == "" {
				status = gin_common.TokenInvalid
			}
			valid := security.TokenValid(authToken, username)
			if !valid {
				status = gin_common.TokenInvalid
			}
			// todo 需要获取用户所拥有的资源并存储
			log.Printf("authenticated user: %v", username)
			context.Set("username", username)
			context.Set("userId", userId)

		} else {
			status = gin_common.Unauthorized
		}
		if status != http.StatusOK {
			gin_common.CreateFail(status, context)
			// 这里必须要加上 Abort()，否则会继续往下执行
			context.Abort()
			return
		}
		// 请求前
		context.Next()
		// 请求后
		// 暂无
	}
}

// GinJWTGetCurrentUsername 通过 gin 的上下文来获取当前登录用户的用户名
func GinJWTGetCurrentUsername(context *gin.Context) string {
	username, _ := context.Get("username")
	if username == nil {
		// 尝试解析 请求头中的 Authorization 信息
		header := context.GetHeader(conf.GlobalJwtConfigProperties.TokenHeader)
		authToken := header[len(conf.GlobalJwtConfigProperties.TokenHead):]
		username, _, err := security.GetUsernameAndUserIdFromToken(authToken)
		if err != nil {
			return ""
		}
		return username
	}
	return username.(string)
}

// GinJWTGetCurrentUserId 通过 gin 的上下文来获取当前登录用户的用户ID
func GinJWTGetCurrentUserId(context *gin.Context) int64 {
	userId, _ := context.Get("userId")
	if userId == nil {
		// 尝试解析 请求头中的 Authorization 信息
		header := context.GetHeader(conf.GlobalJwtConfigProperties.TokenHeader)
		authToken := header[len(conf.GlobalJwtConfigProperties.TokenHead):]
		_, userId, err := security.GetUsernameAndUserIdFromToken(authToken)
		if err != nil {
			return 0
		}
		return userId
	}
	return userId.(int64)
}
