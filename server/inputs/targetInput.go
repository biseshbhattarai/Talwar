package inputs

import "github.com/biseshbhattarai/talwar/models"

type TargetInput struct {
	Name      string          `json:"name" binding:"required"`
	IsScope   bool            `json:"isScope" binding:"required"`
	IsScanned bool            `json:"isScanned" binding:"required"`
	Domains   []models.Domain `json:"domains" binding:"required"`
}
