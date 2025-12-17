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
func (h *UserHandler) Register() {
	
}

// POST

// GET

// PATCH