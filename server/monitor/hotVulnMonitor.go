package monitor

import (
	"Antioxidant/server/common"
	"Antioxidant/server/model"
	"fmt"
	"log"
	"sync"

	"github.com/robfig/cron"
)

/**
监控热点漏洞
*/

/*
HotVulnMonitor
监控
*/
func HotVulnMonitor(d *common.Database, a *common.ApiConfig) {
	// 读取 热点漏洞 关键字
	hv, err := common.ReadHotVuln("./server/repos/HotKey.yaml")
	if err != nil {
		fmt.Println(err)
	}

	for _, keyword := range hv.Keywords {
		log.Println("Hot vuln keyword...", keyword)
		reposResult := d.GithubService.GetRepos(keyword)
		if reposResult == nil {
			log.Println("No hot vuln repo result...")
		}

		sendFlag := false
		content := "# 热点漏洞仓库\n"
		repos := reposResult.Repositories
		num := 0
		for _, repo := range repos {
			// 查询仓库中的热点漏洞仓库
			isQuery, _ := d.QueryHotVuln(*repo.ID)
			if isQuery {
				continue
			}

			// 拼接信息
			num += 1
			content += fmt.Sprintf("  %d. %s\n", num, keyword)
			content += fmt.Sprintf("     [%s](%s)\n", *repo.HTMLURL, *repo.HTMLURL)

			// 数据库插入数据
			h := model.HotVuln{}
			commitTime := d.GithubService.GetLastCommitDatetime(*repo.Owner.Login, *repo.Name)
			h.GenHotVulnData(&repo, keyword, commitTime)
			d.InsertHotVuln(&h)

			// 超过10个就不要了
			sendFlag = true
			if num > 4 {
				break
			}
		}

		if sendFlag {
			text := model.Text{Content: content}
			data := model.MdData{
				Msgtype:  "markdown",
				Markdown: text,
			}

			//  获取 github repo 并发送信息
			a.SendData2QY(&data)
		}
	}
}

/*
RunHotVulnMonitor
启动
*/
func RunHotVulnMonitor(d *common.Database, a *common.ApiConfig, wg sync.WaitGroup) {
	// 定时查询 github repo
	wg.Done()

	c := cron.New()
	c.AddFunc("0 */1 * * * *", func() { // 每小时更新执行一次
		// 比较数据
		log.Println("Compare Data...")

		HotVulnMonitor(d, a)
	})
	c.Start()
	select {}
}
