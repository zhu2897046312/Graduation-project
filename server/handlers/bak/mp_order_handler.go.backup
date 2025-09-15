package handlers

import (
	"server/models/common"
	"server/models/mp"
	"server/service"
	"server/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MpOrderHandler struct {
	service *service.MpOrderService
}

func NewMpOrderHandler(service *service.MpOrderService) *MpOrderHandler {
	return &MpOrderHandler{service: service}
}

// 创建订单
func (h *MpOrderHandler) CreateOrder(c *gin.Context) {
	var order mp.MpOrder
	if err := c.ShouldBindJSON(&order); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.CreateOrder(&order); err != nil {
		Error(c, 11001, err.Error())
		return
	}

	Success(c, order)
}

// 更新订单
func (h *MpOrderHandler) UpdateOrder(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		InvalidParams(c)
		return
	}

	var order mp.MpOrder
	if err := c.ShouldBindJSON(&order); err != nil {
		InvalidParams(c)
		return
	}
	order.ID = common.MyID(utils.ConvertToUint(id))

	if err := h.service.UpdateOrder(&order); err != nil {
		Error(c, 11002, err.Error())
		return
	}

	Success(c, order)
}

// 获取订单详情
func (h *MpOrderHandler) GetOrder(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		InvalidParams(c)
		return
	}

	order, err := h.service.GetOrderByID(common.MyID(utils.ConvertToUint(id)))
	if err != nil {
		Error(c, 11003, "订单不存在")
		return
	}

	Success(c, order)
}

// 获取用户订单列表
func (h *MpOrderHandler) GetOrdersByUser(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil || userID <= 0 {
		InvalidParams(c)
		return
	}

	orders, err := h.service.GetOrdersByUserID(common.MyID(userID))
	if err != nil {
		Error(c, 11004, "获取订单失败")
		return
	}

	Success(c, orders)
}

// 根据状态获取订单列表
func (h *MpOrderHandler) GetOrdersByState(c *gin.Context) {
	state, err := strconv.ParseInt(c.Query("state"), 10, 8)
	if err != nil {
		InvalidParams(c)
		return
	}

	orders, err := h.service.GetOrdersByState(int8(state))
	if err != nil {
		Error(c, 11005, "获取订单失败")
		return
	}

	Success(c, orders)
}

// 更新订单状态
func (h *MpOrderHandler) UpdateOrderState(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		InvalidParams(c)
		return
	}

	var req struct {
		State int8 `json:"state"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.UpdateOrderState(common.MyID(utils.ConvertToUint(id)), req.State); err != nil {
		Error(c, 11006, err.Error())
		return
	}

	Success(c, nil)
}

// 根据第三方支付ID获取订单
func (h *MpOrderHandler) GetOrderByThirdID(c *gin.Context) {
	thirdID := c.Param("third_id")
	if thirdID == "" {
		InvalidParams(c)
		return
	}

	order, err := h.service.GetOrderByThirdID(common.MyID(utils.ConvertToUint(thirdID)))
	if err != nil {
		Error(c, 11007, "订单不存在")
		return
	}

	Success(c, order)
}