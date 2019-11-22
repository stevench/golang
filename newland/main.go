package main

import (
	"github.com/gin-gonic/gin"
	"io"
	. "newland/apis"
	db "newland/database"
	. "newland/middleware"
	"os"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	router.Use(RequestIdMiddleware())

	router.Use(LoadTlsMiddleware())

	router.GET("/", IndexApi)

	router.GET("/person/:id", GetPersonApi)

	router.POST("/person", AddPersonApi)

	router.PUT("/person/:id", ModPersonApi)

	router.DELETE("/person/:id", DelPersonApi)

	router.GET("/persons", GetPersonsApi)

	return router
}

func main() {
	gin.DefaultWriter = io.MultiWriter(os.Stdout)
	defer db.SqlDB.Close()
	router := initRouter()
	router.RunTLS(":8000", "./ssl/server.crt", "./ssl/server.key")
}
