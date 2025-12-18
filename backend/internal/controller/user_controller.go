package controller

import (
	"errors"
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

func respondError(c *gin.Context, err error) {
    switch {
    case errors.Is(err, appErrors.ErrNotFound):
        c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
    case errors.Is(err, appErrors.ErrEmailAlreadyExists):
        c.JSON(http.StatusConflict, gin.H{"error": "Email already registered"})
    case errors.Is(err, appErrors.ErrInvalidCredentials):
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
    default:
        // 内部ログ出力を推奨
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
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
		respondError(c, err)
		return
	}
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
		respondError(c, err)
		return
	}
	c.JSON(http.StatusOK, tokenOutput)
}

// GET
func (h *UserHandler) GetProfile(c *gin.Context) {
	// 本来はMiddlewareでセットされたuserIDを取得する
	userID := c.GetString("userID")

	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	output, err := h.userUseCase.GetProfile(c.Request.Context(), userID)
	if err != nil {
		respondError(c, err)
		return
	}
	c.JSON(http.StatusOK, output)
}

// PATCH
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	userID := c.GetString("userID")

	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	var input dto.UserUpdateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request format",
		})
		return
	}

	output, err := h.userUseCase.UpdateProfile(c.Request.Context(), userID, &input)
	if err != nil {
		respondError(c, err)
		return
	}
	c.JSON(http.StatusOK, output)
}