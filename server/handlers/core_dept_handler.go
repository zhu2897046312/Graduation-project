package handlers

import (
	"github.com/gin-gonic/gin"
	"server/models/core"
	"server/service"
	"strconv"
)

type CoreDeptHandler struct {
	service *service.CoreDeptService
}

func NewCoreDeptHandler(service *service.CoreDeptService) *CoreDeptHandler {
	return &CoreDeptHandler{service: service}
}

// 创建部门
func (h *CoreDeptHandler) CreateDept(c *gin.Context) {
	var dept core.CoreDept
	if err := c.ShouldBindJSON(&dept); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.CreateDept(&dept); err != nil {
		Error(c, 7001, err.Error())
		return
	}

	Success(c, dept)
}

// 更新部门
func (h *CoreDeptHandler) UpdateDept(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		InvalidParams(c)
		return
	}

	var dept core.CoreDept
	if err := c.ShouldBindJSON(&dept); err != nil {
		InvalidParams(c)
		return
	}
	dept.ID = id

	if err := h.service.UpdateDept(&dept); err != nil {
		Error(c, 7002, err.Error())
		return
	}

	Success(c, dept)
}

// 获取部门详情
func (h *CoreDeptHandler) GetDept(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		InvalidParams(c)
		return
	}

	dept, err := h.service.GetDeptByID(id)
	if err != nil {
		Error(c, 7003, "部门不存在")
		return
	}

	Success(c, dept)
}

// 获取子部门
func (h *CoreDeptHandler) GetSubDepts(c *gin.Context) {
	pid, err := strconv.ParseInt(c.Query("pid"), 10, 64)
	if err != nil {
		pid = 0 // 默认获取顶级部门
	}

	if pid < 0 {
		InvalidParams(c)
		return
	}

	depts, err := h.service.GetSubDepts(pid)
	if err != nil {
		Error(c, 7004, "获取子部门失败")
		return
	}

	Success(c, depts)
}

// 删除部门
func (h *CoreDeptHandler) DeleteDept(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		InvalidParams(c)
		return
	}

	if err := h.service.DeleteDept(id); err != nil {
		Error(c, 7005, err.Error())
		return
	}

	Success(c, nil)
}