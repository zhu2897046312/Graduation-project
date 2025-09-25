package admin

import (
	"server/models/common"
	"server/models/sp"
	"server/service"
	"server/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// import "github.com/go-playground/validator/v10"
type SpCategoryHandler struct {
	service *service.SpCategoryService
}

// 树形结构返回结果
type CategoryTreeResult struct {
	Label    string               `json:"label"`
	Value    common.MyID          `json:"value"`
	Node     sp.SpCategory        `json:"node"`
	Children []CategoryTreeResult `json:"children"`
}

type SpProductCreateReq struct {
	PID            common.MyID      `json:"pid"`             // 父级ID（可选）
	Title          string           `json:"title"`           // 商品名称（必填，长度1-200）
	Code           string           `json:"code"`            // 商品编码（可选，长度≤50）
	State          common.MyState   `json:"state"`           // 状态（0下架/1上架，默认为1）
	Icon           string           `json:"icon"`            // 图标URL（可选，需为URL格式）
	Picture        string           `json:"picture"`         // 主图URL（可选，需为URL格式）
	Description    string           `json:"description"`     // 描述（可选，长度≤500）
	SEOTitle       string           `json:"seo_title"`       // SEO标题（可选，长度≤100）
	SEOKeyword     string           `json:"seo_keyword"`     // SEO关键词（可选，长度≤200）
	SEODescription string           `json:"seo_description"` // SEO描述（可选，长度≤300）
	SortNum        common.MySortNum `json:"sort_num"`        // 排序值（可选） validate:"omitempty"
}

// SpProductUpdateReq 更新商品的请求参数
type SpProductUpdateReq struct {
	ID             common.MyID      `json:"id"`              // 商品ID（必填）
	PID            common.MyID      `json:"pid"`             // 父级ID（可选）
	Title          string           `json:"title"`           // 商品名称（必填，长度1-200）
	Code           string           `json:"code"`            // 商品编码（可选，长度≤50）
	State          common.MyState   `json:"state"`           // 状态（0下架/1上架，默认为1）
	Icon           string           `json:"icon"`            // 图标URL（可选，需为URL格式）
	Picture        string           `json:"picture"`         // 主图URL（可选，需为URL格式）
	Description    string           `json:"description"`     // 描述（可选，长度≤500）
	SEOTitle       string           `json:"seo_title"`       // SEO标题（可选，长度≤100）
	SEOKeyword     string           `json:"seo_keyword"`     // SEO关键词（可选，长度≤200）
	SEODescription string           `json:"seo_description"` // SEO描述（可选，长度≤300）
	SortNum        common.MySortNum `json:"sort_num"`        // 排序值（可选）
}

func NewSpCategoryHandler(service *service.SpCategoryService) *SpCategoryHandler {
	return &SpCategoryHandler{service: service}
}

// 创建分类
func (h *SpCategoryHandler) CreateCategory(c *gin.Context) {
	var req SpProductCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}
	category := sp.SpCategory{
		Pid:            req.PID,
		Title:          req.Title,
		Code:           req.Code,
		State:          req.State,
		Icon:           req.Icon,
		Picture:        req.Picture,
		Description:    req.Description,
		SeoTitle:       req.SEOTitle,
		SeoKeyword:     req.SEOKeyword,
		SeoDescription: req.SEODescription,
		SortNum:        req.SortNum,
	}
	if err := h.service.CreateCategory(&category); err != nil {
		Error(c, 22001, err.Error())
		return
	}

	Success(c, category)
}

// 更新分类
func (h *SpCategoryHandler) UpdateCategory(c *gin.Context) {
	var req SpProductUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		Error(c, 22002, "参数校验失败: "+err.Error())
		return
	}
	category := sp.SpCategory{
		ID:             req.ID,
		Pid:            req.PID,
		Title:          req.Title,
		Code:           req.Code,
		State:          req.State,
		Icon:           req.Icon,
		Picture:        req.Picture,
		Description:    req.Description,
		SeoTitle:       req.SEOTitle,
		SeoKeyword:     req.SEOKeyword,
		SeoDescription: req.SEODescription,
		SortNum:        req.SortNum,
	}
	if err := h.service.UpdateCategory(&category); err != nil {
		Error(c, 22002, err.Error())
		return
	}
	Success(c, nil)
}

// 获取分类详情
func (h *SpCategoryHandler) GetCategory(c *gin.Context) {
	id := c.Query("id")
	uintId := utils.ConvertToUint(id)
	if uintId == 0 {
		InvalidParams_1(c, uintId)
		return
	}

	category, err := h.service.GetCategoryByID(common.MyID(uintId))
	if err != nil {
		Error(c, 22003, "分类不存在")
		return
	}

	Success(c, category)
}

// 获取分类树
func (h *SpCategoryHandler) Tree(c *gin.Context) {
	pid, err := strconv.ParseUint(c.Query("pid"), 10, 32)
	if err != nil {
		pid = 0 // 默认从顶级分类开始
	}

	var state *common.MyState
	if stateStr := c.Query("state"); stateStr != "" {
		if stateVal, err := strconv.ParseUint(stateStr, 10, 8); err == nil {
			stateUint8 := common.MyState(stateVal)
			state = &stateUint8
		}
	}

	// 获取所有分类
	categories, err := h.service.GetAllCategories()
	if err != nil {
		Error(c, 22007, "获取分类失败")
		return
	}

	// 构建树形结构
	tree := h.buildTree(categories, common.MyID(pid), state)

	Success(c, tree)
}

// 递归构建树形结构
func (h *SpCategoryHandler) buildTree(categories []*sp.SpCategory, pid common.MyID, state *common.MyState) []CategoryTreeResult {
	var tree []CategoryTreeResult

	for _, category := range categories {
		if category.Pid == pid && (state == nil || category.State == *state) {
			treeResult := CategoryTreeResult{
				Label: category.Title,
				Value: (category.ID),
				Node:  *category,
			}

			// 递归获取子节点
			children := h.buildTree(categories, category.ID, state)
			if len(children) > 0 {
				treeResult.Children = children
			} else {
				treeResult.Children = nil
			}

			tree = append(tree, treeResult)
		}
	}

	return tree
}
