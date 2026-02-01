package controller

import (
	"destinyface/internal/contextkey"
	"destinyface/internal/usecase"
	"net/http"
	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	profileUseCase usecase.ProfileUseCaseInterface
}

func NewProfileHandler(p usecase.ProfileUseCaseInterface) *ProfileHandler {
	return &ProfileHandler{
		profileUseCase: p,
	}
}

// 画像を保存する
func (h *ProfileHandler) SetupImage(c *gin.Context) {
	// 認証済みユーザーIDをコンテキストから取得する
	userID := c.GetString(contextkey.UserID)
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	// マルチパートフォームから画像ファイルを取得
	fileHeader, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "画像ファイルが必要です",
		})
		return
	}

	// UseCaseに渡すためにio.Readerとして開く
	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "ファイルの読み込みに失敗しました",
		})
		return
	}
	defer file.Close()

	// UseCaseの実行をする
	output, err := h.profileUseCase.SetupProfileImage(c.Request.Context() ,userID, file)
	if err != nil {
		respondError(c, err)
		return
	}

	// レスポンスを返す
	c.JSON(http.StatusOK, output)

}