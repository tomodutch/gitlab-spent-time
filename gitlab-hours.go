package main

import (
	"flag"
	"log"
	"regexp"
	"strconv"

	"github.com/go-git/go-git"
	"github.com/xanzy/go-gitlab"
)

func main() {
	projectID := flag.Int("project", 0, "project id")
	apiKey := flag.String("apikey", "", "Api key")
	gitRepoPath := flag.String("repo", "", "Full path to .git directory")
	gitlabURL := flag.String("url", "https://gitlab.com/api/v4", "gitlab url")
	flag.Parse()
	log.Printf("gitlab-hours. project: %v", *projectID)
	re := regexp.MustCompile(`#(?P<issue>\d+)\+(?P<spent>\w+)`)
	gitlabClient, err := gitlab.NewClient(*apiKey, gitlab.WithBaseURL(*gitlabURL))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		return
	}

	repo, err := git.PlainOpen(*gitRepoPath)
	if err != nil {
		log.Fatalf("Failed to find git repository %v", err)
		return
	}

	ref, err := repo.Head()
	if err != nil {
		log.Fatalf("Failed to open HEAD in git repository %v", err)
		return
	}

	cIter, err := repo.Log(&git.LogOptions{From: ref.Hash()})
	latestCommit, err := cIter.Next()

	matches := re.FindAllStringSubmatch(latestCommit.Message, -1)
	for _, timeSpent := range matches {
		issueID, _ := strconv.Atoi(timeSpent[1])
		timeSpent := timeSpent[2]
		log.Printf("Recording time for project: %v issue: %v time: %v", *projectID, issueID, timeSpent)
		_, _, err := gitlabClient.Issues.AddSpentTime(*projectID, issueID, &gitlab.AddSpentTimeOptions{
			Duration: &timeSpent,
		})
		if err != nil {
			log.Fatalf("Failed to save time: %v", err)
		}
	}
}
