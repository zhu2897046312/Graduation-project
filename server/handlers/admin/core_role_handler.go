package admin

import (
	"encoding/json"
	"server/models/core"
	"server/service"
	"server/models/common"
	"server/utils"

	"github.com/gin-gonic/gin"
)

type CoreRoleHandler struct {
	service *service.CoreRoleService
}
type RoleCreateRequest struct {
	Name        string  `json:"role_name"`
	Status      int8    `json:"role_status"`
	Permissions []int32 `json:"permission"`
	Remark      string  `json:"remark"`
}

type RoleUpdateRequest struct {
	ID          interface{}   `json:"id"`
	Name        string  `json:"role_name"`
	Status      int8    `json:"role_status"`
	Permissions []int32 `json:"permission"`
	Remark      string  `json:"remark"`
}

func NewCoreRoleHandler(service *service.CoreRoleService) *CoreRoleHandler {
	return &CoreRoleHandler{service: service}
}

// 创建角色
func (h *CoreRoleHandler) CreateRole(c *gin.Context) {
	var req RoleCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}
	rolePermissions, _ := json.Marshal(req.Permissions)
	rolePermissionsJson := json.RawMessage(rolePermissions)
	role := core.CoreRole{
		RoleName:   req.Name,
		RoleStatus: req.Status,
		Permission: rolePermissionsJson,
		Remark:     req.Remark,
	}
	if err := h.service.CreateRole(&role); err != nil {
		Error(c, 10001, err.Error())
		return
	}

	Success(c, role)
}

// 更新角色
func (h *CoreRoleHandler) UpdateRole(c *gin.Context) {
	var req RoleUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}
	roleID := utils.ConvertToUint(req.ID)
	rolePermissions, _ := json.Marshal(req.Permissions)
	rolePermissionsJson := json.RawMessage(rolePermissions)
	role := core.CoreRole{
		ID :         common.MyID(roleID),
		RoleName:   req.Name,
		RoleStatus: req.Status,
		Permission: rolePermissionsJson,
		Remark:     req.Remark,
	}
	if err := h.service.UpdateRole(&role); err != nil {
		Error(c, 10001, err.Error())
		return
	}

	Success(c, role)
}

// 获取角色详情
func (h *CoreRoleHandler) GetRole(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		InvalidParams(c)
		return
	}
	uID := utils.ConvertToUint(id)

	role, err := h.service.GetRoleByID(common.MyID(uID))
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

func (h *CoreRoleHandler) DeleteRole(c *gin.Context) {
	id :=  c.Query("id")
	if id == "" {
		InvalidParams(c)	
		return
	}
	uID := utils.ConvertToUint(id)
	if err := h.service.DeleteRole(common.MyID(uID)); err != nil {
		Error(c, 10009, err.Error())
		return
	}	
	Success(c, nil)
}
