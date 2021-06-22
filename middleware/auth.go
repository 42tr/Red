package middleware

import (
	"fmt"
	"net/http"
	"red/serializer"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth, err := c.Cookie("Authorization")
		if err != nil {
			return
		}
		claims, err := ParseToken(auth)
		if err != nil {
			return
		}
		c.Set("uid", claims["uid"])
	}
}

const (
	SECRETKEY = "243223ffslsfsldfl412fdsfsdf" //私钥
)

//解析token
func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(SECRETKEY), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

// AuthRequired 需要登录
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if uid, _ := c.Get("uid"); uid != nil {
			c.Next()
			return
		}

		c.JSON(http.StatusUnauthorized, serializer.CheckLogin())
		c.Abort()
	}
}
