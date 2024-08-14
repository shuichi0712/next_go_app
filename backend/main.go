package main

import (
	"github.com/gin-gonic/gin"
	"context"
   "log"
	"next_go_blog/ent"
	_ "github.com/lib/pq"
    "os"
"next_go_blog/ent/user"
    "github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
    	log.Fatal("Error loading .env file")
	}
	godotenv.Load(".env")
	//Ginフレームワークのデフォルトの設定を使用してルータを作成
	router := gin.Default()
	//PostgreSQLに接続
   client, err := ent.Open("postgres", "host=db port=5432 user=postgres dbname=db password=password sslmode=disable")
	localhost := os.Getenv("LOCAL_URL")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", localhost)
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
	router.POST("users/sign_in",func(c *gin.Context){

		// ログインで送られてくるリクエストを型定義
		type SignInRequest struct{
			Email string `json:"email" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		// 変数reqをSignInRequestで定義
		var req SignInRequest

		//reqに取得したデータを格納、変換でエラーが起きた場合はエラーを返して終了
		if err :=c.ShouldBindJSON(&req); err != nil{
			c.JSON(400, gin.H{"error": "Invalid request"})
			return
		}

		// ユーザの検索を行う
		sign_in_user, err := client.User.Query().
			Where(user.EmailEQ(req.Email), user.PasswordEQ(req.Password)).
			First(context.Background())
		
		//エラーを返す
		if err != nil {
			c.JSON(401, gin.H{"error": "invalid credentials"})
			return
		}

		// ログイン成功
		c.JSON(200, gin.H{"user": sign_in_user})

	})
	router.POST("/studytimes", func(c *gin.Context) {
		type StudyTimeRequest struct {
			UserId int `json:"user_id" binding:"required"`
			Title string `json:"title" binding:"required"`
			Hour int `json:"hour" binding:"required"`
			Minite int `json:"minite" binding:"required"`
		}
		var req StudyTimeRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request"})
			return
		}
		newStudyTime, err := client.StudyTime.Create().SetUserID(req.UserId).SetTitle(req.Title).SetHour(req.Hour).SetMinite(req.Minite).Save(context.Background())
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, gin.H{"study_time": newStudyTime})
	})

	// サーバー起動
	router.Run(":8080")
}
