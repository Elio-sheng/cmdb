package admin_v1

import (
	"gin-cmdb/internal/controller"
)

type CommonController struct {
	controller.Api
}

func NewCommonController() *CommonController {
	return &CommonController{}
}
