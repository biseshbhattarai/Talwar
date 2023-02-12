package controllers

import (
	"net/http"
	"strconv"

	"github.com/biseshbhattarai/talwar/models"
	"github.com/gin-gonic/gin"
)

func StartTruffleScan(context *gin.Context) {
	targetID := context.Param("target_id")
	parseIntTargetId, _ := strconv.Atoi(targetID)
	models.StartTruffleScan(parseIntTargetId)
	context.JSON(http.StatusOK, gin.H{"domains": "Completed"})
}

func ListTruffleScan(context *gin.Context) {
	targetID := context.Param("target_id")
	parseIntTargetId, _ := strconv.Atoi(targetID)
	truffleScan, err := models.ListGitScanResults(parseIntTargetId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"truffleScan": truffleScan})
}
