package handlers

import (
	"github.com/gin-gonic/gin"
	"server/service"
	"server/models/cms"
	"strconv"
)

type CmsDocumentVideoHandler struct {
	service *service.CmsDocumentVideoService
}

func NewCmsDocumentVideoHandler(service *service.CmsDocumentVideoService) *CmsDocumentVideoHandler {
	return &CmsDocumentVideoHandler{service: service}
}

// 创建文档视频
func (h *CmsDocumentVideoHandler) CreateVideo(c *gin.Context) {
	documentID, err := strconv.ParseInt(c.Param("document_id"), 10, 64)
	if err != nil || documentID <= 0 {
		InvalidParams(c)
		return
	}

	var video cms.CmsDocumentVideo
	if err := c.ShouldBindJSON(&video); err != nil {
		InvalidParams(c)
		return
	}
	video.DocumentID = documentID

	if err := h.service.CreateDocumentVideo(&video); err != nil {
		Error(c, 5001, err.Error())
		return
	}

	Success(c, video)
}

// 更新文档视频
func (h *CmsDocumentVideoHandler) UpdateVideo(c *gin.Context) {
	documentID, err := strconv.ParseInt(c.Param("document_id"), 10, 64)
	if err != nil || documentID <= 0 {
		InvalidParams(c)
		return
	}

	var video cms.CmsDocumentVideo
	if err := c.ShouldBindJSON(&video); err != nil {
		InvalidParams(c)
		return
	}
	video.DocumentID = documentID

	if err := h.service.UpdateDocumentVideo(&video); err != nil {
		Error(c, 5002, err.Error())
		return
	}

	Success(c, video)
}

// 获取文档视频
func (h *CmsDocumentVideoHandler) GetVideo(c *gin.Context) {
	documentID, err := strconv.ParseInt(c.Param("document_id"), 10, 64)
	if err != nil || documentID <= 0 {
		InvalidParams(c)
		return
	}

	video, err := h.service.GetVideoByDocumentID(documentID)
	if err != nil {
		Error(c, 5003, "视频不存在")
		return
	}

	Success(c, video)
}