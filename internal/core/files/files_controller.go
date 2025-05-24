package files

import (
	"bernardtm/backend/internal/core/shareds"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FilesController interface {
	Create(ctx *gin.Context)
}

type filesController struct {
	service FilesService
}

func NewFilesController(service FilesService) *filesController {
	return &filesController{
		service: service,
	}
}

// Create creates a new file
// @Summary Create a new file
// @Description A new file will be created and uploaded to the server
// @Tags Files
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param file formData file true "The file to upload"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} shareds.ErrorResponse
// @Failure 500 {object} shareds.ErrorResponse
// @Router /files [post]
func (uc *filesController) Create(c *gin.Context) {
	file, err := c.FormFile("file")

	if err != nil {
		c.JSON(http.StatusBadRequest, shareds.ErrorResponse{Message: "File is required"})
		return
	}

	fileStream, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, shareds.ErrorResponse{Message: "Failed to open file"})
		return
	}
	defer fileStream.Close()

	input := FileRequest{
		File: File{
			Name: file.Filename,
		},
	}

	createdId, link, err := uc.service.Create(input, fileStream)

	if err != nil {
		c.JSON(http.StatusInternalServerError, shareds.ErrorResponse{Message: "Error uploading file"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": createdId, "link": link, "message": "File uploaded successfully"})

}
