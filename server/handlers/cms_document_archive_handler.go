package handlers

import (
	"github.com/gin-gonic/gin"
	"server/service"
	"server/models/cms"
	"strconv"
)

type CmsDocumentArchiveHandler struct {
	service *service.CmsDocumentArchiveService
}

func NewCmsDocumentArchiveHandler(service *service.CmsDocumentArchiveService) *CmsDocumentArchiveHandler {
	return &CmsDocumentArchiveHandler{service: service}
}

// 创建文档存档
func (h *CmsDocumentArchiveHandler) CreateArchive(c *gin.Context) {
	documentID, err := strconv.ParseInt(c.Param("document_id"), 10, 64)
	if err != nil || documentID <= 0 {
		InvalidParams(c)
		return
	}

	var archive cms.CmsDocumentArchive
	if err := c.ShouldBindJSON(&archive); err != nil {
		InvalidParams(c)
		return
	}
	archive.DocumentID = documentID

	if err := h.service.CreateArchive(&archive); err != nil {
		Error(c, 3001, err.Error())
		return
	}

	Success(c, archive)
}

// 更新文档存档
func (h *CmsDocumentArchiveHandler) UpdateArchive(c *gin.Context) {
	documentID, err := strconv.ParseInt(c.Param("document_id"), 10, 64)
	if err != nil || documentID <= 0 {
		InvalidParams(c)
		return
	}

	var archive cms.CmsDocumentArchive
	if err := c.ShouldBindJSON(&archive); err != nil {
		InvalidParams(c)
		return
	}
	archive.DocumentID = documentID

	if err := h.service.UpdateArchive(&archive); err != nil {
		Error(c, 3002, err.Error())
		return
	}

	Success(c, archive)
}

// 获取文档存档
func (h *CmsDocumentArchiveHandler) GetArchive(c *gin.Context) {
	documentID, err := strconv.ParseInt(c.Param("document_id"), 10, 64)
	if err != nil || documentID <= 0 {
		InvalidParams(c)
		return
	}

	archive, err := h.service.GetArchiveByDocumentID(documentID)
	if err != nil {
		Error(c, 3003, "存档不存在")
		return
	}

	Success(c, archive)
}