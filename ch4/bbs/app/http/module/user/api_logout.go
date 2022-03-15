package user

import (
	"github.com/gohade/hade/framework/gin"
	provider "gocamp/ch4/bbs/app/provider/user"
)


// Logout 登出
// @Summary 用户登出
// @Description 用户登出接口
// @Accept  json
// @Produce  json
// @Tags user
// @Success 200 {string} Message "登出成功"
// @Router /user/logout [get]
func (api *UserApi) Logout(c *gin.Context) {
	userService := c.MustMake(provider.UserKey).(provider.Service)

	// 验证用户
	cookie,err := c.Cookie("token")
	if cookie == "" {
		c.ISetStatus(404).IText("请输入参数")
		return
	}
	user, err := userService.VerifyLogin(c, cookie)
	if err != nil {
		c.ISetStatus(404).IText("请输入参数")
		return
	}
	// 登出
	if err := userService.Logout(c, user); err != nil {
		c.ISetStatus(500).IText("登出失败")
		return
	}
	c.ISetOkStatus().IText("登出成功")
	return
}