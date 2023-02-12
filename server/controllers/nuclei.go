package controllers

import (
	"net/http"
	"strconv"

	"github.com/biseshbhattarai/talwar/models"
	"github.com/gin-gonic/gin"
)

func StartNucleiScan(context *gin.Context) {
	targetID := context.Param("target_id")
	parseIntTargetId, _ := strconv.Atoi(targetID)
	models.StartNucleiScan(parseIntTargetId)
	context.JSON(http.StatusOK, gin.H{"domains": "Completed"})
}

func ListNucleiScan(context *gin.Context) {
	targetID := context.Param("target_id")
	parseIntTargetId, _ := strconv.Atoi(targetID)
	nucleiScan, err := models.ListNuclei(parseIntTargetId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"nucleiScan": nucleiScan})
}
