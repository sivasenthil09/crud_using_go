package routes

import (
	"net/http"
	"restapi/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, controller controllers.UserController) {

	router.POST("/create", controller.CreateUser)
	router.GET("/get/:name", controller.GetUser)
	router.GET("/getall", controller.GetAll)
	router.PATCH("/update", controller.UpdateUser)
	router.DELETE("/delete/:name", controller.DeleteUser)
}
func Default(router *gin.Engine) {
	router.GET("/api", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Server is Healthy"})
	})
}
