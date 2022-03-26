package common

import (
	"Antioxidant/server/model"
	"log"

	"github.com/google/go-github/github"
)

/**
数据库操作
*/

/*
QueryRepo
查询
*/
func (d *Database) QueryRepo(repoName string) (bool, *model.Repo) {
	repo := &model.Repo{}
	r := d.DB.Where("full_name = ?", repoName).First(repo)
	if err := r.Error; err == nil {
		log.Println("Query data got...", repoName)
		return true, repo
	} else {
		log.Println("No query data...")
		return false, nil
	}
}

/*
InsertRepo
插入数据
*/
func (d *Database) InsertRepo(repo *model.Repo) {
	isQuery, _ := d.QueryRepo(repo.FullName)
	if !isQuery {
		log.Println("Insert repo...", repo.HTMLURL)
		d.DB.Create(repo)
	}
}

/*
UpdateRepo
更新repo
*/
func (d *Database) UpdateRepo(repo *model.Repo, githubRepo *github.Repository) {
	d.DB.Model(repo).Where(
		"html_url = ?", githubRepo.HTMLURL).Update("pushed_at", githubRepo.PushedAt.String())
	log.Println("Update Repo...", repo.FullName)
}

/*
QueryCve
查询 cve 信息
*/
func (d *Database) QueryCve(cveID string) (bool, *model.CVE) {
	cve := &model.CVE{}
	r := d.DB.Where("cve_id = ?", cveID).First(cve)
	if err := r.Error; err == nil {
		log.Println("Query data got...", cveID)
		return true, cve
	} else {
		log.Println("No query data...")
		return false, nil
	}
}

/*
InsertCve
插入数据
*/
func (d *Database) InsertCve(cve *model.CVE) {
	isQuery, _ := d.QueryCve(cve.CveID)
	if !isQuery {
		log.Println("Insert cve...", cve.CveID)
		d.DB.Create(cve)
	}
}
