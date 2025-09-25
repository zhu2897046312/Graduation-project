package admin

import (
	"server/models/common"
	"server/models/sp"
	"server/service"
	"server/utils"

	"github.com/gin-gonic/gin"
)

type SpProdAttributesPageRequest struct {
	Title    string `json:"title"`
	Page     int    `json:"page_no"`
	PageSize int    `json:"page_size"`
}

type SpProdAttributesPageResponse struct {
	Data  []sp.SpProdAttributes `json:"list"`
	Total int64                 `json:"total"`
}

type SpProdAttributesHandler struct {
	service                       *service.SpProdAttributesService
	spProdAttributesValuesService *service.SpProdAttributesValueService
}

type SpProdAttributesCreateRequest struct {
	Title   string           `json:"title"`
	SortNum common.MySortNum `json:"sort_num"`
}

func NewSpProdAttributesHandler(service *service.SpProdAttributesService, spProdAttributesValuesService *service.SpProdAttributesValueService) *SpProdAttributesHandler {
	return &SpProdAttributesHandler{
		service:                       service,
		spProdAttributesValuesService: spProdAttributesValuesService,
	}
}

// 创建商品属性
func (h *SpProdAttributesHandler) CreateAttribute(c *gin.Context) {
	var req SpProdAttributesCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}

	attr := sp.SpProdAttributes{
		Title:   req.Title,
		SortNum: req.SortNum,
	}

	if err := h.service.CreateAttribute(&attr); err != nil {
		Error(c, 28001, err.Error())
		return
	}

	Success(c, attr)
}

// 更新商品属性
func (h *SpProdAttributesHandler) UpdateAttribute(c *gin.Context) {
	var req sp.SpProdAttributes
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}
	attr := sp.SpProdAttributes{
		ID:      req.ID,
		Title:   req.Title,
		SortNum: req.SortNum,
	}

	if err := h.service.UpdateAttribute(&attr); err != nil {
		Error(c, 28002, err.Error())
		return
	}

	Success(c, attr)
}

// 获取属性详情
func (h *SpProdAttributesHandler) GetAttribute(c *gin.Context) {
	id := c.Query("id")
	uintId := utils.ConvertToUint(id)
	if uintId == 0 {
		InvalidParams_1(c, uintId)
		return
	}
	// fmt.Println(uintId)
	attr, err := h.service.GetAttributeByID(common.MyID(uintId))
	if err != nil {
		Error(c, 28003, "属性不存在")
		return
	}

	Success(c, attr)
}

// 获取所有属性
func (h *SpProdAttributesHandler) GetAllAttributes(c *gin.Context) {
	attrs, err := h.service.GetAllAttributes()
	if err != nil {
		Error(c, 28004, "获取属性列表失败")
		return
	}

	Success(c, attrs)
}

// 分页获取属性
func (h *SpProdAttributesHandler) List(c *gin.Context) {
	var req SpProdAttributesPageRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		InvalidParams(c)
		return
	}

	// 设置默认值
	if req.Page == 0 {
		req.Page = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 20
	}

	attrs, total, err := h.service.GetAttributesByPage(req.Title, req.Page, req.PageSize)
	if err != nil {
		Error(c, 28007, "获取属性列表失败")
		return
	}

	// 计算总页数
	totalPages := total / int64(req.PageSize)
	if total%int64(req.PageSize) > 0 {
		totalPages++
	}

	response := SpProdAttributesPageResponse{
		Data:  attrs,
		Total: total,
	}

	Success(c, response)
}

// 删除属性
func (h *SpProdAttributesHandler) DeleteAttribute(c *gin.Context) {
	id := c.Query("id")
	uintId := utils.ConvertToUint(id)
	if uintId == 0 {
		InvalidParams_1(c, uintId)
		return
	}
	prodAttributesID := common.MyID(utils.ConvertToUint(id))
	_, len, _ := h.spProdAttributesValuesService.GetAllByProdAttributesID(sp.SpProdAttributesQueryParams{
		ProdAttributesID: prodAttributesID,
	})
	if len > 0 {
		Error(c, 28006, "属性值不为空，无法删除")
		return
	}
	if err := h.service.DeleteAttribute(common.MyID(uintId)); err != nil {
		Error(c, 28006, err.Error())
		return
	}

	Success(c, nil)
}
