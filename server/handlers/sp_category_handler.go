package handlers

import (
	"github.com/gin-gonic/gin"
	"server/models/sp"
	"server/service"
	"strconv"
)

type SpCategoryHandler struct {
	service *service.SpCategoryService
}

// 树形结构返回结果
type CategoryTreeResult struct {
	Label    string               `json:"label"`
	Value    uint                 `json:"value"`
	Node     sp.SpCategory        `json:"node"`
	Children []CategoryTreeResult `json:"children"`
}

func NewSpCategoryHandler(service *service.SpCategoryService) *SpCategoryHandler {
	return &SpCategoryHandler{service: service}
}

// 创建分类
func (h *SpCategoryHandler) CreateCategory(c *gin.Context) {
	var category sp.SpCategory
	if err := c.ShouldBindJSON(&category); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.CreateCategory(&category); err != nil {
		Error(c, 22001, err.Error())
		return
	}

	Success(c, category)
}

// 更新分类
func (h *SpCategoryHandler) UpdateCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil || id == 0 {
		InvalidParams(c)
		return
	}

	var category sp.SpCategory
	if err := c.ShouldBindJSON(&category); err != nil {
		InvalidParams(c)
		return
	}
	category.ID = uint(id)

	if err := h.service.UpdateCategory(&category); err != nil {
		Error(c, 22002, err.Error())
		return
	}

	Success(c, category)
}

// 获取分类详情
func (h *SpCategoryHandler) GetCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil || id == 0 {
		InvalidParams(c)
		return
	}

	category, err := h.service.GetCategoryByID(uint(id))
	if err != nil {
		Error(c, 22003, "分类不存在")
		return
	}

	Success(c, category)
}

// 获取子分类
func (h *SpCategoryHandler) GetSubCategories(c *gin.Context) {
	pid, err := strconv.ParseUint(c.Query("pid"), 10, 32)
	if err != nil {
		pid = 0 // 默认获取顶级分类
	}

	categories, err := h.service.GetCategoriesByPid(uint(pid))
	if err != nil {
		Error(c, 22004, "获取子分类失败")
		return
	}

	Success(c, categories)
}

// 获取分类树
func (h *SpCategoryHandler) GetCategoryTree(c *gin.Context) {
	pid, err := strconv.ParseUint(c.Query("pid"), 10, 32)
	if err != nil {
		pid = 0 // 默认从顶级分类开始
	}

	var state *uint8
	if stateStr := c.Query("state"); stateStr != "" {
		if stateVal, err := strconv.ParseUint(stateStr, 10, 8); err == nil {
			stateUint8 := uint8(stateVal)
			state = &stateUint8
		}
	}

	// 获取所有分类
	categories, err := h.service.GetAllCategories()
	if err != nil {
		Error(c, 22007, "获取分类失败")
		return
	}

	// 构建树形结构
	tree := h.buildTree(categories, uint(pid), state)

	Success(c, tree)
}

// 递归构建树形结构
func (h *SpCategoryHandler) buildTree(categories []*sp.SpCategory, pid uint, state *uint8) []CategoryTreeResult {
	var tree []CategoryTreeResult
	
	for _, category := range categories {
		if category.Pid == pid && (state == nil || category.State == *state) {
			treeResult := CategoryTreeResult{
				Label: category.Title,
				Value: category.ID,
				Node:  *category,
			}
			
			// 递归获取子节点
			children := h.buildTree(categories, category.ID, state)
			if len(children) > 0 {
				treeResult.Children = children
			} else {
				treeResult.Children = nil
			}
			
			tree = append(tree, treeResult)
		}
	}
	
	return tree
}

// 更新分类状态
func (h *SpCategoryHandler) UpdateCategoryState(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil || id == 0 {
		InvalidParams(c)
		return
	}

	var req struct {
		State uint8 `json:"state"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.UpdateCategoryState(uint(id), req.State); err != nil {
		Error(c, 22005, err.Error())
		return
	}

	Success(c, nil)
}

// 更新分类排序
func (h *SpCategoryHandler) UpdateCategorySortNum(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil || id == 0 {
		InvalidParams(c)
		return
	}

	var req struct {
		SortNum uint16 `json:"sort_num"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.UpdateCategorySortNum(uint(id), req.SortNum); err != nil {
		Error(c, 22006, err.Error())
		return
	}

	Success(c, nil)
}