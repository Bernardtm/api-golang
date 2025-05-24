package status

import (
	"net/http"
	"strconv"

	"bernardtm/backend/internal/core/shareds"

	"github.com/gin-gonic/gin"
)

type StatusController interface {
	GetAll(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Paginate(ctx *gin.Context)
}
type statusController struct {
	service StatusService
}

func NewStatusController(service StatusService) *statusController {
	return &statusController{service: service}
}

// GetAll Get all Status
// @Summary Get all Status
// @Tags Status
// @Produce json
// @Security BearerAuth
// @Success 200 {array} StatusResponse
// @Failure 500 {object} shareds.ErrorResponse
// @Router /status [get]
func (c *statusController) GetAll(ctx *gin.Context) {
	statuses, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, shareds.ErrorResponse{Message: "Error fetching statuses"})
		return
	}
	ctx.JSON(http.StatusOK, statuses)
}

// GetByID gets a status by ID
// @Summary Get Status by ID
// @Tags Status
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID of the Status"
// @Success 200 {object} StatusResponse
// @Failure 400 {object} shareds.ErrorResponse
// @Failure 404 {object} shareds.ErrorResponse
// @Router /status/{id} [get]
func (c *statusController) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")

	status, err := c.service.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, shareds.ErrorResponse{Message: "Status not found"})
		return
	}
	ctx.JSON(http.StatusOK, status)
}

// Create creates a new status
// @Summary Create a new Status
// @Tags Status
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param input body StatusRequest true "Status Data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} shareds.ErrorResponse
// @Failure 500 {object} shareds.ErrorResponse
// @Router /status [post]
func (c *statusController) Create(ctx *gin.Context) {
	var input StatusRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, shareds.ErrorResponse{Message: "Invalid input data"})
		return
	}

	createdID, err := c.service.Create(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, shareds.ErrorResponse{Message: "Error creating status"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": createdID, "message": "Status created successfully"})
}

// Update updates a status existente
// @Summary Update an existing Status
// @Tags Status
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID of the Status"
// @Param input body StatusRequest true "Updated Status Data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} shareds.ErrorResponse
// @Failure 500 {object} shareds.ErrorResponse
// @Router /status/{id} [put]
func (c *statusController) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	var input StatusRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, shareds.ErrorResponse{Message: "Invalid input data"})
		return
	}

	err := c.service.Update(id, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, shareds.ErrorResponse{Message: "Error updating status"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Status updated successfully"})
}

// Delete deletes a status by ID
// @Summary Delete Status by ID
// @Tags Status
// @Param id path string true "ID of the Status"
// @Success 204
// @Security BearerAuth
// @Failure 400 {object} shareds.ErrorResponse
// @Failure 404 {object} shareds.ErrorResponse
// @Router /status/{id} [delete]
func (c *statusController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, shareds.ErrorResponse{Message: "Status not found"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

// Paginate paginate Status
// @Summary Paginate Status
// @Tags Status
// @Produce json
// @Security BearerAuth
// @Param page query int true "Page Number"
// @Param size query int true "Page Size"
// @Success 200 {array} StatusResponse
// @Failure 400 {object} shareds.ErrorResponse
// @Failure 500 {object} shareds.ErrorResponse
// @Router /status/paginate [get]
func (c *statusController) Paginate(ctx *gin.Context) {
	pageParam := ctx.DefaultQuery("page", "1")
	sizeParam := ctx.DefaultQuery("size", "5")

	page, err := strconv.Atoi(pageParam)
	if err != nil || page <= 0 {
		page = 1
	}

	limit, err := strconv.Atoi(sizeParam)
	if err != nil || limit <= 0 || limit > 50 {
		limit = 10
	}

	statuses, err := c.service.Paginate(page, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, shareds.ErrorResponse{Message: "Error fetching paginated statuses"})
		return
	}

	ctx.JSON(http.StatusOK, statuses)
}
