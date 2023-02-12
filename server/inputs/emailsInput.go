package inputs

import "github.com/biseshbhattarai/talwar/models"

type EmailInput struct {
	Email    string          `json:"email" binding:"required"`
	ScanType models.ScanType `json:"scanType" binding:"required"`
}
