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
// 树形结构返回结果
type CoreDeptTreeResult struct {
	Label    string               `json:"label"`
	Value    uint                 `json:"value"`
	Node     core.CoreDept        `json:"node"`
	Children []CoreDeptTreeResult `json:"children"`
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

func (h *CoreDeptHandler) Tree(c *gin.Context) {
	pid, err := strconv.ParseUint(c.Query("pid"), 10, 32)
	if err != nil {
		pid = 0 // 默认从顶级分类开始
	}

	// 获取所有分类
	depts,_, err := h.service.GetAll()
	if err != nil {
		Error(c, 22007, "获取分类失败")
		return
	}

	// 构建树形结构
	tree := h.buildTree(depts, uint(pid))

	Success(c, tree)
}

// 递归构建树形结构
func (h *CoreDeptHandler) buildTree(allDepts []*core.CoreDept, pid uint) []CoreDeptTreeResult {
    var tree []CoreDeptTreeResult

    for _, dept := range allDepts {
        // 只检查PID匹配，不进行状态过滤
        if dept.Pid == int64(pid) {
            node := CoreDeptTreeResult{
                Label: dept.DeptName,
                Value: uint(dept.ID),
                Node:  *dept,
            }

            // 递归获取子节点
            children := h.buildTree(allDepts, uint(dept.ID))
            if len(children) > 0 {
                node.Children = children
            } else {
                node.Children = nil // 与Java版本保持一致，空子树设为nil
            }

            tree = append(tree, node)
        }
    }
    return tree
}