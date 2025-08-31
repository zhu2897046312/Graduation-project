package handlers

import (
	"server/models/cms"
	"server/service"
	"server/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CmsRecommendIndexHandler struct {
	service *service.CmsRecommendIndexService
}

type ListRecommendsIndexRequest struct {
	RecommendID interface{}    `json:"recommend_id"`
	Title       string `json:"title"`
	Page        int    `json:"page_no"`
	PageSize    int    `json:"page_size"`
}

func NewCmsRecommendIndexHandler(service *service.CmsRecommendIndexService) *CmsRecommendIndexHandler {
	return &CmsRecommendIndexHandler{service: service}
}

// 创建推荐索引
func (h *CmsRecommendIndexHandler) CreateIndex(c *gin.Context) {
	var index cms.CmsRecommendIndex
	if err := c.ShouldBindJSON(&index); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.CreateRecommendIndex(&index); err != nil {
		Error(c, 7001, err.Error())
		return
	}

	Success(c, index)
}

// 更新推荐索引
func (h *CmsRecommendIndexHandler) UpdateIndex(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		InvalidParams(c)
		return
	}

	var index cms.CmsRecommendIndex
	if err := c.ShouldBindJSON(&index); err != nil {
		InvalidParams(c)
		return
	}
	index.ID = id

	if err := h.service.UpdateRecommendIndex(&index); err != nil {
		Error(c, 7002, err.Error())
		return
	}

	Success(c, index)
}

// 根据推荐ID获取索引
func (h *CmsRecommendIndexHandler) GetByRecommendID(c *gin.Context) {
	recommendID, err := strconv.Atoi(c.Param("recommend_id"))
	if err != nil || recommendID <= 0 {
		InvalidParams(c)
		return
	}

	indices, err := h.service.GetIndicesByRecommendID(recommendID)
	if err != nil {
		ServerError(c, err)
		return
	}

	Success(c, indices)
}

// 根据状态获取索引
func (h *CmsRecommendIndexHandler) GetByState(c *gin.Context) {
	state, err := strconv.Atoi(c.Query("state"))
	if err != nil {
		InvalidParams(c)
		return
	}

	indices, err := h.service.GetIndicesByState(int8(state))
	if err != nil {
		ServerError(c, err)
		return
	}

	Success(c, indices)
}

// 删除推荐索引
func (h *CmsRecommendIndexHandler) DeleteIndex(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		InvalidParams(c)
		return
	}

	if err := h.service.DeleteRecommendIndex(id); err != nil {
		Error(c, 7003, err.Error())
		return
	}

	Success(c, nil)
}

func (h *CmsRecommendIndexHandler) ListRecommendsIndex(c *gin.Context) {
	var req ListRecommendsIndexRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}
	recommend_id:= utils.ConvertToUint(req.RecommendID)
	recommends, total, err := h.service.ListRecommendsIndex(cms.RecommendIndexQueryParams{
		Title:     req.Title,
		RecommendID: int(recommend_id),
		Page:     req.Page,
		PageSize: req.PageSize,
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
