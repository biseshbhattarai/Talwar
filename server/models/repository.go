package models

import (
	database "github.com/biseshbhattarai/talwar/db"
	"github.com/google/go-github/github"
)

type Repository struct {
	ID             int64 `gorm:"primaryKey"`
	RepoName       string
	RepoUrl        string
	OrganizationID int `gorm:"foreignkey:ID"`
	Organization   Organization
}

func (repository *Repository) Save() (*Repository, error) {
	err := database.DbConn.Create(&repository).Error
	if err != nil {
		return &Repository{}, err
	}
	return repository, nil
}

func AddRepositoriesToOrganization(repositories []*github.Repository, organizationID int) {
	var repository Repository
	for _, githubRepository := range repositories {
		repository = Repository{
			RepoName:       *githubRepository.Name,
			RepoUrl:        *githubRepository.HTMLURL,
			OrganizationID: organizationID,
		}
		repository.Save()
	}
}

func ListRepositoryByOrganizationID(organizationID int) ([]Repository, error) {
	var repositories []Repository
	err := database.DbConn.Where("organization_id = ?", organizationID).Find(&repositories).Error
	if err != nil {
		return []Repository{}, err
	}
	return repositories, nil
}

func ListOrganizationReposByTargetId(targetId int) ([]Repository, error) {
	var repositories []Repository
	err := database.DbConn.Where("organization_id in (select id from organizations where target_id = ?)", targetId).Find(&repositories).Error
	if err != nil {
		return []Repository{}, err
	}
	return repositories, nil
}
