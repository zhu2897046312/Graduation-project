package handlers

import (
	"fmt"
	"server/middleware"
	"server/models/cms"
	"server/models/common"
	"server/service"
	"server/utils"
	"time"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type CmsDocumentHandler struct {
	service *service.CmsDocumentService
	archiveService *service.CmsDocumentArchiveService
}
type ListDocumentRequest struct {
	Title    string `json:"title"`
	Page     int    `json:"page_no"`
	PageSize int    `json:"page_size"`
}

func NewCmsDocumentHandler(service *service.CmsDocumentService, archiveService *service.CmsDocumentArchiveService) *CmsDocumentHandler {
	return &CmsDocumentHandler{
		service: service,
		archiveService: archiveService,
	}
}

type SaveDocumentRequest struct {
	Author         string      `json:"author"`
	Code           string      `json:"code"`
	CategoryID     interface{} `json:"category_id"`
	Cont           string      `json:"cont"`
	DownloadFiles  []string    `json:"download_files"`
	ID             int64       `json:"id"`
	LinkNum        int64       `json:"link_num"`
	LinkType       int64       `json:"link_type"`
	ReadNum        int64       `json:"read_num"`
	SendTime       string      `json:"send_time"`
	SEODescription string      `json:"seo_description"`
	SEOKeyword     string      `json:"seo_keyword"`
	SEOTitle       string      `json:"seo_title"`
	SortNum        int64       `json:"sort_num"`
	Source         string      `json:"source"`
	State          int64       `json:"state"`
	Thumb          string      `json:"thumb"`
	Title          string      `json:"title"`
}

func (h *CmsDocumentHandler) GetDocumentByCode(c *gin.Context) {
	code := c.Query("code")

	documents, err := h.service.GetDocumentByCode(code)
	if err != nil {
		ServerError(c, err)
		return
	}
	myID := common.MyID(uint64(documents.ID))
	archive , err_1 := h.archiveService.GetArchiveByDocumentID(myID)
	if err_1 != nil {
		ServerError(c, err_1)
		return
	}

	Success(c, gin.H{
		"document": documents,
		"cont": archive,
	})
}

func (h *CmsDocumentHandler) ListDocuments(c *gin.Context) {
	var req ListDocumentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}

	documents, total, err := h.service.ListDocuments(req.Page, req.PageSize, req.Title)
	if err != nil {
		ServerError(c, err)
		return
	}

	Success(c, gin.H{
		"list":  documents,
		"total": total,
	})
}

func (h *CmsDocumentHandler) GetAll(c *gin.Context) {
	documents, total, err := h.service.ListDocuments(0,0,"")
	if err != nil {
		ServerError(c, err)
		return
	}

	Success(c, gin.H{
		"list":  documents,
		"total": total,
	})
}

func (h *CmsDocumentHandler) SaveDocument(c *gin.Context) {
	var req SaveDocumentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}

	// 创建文档对象
	document := cms.CmsDocument{
		Title:  req.Title,
		Code:   req.Code,
		Thumb:  req.Thumb,
		State: int8(req.State),
		LinkType: int8(req.LinkType),
		Author: req.Author,
		Source: req.Source,
		AdminID: common.MyID(middleware.GetUserIDFromContext(c)),
		ReadNum: int(req.ReadNum),
		LikeNum: int(req.LinkNum),
		SortNum: int(req.SortNum),
		UpdatedTime: time.Now(),
	}
	archive := cms.CmsDocumentArchive{
		Cont: req.Cont,
		DownloadFiles: nil,
		SeoTitle: req.SEOTitle,
		SeoKeyword: req.SEOKeyword,
		SeoDescription: req.SEODescription,
	}
	if req.ID > 0 {
		err := h.service.UpdateDocument(&document)
		if err != nil {
			ServerError(c, err)
			return
		}
		archive.DocumentID = document.ID
		_, err = h.archiveService.GetArchiveByDocumentID(document.ID)
		
		if err == gorm.ErrRecordNotFound {		
			err = h.archiveService.CreateArchive(&archive)
			if err != nil {
				ServerError(c, err)
				return
			}
		}else{
			err = h.archiveService.UpdateArchive(&archive)
			if err != nil {
				ServerError(c, err)
				return
			}
		}
	}else{
		document.CreatedTime = time.Now()
		h.service.CreateDocument(&document)
		archive.DocumentID = document.ID
		_, err := h.archiveService.GetArchiveByDocumentID(document.ID)
		
		if err == gorm.ErrRecordNotFound {
			err = h.archiveService.CreateArchive(&archive)
			if err != nil {
				ServerError(c, err)
				return
			}
		}else{
			err = h.archiveService.UpdateArchive(&archive)
			if err != nil {
				ServerError(c, err)
				return
			}
		}
	}
	Success(c,nil)
}

func (h *CmsDocumentHandler) DeleteDocument(c *gin.Context) {
	id := c.Query("id")
	uid := utils.ConvertToUint(id)
	if uid <= 0 {
		InvalidParams(c)
		return
	}
	err := h.service.DeleteByID(common.MyID(uid))
	if err != nil {
		ServerError(c, err)
		return
	}
	err = h.archiveService.DeleteByDocumetnID(common.MyID(uid))
	if err != nil {
		fmt.Println(err)
		ServerError(c, err)
		return
	}	
	Success(c, nil)
}