package handlers

import (
	"github.com/gin-gonic/gin"
	"server/service"
	"server/models/cms"
	"strconv"
)

type CmsTagHandler struct {
	service *service.CmsTagService
}

func NewCmsTagHandler(service *service.CmsTagService) *CmsTagHandler {
	return &CmsTagHandler{service: service}
}

// 创建标签
func (h *CmsTagHandler) CreateTag(c *gin.Context) {
	var tag cms.CmsTag
	if err := c.ShouldBindJSON(&tag); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.CreateTag(&tag); err != nil {
		Error(c, 10001, err.Error())
		return
	}

	Success(c, tag)
}

// 更新标签
func (h *CmsTagHandler) UpdateTag(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		InvalidParams(c)
		return
	}

	var tag cms.CmsTag
	if err := c.ShouldBindJSON(&tag); err != nil {
		InvalidParams(c)
		return
	}
	tag.ID = id

	if err := h.service.UpdateTag(&tag); err != nil {
		Error(c, 10002, err.Error())
		return
	}

	Success(c, tag)
}

// 获取标签详情
func (h *CmsTagHandler) GetTag(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		InvalidParams(c)
		return
	}

	tag, err := h.service.GetTagByID(id)
	if err != nil {
		Error(c, 10003, "标签不存在")
		return
	}

	Success(c, tag)
}

// 根据状态获取标签
func (h *CmsTagHandler) GetByState(c *gin.Context) {
	state, err := strconv.Atoi(c.Query("state"))
	if err != nil {
		InvalidParams(c)
		return
	}

	tags, err := h.service.GetTagsByState(int8(state))
	if err != nil {
		ServerError(c, err)
		return
	}

	Success(c, tags)
}

// 增加标签阅读量
func (h *CmsTagHandler) IncrementReadNum(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		InvalidParams(c)
		return
	}

	if err := h.service.IncrementReadNum(id); err != nil {
		Error(c, 10004, err.Error())
		return
	}

	Success(c, nil)
}

// 搜索标签
func (h *CmsTagHandler) SearchTags(c *gin.Context) {
	keyword := c.Query("keyword")
	if keyword == "" {
		InvalidParams(c)
		return
	}

	tags, err := h.service.SearchTags(keyword)
	if err != nil {
		ServerError(c, err)
		return
	}

	Success(c, tags)
}