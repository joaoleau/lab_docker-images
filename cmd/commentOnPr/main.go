package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/google/go-github/v57/github"
	"golang.org/x/oauth2"
)

func main() {
	log.Println("🚀 Iniciando processo de comentário no PR...")

	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Fatal("❌ GITHUB_TOKEN não encontrado nas variáveis de ambiente")
	}

	prNumber := os.Getenv("PR_NUMBER")
	if prNumber == "" {
		log.Fatal("❌ PR_NUMBER não encontrado nas variáveis de ambiente")
	}

	prNum, err := strconv.Atoi(prNumber)
	if err != nil {
		log.Fatalf("❌ PR_NUMBER inválido: %v", err)
	}

	owner := os.Getenv("GITHUB_OWNER")
	if owner == "" {
		log.Fatal("❌ GITHUB_OWNER não encontrado nas variáveis de ambiente")
	}

	repo := os.Getenv("GITHUB_REPO")
	if repo == "" {
		log.Fatal("❌ GITHUB_REPO não encontrado nas variáveis de ambiente")
	}

	pathMdFile := os.Getenv("PATH_MD_FILE")
	if pathMdFile == "" {
		log.Fatal("❌ PATH_MD_FILE não encontrado nas variáveis de ambiente")
	}

	log.Printf("📄 Lendo conteúdo do arquivo: %s", pathMdFile)
	newBody, err := os.ReadFile(pathMdFile)
	if err != nil {
		log.Fatalf("❌ Erro ao ler arquivo .md: %v", err)
	}

	log.Printf("💬 Criando comentário no PR #%d do repositório %s/%s", prNum, owner, repo)

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	comment := &github.IssueComment{
		Body: github.String(string(newBody)),
	}

	createdComment, resp, err := client.Issues.CreateComment(ctx, owner, repo, prNum, comment)
	if err != nil {
		log.Fatalf("❌ Erro ao criar comentário: %v", err)
	}

	log.Printf("✅ Comentário criado com sucesso! URL: %s (Status: %d)", createdComment.GetHTMLURL(), resp.StatusCode)
	fmt.Println("Fim do processo.")
}
