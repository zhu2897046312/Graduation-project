package handlers

import (
	"github.com/gin-gonic/gin"
	"server/models/sp"
	"server/service"
	"strconv"
)

type SpOrderItemHandler struct {
	service *service.SpOrderItemService
}

func NewSpOrderItemHandler(service *service.SpOrderItemService) *SpOrderItemHandler {
	return &SpOrderItemHandler{service: service}
}

// 创建订单项
func (h *SpOrderItemHandler) CreateOrderItem(c *gin.Context) {
	var item sp.SpOrderItem
	if err := c.ShouldBindJSON(&item); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.CreateOrderItem(&item); err != nil {
		Error(c, 23001, err.Error())
		return
	}

	Success(c, item)
}

// 批量创建订单项
func (h *SpOrderItemHandler) BatchCreateOrderItems(c *gin.Context) {
	var items []*sp.SpOrderItem
	if err := c.ShouldBindJSON(&items); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.BatchCreateOrderItems(items); err != nil {
		Error(c, 23002, err.Error())
		return
	}

	Success(c, nil)
}

// 根据订单ID获取订单项
func (h *SpOrderItemHandler) GetItemsByOrder(c *gin.Context) {
	orderID, err := strconv.ParseUint(c.Param("order_id"), 10, 32)
	if err != nil || orderID == 0 {
		InvalidParams(c)
		return
	}

	items, err := h.service.GetItemsByOrderID(uint(orderID))
	if err != nil {
		Error(c, 23003, "获取订单项失败")
		return
	}

	Success(c, items)
}

// 根据产品ID获取订单项
func (h *SpOrderItemHandler) GetItemsByProduct(c *gin.Context) {
	productID, err := strconv.ParseUint(c.Param("product_id"), 10, 32)
	if err != nil || productID == 0 {
		InvalidParams(c)
		return
	}

	items, err := h.service.GetItemsByProductID(uint(productID))
	if err != nil {
		Error(c, 23004, "获取订单项失败")
		return
	}

	Success(c, items)
}

// 根据SKU ID获取订单项
func (h *SpOrderItemHandler) GetItemsBySku(c *gin.Context) {
	skuID, err := strconv.ParseUint(c.Param("sku_id"), 10, 32)
	if err != nil || skuID == 0 {
		InvalidParams(c)
		return
	}

	items, err := h.service.GetItemsBySkuID(uint(skuID))
	if err != nil {
		Error(c, 23005, "获取订单项失败")
		return
	}

	Success(c, items)
}

// 计算产品销售总量
func (h *SpOrderItemHandler) CalculateProductSales(c *gin.Context) {
	productID, err := strconv.ParseUint(c.Param("product_id"), 10, 32)
	if err != nil || productID == 0 {
		InvalidParams(c)
		return
	}

	sales, err := h.service.CalculateProductSales(uint(productID))
	if err != nil {
		Error(c, 23006, "计算销量失败")
		return
	}

	Success(c, gin.H{"sales": sales})
}