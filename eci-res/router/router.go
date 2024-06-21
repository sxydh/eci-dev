package router

import (
	"eci-res/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	eciController := controller.NewEciController()
	r.GET("/eci/list", eciController.ListEci)
	r.POST("/eci", eciController.AddEci)
	r.DELETE("/eci", eciController.DeleteEci)
	r.GET("/term", eciController.ContainerWsHandler)
	return r
}
