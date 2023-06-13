/**
 * @author hongjun500
 * @date 2023/6/4
 * @tool ThinkPadX1隐士
 * Created with GoLand 2022.2
 * Description: 使用 golang-jwt 实现 JWT 认证,并将其封装成 gin 的中间件
 */

package mid

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/conf"
	"github.com/hongjun500/mall-go/internal/gin_common"
	"github.com/hongjun500/mall-go/pkg/jwt"
	"log"
	"net/http"
	"strings"
)

// GinJWTMiddleware 自定义使用 golang-jwt 实现 JWT 认证的 gin 中间件
func GinJWTMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		status := http.StatusOK
		header := context.GetHeader(conf.GlobalJwtConfigProperties.TokenHeader)
		if header != "" && strings.HasPrefix(header, conf.GlobalJwtConfigProperties.TokenHead) {
			authToken := header[len(conf.GlobalJwtConfigProperties.TokenHead):]
			username, userId, err := jwt.GetUsernameAndUserIdFromToken(authToken)
			if err != nil {
				log.Printf("get username from token[%v] is fail %d\n", authToken, err)
				status = gin_common.TokenInvalid
			}
			log.Printf("checking username: %v", username)
			if username == "" {
				status = gin_common.TokenInvalid
			}
			valid := jwt.TokenValid(authToken, username)
			if !valid {
				status = gin_common.TokenInvalid
			}
			log.Printf("authenticated user: %v", username)
			context.Set("username", username)
			context.Set("user_id", userId)

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
	return username.(string)
}

// GinJWTGetCurrentUserId 通过 gin 的上下文来获取当前登录用户的用户ID
func GinJWTGetCurrentUserId(context *gin.Context) int64 {
	userId, _ := context.Get("user_id")
	return userId.(int64)
}
