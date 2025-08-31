package handlers

import (
	"github.com/gin-gonic/gin"
	"server/service"
	"server/models/cms"
	"strconv"
)

type ListRecommendsRequest struct{
	Page       int         `json:"page_no"`
	PageSize   int         `json:"page_size"`
}

type CmsRecommendHandler struct {
	service *service.CmsRecommendService
}

func NewCmsRecommendHandler(service *service.CmsRecommendService) *CmsRecommendHandler {
	return &CmsRecommendHandler{service: service}
}

// 创建推荐内容
func (h *CmsRecommendHandler) CreateRecommend(c *gin.Context) {
	var recommend cms.CmsRecommend
	if err := c.ShouldBindJSON(&recommend); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.CreateRecommend(&recommend); err != nil {
		Error(c, 8001, err.Error())
		return
	}

	Success(c, recommend)
}

// 更新推荐内容
func (h *CmsRecommendHandler) UpdateRecommend(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		InvalidParams(c)
		return
	}

	var recommend cms.CmsRecommend
	if err := c.ShouldBindJSON(&recommend); err != nil {
		InvalidParams(c)
		return
	}
	recommend.ID = int(id)

	if err := h.service.UpdateRecommend(&recommend); err != nil {
		Error(c, 8002, err.Error())
		return
	}

	Success(c, recommend)
}

// 获取有效推荐
func (h *CmsRecommendHandler) GetActiveRecommends(c *gin.Context) {
	recommends, err := h.service.GetActiveRecommends()
	if err != nil {
		ServerError(c, err)
		return
	}

	Success(c, recommends)
}

// 根据状态获取推荐
func (h *CmsRecommendHandler) GetRecommendsByState(c *gin.Context) {
	state, err := strconv.Atoi(c.Query("state"))
	if err != nil {
		InvalidParams(c)
		return
	}

	recommends, err := h.service.GetRecommendsByState(int8(state))
	if err != nil {
		ServerError(c, err)
		return
	}

	Success(c, recommends)
}

func (h *CmsRecommendHandler) ListRecommends(c *gin.Context) {
	var req ListRecommendsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}
	recommends, total, err := h.service.ListRecommends(cms.RecommendQueryParams{
		Page:       req.Page,
		PageSize:   req.PageSize,
	})
	if err != nil {
		ServerError(c, err)
		return
	}	

	Success(c, gin.H{
		"list":  recommends,
		"total": total,
	})
}