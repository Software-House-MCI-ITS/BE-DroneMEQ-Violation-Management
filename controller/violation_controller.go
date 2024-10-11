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
	ViolationController interface {
		Create(ctx *gin.Context)
		GetAllViolation(ctx *gin.Context)
		ViolationById(ctx *gin.Context)
		ViolationByStatus(ctx *gin.Context)
	}

	violationController struct {
		violationService service.ViolationService
	}
)

func NewViolationController(vs service.ViolationService) ViolationController {
	return &violationController{
		violationService: vs,
	}
}

func (c *violationController) Create(ctx *gin.Context) {
	var violation dto.ViolationCreateRequest
	// Check if the request body is valid
	if err := ctx.ShouldBind(&violation); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// Create the violation
	result, err := c.violationService.CreateViolation(ctx.Request.Context(), violation)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_VIOLATION, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_VIOLATION, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *violationController) GetAllViolation(ctx *gin.Context) {
	var req dto.PaginationRequest
	// Check if the request body is valid
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// Get all violations with pagination
	result, err := c.violationService.GetAllViolationWithPagination(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_ALL_VIOLATION, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_ALL_VIOLATION, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *violationController) ViolationById(ctx *gin.Context) {
    // Get pk_id from URL parameter
    violationIdStr := ctx.Param("pk_id")
    
    // Convert pk_id to int64
    violationId, err := strconv.ParseInt(violationIdStr, 10, 64)
    if err != nil {
        // Handle the case where pk_id is not a valid integer
        res := utils.BuildResponseFailed("Invalid pk_id", err.Error(), nil)
        ctx.JSON(http.StatusBadRequest, res)
        return
    }

    // Fetch the violation using the violationService
    result, err := c.violationService.GetViolationById(ctx.Request.Context(), violationId)
    if err != nil {
        res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_VIOLATION_BY_ID, err.Error(), nil)
        ctx.JSON(http.StatusBadRequest, res)
        return
    }

    // Return the result
    res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_VIOLATION_BY_ID, result)
    ctx.JSON(http.StatusOK, res)
}

func (c *violationController) ViolationByStatus(ctx *gin.Context) {
    // Get status from URL parameter
    status := ctx.Param("status")

    // Call the violationService to get violations by status
    result, err := c.violationService.GetViolationByStatus(ctx.Request.Context(), status)
    if err != nil {
        res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_VIOLATION_BY_STATUS, err.Error(), nil)
        ctx.JSON(http.StatusBadRequest, res)
        return
    }

    // Return the result
    res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_VIOLATION_BY_STATUS, result)
    ctx.JSON(http.StatusOK, res)
}
