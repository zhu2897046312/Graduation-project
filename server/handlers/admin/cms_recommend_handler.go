package admin

import (
	"server/models/cms"
	"server/models/common"
	"server/service"
	"server/utils"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ListRecommendsRequest struct {
	Page     int `json:"page_no"`
	PageSize int `json:"page_size"`
}

type CmsRecommendHandler struct {
	service               *service.CmsRecommendService
	recommendIndexService *service.CmsRecommendIndexService
}
type CmsRecommendCreateRequest struct {
	Title       string         `json:"title"`       // 推荐位名称
	Code        string         `json:"code"`        // 编码
	Thumb       string         `json:"thumb"`       // 缩略图
	Description string         `json:"description"` // 描述
	State       common.MyState `json:"state"`       // 状态:1=已发布;2=未发布
	MoreLink    string         `json:"moreLink"`    // 更多链接
}

type CmsRecommendUpdateRequest struct {
	ID          interface{}    `json:"id"`
	Title       string         `json:"title"`       // 推荐位名称
	Code        string         `json:"code"`        // 编码
	Thumb       string         `json:"thumb"`       // 缩略图
	Description string         `json:"description"` // 描述
	State       common.MyState `json:"state"`       // 状态:1=已发布;2=未发布
	MoreLink    string         `json:"moreLink"`    // 更多链接
}

func NewCmsRecommendHandler(service *service.CmsRecommendService, recommendIndexService *service.CmsRecommendIndexService) *CmsRecommendHandler {
	return &CmsRecommendHandler{
		service:               service,
		recommendIndexService: recommendIndexService,
	}
}

// 创建推荐内容
func (h *CmsRecommendHandler) CreateRecommend(c *gin.Context) {
	var req CmsRecommendCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}
	recommend := cms.CmsRecommend{
		Title:       req.Title,
		Code:        req.Code,
		Thumb:       req.Thumb,
		Description: req.Description,
		State:       req.State,
		MoreLink:    req.MoreLink,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}
	if err := h.service.CreateRecommend(&recommend); err != nil {
		Error(c, 8001, err.Error())
		return
	}

	Success(c, nil)
}

// 更新推荐内容
func (h *CmsRecommendHandler) UpdateRecommend(c *gin.Context) {
	var req CmsRecommendUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}
	id := utils.ConvertToUint(req.ID)

	recommend := cms.CmsRecommend{
		ID:          common.MyID(id),
		Title:       req.Title,
		Code:        req.Code,
		Thumb:       req.Thumb,
		Description: req.Description,
		State:       req.State,
		MoreLink:    req.MoreLink,
		UpdatedTime: time.Now(),
	}

	if err := h.service.UpdateRecommend(&recommend); err != nil {
		Error(c, 8002, err.Error())
		return
	}

	Success(c, nil)
}

func (h *CmsRecommendHandler) ListRecommends(c *gin.Context) {
	var req ListRecommendsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}
	recommends, total, err := h.service.ListRecommends(cms.RecommendQueryParams{
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

func (h *CmsRecommendHandler) DeleteRecommendByID(c *gin.Context) {
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
	// 2. 检查是否有关联的推荐索引
	indices, err := h.recommendIndexService.GetIndicesByRecommendID(common.MyID(idUint))
	if err != nil && err != gorm.ErrRecordNotFound {
		ServerError(c, err)
		return
	}

	// 如果有关联数据，不允许删除
	if err == nil && len(indices) > 0 {
		Error(c, 8004, "请先删除关联的推荐内容")
		return
	}

	if err := h.service.DeleteRecommendByID(common.MyID(idUint)); err != nil {
		Error(c, 8003, err.Error())
		return
	}
	Success(c, nil)
}

func (h *CmsRecommendHandler) GetRecommendByID(c *gin.Context) {
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
	recommend, err := h.service.GetRecommendByID(common.MyID(idUint))
	if err != nil {
		ServerError(c, err)
		return
	}
	Success(c, recommend)
}
