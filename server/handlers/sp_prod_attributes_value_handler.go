package handlers

import (
	"server/models/sp"
	"server/service"
	"server/models/common"
	"server/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ListProdAttributesRequest struct {
	ProdAttributesID interface{}   `json:"prod_attributes_id"`
	Title            string `json:"title"`
	Page             int    `json:"page_no"`
	PageSize         int    `json:"page_size"`
}

type SpProdAttributesValueCreateRequest struct {
	Title string `json:"title"`
	SortNum interface{} `json:"sort_num"`
}

type SpProdAttributesValueHandler struct {
	service *service.SpProdAttributesValueService
}

func NewSpProdAttributesValueHandler(service *service.SpProdAttributesValueService) *SpProdAttributesValueHandler {
	return &SpProdAttributesValueHandler{service: service}
}

// 创建属性值
func (h *SpProdAttributesValueHandler) CreateAttributeValue(c *gin.Context) {
	var req SpProdAttributesValueCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}
	value := sp.SpProdAttributesValue{
		Title: req.Title,
		SortNum: uint16(utils.ConvertToUint(req.SortNum)),
	}
	if err := h.service.CreateAttributeValue(&value); err != nil {
		Error(c, 29001, err.Error())
		return
	}

	Success(c, value)
}

// 更新属性值
func (h *SpProdAttributesValueHandler) UpdateAttributeValue(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil || id == 0 {
		InvalidParams(c)
		return
	}

	var value sp.SpProdAttributesValue
	if err := c.ShouldBindJSON(&value); err != nil {
		InvalidParams(c)
		return
	}
	value.ID = common.MyID(id)

	if err := h.service.UpdateAttributeValue(&value); err != nil {
		Error(c, 29002, err.Error())
		return
	}

	Success(c, value)
}

// 根据属性ID获取值列表
func (h *SpProdAttributesValueHandler) List(c *gin.Context) {
	var req ListProdAttributesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}
	ProdAttributesID := common.MyID(utils.ConvertToUint(req.ProdAttributesID))
	attr := sp.SpProdAttributesQueryParams{
		Page:             req.Page,
		PageSize:         req.PageSize,
		ProdAttributesID: ProdAttributesID,
	}
	values, len, err := h.service.ListProdAttributes(attr)
	if err != nil {
		Error(c, 29003, "获取属性值失败")
		return
	}

	Success(c, gin.H{
		"list": values,
		"total":  len,
	})
}

// 获取属性值详情
func (h *SpProdAttributesValueHandler) GetValue(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil || id == 0 {
		InvalidParams(c)
		return
	}

	value, err := h.service.GetValueByID(common.MyID(id))
	if err != nil {
		Error(c, 29004, "属性值不存在")
		return
	}

	Success(c, value)
}

// 批量创建属性值
func (h *SpProdAttributesValueHandler) BatchCreateAttributeValues(c *gin.Context) {
	var values []sp.SpProdAttributesValue
	if err := c.ShouldBindJSON(&values); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.BatchCreateAttributeValues(values); err != nil {
		Error(c, 29005, err.Error())
		return
	}

	Success(c, nil)
}

// 删除属性的所有值
func (h *SpProdAttributesValueHandler) DeleteValuesByAttribute(c *gin.Context) {
	attrID, err := strconv.ParseUint(c.Param("attr_id"), 10, 32)
	if err != nil || attrID == 0 {
		InvalidParams(c)
		return
	}

	if err := h.service.DeleteValuesByAttributeID(common.MyID(attrID)); err != nil {
		Error(c, 29006, err.Error())
		return
	}

	Success(c, nil)
}
