package handlers

import (
	"github.com/gin-gonic/gin"
	"server/models/sp"
	"server/service"
	"strconv"
)

type SpOrderRefundHandler struct {
	service *service.SpOrderRefundService
}

func NewSpOrderRefundHandler(service *service.SpOrderRefundService) *SpOrderRefundHandler {
	return &SpOrderRefundHandler{service: service}
}

// 创建退款记录
func (h *SpOrderRefundHandler) CreateRefund(c *gin.Context) {
	var refund sp.SpOrderRefund
	if err := c.ShouldBindJSON(&refund); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.CreateRefund(&refund); err != nil {
		Error(c, 26001, err.Error())
		return
	}

	Success(c, refund)
}

// 更新退款记录
func (h *SpOrderRefundHandler) UpdateRefund(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil || id == 0 {
		InvalidParams(c)
		return
	}

	var refund sp.SpOrderRefund
	if err := c.ShouldBindJSON(&refund); err != nil {
		InvalidParams(c)
		return
	}
	refund.ID = uint(id)

	if err := h.service.UpdateRefund(&refund); err != nil {
		Error(c, 26002, err.Error())
		return
	}

	Success(c, refund)
}

// 根据订单ID获取退款记录
func (h *SpOrderRefundHandler) GetRefundByOrder(c *gin.Context) {
	orderID, err := strconv.ParseUint(c.Param("order_id"), 10, 32)
	if err != nil || orderID == 0 {
		InvalidParams(c)
		return
	}

	refund, err := h.service.GetRefundByOrderID(uint(orderID))
	if err != nil {
		Error(c, 26003, "退款记录不存在")
		return
	}

	Success(c, refund)
}

// 根据退款单号获取退款记录
func (h *SpOrderRefundHandler) GetRefundByRefundNo(c *gin.Context) {
	refundNo := c.Param("refund_no")
	if refundNo == "" {
		InvalidParams(c)
		return
	}

	refund, err := h.service.GetRefundByRefundNo(refundNo)
	if err != nil {
		Error(c, 26004, "退款记录不存在")
		return
	}

	Success(c, refund)
}

// 更新退款状态
func (h *SpOrderRefundHandler) UpdateRefundStatus(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil || id == 0 {
		InvalidParams(c)
		return
	}

	var req struct {
		Status uint8 `json:"status"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.UpdateRefundStatus(uint(id), req.Status); err != nil {
		Error(c, 26005, err.Error())
		return
	}

	Success(c, nil)
}

// 更新退款金额
func (h *SpOrderRefundHandler) UpdateRefundAmount(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil || id == 0 {
		InvalidParams(c)
		return
	}

	var req struct {
		Amount float64 `json:"amount"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.UpdateRefundAmount(uint(id), req.Amount); err != nil {
		Error(c, 26006, err.Error())
		return
	}

	Success(c, nil)
}