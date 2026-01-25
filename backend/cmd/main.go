package main

import (
	"log"
	"os"
	"time"

	"destinyface/internal/infrastructure/persistence"
	"destinyface/internal/infrastructure/redis"
	"destinyface/internal/presentation/controller"
	"destinyface/internal/presentation/middleware"

	// "destinyface/internal/presentation/middleware"
	"destinyface/internal/usecase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	goredis "github.com/redis/go-redis/v9"
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

	// Redisクライアントの初期化を追加
	rdb := goredis.NewClient(&goredis.Options{
        Addr:     "redis:6379", 
        Password: "",
        DB: 0,
    })
	
	sessionRepo := redis.NewSessionRepository(rdb)

	// 4. 各層の依存注入 (DI)
	userRepo := persistence.NewUserRepository(db) 
	userUseCase := usecase.NewUserUseCase(userRepo, sessionRepo)
	userHandler := controller.NewUserHandler(userUseCase)

	// 5. サーバー設定 (Gin)
	r := gin.Default()

	// CORS設定
	r.Use(cors.New(cors.Config{
		// フロントエンドのURLを許可
        AllowOrigins: []string{"http://localhost:3000"},
        // 許可するHTTPメソッド
        AllowMethods: []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
        // 許可するヘッダー
        AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
        // Cookieなどの認証情報を許可
        AllowCredentials: true,
        // OPTIONSリクエストの結果をキャッシュする時間
        MaxAge: 12 * time.Hour,
	}))

	// --- ルーティング ---
	
	// A. 認証不要ルート
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", userHandler.Register)
		authGroup.POST("/login", userHandler.Login)
	}

	// B. 認証必須ルート (ミドルウェアを適用)
	userGroup := r.Group("/users")

	// jwtServiceではなく、作成したsessionRepoを渡すように変更する
	userGroup.Use(middleware.UserAuthentication(sessionRepo))
	{
		userGroup.GET("/me", userHandler.GetProfile)
		userGroup.PATCH("/me", userHandler.UpdateProfile)

		// 追加
		userGroup.PATCH("/setup/name", userHandler.SetupName)
		userGroup.PATCH("/setup/image", userHandler.SetupImage)
	}

	// 6. 起動
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("🚀 Server started on :%s", port)
	r.Run(":" + port)
}