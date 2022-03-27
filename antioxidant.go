package main

import (
	"Antioxidant/server/common"
	"Antioxidant/server/monitor"
	"fmt"
	"log"
	"sync"
)

func main() {
	log.Println("Start Antioxidant...")

	ac, err := common.ReadConfig("./config.yaml")
	if err != nil {
		fmt.Println(err)
	}

	// 企业微信 webhook
	webhook := common.ApiConfig{QyWebhook: ac.QyWechat}

	// 初始化数据
	log.Println("Init Data...")
	d := common.Database{
		GithubService: &common.GithubService{
			GithubToken: ac.GithubToken,
		},
	}
	d.ConnDB()
	d.GithubService.GetGithubClient()
	d.InitDB()

	var wg sync.WaitGroup
	wg.Add(2)

	// 定时查询 github repo
	go monitor.RunRepoMonitor(&d, &webhook, wg)

	// 定时查询 github cve
	go monitor.RunCVEMonitor(&d, &webhook, wg)

	wg.Wait()
}
