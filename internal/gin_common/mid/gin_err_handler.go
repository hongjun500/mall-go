// @author hongjun500
// @date 2023/6/29 10:11
// @tool ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description: 全局错误处理中间件

package mid

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/gin_common"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// log.Debugf(assert.CollectT{})
				gin_common.CreateFail(c, gin_common.UnknownError)
				c.Abort()
				return
			}
		}()
		c.Next()
	}
}
