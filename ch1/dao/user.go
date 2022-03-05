package dao

import (
	"gocamp/ch1/model"

	"github.com/pkg/errors"
)

// 整个层级结构氛围4层
// 用户交互层ctrl
// 服务聚合层service
// model聚合层dao
// 数据访问层model
// dao层错误应该被包装，model可以直接返回error根因(调用其他包)，不同的dao对model会做聚合
// 若在model发生error，通常是调用三方库导致的错误，需要通过errors.Wrap()包装后返回
// 若在dao调用model方法出现错误，直接返回err即可
// 若在dao处理三方库底层错误导致，需要通过errors.Wrap()包装返回
type UserArticle struct {
	Id       int             `json:"id"`
	Username string          `json:"username"`
	Articles []model.Article `json:"articles"`
}

// 获取用户文章
func GetUserArticle(uid int) (UserArticle, error) {
	var err error
	var user *model.User
	var articles []model.Article
	var userArticle UserArticle
	user, err = model.GetUserName(uid)
	if err != nil {
		return userArticle, err
	}
	rows, err := model.DB.Table("article").Select("*").Where("uid = ?", uid).Rows()
	if err != nil {
		return userArticle, errors.Wrap(err, "dao ：GetUserArticle method article查询错误")
	}
	for rows.Next() {
		err := model.DB.ScanRows(rows, &articles)
		if err != nil {
			return userArticle, errors.Wrap(err, "dao ：GetUserArticle method ScanRows错误")
		}
	}
	userArticle.Id = uid
	userArticle.Username = user.Username
	userArticle.Articles = articles
	return userArticle, nil
}

type Level struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}
type UserLevel struct {
	Uid   int   `json:"uid"`
	Level Level `json:"level"`
}

// 获取用户等级
func GetUserLevels(uid int) (UserLevel, error) {
	level := Level{"LX", "铂金"}
	userLevel := UserLevel{uid, level}
	return userLevel, nil
}
