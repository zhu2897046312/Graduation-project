package handlers

import (
	"github.com/gin-gonic/gin"
	"server/models/shop"
	"server/service"
	"strconv"
)

type ShopTagHandler struct {
	service *service.ShopTagService
}

func NewShopTagHandler(service *service.ShopTagService) *ShopTagHandler {
	return &ShopTagHandler{service: service}
}

// 创建商品标签
func (h *ShopTagHandler) CreateTag(c *gin.Context) {
	var tag shop.ShopTag
	if err := c.ShouldBindJSON(&tag); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.CreateTag(&tag); err != nil {
		Error(c, 21001, err.Error())
		return
	}

	Success(c, tag)
}

// 更新商品标签
func (h *ShopTagHandler) UpdateTag(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		InvalidParams(c)
		return
	}

	var tag shop.ShopTag
	if err := c.ShouldBindJSON(&tag); err != nil {
		InvalidParams(c)
		return
	}
	tag.ID = id

	if err := h.service.UpdateTag(&tag); err != nil {
		Error(c, 21002, err.Error())
		return
	}

	Success(c, tag)
}

// 获取标签详情
func (h *ShopTagHandler) GetTag(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		InvalidParams(c)
		return
	}

	tag, err := h.service.GetTagByID(id)
	if err != nil {
		Error(c, 21003, "标签不存在")
		return
	}

	Success(c, tag)
}

// 根据状态获取标签列表
func (h *ShopTagHandler) GetTagsByState(c *gin.Context) {
	state, err := strconv.ParseInt(c.Query("state"), 10, 8)
	if err != nil {
		InvalidParams(c)
		return
	}

	tags, err := h.service.GetTagsByState(int8(state))
	if err != nil {
		Error(c, 21004, "获取标签列表失败")
		return
	}

	Success(c, tags)
}

// 根据匹配词搜索标签
func (h *ShopTagHandler) SearchTags(c *gin.Context) {
	matchWord := c.Query("match_word")
	if matchWord == "" {
		InvalidParams(c)
		return
	}

	tags, err := h.service.SearchTagsByMatchWord(matchWord)
	if err != nil {
		Error(c, 21005, "搜索标签失败")
		return
	}

	Success(c, tags)
}

// 增加标签阅读量
func (h *ShopTagHandler) IncrementTagReadNum(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		InvalidParams(c)
		return
	}

	if err := h.service.IncrementTagReadNum(id); err != nil {
		Error(c, 21006, err.Error())
		return
	}

	Success(c, nil)
}

// 分页获取标签
func (h *ShopTagHandler) ListTags(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	tags, total, err := h.service.ListTags(page, pageSize)
	if err != nil {
		Error(c, 21007, "获取标签列表失败")
		return
	}

	Success(c, gin.H{
		"list":  tags,
		"total": total,
	})
}