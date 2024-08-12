package main

import (
	"github.com/gin-gonic/gin"
	"context"
   "log"
	"next_go_blog/ent"
	_ "github.com/lib/pq"
)

func main() {
	//Ginフレームワークのデフォルトの設定を使用してルータを作成
	router := gin.Default()
	//PostgreSQLに接続
   client, err := ent.Open("postgres", "host=db port=5432 user=postgres dbname=db password=password sslmode=disable")

	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
				c.AbortWithStatus(200)
				return
		}
		c.Next()
	})
	// ルートハンドラの定義
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
	router.POST("users/signup", func(c *gin.Context) {
		type SignUpRequest struct{
			Email string `json:"email" binding:"required"`
			Name string `json:"name" binding:"required"`
			Password string `json:"password" binding:"required"`
		}
		var req SignUpRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request"})
			return
		}
		newUser, err := client.User.Create().SetEmail(req.Email).SetName(req.Name).SetPassword(req.Password).Save(context.Background())
		// エラーの場合はエラーを返して処理終了。
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error(),"messsage":"sign up missing"})
			return
		}
		// 保存したUserの情報をレスポンスとして返す。
		c.JSON(201, gin.H{"user": newUser})
	})

	// サーバー起動
	router.Run(":8080")
}
