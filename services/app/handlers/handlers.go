package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)


func homeHandler(w http.ResponseWriter, _ *http.Request) {
	snippet := renderChart()
	// resultQueryReadings := queryReadings(db)
	fillTemplate(w, snippet)
	
}

func postTempHandler(c *gin.Context) {
	
}
