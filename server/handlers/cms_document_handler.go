package handlers

import (
	"github.com/gin-gonic/gin"
	"server/service"
	"strconv"
)

type CmsDocumentHandler struct {
	service *service.CmsDocumentService
}

func NewCmsDocumentHandler(service *service.CmsDocumentService) *CmsDocumentHandler {
	return &CmsDocumentHandler{service: service}
}

// 根据分类ID获取文档
func (h *CmsDocumentHandler) GetByCategoryID(c *gin.Context) {
	categoryID, err := strconv.ParseInt(c.Param("category_id"), 10, 64)
	if err != nil || categoryID <= 0 {
		InvalidParams(c)
		return
	}

	documents, err := h.service.GetDocumentsByCategoryID(categoryID)
	if err != nil {
		ServerError(c, err)
		return
	}

	Success(c, documents)
}

// 获取热门文档
func (h *CmsDocumentHandler) GetPopular(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if limit <= 0 || limit > 100 {
		limit = 10
	}

	documents, err := h.service.GetPopularDocuments(limit)
	if err != nil {
		ServerError(c, err)
		return
	}

	Success(c, documents)
}

// 分页获取文档
func (h *CmsDocumentHandler) ListDocuments(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	documents, total, err := h.service.ListDocuments(page, pageSize)
	if err != nil {
		ServerError(c, err)
		return
	}

	Success(c, gin.H{
		"list":  documents,
		"total": total,
	})
}