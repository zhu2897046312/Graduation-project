package handlers

import (
	"github.com/gin-gonic/gin"
	"server/service"
	"server/models/cms"
	"strconv"
)

type CmsCategoryHandler struct {
	service *service.CmsCategoryService
}

func NewCmsCategoryHandler(service *service.CmsCategoryService) *CmsCategoryHandler {
	return &CmsCategoryHandler{service: service}
}

// 创建分类
func (h *CmsCategoryHandler) CreateCategory(c *gin.Context) {
	var category cms.CmsCategory
	if err := c.ShouldBindJSON(&category); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.CreateCategory(&category); err != nil {
		Error(c, 2001, err.Error())
		return
	}

	Success(c, category)
}

// 更新分类
func (h *CmsCategoryHandler) UpdateCategory(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		InvalidParams(c)
		return
	}

	var category cms.CmsCategory
	if err := c.ShouldBindJSON(&category); err != nil {
		InvalidParams(c)
		return
	}
	category.ID = id

	if err := h.service.UpdateCategory(&category); err != nil {
		Error(c, 2002, err.Error())
		return
	}

	Success(c, category)
}

// 获取分类详情
func (h *CmsCategoryHandler) GetCategory(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		InvalidParams(c)
		return
	}

	category, err := h.service.GetCategoryByID(id)
	if err != nil {
		Error(c, 2003, "分类不存在")
		return
	}

	Success(c, category)
}

// 获取子分类
func (h *CmsCategoryHandler) GetSubCategories(c *gin.Context) {
	parentID, _ := strconv.ParseInt(c.Query("parent_id"), 10, 64)
	if parentID < 0 {
		parentID = 0 // 默认为顶级分类
	}

	categories, err := h.service.GetCategoriesByParentID(parentID)
	if err != nil {
		ServerError(c, err)
		return
	}

	Success(c, categories)
}

// 分页获取分类
func (h *CmsCategoryHandler) ListCategories(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	categories, total, err := h.service.ListCategories(page, pageSize)
	if err != nil {
		ServerError(c, err)
		return
	}

	Success(c, gin.H{
		"list":  categories,
		"total": total,
	})
}

// 更新分类排序
func (h *CmsCategoryHandler) UpdateCategorySort(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		InvalidParams(c)
		return
	}

	var req struct {
		SortNum int `json:"sort_num"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}

	if req.SortNum < 0 {
		Error(c, 2004, "排序值不能为负数")
		return
	}

	if err := h.service.UpdateCategorySort(id, req.SortNum); err != nil {
		Error(c, 2005, err.Error())
		return
	}

	Success(c, nil)
}

// 删除分类
func (h *CmsCategoryHandler) DeleteCategory(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		InvalidParams(c)
		return
	}

	if err := h.service.DeleteCategory(id); err != nil {
		Error(c, 2006, err.Error())
		return
	}

	Success(c, nil)
}