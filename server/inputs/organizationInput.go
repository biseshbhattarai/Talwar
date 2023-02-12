package inputs

import "github.com/biseshbhattarai/talwar/models"

type OrganizationRepoInput struct {
	OrgName          string                  `json:"orgName" binding:"required"`
	OrganizationType models.OrganizationType `json:"organizationType" binding:"required"`
	Repositories     []models.Repository     `json:"repoUrl"`
}
