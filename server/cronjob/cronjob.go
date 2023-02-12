package cronjob

import (
	"fmt"

	"github.com/biseshbhattarai/talwar/models"
	"github.com/robfig/cron"
)

func RunScheduler() (*cron.Cron, error) {
	c := cron.New()
	schedule, _ := models.ListScanSchedule()
	for s, _ := range schedule {
		scheduleInterval := fmt.Sprintf("0 */%d * * * *", schedule[s].ScanInterval)
		c.AddFunc(scheduleInterval, func() {
			fmt.Println("Running scheduled scan for target: ", schedule[s].TargetID)
			if schedule[s].ScanType == models.Subdomain {
				fmt.Println("Running subdomain scan for domain: ", schedule[s].Target)
				models.StartScan(int(schedule[s].TargetID))
			}
			if schedule[s].ScanType == models.GithubScan {
				fmt.Println("Running github scan for target: ", schedule[s].TargetID)
				models.StartTruffleScan(int(schedule[s].TargetID))
			}
			if schedule[s].ScanType == models.PortScan {
				fmt.Println("Running nuclei scan for target: ", schedule[s].TargetID)
				models.StartPortScan(int(schedule[s].TargetID))
			}
			if schedule[s].ScanType == models.NucleiScan {
				fmt.Println("Running nuclei scan for target: ", schedule[s].TargetID)
				models.StartNucleiScan(int(schedule[s].TargetID))
			}
		})
	}
	return c, nil
}
