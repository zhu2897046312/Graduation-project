package handlers

import (
	"encoding/json"
	"server/models/core"
	"server/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CoreRoleHandler struct {
	service *service.CoreRoleService
}

func NewCoreRoleHandler(service *service.CoreRoleService) *CoreRoleHandler {
	return &CoreRoleHandler{service: service}
}

// 创建角色
func (h *CoreRoleHandler) CreateRole(c *gin.Context) {
	var role core.CoreRole
	if err := c.ShouldBindJSON(&role); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.CreateRole(&role); err != nil {
		Error(c, 10001, err.Error())
		return
	}

	Success(c, role)
}

// 更新角色
func (h *CoreRoleHandler) UpdateRole(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		InvalidParams(c)
		return
	}

	var role core.CoreRole
	if err := c.ShouldBindJSON(&role); err != nil {
		InvalidParams(c)
		return
	}
	role.ID = id

	if err := h.service.UpdateRole(&role); err != nil {
		Error(c, 10002, err.Error())
		return
	}

	Success(c, role)
}

// 获取角色详情
func (h *CoreRoleHandler) GetRole(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		InvalidParams(c)
		return
	}

	role, err := h.service.GetRoleByID(id)
	if err != nil {
		Error(c, 10003, "角色不存在")
		return
	}

	Success(c, role)
}

// 获取所有角色
func (h *CoreRoleHandler) GetAllRoles(c *gin.Context) {
	roles, err := h.service.GetAllRoles()
	if err != nil {
		Error(c, 10004, "获取角色列表失败")
		return
	}

	Success(c, roles)
}

// 更新角色状态
func (h *CoreRoleHandler) UpdateRoleStatus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		InvalidParams(c)
		return
	}

	var req struct {
		Status int8 `json:"status"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.UpdateRoleStatus(id, req.Status); err != nil {
		Error(c, 10005, err.Error())
		return
	}

	Success(c, nil)
}

// 更新角色权限
func (h *CoreRoleHandler) UpdateRolePermissions(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		InvalidParams(c)
		return
	}

	var req struct {
		Permissions json.RawMessage `json:"permissions"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}

	if len(req.Permissions) == 0 {
		Error(c, 10006, "权限数据不能为空")
		return
	}

	if err := h.service.UpdateRolePermissions(id, []byte(req.Permissions)); err != nil {
		Error(c, 10007, err.Error())
		return
	}

	Success(c, nil)
}

func (h *CoreRoleHandler) List(c *gin.Context) {
	var req struct {
		Page     int `json:"page"`
		PageSize int `json:"page_size"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}
	
	roles, total, err := h.service.List(req.Page, req.PageSize)
	if err != nil {
		Error(c, 10008, "获取角色列表失败")
		return
	}

	Success(c, gin.H{
		"list":  roles,
		"total": total,
	})
}
