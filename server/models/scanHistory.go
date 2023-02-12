package models

import (
	"time"

	database "github.com/biseshbhattarai/talwar/db"
	"github.com/jinzhu/gorm"
)

type ScanType string

const (
	Subdomain    ScanType = "subdomain"
	GithubScan   ScanType = "githubscan"
	LinkedinScan ScanType = "linkedinscan"
	NucleiScan   ScanType = "nucleiscan"
	PortScan     ScanType = "portscan"
)

type ScanHistory struct {
	ID        int64 `gorm:"primaryKey"`
	ScanType  ScanType
	ScannedAt time.Time
	TargetID  int64 `gorm:"foreignkey:ID"`
	Target    Target
}

func (scanHistory *ScanHistory) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ScannedAt", time.Now())
	return nil
}

func (scanHistory *ScanHistory) Save() (*ScanHistory, error) {
	err := database.DbConn.Create(&scanHistory).Error
	if err != nil {
		return &ScanHistory{}, err
	}
	return scanHistory, nil
}

func ListScanHistory() ([]ScanHistory, error) {
	var scanHistory []ScanHistory
	err := database.DbConn.Find(&scanHistory).Scan(&scanHistory).Error
	if err != nil {
		return []ScanHistory{}, err
	}
	return scanHistory, nil
}

func ListScanHistoryByTargetID(targetID int) ([]ScanHistory, error) {
	var scanHistory []ScanHistory
	err := database.DbConn.Raw("SELECT * from scan_histories WHERE target_id=?", targetID).Scan(&scanHistory).Error
	if err != nil {
		return []ScanHistory{}, err
	}
	return scanHistory, nil
}

func ListScanHistoryByScanType(scanType ScanType) ([]ScanHistory, error) {
	var scanHistory []ScanHistory
	err := database.DbConn.Where("scan_type = ?", scanType).Find(&scanHistory).Error
	if err != nil {
		return []ScanHistory{}, err
	}
	return scanHistory, nil
}

func ListScanHistoryByTargetIDAndScanType(targetID int, scanType string) ([]ScanHistory, error) {
	var scanHistory []ScanHistory
	err := database.DbConn.Where("target_id = ? AND scan_type = ?", targetID, scanType).Find(&scanHistory).Error
	if err != nil {
		return []ScanHistory{}, err
	}
	return scanHistory, nil
}

func ListScanHistoryByTargetIDAndScanTypeAndScannedAt(targetID int, scanType ScanType, scannedAt time.Time) ([]ScanHistory, error) {
	var scanHistory []ScanHistory
	err := database.DbConn.Where("target_id = ? AND scan_type = ? AND scanned_at = ?", targetID, scanType, scannedAt).Find(&scanHistory).Error
	if err != nil {
		return []ScanHistory{}, err
	}
	return scanHistory, nil
}
