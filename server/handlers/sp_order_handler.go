package handlers

import (
	"github.com/gin-gonic/gin"
	"server/models/sp"
	"server/service"
	"strconv"
)

type SpOrderHandler struct {
	service *service.SpOrderService
}

func NewSpOrderHandler(service *service.SpOrderService) *SpOrderHandler {
	return &SpOrderHandler{service: service}
}

// 创建订单
func (h *SpOrderHandler) CreateOrder(c *gin.Context) {
	var order sp.SpOrder
	if err := c.ShouldBindJSON(&order); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.CreateOrder(&order); err != nil {
		Error(c, 27001, err.Error())
		return
	}

	Success(c, order)
}

// 更新订单
func (h *SpOrderHandler) UpdateOrder(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil || id == 0 {
		InvalidParams(c)
		return
	}

	var order sp.SpOrder
	if err := c.ShouldBindJSON(&order); err != nil {
		InvalidParams(c)
		return
	}
	order.ID = uint(id)

	if err := h.service.UpdateOrder(&order); err != nil {
		Error(c, 27002, err.Error())
		return
	}

	Success(c, order)
}

// 获取订单详情
func (h *SpOrderHandler) GetOrder(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil || id == 0 {
		InvalidParams(c)
		return
	}

	order, err := h.service.GetOrderByID(uint(id))
	if err != nil {
		Error(c, 27003, "订单不存在")
		return
	}

	Success(c, order)
}

// 根据订单号获取订单
func (h *SpOrderHandler) GetOrderByCode(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		InvalidParams(c)
		return
	}

	order, err := h.service.GetOrderByCode(code)
	if err != nil {
		Error(c, 27004, "订单不存在")
		return
	}

	Success(c, order)
}

// 根据用户ID获取订单列表
func (h *SpOrderHandler) GetOrdersByUser(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 32)
	if err != nil || userID == 0 {
		InvalidParams(c)
		return
	}

	orders, err := h.service.GetOrdersByUserID(uint(userID))
	if err != nil {
		Error(c, 27005, "获取订单列表失败")
		return
	}

	Success(c, orders)
}

// 根据状态获取订单列表
func (h *SpOrderHandler) GetOrdersByState(c *gin.Context) {
	state, err := strconv.ParseUint(c.Query("state"), 10, 8)
	if err != nil {
		InvalidParams(c)
		return
	}

	orders, err := h.service.GetOrdersByState(uint8(state))
	if err != nil {
		Error(c, 27006, "获取订单列表失败")
		return
	}

	Success(c, orders)
}

// 更新订单状态
func (h *SpOrderHandler) UpdateOrderState(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil || id == 0 {
		InvalidParams(c)
		return
	}

	var req struct {
		State uint8 `json:"state"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.UpdateOrderState(uint(id), req.State); err != nil {
		Error(c, 27007, err.Error())
		return
	}

	Success(c, nil)
}

// 更新物流信息
func (h *SpOrderHandler) UpdateDeliveryInfo(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil || id == 0 {
		InvalidParams(c)
		return
	}

	var req struct {
		Company string `json:"company"`
		SN      string `json:"sn"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.UpdateDeliveryInfo(uint(id), req.Company, req.SN); err != nil {
		Error(c, 27008, err.Error())
		return
	}

	Success(c, nil)
}