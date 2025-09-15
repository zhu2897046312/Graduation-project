package handlers

import (
	"server/models/cms"
	"server/models/common"
	"server/service"
	"strconv"

	"github.com/gin-gonic/gin"
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
	archive.DocumentID = common.MyID(documentID)

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
	archive.DocumentID = common.MyID(documentID)

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

	archive, err := h.service.GetArchiveByDocumentID(common.MyID(documentID))
	if err != nil {
		Error(c, 3003, "存档不存在")
		return
	}

	Success(c, archive)
}