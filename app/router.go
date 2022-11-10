package app

import (
	"myproject/api"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router() {
	//設定環境模式 不知道debug跳出來的
	gin.SetMode(gin.ReleaseMode)
	//Default已包含兩個middleware logger() & recovery()
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"}, // 允許的請求方法
		AllowHeaders:     []string{"Origin", "Content-Length", "signature", "Content-type,"},
		AllowAllOrigins:  true,
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))
	//當客戶端執行post方法的請求時會執行後面的函數
	router.POST("/check", api.Check)
	router.POST("/transfer", api.Transfer)
	router.POST("/bet", api.HistoryBet)
	router.POST("/summary", api.HistorySummary)
	router.POST("/report/bet", api.ReportBet)
	router.POST("/report/summary", api.ReportSummary)
	router.POST("/betrecord", api.Betrecord)
	//分組後該成員訪問為http://127.0.0.1:5000/player
	tp := router.Group("player")
	tp.POST("/create", api.CreatePlayer)
	tp.POST("/login", api.Login)
	tp.POST("/deposit", api.Deposit)
	tp.POST("/withdraw", api.Withdraw)
	tp.POST("/balance", api.Balance)
	tp.POST("/logout", api.Logout)
	//錯誤提示
	router.NoRoute(func(c *gin.Context) {
		c.JSON(400, gin.H{"error": "Bad Request"})
	})
	//執行後啟動在port接口5000
	router.Run(":5000")
}
