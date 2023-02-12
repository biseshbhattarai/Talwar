package models

import (
	"fmt"
	"time"

	database "github.com/biseshbhattarai/talwar/db"
	"github.com/biseshbhattarai/talwar/utils"
)

type Port struct {
	ID          int64 `gorm:"primaryKey"`
	PortNumber  string
	SubDomain   SubDomain
	SubDomainID int64 `gorm:"foreignkey:ID"`
}

// scan port by listing subdomains and then scanning them
func ExtractPortsFromSubdomains(subdomains []SubDomain) {

	for _, subdomain := range subdomains {
		go func(subdomain SubDomain) {
			parsedPorts := utils.FindPortAndProtocal(subdomain.Name)
			StorePorts(parsedPorts, int(subdomain.ID))

		}(subdomain)
	}
}

func StartPortScan(targetID int) {
	subdomains, err := ListSubDomains(targetID)
	if err != nil {
		fmt.Println("could not get targets", err)
	}
	ExtractPortsFromSubdomains(subdomains)
	ScanHistory := ScanHistory{
		TargetID:  int64(targetID),
		ScanType:  "port",
		ScannedAt: time.Now().UTC(),
	}
	ScanHistory.Save()
}

// store ports in database
func StorePorts(ports []string, subdomainID int) {
	for _, port := range ports {
		if port != "" {

			port := Port{
				PortNumber:  port,
				SubDomainID: int64(subdomainID),
			}
			database.DbConn.Create(&port)
		}
	}
}

// list ports by target id
func ListPorts(targetID int) ([]Port, error) {
	var ports []Port
	// list ports from target id
	err := database.DbConn.Table("ports").Select("ports.port_number").Joins("INNER JOIN sub_domains ON sub_domains.id = ports.sub_domain_id").Joins("INNER JOIN domains ON domains.id=sub_domains.domain_id").Where("domains.target_id = ?", targetID).Scan(&ports).Error

	if err != nil {
		return []Port{}, err
	}
	return ports, nil
}
