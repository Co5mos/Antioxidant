package model

import (
	"github.com/google/go-github/github"
	"gorm.io/gorm"
)

/**
github 仓库数据表接口
*/

/*
Repo
仓库表
*/
type Repo struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"` // ID
	Name      string // 仓库名称
	RepoType  string // 仓库类型
	URL       string // 仓库地址
	UpdatedAt string // 更新时间
}

/*
GenRepoData
生成Repo数据
*/
func (r *Repo) GenRepoData(githubRepo *github.Repository, repoTypes string) {

	timestamp := githubRepo.GetUpdatedAt()
	r.Name = *githubRepo.Name
	r.RepoType = repoTypes
	r.URL = *githubRepo.HTMLURL
	r.UpdatedAt = timestamp.Time.String()
}
