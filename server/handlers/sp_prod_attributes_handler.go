package handlers

import (
	"github.com/gin-gonic/gin"
	"server/models/sp"
	"server/service"
	"strconv"
)

type SpProdAttributesHandler struct {
	service *service.SpProdAttributesService
}

func NewSpProdAttributesHandler(service *service.SpProdAttributesService) *SpProdAttributesHandler {
	return &SpProdAttributesHandler{service: service}
}

// 创建商品属性
func (h *SpProdAttributesHandler) CreateAttribute(c *gin.Context) {
	var attr sp.SpProdAttributes
	if err := c.ShouldBindJSON(&attr); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.CreateAttribute(&attr); err != nil {
		Error(c, 28001, err.Error())
		return
	}

	Success(c, attr)
}

// 更新商品属性
func (h *SpProdAttributesHandler) UpdateAttribute(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil || id == 0 {
		InvalidParams(c)
		return
	}

	var attr sp.SpProdAttributes
	if err := c.ShouldBindJSON(&attr); err != nil {
		InvalidParams(c)
		return
	}
	attr.ID = uint(id)

	if err := h.service.UpdateAttribute(&attr); err != nil {
		Error(c, 28002, err.Error())
		return
	}

	Success(c, attr)
}

// 获取属性详情
func (h *SpProdAttributesHandler) GetAttribute(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil || id == 0 {
		InvalidParams(c)
		return
	}

	attr, err := h.service.GetAttributeByID(uint(id))
	if err != nil {
		Error(c, 28003, "属性不存在")
		return
	}

	Success(c, attr)
}

// 获取所有属性
func (h *SpProdAttributesHandler) GetAllAttributes(c *gin.Context) {
	attrs, err := h.service.GetAllAttributes()
	if err != nil {
		Error(c, 28004, "获取属性列表失败")
		return
	}

	Success(c, attrs)
}

// 更新属性排序
func (h *SpProdAttributesHandler) UpdateAttributeSortNum(c *gin.Context) {
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

	if err := h.service.UpdateAttributeSortNum(uint(id), req.SortNum); err != nil {
		Error(c, 28005, err.Error())
		return
	}

	Success(c, nil)
}

// 删除属性
func (h *SpProdAttributesHandler) DeleteAttribute(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil || id == 0 {
		InvalidParams(c)
		return
	}

	if err := h.service.DeleteAttribute(uint(id)); err != nil {
		Error(c, 28006, err.Error())
		return
	}

	Success(c, nil)
}