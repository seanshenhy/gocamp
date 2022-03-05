package service

import "gocamp/ch1/dao"

type UserInfo struct {
	UserArticle dao.UserArticle `json:"userArticle"`
	UserLevel   dao.UserLevel   `json:"userLevel"`
}

// 获取用户信息
func GetUserInfo(uid int) (UserInfo, error) {
	var userInfo UserInfo
	// 获取用户文章
	userArticles, err := dao.GetUserArticle(uid)
	if err != nil {
		return userInfo, err
	}
	userInfo.UserArticle = userArticles
	// 获取用户等级
	userLevel, err := dao.GetUserLevels(uid)
	if err != nil {
		return userInfo, err
	}
	userInfo.UserLevel = userLevel
	return userInfo, nil
}
