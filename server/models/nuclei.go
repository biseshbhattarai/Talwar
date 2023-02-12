package models

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
	"time"

	database "github.com/biseshbhattarai/talwar/db"
)

type Nuclei struct {
	Id          int64 `gorm:"primaryKey"`
	TargetID    int64 `gorm:"foreignkey:ID"`
	Target      Target
	Host        string
	Name        string
	Description string
	Severity    string
}

type Result struct {
	Template         string   `json:"template"`
	TemplateURL      string   `json:"template-url"`
	TemplateID       string   `json:"template-id"`
	Info             Info     `json:"info"`
	Type             string   `json:"type"`
	Host             string   `json:"host"`
	MatchedAt        string   `json:"matched-at"`
	ExtractedResults []string `json:"extracted-results"`
	Timestamp        string   `json:"timestamp"`
	MatcherStatus    bool     `json:"matcher-status"`
	MatchedLine      *string  `json:"matched-line"`
}

type Info struct {
	Name           string         `json:"name"`
	Author         []string       `json:"author"`
	Tags           []string       `json:"tags"`
	Description    string         `json:"description"`
	Reference      []string       `json:"reference"`
	Severity       string         `json:"severity"`
	Classification Classification `json:"classification"`
}

type Classification struct {
	CveID *string  `json:"cve-id"`
	CweID []string `json:"cwe-id"`
}

func NucleiScanner(domain string, targetID int) {
	fmt.Println(domain)
	cmd := exec.Command("nuclei", "-u", domain, "-silent", "-json")
	fmt.Println(cmd)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("could not run command", err)
	}
	nucleiScan := string(out)
	parsedNucleiScan := strings.Split(nucleiScan, "\n")

	var result Result
	err = json.Unmarshal([]byte(parsedNucleiScan[0]), &result)
	if err != nil {
		fmt.Println(err)
	}
	if result.Info.Name != "" && result.Info.Severity != "" {
		nuclei := Nuclei{
			Host:        result.Host,
			Name:        result.Info.Name,
			Description: result.Info.Description,
			Severity:    result.Info.Severity,
			TargetID:    int64(targetID),
		}
		database.DbConn.Create(&nuclei)
	}

	fmt.Println(result.Info.Name, result.Info.Description, result.Info.Severity)

}

func ExtractFromNuclei(subdomains []SubDomain, targetID int) {
	for _, domain := range subdomains {
		go func(subdomain SubDomain) {
			if subdomain.Name != "" {
				NucleiScanner(subdomain.Name, targetID)
			}
		}(domain)
	}
}

func StartNucleiScan(targetID int) {
	subdomains, err := ListSubDomains(targetID)
	if err != nil {
		fmt.Println("could not get targets", err)
	}
	ExtractFromNuclei(subdomains, targetID)
	ScanHistory := ScanHistory{
		TargetID:  int64(targetID),
		ScanType:  "nuclei",
		ScannedAt: time.Now().UTC(),
	}
	ScanHistory.Save()
}

func ListNuclei(targetID int) ([]Nuclei, error) {
	var nuclei []Nuclei
	err := database.DbConn.Where("target_id = ?", targetID).Find(&nuclei).Error
	if err != nil {
		return nil, err
	}
	return nuclei, nil
}
