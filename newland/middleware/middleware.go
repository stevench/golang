package middleware

import (
	"github.com/astaxie/beego"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"github.com/unrolled/secure"
)

func RequestIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Request-Id", uuid.Must(uuid.NewV4()).String())
		c.Next()
	}
}

func LoadTlsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		beego.Info("load ssl crt...")
		middleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "localhost:8000",
		})
		err := middleware.Process(c.Writer, c.Request)
		if err != nil {
			beego.Info(err)
			return
		}
		c.Next()
		beego.Info("finished load ssl srt...")
	}
}
