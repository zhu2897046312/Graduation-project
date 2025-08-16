package handlers

import (
	"github.com/gin-gonic/gin"
	"server/service"
	"server/models/cms"
	"strconv"
)

type CmsDocumentTagHandler struct {
	service *service.CmsDocumentTagService
}

func NewCmsDocumentTagHandler(service *service.CmsDocumentTagService) *CmsDocumentTagHandler {
	return &CmsDocumentTagHandler{service: service}
}

// 创建文档标签关联
func (h *CmsDocumentTagHandler) CreateDocumentTag(c *gin.Context) {
	var docTag cms.CmsDocumentTag
	if err := c.ShouldBindJSON(&docTag); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.CreateDocumentTag(&docTag); err != nil {
		Error(c, 4001, err.Error())
		return
	}

	Success(c, docTag)
}

// 删除文档标签关联
func (h *CmsDocumentTagHandler) DeleteDocumentTag(c *gin.Context) {
	documentID, err := strconv.ParseInt(c.Param("document_id"), 10, 64)
	if err != nil || documentID <= 0 {
		InvalidParams(c)
		return
	}

	tagID, err := strconv.ParseInt(c.Param("tag_id"), 10, 64)
	if err != nil || tagID <= 0 {
		InvalidParams(c)
		return
	}

	if err := h.service.DeleteDocumentTag(documentID, tagID); err != nil {
		Error(c, 4002, err.Error())
		return
	}

	Success(c, nil)
}

// 根据文档ID获取标签
func (h *CmsDocumentTagHandler) GetTagsByDocument(c *gin.Context) {
	documentID, err := strconv.ParseInt(c.Param("document_id"), 10, 64)
	if err != nil || documentID <= 0 {
		InvalidParams(c)
		return
	}

	tags, err := h.service.GetTagsByDocumentID(documentID)
	if err != nil {
		ServerError(c, err)
		return
	}

	Success(c, tags)
}

// 根据标签ID获取文档
func (h *CmsDocumentTagHandler) GetDocumentsByTag(c *gin.Context) {
	tagID, err := strconv.ParseInt(c.Param("tag_id"), 10, 64)
	if err != nil || tagID <= 0 {
		InvalidParams(c)
		return
	}

	docTags, err := h.service.GetDocumentsByTagID(tagID)
	if err != nil {
		ServerError(c, err)
		return
	}

	Success(c, docTags)
}