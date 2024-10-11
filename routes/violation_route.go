package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mci-its/backend-service/controller"
)

func Violation(route *gin.Engine, violationController controller.ViolationController) {
	routes := route.Group("/api/violations")
	{
		// User
		routes.POST("", violationController.Create)
		routes.GET("", violationController.GetAllViolation)
		routes.GET("/:pk_id", violationController.ViolationById)
		routes.GET("/status/:status", violationController.ViolationByStatus)
	}
}
