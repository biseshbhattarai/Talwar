package controllers

import (
	"net/http"
	"strconv"

	"github.com/biseshbhattarai/talwar/models"
	"github.com/gin-gonic/gin"
)

//start port scan
func StartPortScan(context *gin.Context) {
	targetID := context.Param("target_id")
	parseIntTargetId, _ := strconv.Atoi(targetID)
	models.StartPortScan(parseIntTargetId)
	context.JSON(http.StatusOK, gin.H{"ports": "Completed"})
}

//list port scan
func ListPortScan(context *gin.Context) {
	targetID := context.Param("target_id")
	parseIntTargetId, _ := strconv.Atoi(targetID)
	portScan, err := models.ListPorts(parseIntTargetId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"portscan": portScan})
}
