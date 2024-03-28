package admin_v1

import "gin-cmdb/internal/controller"

type RoleController struct {
	controller.Api
}

func NewRoleController() *RoleController {
	return &RoleController{}
}
