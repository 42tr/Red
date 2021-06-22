package service

import (
	"red/model"
	"red/serializer"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FinanceService struct {
	StartDate  string `json:"startDate"`
	EndDate    string `json:"endDate"`
	ProjectId  uint   `json:"projectId"`
	UserId     uint   `json:"userId"`
	ApprovalId uint   `json:"menuId"`
	Status     int    `json:"status"`
	Page       Page   `json:"page"`
}

type Page struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
}

type FinanceRsp struct {
	ID         uint    `json:"id"`
	Menu       string  `json:"menu"`
	Project    string  `json:"project"`
	User       string  `json:"user"`
	ApplyDate  string  `json:"applyDate"`
	HandleDate string  `json:"handleDate"`
	Money      float32 `json:"money"`
}

func (service *FinanceService) GetFinanceList(c *gin.Context) serializer.Response {
	sql := "select a.id, b.name as menu, c.name as project, d.nickname as user, a.created_at as apply_date, " +
		"a.updated_at as handle_date, a.money from applies a, approvals b, projects c, users d " +
		"where a.approval_id = b.id and a.project_id = c.id and a.user_id = d.id "
	sql += service.conditionSql()
	from, size := 0, 20
	if service.Page.PageSize > 0 {
		size = service.Page.PageSize
	}
	if service.Page.PageNum > 0 {
		from = service.Page.PageNum * size
	}
	sql += " limit " + strconv.Itoa(from) + ", " + strconv.Itoa(size)

	var rsp []FinanceRsp
	model.DB.Raw(sql).Scan(&rsp)
	return serializer.Success(rsp)
}

type FinanceAggRsp struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (service *FinanceService) GetFinanceAggregation(c *gin.Context) serializer.Response {
	tip := c.Param("type")
	var sql string
	conditionSql := service.conditionSql()
	switch tip {
	case "project":
		sql = "select sum(a.money) as value, b.name from applies a, projects b " +
			"where a.project_id = b.id " + conditionSql + " group by b.name"
	case "time":
		sql = "select date_format(created_at, '%y-%m-%d') as name, sum(money) as value " +
			"from applies a where 1 = 1 " + conditionSql + " group by date_format(created_at, '%y-%m-%d')"
	case "menu":
		sql = "select b.name, sum(a.money) as value from applies a, approvals b " +
			"where a.approval_id = b.id " + conditionSql + " group by b.name"
	case "user":
		sql = "select a.user_id as name, sum(a.money) as value from applies a where 1=1 " +
			conditionSql + " group by user_id"
	default:
		return serializer.ParamErr("聚合类型错误", nil)
	}
	var rsp []FinanceAggRsp
	model.DB.Raw(sql).Scan(&rsp)
	return serializer.Success(rsp)
}

func (service *FinanceService) conditionSql() string {
	var sql string
	if service.ApprovalId > 0 {
		sql += " and a.approval_id = " + strconv.Itoa(int(service.ApprovalId))
	}
	if service.ProjectId > 0 {
		sql += " and a.project_id = " + strconv.Itoa(int(service.ProjectId))
	}
	if len(service.StartDate) > 0 {
		sql += " and a.created_at >= " + service.StartDate
	}
	if len(service.EndDate) > 0 {
		sql += " and a.created_at <= " + service.EndDate
	}
	if service.UserId > 0 {
		sql += " and a.user_id = " + strconv.Itoa(int(service.UserId))
	}
	if service.Status >= 0 {
		sql += " and a.status = " + strconv.Itoa(service.Status)
	}
	return sql
}
