package api

import (
	"encoding/json"
	"fmt"
	"red/serializer"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// CurrentUser 获取当前用户
func CurrentUID(c *gin.Context) uint {
	if uid, _ := c.Get("uid"); uid != nil {
		id, err := strconv.Atoi(uid.(string))
		if err == nil {
			return uint(id)
		}
	}
	return 0
}

// ErrorResponse 返回错误消息
func ErrorResponse(err error) serializer.Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		//for _, e := range ve {
		//	field := conf.T(fmt.Sprintf("Field.%s", e.Field))
		//	tag := conf.T(fmt.Sprintf("Tag.Valid.%s", e.Tag))
		//	return serializer.ParamErr(
		//		fmt.Sprintf("%s%s", field, tag),
		//		err,
		//	)
		//}
		return serializer.ParamErr(
			fmt.Sprint(ve),
			err,
		)
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.ParamErr("JSON类型不匹配", err)
	}

	return serializer.ParamErr("参数错误", err)
}
