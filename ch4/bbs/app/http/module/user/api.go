package user

import (
	"github.com/gohade/hade/framework/gin"
	"gocamp/ch4/bbs/app/provider/user"
)

type UserApi struct {}

// RegisterRoutes 注册路由
func RegisterRoutes(r *gin.Engine) error {
	api := &UserApi{}
	if !r.IsBind(user.UserKey) {
		r.Bind(&user.UserProvider{})
	}
	// 注册
	r.POST("/user/register", api.Register)
	// 登录
	r.POST("/user/login", api.Login)
	// 登出
	r.GET("/user/logout", api.Logout)
	// 注册
	r.GET("/user/register/verify", api.Verify)

	return nil
}