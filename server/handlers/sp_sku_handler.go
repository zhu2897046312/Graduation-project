package handlers

import (
	"github.com/gin-gonic/gin"
	"server/models/sp"
	"server/service"
	"strconv"
)

type SpSkuHandler struct {
	service *service.SpSkuService
}

func NewSpSkuHandler(service *service.SpSkuService) *SpSkuHandler {
	return &SpSkuHandler{service: service}
}

// CreateSku 创建SKU
func (h *SpSkuHandler) CreateSku(c *gin.Context) {
	var sku sp.SpSku
	if err := c.ShouldBindJSON(&sku); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.CreateSku(&sku); err != nil {
		Error(c, 3201, err.Error())
		return
	}

	Success(c, sku)
}

// UpdateSku 更新SKU
func (h *SpSkuHandler) UpdateSku(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		InvalidParams(c)
		return
	}

	var sku sp.SpSku
	if err := c.ShouldBindJSON(&sku); err != nil {
		InvalidParams(c)
		return
	}
	sku.ID = uint(id)

	if err := h.service.UpdateSku(&sku); err != nil {
		Error(c, 3202, err.Error())
		return
	}

	Success(c, sku)
}

// GetSkusByProduct 获取商品SKU列表
func (h *SpSkuHandler) GetSkusByProduct(c *gin.Context) {
	productID, err := strconv.ParseUint(c.Param("product_id"), 10, 64)
	if err != nil || productID == 0 {
		InvalidParams(c)
		return
	}

	skus, err := h.service.GetSkusByProductID(uint(productID))
	if err != nil {
		Error(c, 3203, err.Error())
		return
	}

	Success(c, skus)
}

// UpdateSkuStock 更新SKU库存
func (h *SpSkuHandler) UpdateSkuStock(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		InvalidParams(c)
		return
	}

	var req struct {
		Stock int `json:"stock"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.UpdateStock(uint(id), req.Stock); err != nil {
		Error(c, 3204, err.Error())
		return
	}

	Success(c, nil)
}