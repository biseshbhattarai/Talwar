package controllers

import (
	"net/http"
	"strconv"

	"github.com/biseshbhattarai/talwar/models"
	"github.com/gin-gonic/gin"
)

func ListScanHistory(context *gin.Context) {
	targetID := context.Param("target_id")
	parseIntTargetId, _ := strconv.Atoi(targetID)
	scanHistory, err := models.ListScanHistoryByTargetID(parseIntTargetId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"scan_history": scanHistory})
}

func ListScanHistoryByScanTypeAndTargetID(context *gin.Context) {
	targetID := context.Param("target_id")
	parseIntTargetId, _ := strconv.Atoi(targetID)
	scanType := context.Param("scan_type")
	scanHistory, err := models.ListScanHistoryByTargetIDAndScanType(parseIntTargetId, scanType)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"scan_history": scanHistory})
}
