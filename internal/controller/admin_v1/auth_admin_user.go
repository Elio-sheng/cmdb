package admin_v1

import (
	"gin-cmdb/internal/controller"
	"gin-cmdb/internal/service/admin_auth"
	"gin-cmdb/internal/validator"
	"gin-cmdb/internal/validator/form"
	"github.com/gin-gonic/gin"
)

type AdminUserController struct {
	controller.Api
}

func NewAdminUserController() *AdminUserController {
	return &AdminUserController{}
}

func (api AdminUserController) GetUserInfo(c *gin.Context) {
	result, err := admin_auth.NewAdminUserService().GetUserInfo(c.GetUint("a_uid"))
	if err != nil {
		api.Err(c, err)
		return
	}
	api.Success(c, result)
	return
}

func (api AdminUserController) Add(c *gin.Context) {
	// 初始化参数结构体
	IDForm := form.NewIDForm()
	//// 绑定参数并使用验证器验证参数
	if err := validator.CheckQueryParams(c, &IDForm); err != nil {
		return
	}
	result, err := admin_auth.NewAdminUserService().GetUserInfo(IDForm.ID)
	if err != nil {
		api.Err(c, err)
		return
	}
	api.Success(c, result)
	return
}

func (api AdminUserController) Delete(c *gin.Context) {
	// 初始化参数结构体
	IDForm := form.NewIDForm()
	//// 绑定参数并使用验证器验证参数
	if err := validator.CheckQueryParams(c, &IDForm); err != nil {
		return
	}
	result, err := admin_auth.NewAdminUserService().GetUserInfo(IDForm.ID)
	if err != nil {
		api.Err(c, err)
		return
	}
	api.Success(c, result)
	return
}
