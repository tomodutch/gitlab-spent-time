package main

import (
	"flag"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/go-git/go-git"
	"github.com/xanzy/go-gitlab"
)

func main() {
	projectID := flag.Int("project", 0, "project id")
	apiKey := flag.String("apikey", "", "Api key")
	gitlabURL := flag.String("url", "https://gitlab.com/api/v4", "gitlab url")
	flag.Parse()
	re := regexp.MustCompile(`#(?P<issue>\d+)\+(?P<spent>\w+)`)
	gitlabClient, err := gitlab.NewClient(*apiKey, gitlab.WithBaseURL(*gitlabURL))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	path, _ := os.Getwd()
	repo, _ := git.PlainOpen(path)
	ref, _ := repo.Head()
	cIter, err := repo.Log(&git.LogOptions{From: ref.Hash()})
	latestCommit, err := cIter.Next()

	matches := re.FindAllStringSubmatch(latestCommit.Message, -1)
	for _, timeSpent := range matches {
		issueID, _ := strconv.Atoi(timeSpent[1])
		timeSpent := timeSpent[2]
		_, _, err := gitlabClient.Issues.AddSpentTime(*projectID, issueID, &gitlab.AddSpentTimeOptions{
			Duration: &timeSpent,
		})
		if err != nil {
			log.Fatalf("Failed to save time: %v", err)
		}
	}
}
