package menus

import (
	"bernardtm/backend/internal/core/shareds"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MenusController interface {
	GetAll(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Paginate(ctx *gin.Context)
	GetMenusByUserID(ctx *gin.Context)
}

type menusController struct {
	service MenusService
}

func NewMenusController(service MenusService) *menusController {
	return &menusController{service: service}
}

// GetAll get all menus
// @Summary Get all Menus
// @Tags Menus
// @Produce json
// @Security BearerAuth
// @Success 200 {array} MenusResponse
// @Failure 500 {object} shareds.ErrorResponse
// @Router /menus [get]
func (c *menusController) GetAll(ctx *gin.Context) {
	menus, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, shareds.ErrorResponse{Message: "Error fetching menus"})
		return
	}
	ctx.JSON(http.StatusOK, menus)
}

// GetByID Get Menu by ID
// @Summary Get Menu by ID
// @Tags Menus
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID of the Menu"
// @Success 200 {object} MenusResponse
// @Failure 400 {object} shareds.ErrorResponse
// @Failure 404 {object} shareds.ErrorResponse
// @Router /menus/{id} [get]
func (c *menusController) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")

	menu, err := c.service.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, shareds.ErrorResponse{Message: "Menu not found"})
		return
	}
	ctx.JSON(http.StatusOK, menu)
}

// Create Create a new Menu
// @Summary Create a new Menu
// @Tags Menus
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param input body MenusRequest true "Menu Data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} shareds.ErrorResponse
// @Failure 500 {object} shareds.ErrorResponse
// @Router /menus [post]
func (c *menusController) Create(ctx *gin.Context) {
	var input MenusRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, shareds.ErrorResponse{Message: "Invalid input data"})
		return
	}

	createdID, err := c.service.Create(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, shareds.ErrorResponse{Message: "Error creating menu"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": createdID, "message": "Menu created successfully"})
}

// Update Update an existing Menu
// @Summary Update an existing Menu
// @Tags Menus
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID of the Menu"
// @Param input body MenusRequest true "Updated Menu Data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} shareds.ErrorResponse
// @Failure 500 {object} shareds.ErrorResponse
// @Router /menus/{id} [put]
func (c *menusController) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	var input MenusRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, shareds.ErrorResponse{Message: "Invalid input data"})
		return
	}

	err := c.service.Update(id, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, shareds.ErrorResponse{Message: "Error updating menu"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Menu updated successfully"})
}

// Delete Delete Menu by ID
// @Summary Delete Menu by ID
// @Tags Menus
// @Param id path string true "ID of the Menu"
// @Success 204
// @Security BearerAuth
// @Failure 400 {object} shareds.ErrorResponse
// @Failure 404 {object} shareds.ErrorResponse
// @Router /menus/{id} [delete]
func (c *menusController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, shareds.ErrorResponse{Message: "Menu not found"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

// Paginate Paginate Menus
// @Summary Paginate Menus
// @Tags Menus
// @Produce json
// @Security BearerAuth
// @Param page query int true "Page Number"
// @Param size query int true "Page Size"
// @Success 200 {array} MenusResponse
// @Failure 400 {object} shareds.ErrorResponse
// @Failure 500 {object} shareds.ErrorResponse
// @Router /menus/paginate [get]
func (c *menusController) Paginate(ctx *gin.Context) {
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

	menus, err := c.service.Paginate(page, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, shareds.ErrorResponse{Message: "Error fetching paginated menus"})
		return
	}

	ctx.JSON(http.StatusOK, menus)
}

// Get Menus by User Logged In
// @Summary Get Menus by User Logged in
// @Tags Menus
// @Produce json
// @Security BearerAuth
// @Success 200 {array} MenusResponse
// @Failure 400 {object} shareds.ErrorResponse
// @Failure 500 {object} shareds.ErrorResponse
// @Router /menus/user [get]
func (c *menusController) GetMenusByUserID(ctx *gin.Context) {

	userid, exists := ctx.Get("ID")
	if !exists {
		ctx.JSON(http.StatusBadRequest, shareds.ErrorResponse{Message: "Invalid user"})
		return
	}

	menus, err := c.service.GetMenusByUserID(userid.(string))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, shareds.ErrorResponse{Message: "Error fetching menus by user"})
		return
	}

	ctx.JSON(http.StatusOK, menus)
}
