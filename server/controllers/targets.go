package controllers

import (
	"net/http"
	"strconv"

	"github.com/biseshbhattarai/talwar/inputs"
	"github.com/biseshbhattarai/talwar/models"
	"github.com/gin-gonic/gin"
)

func RegisterTarget(context *gin.Context) {
	var input inputs.TargetInput

	if err := context.BindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	target := models.Target{
		Name:      input.Name,
		IsScope:   input.IsScope,
		IsScanned: false,
		Domains:   input.Domains,
	}
	savedTarget, err := target.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"target": savedTarget})

}

func ListTargets(context *gin.Context) {
	targets, err := models.ListTargets()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"targets": targets})
}

func ListDomainsByTargetID(context *gin.Context) {
	targetID := context.Param("target_id")
	parseIntTargetId, _ := strconv.Atoi(targetID)
	domains, err := models.ListDomainsByTargetID(parseIntTargetId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"domains": domains})
}
