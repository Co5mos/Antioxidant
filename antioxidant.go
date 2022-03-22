package main

import (
	"Antioxidant/server/common"
	"Antioxidant/server/monitor"
	"log"

	"github.com/robfig/cron"
)

func main() {
	log.Println("Start Antioxidant...")

	// github token
	token := common.ThirdPartyToken{GithubToken: ""}

	// 企业微信 webhook
	webhook := common.ApiConfig{
		QyWebhook: ""}

	// 初始化数据
	log.Println("Init Data...")
	d := common.Database{}
	d.ConnDB()
	d.InitDB(&token)

	// 定时
	c := cron.New()
	c.AddFunc("0 */1 * * * *", func() { // 每小时更新执行一次
		// 比较数据
		log.Println("Compare Data...")
		monitor.CompareAllRepo(&d, &webhook)
	})
	c.Start()
	select {}
}
