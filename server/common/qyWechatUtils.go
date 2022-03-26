package common

import (
	"Antioxidant/server/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

/**
企业微信相关接口
*/

type QyResp struct {
	Errcode int
	Errmsg  string
}

/*
GenQyData
生成企业微信 api 数据，简单模式
*/
func (a *ApiConfig) GenQyData(repo *model.Repo) *model.QyData {
	data := model.QyData{
		Msgtype: "text",
		Text:    model.Text{Content: repo.HTMLURL},
	}
	return &data
}

/*
GenRepoQyMdData
生成 Markdown 格式数据
*/
func (a *ApiConfig) GenRepoQyMdData(d *Database) (bool, *model.MdData) {

	sendFlag := false
	content := "# Github Repos Update:\n"
	for _, repo := range d.Repos {
		log.Println("Compare repo...", repo.FullName)
		isGet, githubRepo := d.GithubService.GetRepoInfoByID(repo.RepoID)
		if !isGet {
			continue
		}

		switch {
		case githubRepo == nil:
			log.Println("No github repo...", repo.FullName)
			break
		case repo.PushedAt != githubRepo.PushedAt.String():
			log.Println("Get new pushed...", repo.FullName)

			// 拼接企微内容
			newFiles := d.GithubService.GetGithubRepoPushedData(repo.FullName, repo.PushedAt)
			if newFiles != nil {
				content += "[" + repo.HTMLURL + "](" + repo.HTMLURL + ")\n"

				for i, f := range newFiles {
					content += fmt.Sprintf("      %d. %s\n", i+1, *f)
				}

				sendFlag = true
			}

			// 更新时间
			d.UpdateRepo(repo, githubRepo)
			repo.PushedAt = githubRepo.PushedAt.String()
		default:
			log.Println("No pushed...", repo.FullName)
		}
	}

	log.Println("Send Flag is ...", sendFlag)
	if sendFlag {
		text := model.Text{Content: content}
		data := model.MdData{
			Msgtype:  "markdown",
			Markdown: text,
		}

		return sendFlag, &data
	} else {
		return sendFlag, nil
	}
}

/*
GenCveData
生成 cve 数据
*/
func (a *ApiConfig) GenCveData(cves []*model.CVE) (bool, *model.MdData) {

	sendFlag := false
	content := "# 你有新的CVE，请注意查收: \n\n"

	for _, cve := range cves {
		content += "## " + cve.CveID + "\n"
		content += "  " + *cve.Desc + "\n"

		sendFlag = true
	}

	if sendFlag {
		text := model.Text{Content: content}
		data := model.MdData{
			Msgtype:  "markdown",
			Markdown: text,
		}

		return sendFlag, &data
	} else {
		return sendFlag, nil
	}
}

/*
SendData2QY
发送数据到企业微信
*/
func (a *ApiConfig) SendData2QY(data *model.MdData) {
	bytesData, err := json.Marshal(data)
	if err != nil {
		log.Fatalln(err)
	}

	reader := strings.NewReader(string(bytesData))
	request, err := http.NewRequest("POST", a.QyWebhook, reader)
	if err != nil {
		log.Fatalln(err)
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")

	client := &http.Client{}
	log.Println("Send Data...")
	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	r := &QyResp{}
	err = json.Unmarshal(respBytes, r)
	if err != nil {
		log.Fatalln(err)
	}

	if r.Errcode == 0 {
		log.Println("Send Success...")
	} else {
		log.Println("Send Failed...", *r)
	}
}
