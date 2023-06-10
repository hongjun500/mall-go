// @author hongjun500
// @date 2023/6/10
// @tool ThinkPadX1隐士
// Created with GoLand 2022.2
// Description: gorm 分页

package gorm_common

import (
	"gorm.io/gorm"
)

type GormCommonPage struct {
	// 当前页
	PageNum int `json:"pageNum" form:"page_num" binding:"required" default:"1"`
	// 每页数量
	PageSize int `json:"pageSize" form:"page_size" binding:"required" default:"10"`
	// 总页数
	TotalPage int64 `json:"totalPage" form:"total_page"`
	// 总记录数
	Total int64 `json:"total" form:"total"`
	// 分页数据
	List any `json:"list" form:"list"`
	// 排序
	OrderBy string `json:"orderBy" form:"order_by"`
	// 排序方式
	Sort string `json:"sort" form:"sort"`
}

type CommonPage interface {
	// SetPageNum 设置当前页
	SetPageNum(pageNum int)
	// GetPageNum 获取当前页
	GetPageNum() int

	// SetPageSize 设置每页数量
	SetPageSize(pageSize int)
	// GetPageSize 获取每页数量
	GetPageSize() int

	// SetTotalPage 设置总页数
	SetTotalPage(totalPage int64)
	// GetTotalPage 获取总页数
	GetTotalPage() int64

	// SetTotal 设置总记录数
	SetTotal(total int64)
	// GetTotal 获取总记录数
	GetTotal() int64

	// SetList 设置分页数据
	SetList(list any)
	// GetList 获取分页数据
	GetList() any

	// SetOrderBy 设置排序
	SetOrderBy(orderBy string)
	// GetOrderBy 获取排序
	GetOrderBy() string

	// SetSort 设置排序方式
	SetSort(sort string)
	// GetSort 获取排序方式
	GetSort() string
}

func (g *GormCommonPage) SetPageNum(pageNum int) {
	g.PageNum = pageNum
}

func (g *GormCommonPage) GetPageNum() int {
	return g.PageNum
}

func (g *GormCommonPage) SetPageSize(pageSize int) {
	g.PageSize = pageSize
}

func (g *GormCommonPage) GetPageSize() int {
	return g.PageSize
}

func (g *GormCommonPage) SetTotalPage(totalPage int64) {
	g.TotalPage = totalPage
}

func (g *GormCommonPage) GetTotalPage() int64 {
	return g.TotalPage
}

func (g *GormCommonPage) SetTotal(total int64) {
	g.Total = total
}

func (g *GormCommonPage) GetTotal() int64 {
	return g.Total
}

func (g *GormCommonPage) SetList(list any) {
	g.List = list
}

func (g *GormCommonPage) GetList() any {
	return g.List
}

func (g *GormCommonPage) SetOrderBy(orderBy string) {
	g.OrderBy = orderBy
}
func (g *GormCommonPage) GetOrderBy() string {
	return g.OrderBy
}

func (g *GormCommonPage) SetSort(sort string) {
	g.Sort = sort
}
func (g *GormCommonPage) GetSort() string {
	return g.Sort
}

func NewPage(pageNum, pageSize int, orderBy ...string) CommonPage {
	if orderBy == nil || len(orderBy) == 0 {
		return &GormCommonPage{
			PageNum:  pageNum,
			PageSize: pageSize,
		}
	}
	return &GormCommonPage{
		PageNum:  pageNum,
		PageSize: pageSize,
		OrderBy:  orderBy[0],
		Sort:     orderBy[1],
	}
}

// ExecutePagedQuery 执行分页查询
// db gorm 数据库连接对象
// page 分页对象 pageNum 当前页 pageSize 每页数量 totalPage 总页数 total 总记录数 list 分页数据 orderBy 排序 sort 排序方式
// queryFunc 查询函数 返回值为 *gorm.DB 类型 用于链式调用 gorm 的查询方法
// 例如: db.Model(&model).Where("id = ?", id) 传入的 model 为 result 的指针 用于接收查询结果
// 例如: db.Model(&model).Where("id = ?", id).Find(&model)
// result 结果集
// 返回值 error
func ExecutePagedQuery(db *gorm.DB, page CommonPage, result any, queryFunc func(*gorm.DB) *gorm.DB) error {
	var count int64
	query := queryFunc(db.Model(result))
	if page.GetOrderBy() != "" && page.GetSort() != "" {
		query = query.Order(page.GetOrderBy() + " " + page.GetSort())
	}
	if page.GetOrderBy() != "" && page.GetSort() == "" {
		query = query.Order(page.GetOrderBy() + " " + "desc")
	}
	if err := query.Count(&count).Error; err != nil {
		return err
	}

	totalPage := count / int64(page.GetPageSize())
	if count%int64(page.GetPageSize()) > 0 {
		totalPage++
	}

	if int64(page.GetPageNum()) > totalPage {
		page.SetPageNum(int(totalPage))
	}

	offset := (page.GetPageNum() - 1) * page.GetPageSize()

	query = query.Offset(offset).Limit(page.GetPageSize())

	if err := query.Find(result).Error; err != nil {
		return err
	}

	page.SetTotalPage(totalPage)
	page.SetTotal(count)
	page.SetList(result)

	return nil
}

// ExecutePagedSQLQuery 执行原生 sql 分页查询
// db gorm 数据库连接对象
// page 分页参数 pageNum 当前页 pageSize 每页数量 totalPage 总页数 total 总记录数 list 分页数据 orderBy 排序 sort 排序方式
// result 结果集
// countSQL 查询总数sql
// querySQL 查询数据sql
// countArgs 查询总数参数
// queryArgs 查询数据参数
// 返回值 error
func ExecutePagedSQLQuery(db *gorm.DB, page CommonPage, result interface{}, countSQL, querySQL string, countArgs, queryArgs []interface{}) error {
	var count int64
	if err := db.Raw(countSQL, countArgs...).Count(&count).Error; err != nil {
		return err
	}

	totalPage := count / int64(page.GetPageSize())
	if int(count)%page.GetPageSize() > 0 {
		totalPage++
	}
	if int64(page.GetPageNum()) > totalPage {
		page.SetPageNum(int(totalPage))
	}

	offset := (page.GetPageNum() - 1) * page.GetPageSize()

	query := db.Raw(querySQL, queryArgs...).Offset(offset).Limit(page.GetPageSize())

	if err := query.Find(result).Error; err != nil {
		return err
	}

	page.SetTotalPage(totalPage)
	page.SetTotal(count)
	page.SetList(result)

	return nil
}
