package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/gin_common/mid"
	"github.com/hongjun500/mall-go/internal/services"
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
}

type CoreSearchRouter struct {
	*SearchRouter
}

type CoreRouterInterface interface {
	InitCoreRouter(service *services.CoreAdminService, coreRouter *CoreAdminRouter)
}

func NewCoreAdminRouter(service *services.CoreAdminService) *CoreAdminRouter {
	return &CoreAdminRouter{
		UmsAdminRouter:            NewUmsAdminRouter(service.UmsAdminService),
		UmsMenuRouter:             NewUmsMenuRouter(service.UmsMenuService),
		UmsResourceCategoryRouter: NewUmsResourceCategoryRouter(service.UmsResourceCategoryService),
		UmsResourceRouter:         NewUmsResourceRouter(service.UmsResourceService),
		UmsRoleRouter:             NewUmsRoleRouter(service.UmsRoleService),
		UmsMemberLevelRouter:      NewUmsMemberLevelRouter(service.UmsMemberLevelService),
	}
}

func NewCoreSearchRouter(service *services.CoreSearchService) *CoreSearchRouter {
	return &CoreSearchRouter{
		SearchRouter: NewSearchRouter(service.ProductSearchService),
	}
}

// InitAdminGroupRouter 初始化 Admin 路由组
func InitAdminGroupRouter(coreRouter *CoreAdminRouter, ginEngine *gin.Engine) {
	// docs.SwaggerInfo.Version = "1.0"
	// 必须要写上这一行很奇怪
	// 解释：必须要导入 swagger 的包，即 docs, 不然 swagger 无法生成文档

	// 设置 Swagger 路由
	ginEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
}

func InitSearchGroupRouter(coreSearchRouter *CoreSearchRouter, ginEngine *gin.Engine) {

	// 设置 Swagger 路由
	ginEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	coreSearchRouter.GroupProductRouter(ginEngine.Group("/product"))
}

// UnauthorizedGroupRouter  未授权路由
func (router *CoreAdminRouter) UnauthorizedGroupRouter(routerEngine *gin.Engine) {
	unAuthGroup := routerEngine.Group("/admin")
	{
		// 用户注册
		unAuthGroup.POST("/register", router.UmsAdminService.UmsAdminRegister)
		/*adminGroup.POST("/register", router.UmsAdminRegister) 这样写也可以*/

		// 用户登录
		unAuthGroup.POST("/login", router.UmsAdminService.UmsAdminLogin)
		// 用户登出
		unAuthGroup.POST("/logout", router.UmsAdminService.UmsAdminLogout)
		unAuthGroup.GET("/authTest", router.UmsAdminService.UmsAdminAuthTest)
	}
}
