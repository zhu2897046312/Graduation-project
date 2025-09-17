package handlers

import (
	"fmt"
	"server/models/shop"
	"server/service"
	"server/models/common"
	"server/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type ShopTagHandler struct {
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

func NewShopTagHandler(service *service.ShopTagService,tagMateService *service.ShopTagMateService) *ShopTagHandler {
	return &ShopTagHandler{
		service: service,
		tagMateService: tagMateService,
	}
}

// 创建商品标签
func (h *ShopTagHandler) CreateTag(c *gin.Context) {
	var req CreaeteTagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}
	
	tag := shop.ShopTag{
		Code:           req.Code,
		MatchWord:      req.MatchWord,
		ReadNum:        int(utils.ConvertToUint(req.ReadNum)),
		SortNum:        int(utils.ConvertToUint(req.SortNum)),
		State:          int8(utils.ConvertToUint(req.State)),
		Thumb:          req.Thumb,
		Title:          req.Title,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}
	tagMate := shop.ShopTagMate{
		Content: req.MatchWord,
		SeoTitle: req.SEOTitle,
		SeoDescription: req.SEODescription,
		SeoKeyword: req.SEOKeyword,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}
	if err := h.service.CreateTag(&tag); err != nil {
		Error(c, 21001, err.Error())
		return
	}
	tagMate.ID = tag.ID
	fmt.Println(tag)
	if err := h.tagMateService.CreateTagMate(&tagMate); err != nil {
		Error(c, 21001, err.Error())
		return
	}

	Success(c, nil)
}

// 更新商品标签
func (h *ShopTagHandler) UpdateTag(c *gin.Context) {
	var req UpdateTagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}
	
	tag := shop.ShopTag{
		Code:           req.Code,
		MatchWord:      req.MatchWord,
		ReadNum:        int(utils.ConvertToUint(req.ReadNum)),
		SortNum:        int(utils.ConvertToUint(req.SortNum)),
		State:          int8(utils.ConvertToUint(req.State)),
		Thumb:          req.Thumb,
		Title:          req.Title,
		ID:             common.MyID(req.ID),
		UpdatedTime: time.Now(),
	}
	tagMate := shop.ShopTagMate{
		ID:             common.MyID(req.ID),
		Content: req.MatchWord,
		SeoTitle: req.SEOTitle,
		SeoDescription: req.SEODescription,
		SeoKeyword: req.SEOKeyword,
		UpdatedTime: time.Now(),
	}
	if err := h.service.UpdateTag(&tag); err != nil {
		Error(c, 21001, err.Error())
		return
	}
	if err := h.tagMateService.UpdateTagMate(&tagMate); err != nil {
		Error(c, 21001, err.Error())
		return
	}

	Success(c, nil)
}

// 获取标签详情
func (h *ShopTagHandler) GetTag(c *gin.Context) {
	id := c.Query("id")
	uid := utils.ConvertToUint(id)
	if uid == 0 {
		InvalidParams(c)
		return
	}

	tag, err := h.service.GetTagByID(common.MyID(uid))
	if err != nil {
		Error(c, 21003, "标签不存在")
		return
	}
	tagMate , err_ := h.tagMateService.GetTagMateByID(common.MyID(uid))
	if err_ != nil {
		Error(c, 21003, "标签不存在")
		return
	}
	resonse := gin.H{
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
	Success(c, resonse)
}

// 分页获取标签
func (h *ShopTagHandler) ListTags(c *gin.Context) {
	var req ListTagsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}

	tags, total, err := h.service.ListTags(shop.TagQueryParams{
		State:      req.State,
		Title:      req.Title,
		Page:       req.Page,
		PageSize:   req.PageSize,
	})
	if err != nil {
		Error(c, 21007, "获取标签列表失败")
		return
	}

	Success(c, gin.H{
		"list":  tags,
		"total": total,
	})
}

func (h *ShopTagHandler) DeleteTag(c *gin.Context) {
	id := c.Query("id")
	uid := utils.ConvertToUint(id)
	if uid == 0 {
		InvalidParams(c)
		return
	}
	if err := h.service.DeleteTagByID(common.MyID(uid)); err != nil{
		Error(c, 21008, "删除标签失败")
		return
	}
	if err := h.tagMateService.DeleteTagMate(common.MyID(uid)); err != nil{
		Error(c, 21008, "删除标签失败")
		return
	}

	Success(c, nil)
}		

func (h *ShopTagHandler) GetTagByCode(c *gin.Context) {
	code := c.Query("code")
	if code == ""{
		InvalidParams(c)
	}

	tag, err := h.service.GetTagByCode(code)
	if err != nil {
		Error(c, 21009, "获取标签失败")
		return
	}
	tagMate , err_ := h.tagMateService.GetTagMateByID(common.MyID(tag.ID))
	if err_ != nil {
		Error(c, 21003, "标签不存在")
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
	Success(c, resonse)
}