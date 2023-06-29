// @author hongjun500
// @date 2023/6/28 10:54
// @tool ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description:

package r_mall_search

import (
	"github.com/gin-gonic/gin"
	docs "github.com/hongjun500/mall-go/docs/mall_search"
	"github.com/hongjun500/mall-go/internal/conf"
	"github.com/hongjun500/mall-go/internal/services/s_mall_search"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type CoreSearchRouter struct {
	*ProductSearchRouter
}

func NewCoreSearchRouter(service *s_mall_search.CoreSearchService) *CoreSearchRouter {
	return &CoreSearchRouter{
		ProductSearchRouter: NewSearchRouter(service.ProductSearchService),
	}
}

// InitSearchGroupRouter 搜索服务路由组
func InitSearchGroupRouter(coreSearchRouter *CoreSearchRouter, ginEngine *gin.Engine) {
	// 设置 Swagger 路由
	ginEngine.GET("/swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler,
			ginSwagger.InstanceName(conf.GlobalSearchServerConfigProperties.ApplicationName),
			ginSwagger.URL("http://localhost:"+
				conf.GlobalSearchServerConfigProperties.Port+
				"/swagger/"+conf.GlobalSearchServerConfigProperties.ApplicationName+"/doc.json"),
			ginSwagger.PersistAuthorization(true)))
	docs.SwaggerInfomall_search.Title = conf.GlobalSearchServerConfigProperties.ApplicationName
	// ginEngine.Use(mid.GinJWTMiddleware()).Use(mid.GinCasbinMiddleware())

	coreSearchRouter.GroupProductRouter(ginEngine.Group("/product"))
}
