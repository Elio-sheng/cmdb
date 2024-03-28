package admin_v1

import (
	"gin-cmdb/internal/controller"
	"gin-cmdb/internal/service/admin_auth"
	"gin-cmdb/internal/validator"
	"gin-cmdb/internal/validator/form"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	controller.Api
}

func NewLoginController() *LoginController {
	return &LoginController{}
}

// Login 管理员用户登录
func (api LoginController) Login(c *gin.Context) {
	// 初始化参数结构体
	loginForm := form.NewLoginForm()
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckPostParams(c, &loginForm); err != nil {
		return
	}
	// 实际业务调用
	result, err := admin_auth.NewLoginService().Login(loginForm.UserName, loginForm.PassWord)
	// 根据业务返回值判断业务成功 OR 失败
	if err != nil {
		api.Err(c, err)
		return
	}

	api.Success(c, result)
	return
}
