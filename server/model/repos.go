package model

import (
	"time"

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
	ID       uint   `gorm:"primaryKey"` // ID
	RepoID   int64  // 仓库 ID
	FullName string // 仓库名称
	RepoType string // 仓库类型
	HTMLURL  string // 仓库地址
	PushedAt string // 更新时间
}

/*
GenRepoData
生成Repo数据
*/
func (r *Repo) GenRepoData(githubRepo *github.Repository, repoTypes string) {

	timestamp := githubRepo.GetPushedAt()
	pushedAt := timestamp.Time.Add(8 * time.Hour)
	r.RepoID = *githubRepo.ID
	r.FullName = *githubRepo.FullName
	r.RepoType = repoTypes
	r.HTMLURL = *githubRepo.HTMLURL
	r.PushedAt = pushedAt.String()
}
