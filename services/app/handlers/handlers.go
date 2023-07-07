package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	renderer "github.com/myshkins/gopetwatch/renderer"
)

func homeHandler(w http.ResponseWriter, _ *http.Request) {
	snippet := renderer.RenderChart()
	// resultQueryReadings := queryReadings(db)
	renderer.FillTemplate(w, snippet)
	
}

func postTempHandler(c *gin.Context) {
	
}

