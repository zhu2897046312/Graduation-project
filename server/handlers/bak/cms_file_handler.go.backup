package handlers

import (
	"github.com/gin-gonic/gin"
	"server/service"
	"server/models/cms"
	"strconv"
)

type CmsFileHandler struct {
	service *service.CmsFileService
}

func NewCmsFileHandler(service *service.CmsFileService) *CmsFileHandler {
	return &CmsFileHandler{service: service}
}

// 创建文件记录
func (h *CmsFileHandler) CreateFile(c *gin.Context) {
	var file cms.CmsFile
	if err := c.ShouldBindJSON(&file); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.CreateFile(&file); err != nil {
		Error(c, 6001, err.Error())
		return
	}

	Success(c, file)
}

// 更新文件记录
func (h *CmsFileHandler) UpdateFile(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		InvalidParams(c)
		return
	}

	var file cms.CmsFile
	if err := c.ShouldBindJSON(&file); err != nil {
		InvalidParams(c)
		return
	}
	file.ID = id

	if err := h.service.UpdateFile(&file); err != nil {
		Error(c, 6002, err.Error())
		return
	}

	Success(c, file)
}

// 获取文件详情
func (h *CmsFileHandler) GetFile(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		InvalidParams(c)
		return
	}

	file, err := h.service.GetFileByID(id)
	if err != nil {
		Error(c, 6003, "文件不存在")
		return
	}

	Success(c, file)
}

// 根据MD5获取文件
func (h *CmsFileHandler) GetFileByMD5(c *gin.Context) {
	md5 := c.Query("md5")
	if md5 == "" {
		InvalidParams(c)
		return
	}

	file, err := h.service.GetFileByMD5(md5)
	if err != nil {
		Error(c, 6004, "文件不存在")
		return
	}

	Success(c, file)
}