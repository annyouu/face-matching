package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"destinyface/internal/domain/entity"
	"destinyface/internal/usecase"
	"destinyface/internal/usecase/dto"
)

type UserHandler struct {
	userUseCase usecase.UserUseCaseInterface
}

func NewUserHandler(u usecase.UserUseCaseInterface) *UserHandler {
	return &UserHandler{
		userUseCase: u,
	}
}

// POST
func (h *UserHandler) Register(c *gin.Context) {
	var input dto.UserRegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request format",
		})
		return
	}
	
	output, err := h.userUseCase.Register(c.Request.Context(), &input)
	if err != nil {
		
	}

}

// POST
func (h *UserHandler) Login() {

}

// GET
func (h *UserHandler) GetProfile()

// PATCH
func (h *UserHandler) UpdateProfile()