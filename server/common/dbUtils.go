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
queryRepo
查询
*/
func (d *Database) queryRepo(repoName string) (bool, *model.Repo) {
	repo := model.Repo{}
	r := d.DB.Where("Name = ?", repoName).First(&repo)
	if err := r.Error; err == nil {
		log.Println("Query data got...", repoName)
		return true, &repo
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
	isQuery, _ := d.queryRepo(repo.Name)
	if !isQuery {
		log.Println("Insert repo...", repo.URL)
		d.DB.Create(repo)
	}
}

/*
UpdateRepo
更新repo
*/
func (d *Database) UpdateRepo(repo *model.Repo, githubRepo *github.Repository) {
	d.DB.Model(repo).Where(
		"URL = ?", githubRepo.HTMLURL).Update("updated_at", githubRepo.UpdatedAt.String())
	log.Println("Update Repo...", repo.Name)
}
