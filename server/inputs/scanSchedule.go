package inputs

import (
	"github.com/biseshbhattarai/talwar/models"
)

type ScanScheduleInput struct {
	ScanType     models.ScanType `json:"scanType" binding:"required"`
	ScanStart    string          `json:"scanStart" binding:"required"`
	ScanInterval int             `json:"scanInterval" binding:"required"`
}
