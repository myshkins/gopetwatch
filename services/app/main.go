package main

import (
	"os"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"

	"github.com/myshkins/gopetwatch/database"
)

func init() {
	log.SetOutput(os.Stdout)

	file, err := os.OpenFile("gopetwatch.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

	log.SetLevel(log.InfoLevel)
}

func main() {
	r := gin.Default()
  		
	
	log.Info("Starting server...")
	err := http.ListenAndServe(":8081", r)
	if err != nil {
		log.Fatal("Server failed to start", err)
	}
}

func loadDatabase() {
	database.Connect()
}
