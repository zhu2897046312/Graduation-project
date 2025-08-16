package handlers

import (
	"github.com/gin-gonic/gin"
	"server/models/sp"
	"server/service"
	"strconv"
)

type SpProductHandler struct {
	service *service.SpProductService
}

func NewSpProductHandler(service *service.SpProductService) *SpProductHandler {
	return &SpProductHandler{service: service}
}

// CreateProduct 创建商品
func (h *SpProductHandler) CreateProduct(c *gin.Context) {
	var product sp.SpProduct
	if err := c.ShouldBindJSON(&product); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.CreateProduct(&product); err != nil {
		Error(c, 3101, err.Error())
		return
	}

	Success(c, product)
}

// UpdateProduct 更新商品
func (h *SpProductHandler) UpdateProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		InvalidParams(c)
		return
	}

	var product sp.SpProduct
	if err := c.ShouldBindJSON(&product); err != nil {
		InvalidParams(c)
		return
	}
	product.ID = uint(id)

	if err := h.service.UpdateProduct(&product); err != nil {
		Error(c, 3102, err.Error())
		return
	}

	Success(c, product)
}

// GetProduct 获取商品详情
func (h *SpProductHandler) GetProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		InvalidParams(c)
		return
	}

	product, err := h.service.GetProductByID(uint(id))
	if err != nil {
		Error(c, 3103, "商品不存在")
		return
	}

	Success(c, product)
}

// ListProducts 分页获取商品
func (h *SpProductHandler) ListProducts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	products, total, err := h.service.ListProducts(page, pageSize)
	if err != nil {
		ServerError(c, err)
		return
	}

	Success(c, gin.H{
		"list":  products,
		"total": total,
	})
}

// UpdateStock 更新库存
func (h *SpProductHandler) UpdateStock(c *gin.Context) {
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
		Error(c, 3104, err.Error())
		return
	}

	Success(c, nil)
}