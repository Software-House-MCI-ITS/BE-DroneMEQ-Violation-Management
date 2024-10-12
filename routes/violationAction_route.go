package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mci-its/backend-service/controller"
)

func ViolationAction(route *gin.Engine, violationActionController controller.ViolationActionController) {
	routes := route.Group("/api/violations")
	{
		routes.POST("/:pk_id/actions", violationActionController.CreateVsAction)
	}
}