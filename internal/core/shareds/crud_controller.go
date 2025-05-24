package shareds

import (
	"github.com/gin-gonic/gin"
)

// CrudController defines the interface for standard CRUD operations
type CrudController interface {
	GetAll(ctx *gin.Context)   // Fetch all records
	GetByID(ctx *gin.Context)  // Fetch a single record by ID
	Paginate(ctx *gin.Context) // Fetch paginated records
	Create(ctx *gin.Context)   // Create a new record
	Update(ctx *gin.Context)   // Update an existing record
	Delete(ctx *gin.Context)   // Delete a record by ID
}
