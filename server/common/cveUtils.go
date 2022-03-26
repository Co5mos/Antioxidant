package common

import (
	"Antioxidant/server/model"
	"bytes"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"

	"github.com/go-resty/resty/v2"

	"github.com/google/go-github/github"
)

/**
CVE 查询方法
*/

type CVEService struct {
}

/*
GetThisYearCves
抓取本年度的cve
*/
func (c *CVEService) GetThisYearCves(result *github.RepositoriesSearchResult) []*model.CVE {
	var cves []*model.CVE
	today := time.Now().Format("2006-01-02")

	repos := result.Repositories
	for _, repo := range repos {
		// 正则匹配CVE编号
		repoName := strings.ToUpper(*repo.Name)
		reg := regexp.MustCompile(`CVE-\d+-\d+`)
		if reg == nil {
			log.Println("No cve info...", repoName)
			continue
		}

		cve := reg.FindAllStringSubmatch(repoName, -1)
		if cve == nil {
			continue
		}

		// 查询数据库中是否已经存在该CVE

		pushedAt := repo.PushedAt.Format("2006-01-02")
		if pushedAt == today {
			isCve, cveDesc := c.GetCVEFromOrg(cve[0][0])
			if !isCve {
				continue
			}

			cves = append(cves, &model.CVE{
				CveID:    cve[0][0],
				PushedAt: pushedAt,
				URL:      *repo.HTMLURL,
				Desc:     cveDesc,
			})
		}
	}

	return cves
}

/*
GetCVEFromOrg
cve 官网查询当前 cve 编号是否存在
*/
func (c *CVEService) GetCVEFromOrg(cve string) (bool, *string) {
	url := "https://cve.mitre.org/cgi-bin/cvename.cgi?name=" + cve

	// 创建请求
	client := resty.New()
	resp, err := client.R().Get(url)
	log.Println("Request CVE org...", resp.StatusCode())
	if err != nil || resp.StatusCode() != 200 {
		log.Println("No cve from org...")
		return false, nil
	}

	// html解析
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(resp.Body()))
	if err != nil {
		log.Println("Parsed cve org failed...")
		return false, nil
	}

	var desc string
	doc.Find("#GeneratedTable > table > tbody > tr:nth-child(4) > td").Each(func(i int, s *goquery.Selection) {
		desc = s.Text()
	})

	if len(desc) > 0 {
		return true, &desc
	} else {
		return false, nil
	}
}
