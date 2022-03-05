package model

import (
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Article struct {
	ID        int            `gorm:"primarykey"json:"id"`
	CreatedAt time.Time      `gorm:"type:varchar(20)" json:"created_at"`
	UpdatedAt time.Time      `gorm:"type:varchar(20)" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Title     string         `gorm:"type:varchar(20);not null" json:"title"`
	Cid       int            `gorm:"type:int(10);not null" json:"cid"`
	Uid       int            `gorm:"type:int(10);not null" json:"uid"`
	Desc      string         `gorm:"type:varchar(200)" json:"desc"`
	Content   string         `gorm:"type:varchar(200)" json:"content"`
	Img       string         `gorm:"type:varchar(200)" json:"img"`
}

func (Article) TableName() string {
	return "article"
}

func GetArticles(uid int) ([]Article, error) {
	var articles []Article

	rows, err := db.Table("article").Select("*").Where("uid = ?", uid).Rows()
	defer rows.Close()
	if err != nil {
		return articles, errors.Wrap(err, "model：GetArticles 获取文章异常")
	}

	for rows.Next() {
		err := db.ScanRows(rows, &articles)
		if err != nil {
			return articles, errors.Wrap(err, "model：GetArticles 获取文章异常")
		}
	}

	return articles, nil
}
