package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mci-its/backend-service/controller"
)

func ViolationNote(route *gin.Engine, violationNoteController controller.ViolationNoteController) {
	routes := route.Group("/api/violations")
	{
		routes.POST("/:pk_id/notes", violationNoteController.CreateVsNote)
	}
}