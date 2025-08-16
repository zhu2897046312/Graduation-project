package handlers

import (
	"github.com/gin-gonic/gin"
	"server/models/core"
	"server/service"
	"strconv"
)

type CoreAdminHandler struct {
	service *service.CoreAdminService
}

func NewCoreAdminHandler(service *service.CoreAdminService) *CoreAdminHandler {
	return &CoreAdminHandler{service: service}
}

// 创建管理员
func (h *CoreAdminHandler) CreateAdmin(c *gin.Context) {
	var admin core.CoreAdmin
	if err := c.ShouldBindJSON(&admin); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.CreateAdmin(&admin); err != nil {
		Error(c, 5001, err.Error())
		return
	}

	Success(c, admin)
}

// 更新管理员
func (h *CoreAdminHandler) UpdateAdmin(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		InvalidParams(c)
		return
	}

	var admin core.CoreAdmin
	if err := c.ShouldBindJSON(&admin); err != nil {
		InvalidParams(c)
		return
	}
	admin.ID = id

	if err := h.service.UpdateAdmin(&admin); err != nil {
		Error(c, 5002, err.Error())
		return
	}

	Success(c, admin)
}

// 获取管理员详情
func (h *CoreAdminHandler) GetAdmin(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		InvalidParams(c)
		return
	}

	admin, err := h.service.GetAdminByID(id)
	if err != nil {
		Error(c, 5003, "管理员不存在")
		return
	}

	Success(c, admin)
}

// 更新管理员状态
func (h *CoreAdminHandler) UpdateAdminStatus(c *gin.Context) {
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

	if err := h.service.UpdateAdminStatus(id, req.Status); err != nil {
		Error(c, 5004, err.Error())
		return
	}

	Success(c, nil)
}

// 更新管理员密码
func (h *CoreAdminHandler) UpdateAdminPassword(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		InvalidParams(c)
		return
	}

	var req struct {
		NewPassword string `json:"new_password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.UpdateAdminPassword(id, req.NewPassword); err != nil {
		Error(c, 5005, err.Error())
		return
	}

	Success(c, nil)
}