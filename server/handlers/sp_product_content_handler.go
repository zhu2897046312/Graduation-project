package handlers

import (
	"github.com/gin-gonic/gin"
	"server/models/sp"
	"server/service"
	"strconv"
)

type SpProductContentHandler struct {
	service *service.SpProductContentService
}

func NewSpProductContentHandler(service *service.SpProductContentService) *SpProductContentHandler {
	return &SpProductContentHandler{service: service}
}

// 创建商品内容
func (h *SpProductContentHandler) CreateContent(c *gin.Context) {
	var content sp.SpProductContent
	if err := c.ShouldBindJSON(&content); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.CreateContent(&content); err != nil {
		Error(c, 30001, err.Error())
		return
	}

	Success(c, content)
}

// 更新商品内容
func (h *SpProductContentHandler) UpdateContent(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil || id == 0 {
		InvalidParams(c)
		return
	}

	var content sp.SpProductContent
	if err := c.ShouldBindJSON(&content); err != nil {
		InvalidParams(c)
		return
	}
	content.ID = uint(id)

	if err := h.service.UpdateContent(&content); err != nil {
		Error(c, 30002, err.Error())
		return
	}

	Success(c, content)
}

// 根据商品ID获取内容
func (h *SpProductContentHandler) GetContentByProduct(c *gin.Context) {
	productID, err := strconv.ParseUint(c.Param("product_id"), 10, 32)
	if err != nil || productID == 0 {
		InvalidParams(c)
		return
	}

	content, err := h.service.GetContentByProductID(uint(productID))
	if err != nil {
		Error(c, 30003, "商品内容不存在")
		return
	}

	Success(c, content)
}

// 更新SEO信息
func (h *SpProductContentHandler) UpdateSEO(c *gin.Context) {
	productID, err := strconv.ParseUint(c.Param("product_id"), 10, 32)
	if err != nil || productID == 0 {
		InvalidParams(c)
		return
	}

	var req struct {
		Title       string `json:"title"`
		Keywords    string `json:"keywords"`
		Description string `json:"description"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}

	if req.Title == "" || req.Keywords == "" || req.Description == "" {
		Error(c, 30004, "SEO信息不能为空")
		return
	}

	if err := h.service.UpdateSEO(uint(productID), req.Title, req.Keywords, req.Description); err != nil {
		Error(c, 30005, err.Error())
		return
	}

	Success(c, nil)
}

// 更新商品内容文本
func (h *SpProductContentHandler) UpdateContentText(c *gin.Context) {
	productID, err := strconv.ParseUint(c.Param("product_id"), 10, 32)
	if err != nil || productID == 0 {
		InvalidParams(c)
		return
	}

	var req struct {
		Content string `json:"content"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}

	if req.Content == "" {
		Error(c, 30006, "内容不能为空")
		return
	}

	if err := h.service.UpdateContentText(uint(productID), req.Content); err != nil {
		Error(c, 30007, err.Error())
		return
	}

	Success(c, nil)
}