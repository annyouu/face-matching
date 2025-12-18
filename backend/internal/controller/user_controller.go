package controller

import (
	stdErrors "errors"
	"net/http"

	"github.com/gin-gonic/gin"
	appErrors "destinyface/internal/errors"
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
		if stdErrors.Is(err, appErrors.ErrEmailAlreadyExists) {
			c.JSON(http.StatusConflict, gin.H{
				"error": "Email already registered",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}
	// 201
	c.JSON(http.StatusCreated, output)

}

// POST
func (h *UserHandler) Login(c *gin.Context) {
	var input dto.UserLoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request format",
		})
		return
	}

	tokenOutput, err := h.userUseCase.Login(c.Request.Context(), &input)
	if err != nil {
		if stdErrors.Is(err, appErrors.ErrEmailAlreadyExists) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid email or password",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}
	// 201
	c.JSON(http.StatusOK, tokenOutput)
}

// GET
func (h *UserHandler) GetProfile(c *gin.Context) {
	// 本来はMiddlewareでセットされたuserIDを取得する
}

// PATCH
func (h *UserHandler) UpdateProfile(c *gin.Context) {

}