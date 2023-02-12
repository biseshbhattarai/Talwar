package controllers

import (
	"net/http"
	"strconv"

	"github.com/biseshbhattarai/talwar/models"
	"github.com/gin-gonic/gin"
)

func StartScan(context *gin.Context) {
	targetID := context.Param("target_id")
	parseIntTargetId, _ := strconv.Atoi(targetID)
	models.StartScan(parseIntTargetId)
	context.JSON(http.StatusOK, gin.H{"domains": "Completed"})
}

func ListSubDomainsByTargetID(context *gin.Context) {
	targetID := context.Param("target_id")
	parseIntTargetId, _ := strconv.Atoi(targetID)
	subDomains, err := models.ListSubDomains(parseIntTargetId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"subdomains": subDomains})
}
