package server

import (
	"red/api"
	"red/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())

	r.GET("/ping", api.Ping)

	a := r.Group("api")
	{
		a.GET("dic", api.GetDic)
	}

	project := r.Group("project")
	project.Use(middleware.AuthRequired())
	{
		project.POST("", api.AddProject)
		project.GET("list", api.GetAllProjectList)
		project.GET("list/:area", api.GetProjectList)
		project.PUT(":id", api.ChangeProject)
		project.DELETE(":id", api.DeleteProject)
		project.GET("income/detail/:type/:value", api.GetIncomeList)
		project.POST("income/detail", api.AddIncome)
		project.GET("income/statistics/:type", api.GetIncomeStatistics)
		project.GET("party/list", api.GetPartyList)
		project.GET("excel", api.GetExcel)
	}

	approval := r.Group("approval")
	approval.Use(middleware.AuthRequired())
	{
		approval.GET("menu", api.GetApprovalMenu)
		approval.GET("list/:status", api.GetApplyList)
		approval.GET("detail/:id", api.GetApplyDetail)
		approval.POST("remark/:id", api.AddRemark)
		approval.PUT(":id/:status", api.ChangeApplyStatus)
		approval.POST("", api.AddApply)
	}

	finance := r.Group("finance")
	finance.Use(middleware.AuthRequired())
	{
		finance.POST("list", api.GetFinanceList)
		finance.POST("aggregation/:type", api.GetFinanceAggregation)
	}

	return r
}
