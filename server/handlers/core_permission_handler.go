package handlers

import (
	"github.com/gin-gonic/gin"
	"server/models/core"
	"server/service"
	"strconv"
)

type CorePermissionHandler struct {
	service *service.CorePermissionService
}

func NewCorePermissionHandler(service *service.CorePermissionService) *CorePermissionHandler {
	return &CorePermissionHandler{service: service}
}

// 创建权限
func (h *CorePermissionHandler) CreatePermission(c *gin.Context) {
	var permission core.CorePermission
	if err := c.ShouldBindJSON(&permission); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.CreatePermission(&permission); err != nil {
		Error(c, 8001, err.Error())
		return
	}

	Success(c, permission)
}

// 更新权限
func (h *CorePermissionHandler) UpdatePermission(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		InvalidParams(c)
		return
	}

	var permission core.CorePermission
	if err := c.ShouldBindJSON(&permission); err != nil {
		InvalidParams(c)
		return
	}
	permission.ID = id

	if err := h.service.UpdatePermission(&permission); err != nil {
		Error(c, 8002, err.Error())
		return
	}

	Success(c, permission)
}

// 获取权限详情
func (h *CorePermissionHandler) GetPermission(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		InvalidParams(c)
		return
	}

	permission, err := h.service.GetPermissionByID(id)
	if err != nil {
		Error(c, 8003, "权限不存在")
		return
	}

	Success(c, permission)
}

// 根据Code获取权限
func (h *CorePermissionHandler) GetPermissionByCode(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		InvalidParams(c)
		return
	}

	permission, err := h.service.GetPermissionByCode(code)
	if err != nil {
		Error(c, 8004, "权限不存在")
		return
	}

	Success(c, permission)
}