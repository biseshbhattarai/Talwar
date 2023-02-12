package models

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
	"time"

	database "github.com/biseshbhattarai/talwar/db"
)

type Git struct {
	Id         int64 `gorm:"primaryKey"`
	TargetID   int64 `gorm:"foreignkey:ID; on_delete:cascade"`
	Target     Target
	Repository string
	Commit     string
	File       string
	Email      string
	Redacted   string
}

type SourceMetadata struct {
	Data struct {
		Git struct {
			Commit     string `json:"commit"`
			File       string `json:"file"`
			Email      string `json:"email"`
			Repository string `json:"repository"`
			Timestamp  string `json:"timestamp"`
			Line       int    `json:"line"`
		} `json:"Git"`
	} `json:"Data"`
}
type StructuredData struct {
	Account string `json:"account"`
	Arn     string `json:"arn"`
	UserID  string `json:"user_id"`
}

type Data struct {
	SourceMetadata SourceMetadata `json:"SourceMetadata"`
	SourceID       int            `json:"SourceID"`
	SourceType     int            `json:"SourceType"`
	SourceName     string         `json:"SourceName"`
	DetectorType   int            `json:"DetectorType"`
	DetectorName   string         `json:"DetectorName"`
	DecoderName    string         `json:"DecoderName"`
	Verified       bool           `json:"Verified"`
	Raw            string         `json:"Raw"`
	Redacted       string         `json:"Redacted"`
	ExtraData      interface{}    `json:"ExtraData"`
	StructuredData StructuredData `json:"StructuredData"`
}

func GitScan(repoUrl string, targetID int) {
	fmt.Println(repoUrl)
	// trufflehog git https://github.com/trufflesecurity/trufflehog.git --json --only-verified
	cmd := exec.Command("trufflehog", "git", repoUrl, "--json")
	out, err := cmd.Output()
	fmt.Println(cmd)
	if err != nil {
		fmt.Println("could not run command", err)
	}
	gitScan := string(out)
	parsedGitScan := strings.Split(gitScan, "\n")
	fmt.Println(gitScan)
	fmt.Println(parsedGitScan)
	if len(parsedGitScan) != 0 {
		var gitScanData Data
		err = json.Unmarshal([]byte(parsedGitScan[0]), &gitScanData)
		if err != nil {
			fmt.Println(err)
		}
		if gitScanData.SourceMetadata.Data.Git.File != "" && gitScanData.SourceMetadata.Data.Git.Email != "" {
			git := Git{
				TargetID:   int64(targetID),
				Repository: gitScanData.SourceMetadata.Data.Git.Repository,
				Commit:     gitScanData.SourceMetadata.Data.Git.Commit,
				File:       gitScanData.SourceMetadata.Data.Git.File,
				Email:      gitScanData.SourceMetadata.Data.Git.Email,
				Redacted:   gitScanData.Redacted,
			}
			database.DbConn.Create(&git)
		}

	}

}

func ExtractFromTruffle(repository []Repository, targetID int) {
	for _, repo := range repository {
		go func(repo Repository) {
			GitScan(repo.RepoUrl, targetID)
		}(repo)
	}
}

func StartTruffleScan(targetID int) {
	repositories, err := ListOrganizationReposByTargetId(targetID)
	if err != nil {
		fmt.Println("could not get targets", err)
	}
	ExtractFromTruffle(repositories, targetID)
	ScanHistory := ScanHistory{
		TargetID:  int64(targetID),
		ScanType:  "githubscan",
		ScannedAt: time.Now().UTC(),
	}
	ScanHistory.Save()
}

func ListGitScanResults(targetID int) ([]Git, error) {
	var git []Git
	err := database.DbConn.Where("target_id = ?", targetID).Find(&git).Error
	if err != nil {
		return nil, err
	}
	return git, nil
}
