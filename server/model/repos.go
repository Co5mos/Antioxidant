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
HotVuln
热点漏洞
*/
type HotVuln struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"` // ID
	HotKey   string // 热词
	RepoID   int64  // 仓库 ID
	FullName string // 仓库名称
	HTMLURL  string // 仓库地址
	CommitAt string // 更新时间
}

/*
GenRepoData
生成Repo数据
*/
func (r *Repo) GenRepoData(githubRepo *github.Repository, repoTypes string, commitTime time.Time) {

	r.RepoID = *githubRepo.ID
	r.FullName = *githubRepo.FullName
	r.RepoType = repoTypes
	r.HTMLURL = *githubRepo.HTMLURL
	r.PushedAt = commitTime.String()
}

/*
GenHotVulnData
生成热点漏洞数据
*/
func (h *HotVuln) GenHotVulnData(githubRepo *github.Repository, hotKey string, commitTime time.Time) {

	h.RepoID = *githubRepo.ID
	h.HotKey = hotKey
	h.FullName = *githubRepo.FullName
	h.HTMLURL = *githubRepo.HTMLURL
	h.CommitAt = commitTime.String()
}
