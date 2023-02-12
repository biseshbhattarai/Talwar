package models

import (
	"fmt"

	database "github.com/biseshbhattarai/talwar/db"
	"github.com/biseshbhattarai/talwar/utils"
)

type OrganizationType string

const (
	Github   OrganizationType = "github"
	Linkedin OrganizationType = "linkedin"
)

type Organization struct {
	ID               int64 `gorm:"primaryKey"`
	OrgName          string
	OrganizationType OrganizationType
	Repository       []Repository
	Target           Target
	TargetID         int64 `gorm:"foreignkey:ID"`
}

func (organization *Organization) Save() (*Organization, error) {
	err := database.DbConn.Create(&organization).Error
	if err != nil {
		fmt.Println(err)
		return &Organization{}, err
	}
	return organization, nil
}

func FilterOrganizationsByType(organizations []Organization, organizationType OrganizationType) []Organization {
	var filteredOrganizations []Organization
	for _, organization := range organizations {
		if organization.OrganizationType == organizationType {
			filteredOrganizations = append(filteredOrganizations, organization)
		}
	}
	return filteredOrganizations
}

func ListOrganizations() ([]Organization, error) {
	var organizations []Organization
	err := database.DbConn.Find(&organizations).Error
	if err != nil {
		return []Organization{}, err
	}
	return organizations, nil
}

func ListOrganizationsByTargetID(targetID int) ([]Organization, error) {
	var organizations []Organization
	err := database.DbConn.Where("target_id = ?", targetID).Find(&organizations).Error
	if err != nil {
		return []Organization{}, err
	}
	return organizations, nil
}

func StartSearchRepositories(targetID int) {
	organizations, err := ListOrganizationsByTargetID(targetID)
	if err != nil {
		return
	}
	filteredOrganizations := FilterOrganizationsByType(organizations, Github)
	for _, organization := range filteredOrganizations {
		go func(organization Organization) {
			repositories := utils.FindOrgRepo(organization.OrgName)
			AddRepositoriesToOrganization(repositories, int(organization.ID))
		}(organization)
	}
}
