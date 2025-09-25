package client

import (
	"server/models/shop"
	"server/service"
	"server/models/common"
	"server/utils"

	"github.com/gin-gonic/gin"
)

type ClientShopTagHandler struct {
	service *service.ShopTagService
	tagMateService *service.ShopTagMateService
}

type ListTagsRequest struct {
	// CategoryID interface{} `json:"category_id"` // 使用interface{}接收任何类型
	State      int `json:"status"`
	Title      string `json:"title"`
	Page       int `json:"page_no"`
	PageSize   int `json:"page_size"`
}

type CreaeteTagRequest struct {
	Code           string `json:"code"`
	MatchWord      string `json:"match_word"`
	ReadNum        interface{} `json:"read_num"`
	SEODescription string `json:"seo_description"`
	SEOKeyword     string `json:"seo_keyword"`
	SEOTitle       string `json:"seo_title"`
	SortNum        interface{} `json:"sort_num"`
	State          interface{}  `json:"state"`
	Thumb          string `json:"thumb"`
	Title          string `json:"title"`
}

type UpdateTagRequest struct {
	ID             int `json:"id"`
	Code           string `json:"code"`
	MatchWord      string `json:"match_word"`
	ReadNum        interface{} `json:"read_num"`
	SEODescription string `json:"seo_description"`
	SEOKeyword     string `json:"seo_keyword"`
	SEOTitle       string `json:"seo_title"`
	SortNum        interface{} `json:"sort_num"`
	State          interface{}  `json:"state"`
	Thumb          string `json:"thumb"`
	Title          string `json:"title"`
}

func NewClientShopTagHandler(service *service.ShopTagService,tagMateService *service.ShopTagMateService) *ClientShopTagHandler {
	return &ClientShopTagHandler{
		service: service,
		tagMateService: tagMateService,
	}
}

// 分页获取标签
func (h *ClientShopTagHandler) ListTags(c *gin.Context) {
	var req ListTagsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.InvalidParams(c)
		return
	}

	tags, total, err := h.service.ListTags(shop.TagQueryParams{
		State:      common.MyState(req.State),
		Title:      req.Title,
		Page:       req.Page,
		PageSize:   req.PageSize,
	})
	if err != nil {
		utils.Error(c, 21007, "获取标签列表失败")
		return
	}

	utils.Success(c, gin.H{
		"list":  tags,
		"total": total,
	})
}


func (h *ClientShopTagHandler) GetTagByCode(c *gin.Context) {
	code := c.Query("code")
	if code == ""{
		utils.InvalidParams(c)
	}

	tag, err := h.service.GetTagByCode(code)
	if err != nil {
		utils.Error(c, 21009, "获取标签失败")
		return
	}
	tagMate , err_ := h.tagMateService.GetTagMateByID(common.MyID(tag.ID))
	if err_ != nil {
		utils.Error(c, 21003, "标签不存在")
		return
	}
	resonse := gin.H{
		"id": tag.ID,
		"seo_title": tagMate.SeoTitle,	
		"seo_description": tagMate.SeoDescription,
		"seo_keyword": tagMate.SeoKeyword,
		"match_word": tagMate.Content,
		"code": tag.Code,	
		"read_num": tag.ReadNum,
		"sort_num": tag.SortNum,
		"state": tag.State,
		"thumb": tag.Thumb,
		"title": tag.Title,
	}
	utils.Success(c, resonse)
}