package middleware

import (
	"errors"
	"net/http"
	"red/serializer"
	"red/util/jwt"

	"github.com/gin-gonic/gin"
)

// AuthRequired 需要登录
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := jwt.Get(c)
		if err == nil {
			c.Set("uid", uid)
			c.Next()
			return
		}

		c.JSON(200, serializer.Err(http.StatusUnauthorized, "请先登录", errors.New("请先登录")))
		c.Abort()
	}
}
