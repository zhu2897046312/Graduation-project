package handlers

import (
	"github.com/gin-gonic/gin"
	"server/models/core"
	"server/models/common"
	"server/service"
	"strconv"
)

type CoreAdminRoleIndexHandler struct {
	service *service.CoreAdminRoleIndexService
}

func NewCoreAdminRoleIndexHandler(service *service.CoreAdminRoleIndexService) *CoreAdminRoleIndexHandler {
	return &CoreAdminRoleIndexHandler{service: service}
}

// 创建管理员-角色关联
func (h *CoreAdminRoleIndexHandler) CreateAdminRole(c *gin.Context) {
	var index core.CoreAdminRoleIndex
	if err := c.ShouldBindJSON(&index); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.CreateAdminRole(&index); err != nil {
		Error(c, 4001, err.Error())
		return
	}

	Success(c, index)
}

// 删除管理员-角色关联
func (h *CoreAdminRoleIndexHandler) DeleteAdminRole(c *gin.Context) {
	adminID, err := strconv.ParseInt(c.Query("admin_id"), 10, 64)
	if err != nil || adminID <= 0 {
		InvalidParams(c)
		return
	}

	roleID, err := strconv.ParseInt(c.Query("role_id"), 10, 64)
	if err != nil || roleID <= 0 {
		InvalidParams(c)
		return
	}

	if err := h.service.DeleteAdminRole(common.MyID(adminID), common.MyID(roleID)); err != nil {
		Error(c, 4002, err.Error())
		return
	}

	Success(c, nil)
}

// 获取管理员角色
func (h *CoreAdminRoleIndexHandler) GetAdminRoles(c *gin.Context) {
	adminID, err := strconv.ParseInt(c.Param("admin_id"), 10, 64)
	if err != nil || adminID <= 0 {
		InvalidParams(c)
		return
	}

	roles, err := h.service.GetRolesByAdminID(common.MyID(adminID))
	if err != nil {
		Error(c, 4003, "获取角色失败")
		return
	}

	Success(c, roles)
}

// 删除管理员所有角色
func (h *CoreAdminRoleIndexHandler) DeleteAllAdminRoles(c *gin.Context) {
	adminID, err := strconv.ParseInt(c.Param("admin_id"), 10, 64)
	if err != nil || adminID <= 0 {
		InvalidParams(c)
		return
	}

	if err := h.service.DeleteAllRolesByAdminID(common.MyID(adminID)); err != nil {
		Error(c, 4004, err.Error())
		return
	}

	Success(c, nil)
}