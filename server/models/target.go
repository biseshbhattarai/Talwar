package models

import (
	database "github.com/biseshbhattarai/talwar/db"
)

type Target struct {
	ID        int64 `gorm:"primaryKey"`
	Name      string
	IsScope   bool
	Domains   []Domain
	IsScanned bool
}

func (target *Target) Save() (*Target, error) {
	err := database.DbConn.Create(&target).Error
	if err != nil {
		return &Target{}, err
	}
	return target, nil
}

func ListTargets() ([]Target, error) {
	//include domains
	var targets []Target
	err := database.DbConn.Find(&targets).Error
	if err != nil {
		return []Target{}, err
	}
	return targets, nil
}

func ListTargetsAndDomains() ([]Target, error) {
	var targets []Target
	err := database.DbConn.Preload("Domains").Find(&targets).Error
	if err != nil {
		return []Target{}, err
	}
	return targets, nil
}
