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
	RecommendID interface{} `json:"recommend_id"`
	Title       string      `json:"title"`
	Page        int         `json:"page_no"`
	PageSize    int         `json:"page_size"`
}

type CmsRecommendIndexCreateReqest struct {
	RecommendID interface{} `json:"recommend_id"` // 所属推荐位
	Title       string      `json:"title" `       // 名称
	Thumb       string      `json:"thumb"`        // 缩略图
	Link        string      `json:"link" `        // 链接
	State       interface{} `json:"state" `       // 状态:1=已发布;2=未发布
	ProductID   interface{} `json:"product_id"`   // 商品id
	DocumentID  interface{} `json:"document_id"`  // 文档id
	SortNum     interface{} `json:"sort_num"`     // 排序
}

type CmsRecommendIndexUpdateReqest struct {
	ID          interface{} `json:"id"`
	RecommendID interface{} `json:"recommend_id"` // 所属推荐位
	Title       string      `json:"title" `       // 名称
	Thumb       string      `json:"thumb"`        // 缩略图
	Link        string      `json:"link" `        // 链接
	State       interface{} `json:"state" `       // 状态:1=已发布;2=未发布
	ProductID   interface{} `json:"product_id"`   // 商品id
	DocumentID  interface{} `json:"document_id"`  // 文档id
	SortNum     interface{} `json:"sort_num"`     // 排序
}

func NewCmsRecommendIndexHandler(service *service.CmsRecommendIndexService) *CmsRecommendIndexHandler {
	return &CmsRecommendIndexHandler{service: service}
}

// 创建推荐索引
func (h *CmsRecommendIndexHandler) CreateIndex(c *gin.Context) {
	var req CmsRecommendIndexCreateReqest
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}

	index := cms.CmsRecommendIndex{
		RecommendID: int(utils.ConvertToUint(req.RecommendID)),
		Title:       req.Title,
		Thumb:       req.Thumb,
		Link:        req.Link,
		State:       int8(utils.ConvertToUint(req.State)),
		ProductID:   int(utils.ConvertToUint(req.ProductID)),
		DocumentID:  int(utils.ConvertToUint(req.DocumentID)),
		SortNum:     int(utils.ConvertToUint(req.SortNum)),
	}

	if err := h.service.CreateRecommendIndex(&index); err != nil {
		Error(c, 7001, err.Error())
		return
	}

	Success(c, nil)
}

// 更新推荐索引
func (h *CmsRecommendIndexHandler) UpdateIndex(c *gin.Context) {
	var req CmsRecommendIndexUpdateReqest
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}

	index := cms.CmsRecommendIndex{
		ID:          int(utils.ConvertToUint(req.ID)),
		RecommendID: int(utils.ConvertToUint(req.RecommendID)),
		Title:       req.Title,
		Thumb:       req.Thumb,
		Link:        req.Link,
		State:       int8(utils.ConvertToUint(req.State)),
		ProductID:   int(utils.ConvertToUint(req.ProductID)),
		DocumentID:  int(utils.ConvertToUint(req.DocumentID)),
		SortNum:     int(utils.ConvertToUint(req.SortNum)),
	}

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

func (h *CmsRecommendIndexHandler) ListRecommendsIndex(c *gin.Context) {
	var req ListRecommendsIndexRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}
	recommend_id := utils.ConvertToUint(req.RecommendID)
	recommends, total, err := h.service.ListRecommendsIndex(cms.RecommendIndexQueryParams{
		Title:       req.Title,
		RecommendID: int(recommend_id),
		Page:        req.Page,
		PageSize:    req.PageSize,
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

func (h *CmsRecommendIndexHandler) DeleteRecommendIndexByID(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		InvalidParams(c)
		return
	}
	idUint := utils.ConvertToUint(id)
	if idUint <= 0 {
		InvalidParams(c)
		return
	}
	if err := h.service.DeleteRecommendIndex(int(idUint)); err != nil {
		Error(c, 7003, err.Error())
		return
	}
	Success(c, nil)
}

func (h *CmsRecommendIndexHandler) GetRecommendIndexByID(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		InvalidParams(c)
		return
	}
	idUint := utils.ConvertToUint(id)
	if idUint <= 0 {
		InvalidParams(c)
		return
	}
	index, err := h.service.GetRecommendIndexByID(int(idUint))
	if err != nil {
		ServerError(c, err)
		return
	}
	Success(c, index)
}	