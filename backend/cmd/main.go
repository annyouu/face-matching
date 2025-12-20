package main

import (
	"log"
	"os"

	"destinyface/internal/presentation/controller"
	"destinyface/internal/infrastructure/auth"
	"destinyface/internal/infrastructure/persistence"
	"destinyface/internal/presentation/middleware"
	"destinyface/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 1. 環境変数の読み込み
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	// 2. DB接続
	db, err := persistence.InitDB()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer db.Close()
	log.Println("✅ Database connected")

	// 3. インフラ層（技術的道具）の準備
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "default_secret" // MVP開発用。本番では必ず設定する
	}
	jwtService := auth.NewJWTService(jwtSecret)

	// 4. 各層の依存注入 (DI)
	userRepo := persistence.NewUserRepository(db) 
	userUseCase := usecase.NewUserUseCase(userRepo, jwtService)
	userHandler := controller.NewUserHandler(userUseCase)

	// 5. サーバー設定 (Gin)
	r := gin.Default()

	// --- ルーティング ---
	
	// A. 認証不要ルート
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", userHandler.Register)
		authGroup.POST("/login", userHandler.Login)
	}

	// B. 認証必須ルート (ミドルウェアを適用)
	userGroup := r.Group("/users")
	userGroup.Use(middleware.UserAuthentication(jwtService))
	{
		userGroup.GET("/me", userHandler.GetProfile)
		userGroup.PATCH("/me", userHandler.UpdateProfile)
	}

	// 6. 起動
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("🚀 Server started on :%s", port)
	r.Run(":" + port)
}