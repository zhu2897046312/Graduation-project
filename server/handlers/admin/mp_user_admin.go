package admin

import (
	"server/models/mp"
	"server/service"
	"server/utils"
	"server/models/common"

	"github.com/gin-gonic/gin"
)

type ListMpUserRequest struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type MpUserHandler struct {
	service *service.MpUserService
}

func NewMpUserHandler(service *service.MpUserService) *MpUserHandler {
	return &MpUserHandler{service: service}
}

func (h *MpUserHandler) List(c *gin.Context) {
	var req ListMpUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		Error(c, 400, "参数错误")
		return
	}
	users, total, err := h.service.ListMpUser(mp.MpUserListParam{
		Page: req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		Error(c, 500, "获取用户列表失败")
		return
	}
	Success(c, gin.H{
		"list": users,
		"total": total,
	})
}

func (h *MpUserHandler) DeleteMpUser(c *gin.Context) {
	id  := c.Query("id")
	if id == "" {
		InvalidParams(c)
		return
	}
	uID := utils.ConvertToUint(id)
	if uID == 0 {
		InvalidParams(c)
		return
	}
	
	err := h.service.DeleteMpUser(common.MyID(uID))
	if err != nil {
		Error(c, 500, "删除用户失败")
		return
	}
	Success(c, "删除用户成功")
}


