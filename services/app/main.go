package main

import (
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"

	"github.com/myshkins/gopetwatch/database"
	"github.com/myshkins/gopetwatch/logger"
	"github.com/myshkins/gopetwatch/handlers"
)

func init() {
	logger.InitLogger()
	logger.Log.SetOutput(os.Stdout)

	file, err := os.OpenFile("gopetwatch.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		logger.Log.SetOutput(file)
	} else {
		logger.Log.Info("Failed to logger.Log.to file, using default stderr")
	}
  
	logger.Log.SetLevel(logrus.DebugLevel)
	loadDatabase()
}

func main() {
	r := gin.Default()
  r.LoadHTMLFiles("templates/index.html")  		

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		os.Getenv("GIN_USER_ID"): os.Getenv("GIN_PW"),
		}))

	r.GET("/", handlers.HomeHandler)
	authorized.GET("/test", handlers.TestHandler)
	authorized.POST("/post", handlers.PostTempHandler)
	
	r.Run(":8081")

	logger.Log.Info("Starting server...")
}

func loadDatabase() {
	database.Connect()
	database.CreateTable()
	database.SeedDatabase()
}
//
