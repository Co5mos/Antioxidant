package monitor

import (
	"Antioxidant/server/common"
	"log"
	"sync"

	"github.com/robfig/cron"
)

/**
监控 github
*/

/*
RunRepoMonitor
启动 repo 监控
*/
func RunRepoMonitor(d *common.Database, a *common.ApiConfig, wg sync.WaitGroup) {
	// 定时查询 github repo
	wg.Done()

	c := cron.New()
	c.AddFunc("0 */10 * * * *", func() { // 每小时更新执行一次
		log.Println("Compare Data...")

		//  获取 github repo 并发送信息
		a.GenRepoQyMdData(d)
	})
	c.Start()
	select {}
}
