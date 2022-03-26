package model

import "gorm.io/gorm"

/**
CVE 表
*/

type CVE struct {
	gorm.Model
	ID       uint    `gorm:"primaryKey"` // ID
	CveID    string  // cve 编号
	PushedAt string  // 发布时间
	URL      string  // cve url
	Desc     *string // cve 描述
}
