package user

import (
	"fmt"
	"github.com/gohade/hade/framework/gin"
	provider "gocamp/ch4/bbs/app/provider/user"
)


// Verify 验证
// @Summary 用户验证
// @Description 用户验证接口
// @Accept  json
// @Produce  json
// @Tags user
// @Param verifyParam body verifyParam true "验证参数"
// @Success 200 {string} Message "验证成功"
// @Router /user/register/verify [get]
func (api *UserApi) Verify(c *gin.Context) {
	userService := c.MustMake(provider.UserKey).(provider.Service)

	token := c.Query("token")
	if token == "" {
		c.ISetStatus(404).IText("参数错误")
		return
	}
	status , err := userService.VerifyRegister(c, token)
	fmt.Println(status , err)
	if status == false || err != nil {
		c.ISetStatus(500).IText("没有预注册，请先注册")
	}
	c.ISetOkStatus().IText("验证成功")
}