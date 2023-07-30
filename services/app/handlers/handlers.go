package handlers

import (
	"net/http"
  // "encoding/json"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/myshkins/gopetwatch/database"
	"github.com/myshkins/gopetwatch/logger"
	renderer "github.com/myshkins/gopetwatch/renderer"
)

func HomeHandler(c *gin.Context) {
	s := renderer.RenderChart()
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "chanch watch",
		"snippet": s,
		})
	// snippet := renderer.RenderChart()
	// resultQueryReadings := queryReadings(db)
}

func PostTempHandler(c *gin.Context) {
	var tempPoint database.Reading
	if err := c.ShouldBindJSON(&tempPoint); err != nil {
		logger.Log.Warn("shouldBind error: ", err)
	} else {
		logger.Log.Info("shouldBind success?:")
		logger.Log.Info(tempPoint.ReadingTimestamp)
		logger.Log.Infof("temp: %f, time: %d", tempPoint.Temperature, &tempPoint.ReadingTimestamp)
	}

	_, err := database.InsertRow(tempPoint)
	if err != nil {
    logger.Log.Warn(err)
	}
}
// hello
