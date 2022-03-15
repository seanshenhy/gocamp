package user

import (
	"github.com/gohade/hade/framework/gin"
	provider "gocamp/ch4/bbs/app/provider/user"
)

type loginParam struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,gte=6"`
}

// Login 登录
// @Summary 用户登录
// @Description 用户登录接口
// @Accept  json
// @Produce  json
// @Tags user
// @Param loginParam body loginParam true "登录参数"
// @Success 200 {string} Message "登录成功"
// @Router /user/login [post]
func (api *UserApi) Login(c *gin.Context) {
	userService := c.MustMake(provider.UserKey).(provider.Service)

	// 参数校验
	param := &loginParam{}
	if err := c.ShouldBind(param);err != nil{
		c.ISetStatus(404).IText("参数错误")
		return
	}
	// model对象
	model := &provider.User{
		UserName: param.UserName ,
		Password:  param.Password,
	}

	// 查询是否存在
	userWithToken, err := userService.Login(c, model)
	if err != nil {
		c.ISetStatus(500).IText(err.Error())
		return
	}
	if userWithToken == nil {
		c.ISetStatus(500).IText("用户不存在")
		return
	}

	c.ISetOkStatus().IText(userWithToken.Token)
	return
}