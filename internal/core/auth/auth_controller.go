package auth

import (
	"bernardtm/backend/internal/core/shareds"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(c *gin.Context)
	Login2Step(c *gin.Context)
	RequestPasswordReset(c *gin.Context)
	ResetPassword(c *gin.Context)
}

type authController struct {
	service AuthService
}

func NewAuthController(service AuthService) *authController {
	return &authController{service: service}
}

// Login autenticate a user
// @Summary Login a user
// @Description Authenticate the user
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param login body LoginRequest true "Login info"
// @Success 200 {object} TokenResponse
// @Failure 400 {object} shareds.ErrorResponse
// @Router /auth/login [post]
func (uc *authController) Login(c *gin.Context) {
	var loginRequest LoginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, shareds.ErrorResponse{Message: "Invalid data"})
		return
	}

	// auth logic
	token, err := uc.service.Login(loginRequest.Email, loginRequest.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, shareds.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, TokenResponse{Token: token})
}

// Validate 2FA code
// @Summary Validate login 2FA code
// @Description Validate login 2FA code and generate a JWT token for API access
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer token" default(Bearer <token>)
// @Param login body Login2StepRequest true "2-Step Login info"
// @Success 200 {object} TokenResponse
// @Failure 400 {object} shareds.ErrorResponse
// @Router /auth/login/verify [post]
func (uc *authController) Login2Step(c *gin.Context) {
	var loginRequest Login2StepRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, shareds.ErrorResponse{Message: "Invalid data"})
		return
	}

	twoFactorCodeID, exists := c.Get("ID")
	if !exists {
		c.JSON(http.StatusBadRequest, shareds.ErrorResponse{Message: "Invalid user"})
		return
	}

	// validate 2fa code and generate jwt token
	loginResponse, err := uc.service.Login2Step(twoFactorCodeID.(string), loginRequest.OTP)
	if err != nil {
		c.JSON(http.StatusBadRequest, shareds.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, loginResponse)
}

// Request Password Reset
// @Summary Request password reset
// @Description Request password reset and send a verification code
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param login body RecoverPasswordRequest true "Password request info"
// @Success 200 {object} shareds.MessageResponse
// @Failure 400 {object} shareds.ErrorResponse
// @Router /auth/login/request-password-reset [post]
func (uc *authController) RequestPasswordReset(c *gin.Context) {
	var requestData RecoverPasswordRequest

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, shareds.ErrorResponse{Message: "Invalid data"})
		return
	}
	// send email with recovery link and verification code
	err := uc.service.RequestPasswordReset(requestData)
	if err != nil {
		c.JSON(http.StatusBadRequest, shareds.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, shareds.MessageResponse{Message: "Password reset request sent successfully"})
}

// Reset Password
// @Summary Reset password
// @Description Reset password after validation, and send a email informing the password has changed
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer token" default(Bearer <token>)
// @Param login body PasswordResetRequest true "Password reset info"
// @Success 200 {object} shareds.MessageResponse
// @Failure 400 {array} shareds.ErrorResponse
// @Router /auth/login/reset-password [post]
func (uc *authController) ResetPassword(c *gin.Context) {
	var requestData PasswordResetRequest

	// validar token no middleware
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, shareds.ErrorResponse{Message: "Invalid data"})
		return
	}

	id, exists := c.Get("ID")
	if !exists {
		c.JSON(http.StatusBadRequest, shareds.ErrorResponse{Message: "Invalid user"})
		return
	}

	userUUID, ok := id.(string)
	if ok {
		errorsList := uc.service.ResetPassword(userUUID, requestData)
		if errorsList != nil {
			var errorResponses []shareds.ErrorResponse
			for _, err := range errorsList {
				errorResponses = append(errorResponses, shareds.ErrorResponse{Message: err.Error()})
			}
			c.JSON(http.StatusBadRequest, errorResponses)
			return
		}

		c.JSON(http.StatusOK, shareds.MessageResponse{Message: "Password reset successfully"})
	} else {
		c.JSON(http.StatusBadRequest, shareds.MessageResponse{Message: "Password reset failed"})
	}

}
