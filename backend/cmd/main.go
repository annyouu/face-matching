package main


import (
	"log"
	"github.com/joho/godotenv"
	"destinyface/internal/infrastructure/persistence"
)

func main() {
	// 1. .envファイルを読み込む
	if err := godotenv.Load(); err != nil {
		// .envファイルのパスが間違っている、またはファイルがない場合に発生
		log.Fatalf("Error loading .env file: %v. Please check file path.", err)
	}

	log.Println("Starting DB connection test...")
	
	// 2. DB接続の試行
	db, err := persistence.InitDB()
	if err != nil {
		// 接続失敗の場合、ここでプログラムが停止します
		log.Fatalf("Database connection failed: %v", err)
	}
	// 3. 成功したら接続を閉じる
	defer db.Close()
	
	log.Println("✅ PostgreSQL接続に成功しました！バックエンド開発を継続できます。")
}