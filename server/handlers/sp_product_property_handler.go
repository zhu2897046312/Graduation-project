package handlers

import (
	"github.com/gin-gonic/gin"
	"server/models/sp"
	"server/service"
	"strconv"
)

type SpProductPropertyHandler struct {
	service *service.SpProductPropertyService
}

func NewSpProductPropertyHandler(service *service.SpProductPropertyService) *SpProductPropertyHandler {
	return &SpProductPropertyHandler{service: service}
}

// CreateProperty 创建商品属性
func (h *SpProductPropertyHandler) CreateProperty(c *gin.Context) {
	var property sp.SpProductProperty
	if err := c.ShouldBindJSON(&property); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.CreateProperty(&property); err != nil {
		Error(c, 3001, err.Error())
		return
	}

	Success(c, property)
}

// UpdateProperty 更新商品属性
func (h *SpProductPropertyHandler) UpdateProperty(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		InvalidParams(c)
		return
	}

	var property sp.SpProductProperty
	if err := c.ShouldBindJSON(&property); err != nil {
		InvalidParams(c)
		return
	}
	property.ID = uint(id)

	if err := h.service.UpdateProperty(&property); err != nil {
		Error(c, 3002, err.Error())
		return
	}

	Success(c, property)
}

// GetPropertiesByProduct 获取商品属性列表
func (h *SpProductPropertyHandler) GetPropertiesByProduct(c *gin.Context) {
	productID, err := strconv.ParseUint(c.Param("product_id"), 10, 64)
	if err != nil || productID == 0 {
		InvalidParams(c)
		return
	}

	properties, err := h.service.GetPropertiesByProductID(uint(productID))
	if err != nil {
		Error(c, 3003, err.Error())
		return
	}

	Success(c, properties)
}

// DeletePropertiesByProduct 删除商品所有属性
func (h *SpProductPropertyHandler) DeletePropertiesByProduct(c *gin.Context) {
	productID, err := strconv.ParseUint(c.Param("product_id"), 10, 64)
	if err != nil || productID == 0 {
		InvalidParams(c)
		return
	}

	if err := h.service.DeletePropertiesByProductID(uint(productID)); err != nil {
		Error(c, 3004, err.Error())
		return
	}

	Success(c, nil)
}