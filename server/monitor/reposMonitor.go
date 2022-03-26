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

	c1 := cron.New()
	c1.AddFunc("0 */1 * * * *", func() { // 每小时更新执行一次
		// 比较数据
		log.Println("Compare Data...")

		//  获取 github repo 信息
		sendFlag, data := a.GenRepoQyMdData(d)
		if sendFlag {
			log.Println("Send Qy data...")
			a.SendData2QY(data)
		}
	})
	c1.Start()
	select {}
}
