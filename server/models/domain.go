package models

import (
	"fmt"
	"time"

	database "github.com/biseshbhattarai/talwar/db"
	"github.com/biseshbhattarai/talwar/utils"
)

type Domain struct {
	ID        int64 `gorm:"primaryKey"`
	Name      string
	TargetID  int64 `gorm:"foreignkey:ID"`
	Target    Target
	SubDomain []SubDomain
	IsScanned bool
}

func ListDomains() ([]Domain, error) {
	var domains []Domain
	err := database.DbConn.Find(&domains).Error
	if err != nil {
		return []Domain{}, err
	}
	return domains, nil
}

func ListDomainsByTargetID(targetID int) ([]Domain, error) {
	//list domains by target id
	var domains []Domain
	err := database.DbConn.Where("target_id = ?", targetID).Find(&domains).Error
	if err != nil {
		return []Domain{}, err
	}
	return domains, nil
}

func ExtractSubdomainsFromDomains(domains []Domain) {
	for _, domain := range domains {
		go func(domain Domain) {
			if !domain.IsScanned {
				parsedSubDomains := utils.RunSubfinder(domain.Name)
				StoreSubdomains(parsedSubDomains, int(domain.ID))
				database.DbConn.Model(&domain).Update("is_scanned", true)
			}
		}(domain)
	}
	StartSendEmails("Your subdomain scan is completed")

}

func ListDomainsByTargetIDAndIsScannedFalse(targetID int) ([]Domain, error) {
	var domains []Domain
	err := database.DbConn.Where("target_id = ? AND is_scanned = ?", targetID, 0).Find(&domains).Error
	if err != nil {
		return []Domain{}, err
	}
	return domains, nil
}

func StartScan(targetID int) {
	targets, err := ListDomainsByTargetIDAndIsScannedFalse(targetID)
	if err != nil {
		fmt.Println("could not get targets", err)
	}
	ExtractSubdomainsFromDomains(targets)
	ScanHistory := ScanHistory{
		TargetID:  int64(targetID),
		ScanType:  "subdomain",
		ScannedAt: time.Now().UTC(),
	}
	savedScanHistory, _ := ScanHistory.Save()
	if err != nil {
		fmt.Println("could not save scan history", err)
	}
	fmt.Println("scan history saved", savedScanHistory)
}

func StoreSubdomains(subdomains []string, domainID int) {
	for _, subdomain := range subdomains {
		subdomain := SubDomain{
			Name:     subdomain,
			DomainID: int64(domainID),
		}
		database.DbConn.Create(&subdomain)
	}
}

func ListSubDomains(targetId int) ([]SubDomain, error) {
	var subDomains []SubDomain
	err := database.DbConn.Table("sub_domains").Select("sub_domains.name,sub_domains.id").Joins("INNER JOIN domains ON domains.id = sub_domains.domain_id").Where("domains.target_id = ?", targetId).Scan(&subDomains).Error
	if err != nil {
		fmt.Println("could not get subdomains", err)
	}
	return subDomains, err
}
