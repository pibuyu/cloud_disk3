package models

import "time"

type ShareBasic struct {
	Id                 int
	Identity           string
	UserIdentity       string
	RepositoryIdentity string
	ExpireTime         int
	ClickNum           int
	CreatedAt          time.Time `xorm:"created"`
	UpdatedAt          time.Time `xorm:"updated"`
	DeletedAt          time.Time `xorm:"deleted"`
}

func (table ShareBasic) TableName() string {
	return "share_basic"

}
