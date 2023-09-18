// @author hongjun500
// @date 2023/6/28 13:10
// @tool ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description:

package r_mall_admin

import (
	"github.com/gin-gonic/gin"
	docs "github.com/hongjun500/mall-go/docs/mall_admin"
	"github.com/hongjun500/mall-go/internal/conf"
	"github.com/hongjun500/mall-go/internal/gin_common/mid"
	"github.com/hongjun500/mall-go/internal/services/s_mall_admin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type CoreAdminRouter struct {
	*UmsAdminRouter
	*UmsMenuRouter
	*UmsResourceCategoryRouter
	*UmsResourceRouter
	*UmsRoleRouter
	*UmsMemberLevelRouter
	*PmsProductCategoryRouter
	*PmsProductAttributeCategoryRouter
	*PmsProductAttributeRouter
	*PmsBrandRouter
}

func NewCoreAdminRouter(service *s_mall_admin.CoreAdminService) *CoreAdminRouter {
	return &CoreAdminRouter{
		UmsAdminRouter:                    NewUmsAdminRouter(service.UmsAdminService),
		UmsMenuRouter:                     NewUmsMenuRouter(service.UmsMenuService),
		UmsResourceCategoryRouter:         NewUmsResourceCategoryRouter(service.UmsResourceCategoryService),
		UmsResourceRouter:                 NewUmsResourceRouter(service.UmsResourceService),
		UmsRoleRouter:                     NewUmsRoleRouter(service.UmsRoleService),
		UmsMemberLevelRouter:              NewUmsMemberLevelRouter(service.UmsMemberLevelService),
		PmsProductCategoryRouter:          NewPmsProductCategoryRouter(service.PmsProductCategoryService),
		PmsProductAttributeCategoryRouter: NewPmsProductAttributeCategoryRouter(service.PmsProductAttributeCategoryService),
		PmsProductAttributeRouter:         NewPmsProductAttributeRouter(service.PmsProductAttributeService),
		PmsBrandRouter:                    NewPmsBrandRouter(service.PmsBrandService),
	}
}

// InitAdminGroupRouter 初始化 Admin 路由组
func InitAdminGroupRouter(coreRouter *CoreAdminRouter, ginEngine *gin.Engine) {
	// docs.SwaggerInfo.Version = "1.0"
	// 必须要写上这一行很奇怪
	// 解释：必须要导入 swagger 的包，即 docs, 不然 swagger 无法生成文档

	// 设置 Swagger 路由
	ginEngine.GET("/swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler,
			ginSwagger.InstanceName(conf.GlobalAdminServerConfigProperties.ApplicationName),
			ginSwagger.URL("http://localhost:"+
				conf.GlobalAdminServerConfigProperties.Port+
				"/swagger/"+conf.GlobalAdminServerConfigProperties.ApplicationName+"/doc.json"),
			ginSwagger.PersistAuthorization(true)))
	docs.SwaggerInfomall_admin.Title = conf.GlobalAdminServerConfigProperties.ApplicationName

	// 注册无需认证的路由
	coreRouter.UnauthorizedGroupRouter(ginEngine)
	ginEngine.Use(mid.GinJWTMiddleware()).Use(mid.GinCasbinMiddleware())
	// 注册多个路由组
	coreRouter.GroupUmsAdminRouter(ginEngine.Group("/admin"))
	coreRouter.GroupUmsMenuRouter(ginEngine.Group("/menu"))
	coreRouter.GroupUmsResourceCategoryRouter(ginEngine.Group("/resourceCategory"))
	coreRouter.GroupUmsResourceRouter(ginEngine.Group("/resource"))
	coreRouter.GroupUmsRoleRouter(ginEngine.Group("/role"))
	coreRouter.GroupUmsMemberLevelRouter(ginEngine.Group("/memberLevel"))
	coreRouter.GroupPmsProductCategoryRouter(ginEngine.Group("/productCategory"))
	coreRouter.GroupPmsProductAttributeCategoryRouter(ginEngine.Group("/productAttribute/category"))
	coreRouter.GroupPmsProductAttributeRouter(ginEngine.Group("/productAttribute"))
	coreRouter.GroupPmsBrandRouter(ginEngine.Group("/brand"))
}
