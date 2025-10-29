package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Error   string      `json:"error"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

type SuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

func RespondWithError(c *gin.Context, statusCode int, message string, details interface{}) {
	c.JSON(statusCode, ErrorResponse{
		Error:   http.StatusText(statusCode),
		Message: message,
		Details: details,
	})
}

func RespondWithSuccess(c *gin.Context, statusCode int, data interface{}, message string) {
	response := SuccessResponse{
		Success: true,
		Data:    data,
	}

	if message != "" {
		response.Message = message
	}

	c.JSON(statusCode, response)
}

func RespondWithValidationError(c *gin.Context, details interface{}) {
	RespondWithError(c, http.StatusBadRequest, "Validation failed", details)
}

func RespondNotFound(c *gin.Context, message string) {
	if message == "" {
		message = "Resource not found"
	}
	RespondWithError(c, http.StatusNotFound, message, nil)
}

func RespondUnauthorized(c *gin.Context, message string) {
	if message == "" {
		message = "Unauthorized access"
	}
	RespondWithError(c, http.StatusUnauthorized, message, nil)
}

func RespondForbidden(c *gin.Context, message string) {
	if message == "" {
		message = "Access forbidden"
	}
	RespondWithError(c, http.StatusForbidden, message, nil)
}

func RespondInternalError(c *gin.Context, message string) {
	if message == "" {
		message = "Internal server error"
	}
	RespondWithError(c, http.StatusInternalServerError, message, nil)
}
