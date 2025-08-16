package handlers

import (
	"github.com/gin-gonic/gin"
	"server/service"
	"strconv"
)

type CmsCommentHandler struct {
	service *service.CmsCommentService
}

func NewCmsCommentHandler(service *service.CmsCommentService) *CmsCommentHandler {
	return &CmsCommentHandler{service: service}
}

// 根据文档ID获取评论
func (h *CmsCommentHandler) GetByDocumentID(c *gin.Context) {
	documentID, err := strconv.ParseInt(c.Param("document_id"), 10, 64)
	if err != nil || documentID <= 0 {
		InvalidParams(c)
		return
	}

	comments, err := h.service.GetCommentsByDocumentID(documentID)
	if err != nil {
		ServerError(c, err)
		return
	}

	Success(c, comments)
}

// 根据用户ID获取评论
func (h *CmsCommentHandler) GetByUserID(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil || userID <= 0 {
		InvalidParams(c)
		return
	}

	comments, err := h.service.GetCommentsByUserID(userID)
	if err != nil {
		ServerError(c, err)
		return
	}

	Success(c, comments)
}

// 获取顶级评论
func (h *CmsCommentHandler) GetTopLevel(c *gin.Context) {
	comments, err := h.service.GetTopLevelComments()
	if err != nil {
		ServerError(c, err)
		return
	}

	Success(c, comments)
}

// 获取评论回复
func (h *CmsCommentHandler) GetReplies(c *gin.Context) {
	commentID, err := strconv.ParseInt(c.Param("comment_id"), 10, 64)
	if err != nil || commentID <= 0 {
		InvalidParams(c)
		return
	}

	replies, err := h.service.GetCommentReplies(commentID)
	if err != nil {
		ServerError(c, err)
		return
	}

	Success(c, replies)
}

// 分页获取评论
func (h *CmsCommentHandler) ListComments(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	comments, total, err := h.service.ListComments(page, pageSize)
	if err != nil {
		ServerError(c, err)
		return
	}

	Success(c, gin.H{
		"list":  comments,
		"total": total,
	})
}