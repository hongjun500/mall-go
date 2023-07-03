//	@author	hongjun500
//	@date	2023/6/10
//	@tool	ThinkPadX1隐士
// Created with GoLand 2022.2
// Description: gorm 分页

package internal

import (
	"github.com/hongjun500/mall-go/pkg"
	"gorm.io/gorm"
)

type GormPage struct {
	Db *gorm.DB
	// QueryFunc 查询函数 返回值为 *gorm.DB 类型 用于链式调用 gorm 的查询方法
	// 例如: db.Model(&model).Where("id = ?", id) 传入的 model 为 result 的指针 用于接收查询结果
	// 例如: db.Model(&model).Where("id = ?", id).Find(&model)
	QueryFunc func(*gorm.DB) *gorm.DB
	*pkg.CommonPage
}

type NativeSqlPage struct {
	Db *gorm.DB
	// CountSQL 查询总数sql 例如: select count(*) from table where id = ?
	// QuerySQL 查询数据sql 例如: select * from table where id = ?
	CountSQL, QuerySQL string
	// CountArgs 查询总数参数
	// QueryArgs 查询数据参数
	CountArgs, QueryArgs []any
	*pkg.CommonPage
}

func (page *GormPage) Paginate() error {
	var count int64
	query := page.QueryFunc(page.Db.Model(page.List))

	if page.OrderBy != "" && page.Sort != "" {
		query = query.Order(page.OrderBy + " " + page.Sort)
	}
	if page.OrderBy != "" && page.Sort == "" {
		query = query.Order(page.OrderBy + " " + "desc")
	}
	if err := query.Count(&count).Error; err != nil {
		return err
	}
	totalPage := count / int64(page.PageSize)
	if count%int64(page.PageSize) > 0 {
		totalPage++
	}

	if int64(page.PageNum) > totalPage {
		page.PageNum = int(totalPage)
	}

	offset := (page.PageNum - 1) * page.PageSize

	query = query.Offset(offset).Limit(page.PageSize)

	if err := query.Find(page.List).Error; err != nil {
		return err
	}

	page.TotalPage = totalPage
	page.Total = count
	return nil
}

func (page *NativeSqlPage) Paginate() error {
	var count int64
	if err := page.Db.Raw(page.CountSQL, page.CountArgs...).Count(&count).Error; err != nil {
		return err
	}

	totalPage := count / int64(page.PageSize)
	if int(count)%page.PageSize > 0 {
		totalPage++
	}
	if int64(page.PageNum) > totalPage {
		page.PageNum = int(totalPage)
	}

	offset := (page.PageNum - 1) * page.PageSize

	query := page.Db.Raw(page.QuerySQL, page.QueryArgs...).Offset(offset).Limit(page.PageSize)

	if err := query.Find(page.List).Error; err != nil {
		return err
	}
	page.TotalPage = totalPage
	page.Total = count
	return nil
}

func NewGormPage(db *gorm.DB, pageNum, pageSize int, orderBy ...string) *GormPage {
	if orderBy == nil || len(orderBy) == 0 {
		return &GormPage{
			db,
			nil,
			&pkg.CommonPage{
				PageNum:  pageNum,
				PageSize: pageSize,
			},
		}
	}
	return &GormPage{
		db,
		nil,
		&pkg.CommonPage{
			PageNum:  pageNum,
			PageSize: pageSize,
			OrderBy:  orderBy[0],
			Sort:     orderBy[1],
		},
	}
}

func NewNativeSqlPage(db *gorm.DB, pageNum, pageSize int, orderBy ...string) *NativeSqlPage {
	if orderBy == nil || len(orderBy) == 0 {
		return &NativeSqlPage{
			db,
			"",
			"",
			nil,
			nil,
			&pkg.CommonPage{
				PageNum:  pageNum,
				PageSize: pageSize,
			},
		}
	}
	return &NativeSqlPage{
		db,
		"",
		"",
		nil,
		nil,
		&pkg.CommonPage{
			PageNum:  pageNum,
			PageSize: pageSize,
			OrderBy:  orderBy[0],
			Sort:     orderBy[1],
		},
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
/*func ExecutePagedQuery(db *gorm.DB, page CommonPage, result any, queryFunc func(*gorm.DB) *gorm.DB) error {
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
}*/
/*
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
}*/
