package utils

import (
	"context"
	"fmt"
	"os/exec"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func RunSubfinder(domain string) []string {
	cmd := exec.Command("subfinder", "-d", domain)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("could not run command", err)
	}
	subdomains := string(out)
	parsedSubdomains := strings.Split(subdomains, "\n")
	return parsedSubdomains

}

func FindOrgRepo(orgName string) []*github.Repository {
	ts := oauth2.StaticTokenSource(
		// replace with your github token
		&oauth2.Token{AccessToken: "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"},
	)
	tc := oauth2.NewClient(context.Background(), ts)
	client := github.NewClient(tc)
	repos, _, err := client.Repositories.ListByOrg(context.Background(), orgName, nil)
	if err != nil {
		fmt.Println("could not get repos", err)
	}
	fmt.Println(repos)
	return repos
}

func FindProtocal(domain string) string {
	cmd := exec.Command("httpx", "-p", "-host", domain)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("could not run command", err)
	}
	protocal := string(out)
	parsedProtocal := strings.Split(protocal, "\n")
	fmt.Println(parsedProtocal)
	return "protocal"
}

func FindPortAndProtocal(domain string) []string {
	cmd := exec.Command("naabu", "-host", domain, "-silent")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("could not run command", err)
	}
	portAndProtocal := string(out)
	parsedPortAndProtocal := strings.Split(portAndProtocal, "\n")
	return parsedPortAndProtocal
}
