package handlers

import (
	"github.com/gin-gonic/gin"
	"server/models/shop"
	"server/service"
	"strconv"
)

type ShopTagMateHandler struct {
	service *service.ShopTagMateService
}

func NewShopTagMateHandler(service *service.ShopTagMateService) *ShopTagMateHandler {
	return &ShopTagMateHandler{service: service}
}

// 创建标签元数据
func (h *ShopTagMateHandler) CreateTagMate(c *gin.Context) {
	var mate shop.ShopTagMate
	if err := c.ShouldBindJSON(&mate); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.CreateTagMate(&mate); err != nil {
		Error(c, 20001, err.Error())
		return
	}

	Success(c, mate)
}

// 更新标签元数据
func (h *ShopTagMateHandler) UpdateTagMate(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		InvalidParams(c)
		return
	}

	var mate shop.ShopTagMate
	if err := c.ShouldBindJSON(&mate); err != nil {
		InvalidParams(c)
		return
	}
	mate.ID = id

	if err := h.service.UpdateTagMate(&mate); err != nil {
		Error(c, 20002, err.Error())
		return
	}

	Success(c, mate)
}

// 获取标签元数据
func (h *ShopTagMateHandler) GetTagMate(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		InvalidParams(c)
		return
	}

	mate, err := h.service.GetTagMateByID(id)
	if err != nil {
		Error(c, 20003, "元数据不存在")
		return
	}

	Success(c, mate)
}

// 更新标签SEO信息
func (h *ShopTagMateHandler) UpdateTagSEO(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
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

	if req.Title == "" && req.Keywords == "" && req.Description == "" {
		Error(c, 20004, "至少需要更新一个SEO字段")
		return
	}

	if err := h.service.UpdateTagSEO(id, req.Title, req.Keywords, req.Description); err != nil {
		Error(c, 20005, err.Error())
		return
	}

	Success(c, nil)
}

// 更新标签内容
func (h *ShopTagMateHandler) UpdateTagContent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
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
		Error(c, 20006, "内容不能为空")
		return
	}

	if err := h.service.UpdateTagContent(id, req.Content); err != nil {
		Error(c, 20007, err.Error())
		return
	}

	Success(c, nil)
}