package service

import (
	"red/model"
	"red/serializer"

	"github.com/gin-gonic/gin"
)

func GetProjectList(c *gin.Context) []serializer.Project {
	area := c.Param("area")

	projects := make([]model.Project, 0)
	model.DB.Where("area = ?", area).Find(&projects)

	projectIds := make([]uint, 0)
	for _, project := range projects {
		projectIds = append(projectIds, project.ID)
	}

	incomes := make([]model.Income, 0)
	model.DB.Where("project_id in (?)", projectIds).Find(&incomes)

	return serializer.BuildProjects(projects, model.GetPartyMap(), incomes)
}

func DeleteProject(c *gin.Context) serializer.Response {
	id := c.Param("id")

	model.DB.Where("id = ?", id).Delete(&model.Project{})
	return serializer.Success()
}

func GetIncomeList(c *gin.Context) serializer.Response {
	tp, value := c.Param("type"), c.Param("value")
	incomes := make([]model.Income, 0)
	if tp == "project" {
		model.DB.Where("project_id = ?", value).Find(&incomes)
	} else {
		sql := "select a.* from incomes a join projects b on b.deleted_at is null and a.project_id = b.id and b.area = '" + value + "'"
		model.DB.Raw(sql).Scan(&incomes)
	}
	return serializer.Success(incomes)
}

func GetProjects(c *gin.Context) []serializer.Project {
	projects := make([]model.Project, 0)
	model.DB.Find(&projects)

	projectIds := make([]uint, 0)
	for _, project := range projects {
		projectIds = append(projectIds, project.ID)
	}

	incomes := make([]model.Income, 0)
	model.DB.Where("project_id in (?)", projectIds).Find(&incomes)

	return serializer.BuildProjects(projects, model.GetPartyMap(), incomes)
}

func IncomeStatistics(c *gin.Context) serializer.Response {
	var sql string
	switch c.Param("type") {
	case "project":
		sql = "select b.id, b.name, sum(a.amount) as value from incomes a, projects b where a.project_id = b.id and b.deleted_at is null group by b.id, b.name"
	case "area":
		sql = "select b.area as name, sum(a.amount) as value from incomes a, projects b where a.project_id = b.id and b.deleted_at is null group by b.area"
	}
	rsp := make([]FinanceAggRsp, 0)
	model.DB.Raw(sql).Scan(&rsp)
	return serializer.Success(rsp)
}

func GetPartyList() serializer.Response {
	parties := make([]model.Party, 0)
	model.DB.Find(&parties)

	return serializer.Success(parties)
}
