package common

import (
	"Antioxidant/server/model"
	"encoding/json"
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
		Text:    model.Text{Content: repo.URL},
	}
	return &data
}

/*
GenQyMdData
生成 Markdown 格式数据
*/
func (a *ApiConfig) GenQyMdData(d *Database) (bool, *model.MdData) {

	sendFlag := false
	content := "# Github Repos Update:\n"
	for _, repo := range d.Repos {
		log.Println("Compare repo...", repo.Name)
		githubRepo := d.Tokens.GetGithubRepoInfo(repo.Name)
		if githubRepo.Name != nil {
			if repo.UpdatedAt != githubRepo.UpdatedAt.String() {
				content = content + "* [" + repo.URL + "](" + repo.URL + ")\n"

				sendFlag = true

				// 更新时间
				d.UpdateRepo(repo, githubRepo)
			}
		}
	}

	text := model.Text{Content: content}
	data := model.MdData{
		Msgtype:  "markdown",
		Markdown: text,
	}

	if sendFlag {
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
