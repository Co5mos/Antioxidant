package common

import (
	"Antioxidant/server/model"
	"log"
	"strings"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	DB     *gorm.DB
	Repos  []*model.Repo
	Tokens *ThirdPartyToken
}

/*
yaml2db
将yaml数据更新到数据库
*/
func (d *Database) yaml2db(repoURL string, rtype string) {
	repoURLSplit := strings.Split(repoURL, "/")
	repoName := repoURLSplit[len(repoURLSplit)-2] + "/" + repoURLSplit[len(repoURLSplit)-1]

	// 判断库里是否存在
	isQuery, repo := d.queryRepo(repoName)
	if !isQuery {
		githubRepo := d.Tokens.GetGithubRepoInfo(repoName)
		if githubRepo.Name != nil {
			// 写入数据库
			r := model.Repo{}
			r.GenRepoData(githubRepo, rtype)
			d.Repos = append(d.Repos, &r)
			d.InsertRepo(&r)
		} else {
			log.Println("No Repo...")
		}
	} else {
		d.Repos = append(d.Repos, repo)
	}
}

/*
SaveYamlData
将yaml文件中的内容保存到数据库
*/
func (d *Database) SaveYamlData() {
	// 读取 yaml 文件
	repoUrls, err := ReadYaml("./server/repos/Repos.yaml")
	if err != nil {
		log.Println(err)
	}

	// 将 yaml 数据更新到数据库
	// TODO 读取不同的类型
	for _, vrepo := range repoUrls.Vuln {
		d.yaml2db(vrepo, "vuln")
	}
	for _, trepo := range repoUrls.Tool {
		d.yaml2db(trepo, "tool")
	}
}

/*
ConnDB
连接数据库
*/
func (d *Database) ConnDB() {
	// 连接数据库
	db, err := gorm.Open(sqlite.Open("repo.DB"), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}

	d.DB = db
}

/*
InitDB
初始化数据库
*/
func (d *Database) InitDB(tokens *ThirdPartyToken) {

	// 迁移文件
	if !d.DB.Migrator().HasTable(&model.Repo{}) {
		log.Println("Create Repo Table...")
		d.DB.Migrator().CreateTable(&model.Repo{})
	}

	// 检测表结构
	d.DB.AutoMigrate(&model.Repo{})
	d.Tokens = tokens

	// 读取 yaml 文件
	d.SaveYamlData()
}
