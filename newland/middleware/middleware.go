package middleware

import (
	"github.com/astaxie/beego"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

func RequestIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Request-Id", uuid.Must(uuid.NewV4()).String())
		c.Next()
	}
}

func Middleware2() gin.HandlerFunc {
	return func(c *gin.Context) {
		beego.Info("entry middleware2...")
		c.Next()
		beego.Info("exit middleware2...")
	}
}
