package client

import (
	"server/models/common"
	"server/service"
	"server/utils"

	"github.com/gin-gonic/gin"
)

type ClientCmsDocumentHandler struct {
	service *service.CmsDocumentService
	archiveService *service.CmsDocumentArchiveService
}
type ListDocumentRequest struct {
	Title    string `json:"title"`
	Page     int    `json:"page_no"`
	PageSize int    `json:"page_size"`
}

func NewClientCmsDocumentHandler(service *service.CmsDocumentService, archiveService *service.CmsDocumentArchiveService) *ClientCmsDocumentHandler {
	return &ClientCmsDocumentHandler{
		service: service,
		archiveService: archiveService,
	}
}

func (h *ClientCmsDocumentHandler) GetDocumentByCode(c *gin.Context) {
	code := c.Query("code")

	documents, err := h.service.GetDocumentByCode(code)
	if err != nil {
		utils.ServerError(c, err)
		return
	}
	myID := common.MyID(uint64(documents.ID))
	archive , err_1 := h.archiveService.GetArchiveByDocumentID(myID)
	if err_1 != nil {
		utils.ServerError(c, err_1)
		return
	}

	utils.Success(c, gin.H{
		"document": documents,
		"cont": archive,
	})
}

func (h *ClientCmsDocumentHandler) GetAll(c *gin.Context) {
	documents, total, err := h.service.ListDocuments(0,0,"")
	if err != nil {
		utils.ServerError(c, err)
		return
	}

	utils.Success(c, gin.H{
		"list":  documents,
		"total": total,
	})
}
