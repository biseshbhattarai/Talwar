package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/biseshbhattarai/talwar/inputs"
	"github.com/biseshbhattarai/talwar/models"
	"github.com/gin-gonic/gin"
)

func RegisterScanSchedule(context *gin.Context) {
	var input inputs.ScanScheduleInput
	targetId := context.Param("target_id")
	parsedTargetId, _ := strconv.ParseInt(targetId, 10, 64)
	//parse scan Start time json unmarshal

	if err := context.BindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(input.ScanStart)
	var scanStart time.Time
	if err := scanStart.UnmarshalText([]byte(input.ScanStart)); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	scanSchedule := models.ScanSchedule{
		TargetID:     parsedTargetId,
		ScanType:     input.ScanType,
		ScanInterval: input.ScanInterval,
		ScanStart:    scanStart,
	}

	savedScanSchedule, err := scanSchedule.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"scanSchedule": savedScanSchedule})
}

func ListScanSchedules(context *gin.Context) {
	scanSchedules, err := models.ListScanSchedule()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"scanSchedules": scanSchedules})
}
