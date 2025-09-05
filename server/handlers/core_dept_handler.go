package handlers

import (
	"encoding/json"
	"server/models/core"
	"server/service"
	"server/utils"
	"strconv"

	"github.com/gin-gonic/gin"
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

type DeptCreateRequest struct {
	DeptName       string   `json:"dept_name"`
	ConnectName    string   `json:"connect_name"`
	ConnectMobile  string   `json:"connect_mobile"`
	ConnectAddress string   `json:"connect_address"`
	LeaderName     string   `json:"leader_name"`
	Thumb          string   `json:"thumb"`
	Content        string   `json:"content"`
	OrganizeInfo   []string `json:"organize"`
	Level          int8     `json:"level"`
	SortNum        int8     `json:"sort_num"`
	Remark         string   `json:"remark"`
	Pid            int64    `json:"pid"`
}

type DeptUpdateRequest struct {
	ID             int64    `json:"id"`
	DeptName       string   `json:"dept_name"`
	ConnectName    string   `json:"connect_name"`
	ConnectMobile  string   `json:"connect_mobile"`
	ConnectAddress string   `json:"connect_address"`
	LeaderName     string   `json:"leader_name"`
	Thumb          string   `json:"thumb"`
	Content        string   `json:"content"`
	OrganizeInfo   []string `json:"organize"`
	Level          int8     `json:"level"`
	SortNum        int8     `json:"sort_num"`
	Remark         string   `json:"remark"`
	Pid            int64    `json:"pid"`
}

func NewCoreDeptHandler(service *service.CoreDeptService) *CoreDeptHandler {
	return &CoreDeptHandler{service: service}
}

// 创建部门
func (h *CoreDeptHandler) CreateDept(c *gin.Context) {
	var req DeptCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}

	if req.Pid > 0{
		subDepts,err:=h.service.GetSubDepts(req.Pid)
		if err == nil && len(subDepts) <= 0 {
			Error(c, 7000, "父级部门不存在")
			return
		}
	}

	organize,_ := json.Marshal(req.OrganizeInfo)
	organizeJson := json.RawMessage(organize)
	dept := core.CoreDept{
		Pid: req.Pid,
		DeptName: req.DeptName,
		ConnectName: req.ConnectName,
		ConnectMobile: req.ConnectMobile,
		ConnectAddress: req.ConnectAddress,
		Organize: organizeJson,
		Level: int16(req.Level),
		SortNum: int(req.SortNum),
		Remark: req.Remark,
	}
	if err := h.service.CreateDept(&dept); err != nil {
		Error(c, 7001, err.Error())
		return
	}

	Success(c, dept)
}

// 更新部门
func (h *CoreDeptHandler) UpdateDept(c *gin.Context) {
	var req DeptUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}

	if req.Pid > 0{
		subDepts,err:=h.service.GetSubDepts(req.Pid)
		if err == nil && len(subDepts) <= 0 {
			Error(c, 7000, "父级部门不存在")
			return
		}
	}

	organize,_ := json.Marshal(req.OrganizeInfo)
	organizeJson := json.RawMessage(organize)
	dept := core.CoreDept{
		ID : req.ID,
		Pid: req.Pid,
		DeptName: req.DeptName,
		ConnectName: req.ConnectName,
		ConnectMobile: req.ConnectMobile,
		ConnectAddress: req.ConnectAddress,
		Organize: organizeJson,
		Level: int16(req.Level),
		SortNum: int(req.SortNum),
		Remark: req.Remark,
	}
	if err := h.service.UpdateDept(&dept); err != nil {
		Error(c, 7001, err.Error())
		return
	}

	Success(c, dept)
}

// 获取部门详情
func (h *CoreDeptHandler) GetDept(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		InvalidParams(c)
		return
	}
	uID := utils.ConvertToUint(id)
	if uID == 0 {
		InvalidParams(c)
		return
	}
	dept, err := h.service.GetDeptByID(int64(uID))
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
	var req DeptUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}
	uID := utils.ConvertToUint(req.ID)
	if uID == 0 {
		InvalidParams(c)
		return
	}
	id := int64(uID)

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
	depts, _, err := h.service.GetAll()
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
