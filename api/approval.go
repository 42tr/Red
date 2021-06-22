package api

import (
	"red/service"

	"github.com/gin-gonic/gin"
)

func GetApprovalMenu(c *gin.Context) {
	res := service.GetApprovalMenu(c)
	c.JSON(200, res)
}

func GetApplyList(c *gin.Context) {
	uid := CurrentUID(c)
	res := service.GetApplyList(c, uid)
	c.JSON(200, res)
}

func GetApplyDetail(c *gin.Context) {
	res := service.GetApplyDetail(c)
	c.JSON(200, res)
}

func AddRemark(c *gin.Context) {
	uid := CurrentUID(c)
	var service service.RemarkService
	if err := c.ShouldBind(&service); err == nil {
		res := service.AddRemark(c, uid)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func ChangeApplyStatus(c *gin.Context) {
	res := service.ChangeApplyStatus(c)
	c.JSON(200, res)
}

func AddApply(c *gin.Context) {
	var service service.AddApplyService
	if err := c.ShouldBind(&service); err == nil {
		res := service.AddApply(CurrentUID(c))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
