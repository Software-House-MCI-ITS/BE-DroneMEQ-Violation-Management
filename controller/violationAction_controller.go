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
	ViolationActionController interface {
		CreateVsAction(ctx *gin.Context)
	}

	violationActionController struct {
		violationActionService service.ViolationActionService
	}
)

func NewViolationActionController(vs service.ViolationActionService) ViolationActionController {
	return &violationActionController{
		violationActionService: vs,
	}
}

func (c *violationActionController) CreateVsAction(ctx *gin.Context) {
	violationIdStr := ctx.Param("pk_id")

	violationId, err := strconv.ParseInt(violationIdStr, 10, 64)
	if err != nil {
		res := utils.BuildResponseFailed("Invalid pk_id", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var violationAction dto.ViolationActionCreateRequest
	// Check if the request body is valid
	if err := ctx.ShouldBind(&violationAction); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// Create the violation action
	result, err := c.violationActionService.CreateViolationAction(ctx.Request.Context(), violationAction, violationId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_VIOLATION_ACTION, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	//return the result
	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_VIOLATION_ACTION, result)
	ctx.JSON(http.StatusOK, res)
}
