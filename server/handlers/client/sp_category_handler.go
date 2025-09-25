package client

import (
	"server/models/common"
	"server/models/sp"
	"server/service"
	"server/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// import "github.com/go-playground/validator/v10"
type ClientSpCategoryHandler struct {
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
	PID            uint   `json:"pid"`             // 父级ID（可选）
	Title          string `json:"title"`           // 商品名称（必填，长度1-200）
	Code           string `json:"code"`            // 商品编码（可选，长度≤50）
	State          uint8  `json:"state"`           // 状态（0下架/1上架，默认为1）
	Icon           string `json:"icon"`            // 图标URL（可选，需为URL格式）
	Picture        string `json:"picture"`         // 主图URL（可选，需为URL格式）
	Description    string `json:"description"`     // 描述（可选，长度≤500）
	SEOTitle       string `json:"seo_title"`       // SEO标题（可选，长度≤100）
	SEOKeyword     string `json:"seo_keyword"`     // SEO关键词（可选，长度≤200）
	SEODescription string `json:"seo_description"` // SEO描述（可选，长度≤300）
	SortNum        uint16 `json:"sort_num"`        // 排序值（可选） validate:"omitempty"
}

// SpProductUpdateReq 更新商品的请求参数
type SpProductUpdateReq struct {
	ID             uint   `json:"id"`              // 商品ID（必填）
	PID            uint   `json:"pid"`             // 父级ID（可选）
	Title          string `json:"title"`           // 商品名称（必填，长度1-200）
	Code           string `json:"code"`            // 商品编码（可选，长度≤50）
	State          uint8  `json:"state"`           // 状态（0下架/1上架，默认为1）
	Icon           string `json:"icon"`            // 图标URL（可选，需为URL格式）
	Picture        string `json:"picture"`         // 主图URL（可选，需为URL格式）
	Description    string `json:"description"`     // 描述（可选，长度≤500）
	SEOTitle       string `json:"seo_title"`       // SEO标题（可选，长度≤100）
	SEOKeyword     string `json:"seo_keyword"`     // SEO关键词（可选，长度≤200）
	SEODescription string `json:"seo_description"` // SEO描述（可选，长度≤300）
	SortNum        uint16 `json:"sort_num"`        // 排序值（可选）
}

func NewClientSpCategoryHandler(service *service.SpCategoryService) *ClientSpCategoryHandler {
	return &ClientSpCategoryHandler{service: service}
}

// 获取分类详情
func (h *ClientSpCategoryHandler) GetCategory(c *gin.Context) {
	id := c.Query("id")
	uintId := utils.ConvertToUint(id)
	if uintId == 0 {
		utils.InvalidParams_1(c, uintId)
		return
	}

	category, err := h.service.GetCategoryByID(common.MyID(uintId))
	if err != nil {
		utils.Error(c, 22003, "分类不存在")
		return
	}

	utils.Success(c, category)
}

func (h *ClientSpCategoryHandler) GetCategoryByCode(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		utils.InvalidParams(c)
		return
	}

	category, err := h.service.GetCategoryByCode(code)
	if err != nil {
		utils.Error(c, 22003, "分类不存在")
		return
	}

	utils.Success(c, category)
}

func (h *ClientSpCategoryHandler) GetCategoryParents(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		utils.InvalidParams(c)
		return
	}

	category, err := h.service.GetParents(code)
	if err != nil {
		utils.Error(c, 22003, "分类不存在")
		return
	}

	utils.Success(c, category)
}

// 获取分类树
func (h *ClientSpCategoryHandler) GetCategoryTree(c *gin.Context) {
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
		utils.Error(c, 22007, "获取分类失败")
		return
	}

	// 构建树形结构
	tree := h.buildTree(categories, common.MyID(pid), state)

	utils.Success(c, tree)
}

// 递归构建树形结构
func (h *ClientSpCategoryHandler) buildTree(categories []*sp.SpCategory, pid common.MyID, state *common.MyState) []CategoryTreeResult {
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