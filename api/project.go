package api

import (
	"net/http"
	"red/model"
	"red/serializer"
	"red/service"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
)

func AddProject(c *gin.Context) {
	var service service.ProjectAddService
	uid := CurrentUID(c)
	if err := c.ShouldBind(&service); err == nil {
		res := service.AddProject(c, uid)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func GetAllProjectList(c *gin.Context) {
	projects := service.GetAllProjectList(c)
	c.JSON(http.StatusOK, serializer.BuildCommonProjects(projects))
}

func GetProjectList(c *gin.Context) {
	res := service.GetProjectList(c)
	c.JSON(200, serializer.Success(res))
}

func ChangeProject(c *gin.Context) {
	var service service.ProjectAddService
	if err := c.ShouldBind(&service); err == nil {
		res := service.ChangeProject(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func DeleteProject(c *gin.Context) {
	res := service.DeleteProject(c)
	c.JSON(200, res)
}

func GetIncomeList(c *gin.Context) {
	res := service.GetIncomeList(c)
	c.JSON(200, res)
}

func AddIncome(c *gin.Context) {
	var income model.Income
	if err := c.ShouldBind(&income); err == nil {
		income.CreateBy = CurrentUID(c)
		model.DB.Create(&income)
		c.JSON(200, serializer.Success())
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func GetIncomeStatistics(c *gin.Context) {
	c.JSON(200, service.IncomeStatistics(c))
}

func GetPartyList(c *gin.Context) {
	c.JSON(200, service.GetPartyList())
}

func GetExcel(c *gin.Context) {
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+"项目信息.xlsx")
	c.Header("Content-Transfer-Encoding", "binary")

	projects := service.GetProjects(c)

	m := make(map[string][]serializer.Project)
	for _, project := range projects {
		if m[project.Area] == nil {
			m[project.Area] = []serializer.Project{project}
		} else {
			m[project.Area] = append(m[project.Area], project)
		}
	}

	f := excelize.NewFile()
	for area, list := range m {
		s := area
		f.NewSheet(s)
		f.SetRowHeight(s, 1, 50)
		for i := 2; i <= 100; i++ {
			f.SetRowHeight(s, i, 40)
		}
		f.SetColWidth(s, "A", "A", 10)
		f.SetColWidth(s, "B", "C", 48)
		f.SetColWidth(s, "D", "I", 23)

		headStyle, _ := f.NewStyle(`{"alignment":{"horizontal": "center", "vertical": "center"},"font":{"bold":true,"size":22}}`)
		f.SetCellStyle(s, "A1", "A1", headStyle)
		style, _ := f.NewStyle(`{"alignment":{"horizontal": "center", "vertical": "center"},"font":{"size":16}}`)
		f.SetCellStyle(s, "A2", "I100", style)

		f.MergeCell(s, "A1", "I1")
		f.SetCellValue(s, "A1", s)

		f.SetCellValue(s, "A2", "序号")
		f.SetCellValue(s, "B2", "甲方全称")
		f.SetCellValue(s, "C2", "项目名称")
		f.SetCellValue(s, "D2", "项目类别")
		f.SetCellValue(s, "E2", "中标价（万元）")
		f.SetCellValue(s, "F2", "代理费（元）")
		f.SetCellValue(s, "G2", "清单编制费（元）")
		f.SetCellValue(s, "H2", "合计（元）")
		f.SetCellValue(s, "I2", "报名费（元）")

		for i, project := range list {
			index := strconv.Itoa(i + 3)
			f.SetCellValue(s, "A"+index, i+1)
			f.SetCellValue(s, "B"+index, project.Party)
			f.SetCellValue(s, "C"+index, project.Name)
			f.SetCellValue(s, "D"+index, project.Type)
			f.SetCellValue(s, "E"+index, project.Price)
			f.SetCellValue(s, "F"+index, project.Price1)
			f.SetCellValue(s, "G"+index, project.Price2)
			f.SetCellValue(s, "H"+index, project.Price1+project.Price2)
			f.SetCellValue(s, "I"+index, project.Price3)
		}
	}

	f.DeleteSheet("Sheet1")
	_ = f.Write(c.Writer)
}
