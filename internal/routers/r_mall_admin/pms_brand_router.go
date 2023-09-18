// @author hongjun500
// @date 2023/7/26 14:11
// @tool ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description:

package r_mall_admin

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/gin_common"
	"github.com/hongjun500/mall-go/internal/request/admin_dto"
	"github.com/hongjun500/mall-go/internal/request/base_dto"
	"github.com/hongjun500/mall-go/internal/services/s_mall_admin"
)

type PmsBrandRouter struct {
	s_mall_admin.PmsBrandService
}

func NewPmsBrandRouter(service s_mall_admin.PmsBrandService) *PmsBrandRouter {
	return &PmsBrandRouter{PmsBrandService: service}
}

func (router *PmsBrandRouter) GroupPmsBrandRouter(routerGroup *gin.RouterGroup) {
	{
		// 获取全部品牌列表
		routerGroup.GET("/listAll", router.listAll)
		// 创建品牌
		routerGroup.POST("/create", router.create)
		// 更新品牌
		routerGroup.POST("/update/:id", router.update)
		// 删除品牌
		routerGroup.GET("/delete/:id", router.delete)
		// 分页获取品牌列表
		routerGroup.GET("/list", router.list)
		// 获取品牌详情
		routerGroup.GET("/:id", router.detail)
		// 批量删除品牌
		routerGroup.POST("/delete/batch", router.deleteBatch)
		// 批量更新显示状态
		routerGroup.POST("/update/showStatus", router.updateShowStatus)
		// 批量更新厂家制造商状态
		routerGroup.POST("/update/factoryStatus", router.updateFactoryStatus)
	}
}

// listAll 获取全部品牌列表
//
//	@Summary		获取全部品牌列表
//	@Description	获取全部品牌列表
//	@Tags			商品品牌管理
//	@Accept			json
//	@Produce		json
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router /brand/listAll [get]
func (router *PmsBrandRouter) listAll(context *gin.Context) {
	listAll, err := router.PmsBrandService.ListAll()
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, listAll)
}

// create 创建品牌
//
//	@Summary		创建品牌
//	@Description	创建品牌
//	@Tags			商品品牌管理
//	@Accept			json
//	@Produce		json
//	@Security		GinJWTMiddleware
//	@Param			pmsBrand	body		admin_dto.PmsBrandDTO	true		"品牌信息"
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router /brand/create [post]
func (router *PmsBrandRouter) create(context *gin.Context) {
	var pmsBrandCreateDto admin_dto.PmsBrandDTO
	if err := context.ShouldBindJSON(&pmsBrandCreateDto); err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	rows, err := router.PmsBrandService.Create(pmsBrandCreateDto)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, rows)
}

// update 更新品牌
//
//	@Summary		更新品牌
//	@Description	更新品牌
//	@Tags			商品品牌管理
//	@Accept			json
//	@Produce		json
//	@Security		GinJWTMiddleware
//	@Param			id	path	int	true		"品牌id"
//	@Param			pmsBrandParam	body		admin_dto.PmsBrandDTO	true		"品牌信息"
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router /brand/update/{id} [post]
func (router *PmsBrandRouter) update(context *gin.Context) {
	var pathId base_dto.PathVariableDTO
	var pmsBrandUpdateDto admin_dto.PmsBrandDTO
	err := context.ShouldBindUri(&pathId)
	if err = context.ShouldBindJSON(&pmsBrandUpdateDto); err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	rows, err := router.PmsBrandService.Update(pathId.Id, pmsBrandUpdateDto)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, rows)
}

// delete 删除品牌
//
//	@Summary		删除品牌
//	@Description	删除品牌
//	@Tags			商品品牌管理
//	@Accept			json
//	@Produce		json
//	@Security		GinJWTMiddleware
//	@Param			id	path	int	true		"品牌id"
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router /brand/delete/{id} [get]
func (router *PmsBrandRouter) delete(context *gin.Context) {
	var pathId base_dto.PathVariableDTO
	if err := context.ShouldBindUri(&pathId); err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	rows, err := router.PmsBrandService.Delete(pathId.Id)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, rows)
}

