package handlers

import (
	"github.com/gin-gonic/gin"
	"server/models/mp"
	"server/service"
	"strconv"
)

type MpProductHandler struct {
	service *service.MpProductService
}

func NewMpProductHandler(service *service.MpProductService) *MpProductHandler {
	return &MpProductHandler{service: service}
}

// 创建产品
func (h *MpProductHandler) CreateProduct(c *gin.Context) {
	var product mp.MpProduct
	if err := c.ShouldBindJSON(&product); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.CreateProduct(&product); err != nil {
		Error(c, 13001, err.Error())
		return
	}

	Success(c, product)
}

// 更新产品
// func (h *MpProductHandler) UpdateProduct(c *gin.Context) {
// 	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
// 	if err != nil || id <= 0 {
// 		InvalidParams(c)
// 		return
// 	}

// 	var product mp.MpProduct
// 	if err := c.ShouldBindJSON(&product); err != nil {
// 		InvalidParams(c)
// 		return
// 	}
// 	product.ID = id

// 	if err := h.service.UpdateProduct(&product); err != nil {
// 		Error(c, 13002, err.Error())
// 		return
// 	}

// 	Success(c, product)
// }

// // 获取产品详情
// func (h *MpProductHandler) GetProduct(c *gin.Context) {
// 	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
// 	if err != nil || id <= 0 {
// 		InvalidParams(c)
// 		return
// 	}

// 	product, err := h.service.GetProductByID(id)
// 	if err != nil {
// 		Error(c, 13003, "产品不存在")
// 		return
// 	}

// 	Success(c, product)
// }

// 根据类型获取产品列表
func (h *MpProductHandler) GetProductsByType(c *gin.Context) {
	productType, err := strconv.ParseInt(c.Query("type"), 10, 8)
	if err != nil {
		InvalidParams(c)
		return
	}

	products, err := h.service.GetProductsByType(int8(productType))
	if err != nil {
		Error(c, 13004, "获取产品失败")
		return
	}

	Success(c, products)
}

// 根据终端类型获取产品列表
func (h *MpProductHandler) GetProductsByTerminal(c *gin.Context) {
	terminalType, err := strconv.ParseInt(c.Query("terminal"), 10, 8)
	if err != nil {
		InvalidParams(c)
		return
	}

	products, err := h.service.GetProductsByTerminal(int8(terminalType))
	if err != nil {
		Error(c, 13005, "获取产品失败")
		return
	}

	Success(c, products)
}

// 根据产品代码获取产品
func (h *MpProductHandler) GetProductByCode(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		InvalidParams(c)
		return
	}

	product, err := h.service.GetProductByCode(code)
	if err != nil {
		Error(c, 13006, "产品不存在")
		return
	}

	Success(c, product)
}

// // 更新产品状态
// func (h *MpProductHandler) UpdateProductState(c *gin.Context) {
// 	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
// 	if err != nil || id <= 0 {
// 		InvalidParams(c)
// 		return
// 	}

// 	var req struct {
// 		State int `json:"state"`
// 	}
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		InvalidParams(c)
// 		return
// 	}

// 	if err := h.service.UpdateProductState(id, req.State); err != nil {
// 		Error(c, 13007, err.Error())
// 		return
// 	}

// 	Success(c, nil)
// }