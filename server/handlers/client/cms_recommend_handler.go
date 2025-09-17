package client

import (
	"server/models/cms"
	"server/service"
	"server/utils"

	"github.com/gin-gonic/gin"
)

type ListRecommendsRequest struct {
	Page     int `json:"page_no"`
	PageSize int `json:"page_size"`
}

type ClientCmsRecommendHandler struct {
	service *service.CmsRecommendService
	recommendIndexService *service.CmsRecommendIndexService
}
type CmsRecommendCreateRequest struct {
	Title       string `json:"title"`       // 推荐位名称
	Code        string `json:"code"`        // 编码
	Thumb       string `json:"thumb"`       // 缩略图
	Description string `json:"description"` // 描述
	State       int    `json:"state"`       // 状态:1=已发布;2=未发布
	MoreLink    string `json:"moreLink"`    // 更多链接
}

type CmsRecommendUpdateRequest struct {
	ID          interface{} `json:"id"`
	Title       string      `json:"title"`       // 推荐位名称
	Code        string      `json:"code"`        // 编码
	Thumb       string      `json:"thumb"`       // 缩略图
	Description string      `json:"description"` // 描述
	State       int         `json:"state"`       // 状态:1=已发布;2=未发布
	MoreLink    string      `json:"moreLink"`    // 更多链接
}

func NewClientCmsRecommendHandler(service *service.CmsRecommendService,recommendIndexService *service.CmsRecommendIndexService) *ClientCmsRecommendHandler {
	return &ClientCmsRecommendHandler{
		service: service,
		recommendIndexService: recommendIndexService,
	}
}

func (h *ClientCmsRecommendHandler) ListRecommendsIndex(c *gin.Context) {
	type IndexReq struct{
		Code string `json:"code"`
	}
	var req IndexReq
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.InvalidParams(c)
		return
	}
	recommend , err := h.service.GetRecommendByCode(req.Code)
	if err != nil {
		utils.ServerError(c, err)
		return
	}
	

	recommendsIndex, total, err_1 := h.recommendIndexService.ListRecommendsIndex(cms.RecommendIndexQueryParams{
		RecommendID: recommend.ID,
	})
	if err_1 != nil {
		utils.ServerError(c, err_1)
		return
	}

	utils.Success(c, gin.H{
		"list":  recommendsIndex,
		"total": total,
	})
}