// list 分页获取品牌列表
//
//	@Summary		分页获取品牌列表
//	@Description	分页获取品牌列表
//	@Tags			商品品牌管理
//	@Accept			json
//	@Produce		json
//	@Security		GinJWTMiddleware
//	@Param			pageNum	query	int	true		"页码"
//	@Param			pageSize	query	int	true		"每页数量"
//	@Param			keyword	query	string	false		"关键字"
//	@Param			showStatus	query	int	false		"显示状态"
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router /brand/list [get]
func (router *PmsBrandRouter) list(context *gin.Context) {
	var pageDto admin_dto.PsmBrandPageDTO
	if err := context.ShouldBind(&pageDto); err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	list, err := router.PmsBrandService.ListBrand(pageDto.Keyword, pageDto.ShowStatus, pageDto.PageNum, pageDto.PageSize)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, list)
}

// detail 获取品牌详情
//
//	@Summary		获取品牌详情
//	@Description	获取品牌详情
//	@Tags			商品品牌管理
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true		"品牌id"
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router /brand/{id} [get]
func (router *PmsBrandRouter) detail(context *gin.Context) {
	var pathId base_dto.PathVariableDTO
	if err := context.ShouldBindUri(&pathId); err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	detail, err := router.PmsBrandService.GetBrand(pathId.Id)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, detail)
}

// deleteBatch 批量删除品牌
//
//	@Summary		批量删除品牌
//	@Description	批量删除品牌
//	@Tags			商品品牌管理
//	@Accept			json
//	@Produce		json
//	@Param			ids formData []int true "ids"
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router /brand/delete/batch [post]
func (router *PmsBrandRouter) deleteBatch(context *gin.Context) {
	var ids base_dto.IdsDTO
	if err := context.ShouldBind(&ids); err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	rows, err := router.PmsBrandService.DeleteBatch(ids.Ids)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, rows)
}

// updateShowStatus 批量更新显示状态
//
//	@Summary		批量更新显示状态
//	@Description	批量更新显示状态
//	@Tags			商品品牌管理
//	@Accept			json
//	@Produce		json
//	@Param			ids formData []int true "ids"
//	@Param			showStatus formData int true "showStatus"
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router /brand/update/showStatus [post]
func (router *PmsBrandRouter) updateShowStatus(context *gin.Context) {
	var ids base_dto.IdsDTO
	showStatus := context.PostForm("showStatus")
	if err := context.ShouldBind(&ids); err != nil && showStatus == "" {
		gin_common.CreateFail(context, err)
		return
	}
	showStatusInt, _ := strconv.ParseInt(showStatus, 10, 64)
	rows, err := router.PmsBrandService.UpdateShowStatus(ids.Ids, int(showStatusInt))
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, rows)
}

// updateFactoryStatus 批量更新厂家制造商状态
//
//	@Summary		批量更新显示状态
//	@Description	批量更新显示状态
//	@Tags			商品品牌管理
//	@Accept			json
//	@Produce		json
//	@Param			ids formData []int true "ids"
//	@Param			factoryStatus formData int true "factoryStatus"
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router /brand/update/factoryStatus [post]
func (router *PmsBrandRouter) updateFactoryStatus(context *gin.Context) {
	var ids base_dto.IdsDTO
	factoryStatus := context.PostForm("factoryStatus")
	if err := context.ShouldBind(&ids); err != nil && factoryStatus == "" {
		gin_common.CreateFail(context, err)
		return
	}
	factoryStatusInt, _ := strconv.ParseInt(factoryStatus, 10, 64)
	rows, err := router.PmsBrandService.UpdateFactoryStatus(ids.Ids, int(factoryStatusInt))
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, rows)
}
