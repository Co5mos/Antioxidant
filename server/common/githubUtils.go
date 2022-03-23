package common

import (
	"context"
	"log"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

/**
github 信息获取
*/

/*
获取 github client
*/
func (t *ThirdPartyToken) getGithubClient() (*github.Client, context.Context) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: t.GithubToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	return client, ctx
}

/*
GetGithubRepoInfo
查询仓库信息
*/
func (t *ThirdPartyToken) GetGithubRepoInfo(repoName string) *github.Repository {
	client, ctx := t.getGithubClient()

	searchOpt := github.SearchOptions{Sort: "stars"}
	rs, _, err := client.Search.Repositories(ctx, repoName, &searchOpt)
	if err != nil {
		log.Println(err)
	}

	if len(rs.Repositories) > 0 {
		return &rs.Repositories[0]
	} else {
		return &github.Repository{}
	}

}
