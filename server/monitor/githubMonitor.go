package monitor

import (
	"Antioxidant/server/common"
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

/**
监控 github
*/

/*
StarRepo
抓取配置的 github 中的 star 更新
*/
func StarRepo(token *common.ThirdPartyToken) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token.GithubToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// 获取当前用户的所有 starred repos
	listOpt := github.ActivityListStarredOptions{
		Sort:        "",
		Direction:   "",
		ListOptions: github.ListOptions{},
	}
	repos, _, _ := client.Activity.ListStarred(ctx, "co5mos", &listOpt)
	fmt.Println(len(repos))
	for _, repo := range repos {
		r := repo.GetRepository()
		fmt.Println(*r.Name)
		fmt.Println(*r.URL)
		fmt.Println(r.GetPushedAt())
	}
}

/*
CompareAllRepo
比较所有repo的更新时间
*/
func CompareAllRepo(d *common.Database, apiConfig *common.ApiConfig) {
	//  获取 github repo 信息
	sendFlag, data := apiConfig.GenQyMdData(d)
	if sendFlag {
		log.Println("Send Qy data...")
		apiConfig.SendData2QY(data)
	}
}
