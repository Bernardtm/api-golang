package users

import (
	"bernardtm/backend/internal/core/shareds"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UsersController interface {
	GetAll(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Paginate(ctx *gin.Context)
}
type usersController struct {
	service UsersService
}

func NewUsersController(service UsersService) *usersController {
	return &usersController{service: service}
}

// GetAll get all usuários
// @Summary Get all Users
// @Tags Users
// @Produce json
// @Security BearerAuth
// @Success 200 {array} UserResponse
// @Failure 500 {object} shareds.ErrorResponse
// @Router /users [get]
func (c *usersController) GetAll(ctx *gin.Context) {
	users, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, shareds.ErrorResponse{Message: "Error fetching users"})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

// GetByID gets a usuário by ID
// @Summary Get User by ID
// @Tags Users
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID of the User"
// @Success 200 {object} UserResponse
// @Failure 400 {object} shareds.ErrorResponse
// @Failure 404 {object} shareds.ErrorResponse
// @Router /users/{id} [get]
func (c *usersController) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := c.service.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, shareds.ErrorResponse{Message: "User not found"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// Create creates a new usuário
// @Summary Create a new User
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param input body UserRequest true "User Data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} shareds.ErrorResponse
// @Failure 500 {object} shareds.ErrorResponse
// @Router /users [post]
func (c *usersController) Create(ctx *gin.Context) {
	var input UserRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, shareds.ErrorResponse{Message: err.Error()})
		return
	}

	createdID, err := c.service.Create(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, shareds.ErrorResponse{Message: "Error creating user"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": createdID, "message": "User created successfully"})
}

// Update updates a usuário existente
// @Summary Update an existing User
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID of the User"
// @Param input body UserRequest true "Updated User Data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} shareds.ErrorResponse
// @Failure 404 {object} shareds.ErrorResponse
// @Failure 500 {object} shareds.ErrorResponse
// @Router /users/{id} [put]
func (c *usersController) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	var input UserRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, shareds.ErrorResponse{Message: "Invalid input data"})
		return
	}

	err := c.service.Update(id, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, shareds.ErrorResponse{Message: "Error updating user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// Delete deletes a usuário by ID
// @Summary Delete User by ID
// @Tags Users
// @Param id path string true "ID of the User"
// @Success 204
// @Security BearerAuth
// @Failure 400 {object} shareds.ErrorResponse
// @Failure 404 {object} shareds.ErrorResponse
// @Router /users/{id} [delete]
func (c *usersController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, shareds.ErrorResponse{Message: "User not found"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

// Paginate returns a paginated list de usuários
// @Summary Paginate Users
// @Tags Users
// @Produce json
// @Security BearerAuth
// @Param page query int true "Page Number"
// @Param size query int true "Page Size"
// @Success 200 {array} UserResponse
// @Failure 400 {object} shareds.ErrorResponse
// @Failure 500 {object} shareds.ErrorResponse
// @Router /users/paginate [get]
func (c *usersController) Paginate(ctx *gin.Context) {
	pageParam := ctx.DefaultQuery("page", "1")
	sizeParam := ctx.DefaultQuery("size", "5")

	page, err := strconv.Atoi(pageParam)
	if err != nil || page <= 0 {
		ctx.JSON(http.StatusBadRequest, shareds.ErrorResponse{Message: "Invalid page parameter"})
		return
	}

	limit, err := strconv.Atoi(sizeParam)
	if err != nil || limit <= 0 || limit > 50 {
		ctx.JSON(http.StatusBadRequest, shareds.ErrorResponse{Message: "Invalid limit parameter"})
		return
	}

	users, err := c.service.Paginate(page, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, shareds.ErrorResponse{Message: "Error fetching paginated users"})
		return
	}

	ctx.JSON(http.StatusOK, users)
}
