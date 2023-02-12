package migrations

import (
	database "github.com/biseshbhattarai/talwar/db"
	"github.com/biseshbhattarai/talwar/models"
)

func Migrate() {
	database.DbConn.AutoMigrate(&models.Target{}, &models.Domain{}, &models.SubDomain{}, &models.Organization{}, &models.Repository{}, &models.ScanSchedule{}, &models.ScanHistory{}, &models.Port{}, &models.Nuclei{}, &models.Git{}, &models.Email{})
}
