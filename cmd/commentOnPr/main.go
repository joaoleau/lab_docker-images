package main

import (
	"context"
	"os"
	"strconv"
	"github.com/google/go-github/v57/github"
	"golang.org/x/oauth2"
)

func main() {
	token := os.Getenv("GITHUB_TOKEN")
	prNumber := os.Getenv("PR_NUMBER")
	prNum, _ := strconv.Atoi(prNumber)
	
	owner := os.Getenv("GITHUB_OWNER")
	repo := os.Getenv("GITHUB_REPO")

	pathMdFile := os.Getenv("PATH_MD_FILE")
	newBody, _ := os.ReadFile(pathMdFile)

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	comment := &github.IssueComment{
		Body: github.String(string(newBody)),
	}
	_, _, _ = client.Issues.CreateComment(ctx, owner, repo, prNum, comment)
}
