package handler

import (
	"bwa-startup/helper"
	"bwa-startup/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler  {
	return &userHandler{userService: userService}
}

func (h *userHandler) RegisterUser (c *gin.Context)  {
	var input user.RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.JsonResponse("Register failed", http.StatusUnprocessableEntity, "error",false, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}


	result, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.JsonResponse("Register failed", http.StatusBadRequest, "error",false, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(result, "_")

	response := helper.JsonResponse("Account has been registered", http.StatusOK, "success",true, formatter)

	c.JSON(http.StatusOK, response)
}