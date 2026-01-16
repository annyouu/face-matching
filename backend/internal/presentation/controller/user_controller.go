package controller

import (
	"errors"
	"log"
	"net/http"

	"destinyface/internal/contextkey"
	appErrors "destinyface/internal/errors"
	"destinyface/internal/usecase"
	"destinyface/internal/usecase/dto"

	"github.com/gin-gonic/gin"
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
	
	log.Printf("[Handler Error] Detail: %+v", err)
    switch {
    case errors.Is(err, appErrors.ErrNotFound):
        c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
    case errors.Is(err, appErrors.ErrEmailAlreadyExists):
        c.JSON(http.StatusConflict, gin.H{"error": "Email already registered"})
    case errors.Is(err, appErrors.ErrInvalidCredentials):
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
    default:
        c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
			"debug_message": err.Error(),
		})
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

// PATCH /users/setup/name
func (h *UserHandler) SetupName(c *gin.Context) {
	userID := c.GetString(contextkey.UserID)
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	var input dto.UserSetupNameInput
	if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
        return
    }

	// UseCaseの呼び出し
	output, err := h.userUseCase.SetupName(c.Request.Context(), userID, &input)
	if err != nil {
		respondError(c, err)
		return
	}

	c.JSON(http.StatusOK, output)
}

// PATCH /users/setup/image
func (h *UserHandler) SetupImage(c *gin.Context) {
	userID := c.GetString(contextkey.UserID)
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	var input dto.UserSetupImageInput
	if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
        return
    }

	// UseCaseの呼び出し
	output, err := h.userUseCase.SetupImage(c.Request.Context(), userID, &input)
	if err != nil {
		respondError(c, err)
		return
	}

	c.JSON(http.StatusOK, output)
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
		// respondError(c, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tokenOutput)
}

// GET
func (h *UserHandler) GetProfile(c *gin.Context) {
	// Redisに保存されたuserIDを取得する
	userID := c.GetString(contextkey.UserID)

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
	userID := c.GetString(contextkey.UserID)

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