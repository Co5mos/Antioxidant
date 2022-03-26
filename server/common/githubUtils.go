package common

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type GithubService struct {
	GithubToken string
	Client      *github.Client
	ctx         *context.Context
}

/**
github 信息获取
*/

/*
GetGithubClient
获取 github client
*/
func (g *GithubService) GetGithubClient() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: g.GithubToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	g.Client = client
	g.ctx = &ctx
}

/*
GetGithubRepoInfo
查询仓库信息
*/
func (g *GithubService) GetGithubRepoInfo(owner string, repoName string) *github.Repository {

	rs, resp, err := g.Client.Repositories.Get(*g.ctx, owner, repoName)
	log.Println("Github repo by owner resp status code...", resp.StatusCode)

	if err != nil {
		log.Println(err)
	}

	log.Println(fmt.Sprintf("Get repo from github... %s/%s", owner, repoName))
	return rs
}

/*
GetRepoInfoByID
通过仓库 id 获取仓库信息
*/
func (g *GithubService) GetRepoInfoByID(repoID int64) (bool, *github.Repository) {
	repo, resp, err := g.Client.Repositories.GetByID(*g.ctx, repoID)
	log.Println("Github repo by id resp status code...", resp.StatusCode)

	if err != nil {
		log.Println(err)
		return false, nil
	} else {
		return true, repo
	}
}

/*
GetGithubRepoPushedData
获取 github repo 最新的 push 数据
*/
func (g *GithubService) GetGithubRepoPushedData(repoFullName, pushedAt string) []*string {
	repoFullNameSplit := strings.Split(repoFullName, "/")
	owner := repoFullNameSplit[len(repoFullNameSplit)-2]
	repo := repoFullNameSplit[len(repoFullNameSplit)-1]

	// 获取最新一次push的commits
	since, err := time.Parse("2006-01-02 15:04:05", pushedAt)
	if err != nil {
		log.Println(err)
	}
	commits, resp, err := g.Client.Repositories.ListCommits(*g.ctx, owner, repo, &github.CommitsListOptions{Since: since})
	log.Println("Github commits resp status code...", resp.StatusCode)

	var addedFiles []*string
	for _, commit := range commits {
		commitSHA := commit.SHA

		// 获取 commit files
		repoCommit, resp1, err1 := g.Client.Repositories.GetCommit(*g.ctx, owner, repo, *commitSHA)
		if err1 != nil {
			log.Println(err1)
		}
		log.Println("Github commit resp status code...", resp1.StatusCode)
		repoCommitFiles := repoCommit.Files

		for _, f := range repoCommitFiles {
			// 添加文件
			if *f.Status == "added" {
				log.Println("New pushed file...", *f.Filename)
				addedFiles = append(addedFiles, f.Filename)
			}
		}

	}

	return addedFiles
}

/*
GetRepos
检索仓库
*/
func (g *GithubService) GetRepos(keyword string) *github.RepositoriesSearchResult {
	result, resp, err := g.Client.Search.Repositories(*g.ctx, keyword, &github.SearchOptions{Sort: "updated"})
	if err != nil {
		log.Println(err)
	}
	log.Println("Github repos resp status code...", resp.StatusCode)
	return result
}
