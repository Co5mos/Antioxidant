package monitor

import (
	"Antioxidant/server/common"
	"Antioxidant/server/model"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/robfig/cron"
)

/**
监控最新的 cve
*/

/*
RunCVEMonitor
启动 cve 监控
*/
func RunCVEMonitor(d *common.Database, a *common.ApiConfig, wg sync.WaitGroup) {
	wg.Done()

	// 定时操作
	c := cron.New()
	c.AddFunc("0 */10 * * * *", func() {
		log.Println("Get cve info...")

		// 初始化操作
		year := time.Now().Year()
		keyword := "CVE-" + strconv.Itoa(year)
		log.Println("Get cve from github by...", keyword)

		// github 查询
		reposResult := d.GithubService.GetRepos("CVE-" + strconv.Itoa(year))
		if reposResult == nil {
			log.Println("No repos result...")
		}

		// 获取当年的cve
		var cves []*model.CVE
		cs := common.CVEService{}
		today := time.Now().Format("2006-01-02")
		log.Println("Today is ...", today)
		sendFlag := false
		content := "# 你有新的CVE，请注意查收: \n\n"

		repos := reposResult.Repositories
		for _, repo := range repos {
			// 正则匹配CVE编号
			repoName := strings.ToUpper(*repo.Name)
			reg := regexp.MustCompile(`(CVE-\d+-\d+)`)
			if reg == nil {
				log.Println("No cve info...", repoName)
				continue
			}

			cveIDs := reg.FindAllStringSubmatch(repoName, -1)
			if cveIDs == nil {
				continue
			}

			cveID := cveIDs[0][0]
			// 查询数据库中是否已经存在该CVE
			isQuery, err := d.QueryCve(cveID)
			if isQuery || err != nil {
				continue
			}

			// 官网查询CVE正确性
			pushedAt := repo.PushedAt.Add(8 * time.Hour).Format("2006-01-02")
			log.Println("CVE repo pushed at...", pushedAt)
			if pushedAt == today {
				isCve, cveDesc := cs.GetCVEFromOrg(cveIDs[0][0])
				if !isCve {
					continue
				}

				cve := &model.CVE{
					CveID:    cveID,
					PushedAt: pushedAt,
					URL:      *repo.HTMLURL,
					Desc:     cveDesc,
				}

				// 将cve信息入库
				d.InsertCve(cve)
				cves = append(cves, cve)

				// 组织发送数据
				content += "## " + cve.CveID + "\n"
				content += fmt.Sprintf("[%s](%s)", cve.URL, cve.URL) + "\n"
				content += "> " + *cve.Desc + "\n"

				sendFlag = true
			}
		}

		// 发送数据
		if sendFlag {
			text := model.Text{Content: content}
			data := model.MdData{
				Msgtype:  "markdown",
				Markdown: text,
			}

			log.Println("Send Qy data...")
			a.SendData2QY(&data)
		}
	})
	c.Start()
	select {}
}
