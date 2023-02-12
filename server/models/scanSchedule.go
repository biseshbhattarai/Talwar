package models

import (
	"time"

	database "github.com/biseshbhattarai/talwar/db"
)

type ScanSchedule struct {
	ID           int64 `gorm:"primaryKey"`
	ScanType     ScanType
	ScanStart    time.Time
	ScanInterval int
	TargetID     int64 `gorm:"foreignkey:ID"`
	Target       Target
}

func (scanSchedule *ScanSchedule) Save() (*ScanSchedule, error) {
	err := database.DbConn.Create(&scanSchedule).Error
	if err != nil {
		return &ScanSchedule{}, err
	}
	return scanSchedule, nil
}

func ListScanSchedule() ([]ScanSchedule, error) {
	var scanSchedule []ScanSchedule
	//list targets with their scan schedules using inner join
	err := database.DbConn.Preload("Target").Find(&scanSchedule).Error
	if err != nil {
		panic(err)
	}
	return scanSchedule, nil
}
