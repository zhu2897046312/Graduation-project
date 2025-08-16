package handlers

import (
	"github.com/gin-gonic/gin"
	"server/models/sp"
	"server/service"
	"strconv"
)

type SpUserCartHandler struct {
	service *service.SpUserCartService
}

func NewSpUserCartHandler(service *service.SpUserCartService) *SpUserCartHandler {
	return &SpUserCartHandler{service: service}
}

// AddToCart 添加到购物车
func (h *SpUserCartHandler) AddToCart(c *gin.Context) {
	var cart sp.SpUserCart
	if err := c.ShouldBindJSON(&cart); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.AddToCart(&cart); err != nil {
		Error(c, 3401, err.Error())
		return
	}

	Success(c, cart)
}

// UpdateCartItem 更新购物车项
func (h *SpUserCartHandler) UpdateCartItem(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		InvalidParams(c)
		return
	}

	var cart sp.SpUserCart
	if err := c.ShouldBindJSON(&cart); err != nil {
		InvalidParams(c)
		return
	}
	cart.ID = uint(id)

	if err := h.service.UpdateCartItem(&cart); err != nil {
		Error(c, 3402, err.Error())
		return
	}

	Success(c, cart)
}

// GetCartItems 获取购物车列表
func (h *SpUserCartHandler) GetCartItems(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil || userID == 0 {
		InvalidParams(c)
		return
	}

	cartItems, err := h.service.GetCartItemsByUserID(uint(userID))
	if err != nil {
		Error(c, 3403, err.Error())
		return
	}

	Success(c, cartItems)
}

// DeleteCartItem 删除购物车项
func (h *SpUserCartHandler) DeleteCartItem(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		InvalidParams(c)
		return
	}

	if err := h.service.DeleteCartItem(uint(id)); err != nil {
		Error(c, 3404, err.Error())
		return
	}

	Success(c, nil)
}

// ClearCart 清空购物车
func (h *SpUserCartHandler) ClearCart(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil || userID == 0 {
		InvalidParams(c)
		return
	}

	if err := h.service.ClearCart(uint(userID)); err != nil {
		Error(c, 3405, err.Error())
		return
	}

	Success(c, nil)
}