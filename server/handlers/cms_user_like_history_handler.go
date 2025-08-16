package handlers

import (
	"github.com/gin-gonic/gin"
	"server/models/cms"
	"server/service"
	"strconv"
)

type CmsUserLikeHistoryHandler struct {
	service *service.CmsUserLikeHistoryService
}

func NewCmsUserLikeHistoryHandler(service *service.CmsUserLikeHistoryService) *CmsUserLikeHistoryHandler {
	return &CmsUserLikeHistoryHandler{service: service}
}

// 创建点赞记录
func (h *CmsUserLikeHistoryHandler) CreateLikeHistory(c *gin.Context) {
	var history cms.CmsUserLikeHistory
	if err := c.ShouldBindJSON(&history); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.CreateLikeHistory(&history); err != nil {
		Error(c, 3001, err.Error())
		return
	}

	Success(c, history)
}

// 更新点赞记录
func (h *CmsUserLikeHistoryHandler) UpdateLikeHistory(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		InvalidParams(c)
		return
	}

	var history cms.CmsUserLikeHistory
	if err := c.ShouldBindJSON(&history); err != nil {
		InvalidParams(c)
		return
	}
	history.ID = id

	if err := h.service.UpdateLikeHistory(&history); err != nil {
		Error(c, 3002, err.Error())
		return
	}

	Success(c, history)
}

// 获取用户点赞记录
func (h *CmsUserLikeHistoryHandler) GetLikeHistoryByUser(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil || userID <= 0 {
		InvalidParams(c)
		return
	}

	histories, err := h.service.GetLikeHistoryByUserID(userID)
	if err != nil {
		Error(c, 3003, "获取点赞记录失败")
		return
	}

	Success(c, histories)
}

// 检查用户是否点赞
func (h *CmsUserLikeHistoryHandler) CheckUserLiked(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil || userID <= 0 {
		InvalidParams(c)
		return
	}

	documentID, err := strconv.ParseInt(c.Query("document_id"), 10, 64)
	if err != nil || documentID <= 0 {
		InvalidParams(c)
		return
	}

	liked, err := h.service.CheckUserLikedDocument(userID, documentID)
	if err != nil {
		Error(c, 3004, "检查点赞状态失败")
		return
	}

	Success(c, gin.H{"liked": liked})
}

// 获取文档点赞数
func (h *CmsUserLikeHistoryHandler) GetLikeCount(c *gin.Context) {
	documentID, err := strconv.ParseInt(c.Param("document_id"), 10, 64)
	if err != nil || documentID <= 0 {
		InvalidParams(c)
		return
	}

	count, err := h.service.GetLikeCountByDocumentID(documentID)
	if err != nil {
		Error(c, 3005, "获取点赞数失败")
		return
	}

	Success(c, gin.H{"count": count})
}