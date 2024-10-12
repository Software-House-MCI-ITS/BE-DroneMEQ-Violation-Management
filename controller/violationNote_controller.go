package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mci-its/backend-service/dto"
	"github.com/mci-its/backend-service/service"
	"github.com/mci-its/backend-service/utils"
)
type (
	ViolationNoteController interface {
		CreateVsNote(ctx *gin.Context)
	}
	violationNoteController struct {
		violationNoteService service.ViolationNoteService
	}
)

func NewViolationNoteController(vs service.ViolationNoteService) ViolationNoteController {
	return &violationNoteController{
		violationNoteService: vs,
	}
}

func (c *violationNoteController) CreateVsNote(ctx *gin.Context) {
	violationIdStr := ctx.Param("pk_id")

	violationId, err := strconv.ParseInt(violationIdStr, 10, 64)
	if err != nil {
		res := utils.BuildResponseFailed("Invalid pk_id", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var violationNote dto.ViolationNoteCreateRequest
	// Check if the request body is valid
	if err := ctx.ShouldBind(&violationNote); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// Create the violation note
	result, err := c.violationNoteService.CreateViolationNote(ctx.Request.Context(), violationNote, violationId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_VIOLATION_NOTE, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	//return the result
	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_VIOLATION_NOTE, result)
	ctx.JSON(http.StatusOK, res)
}