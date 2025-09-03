package handlers

import (
	"server/models/sp"
	"server/service"
	"server/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ListSpOrderRefundRequest struct {
	RefundNo  string `json:"refund_no"`
	OrderCode string `json:"order_code"`
	Status    uint   `json:"status"`
	NikeName  string `json:"nikename"`
	Email     string `json:"email"`
}

type SpOrderRefundListVo struct {
	ID           uint    `json:"id" schema:"主键"`
	RefundNo     string  `json:"refund_no" schema:"退款单号"`
	Status       uint8   `json:"status" schema:"退款状态"`
	RefundAmount float64 `json:"refund_amount" schema:"退款金额"`
	RefundTime   string  `json:"refund_time" schema:"退款时间"`
	OrderCode    string  `json:"order_code" schema:"订单编号"`
	Reason       string  `json:"reason" schema:"退款原因"`
	Name         string  `json:"name" schema:"用户昵称"`
	Email        string  `json:"email" schema:"用户邮箱"`
}
type SpOrderRefundHandler struct {
	service      *service.SpOrderRefundService
	orderService *service.SpOrderService
}

func NewSpOrderRefundHandler(service *service.SpOrderRefundService, orderService *service.SpOrderService) *SpOrderRefundHandler {
	return &SpOrderRefundHandler{
		service:      service,
		orderService: orderService,
	}
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
	req := c.Query("id")
	if req == "" {
		InvalidParams(c)
		return
	}
	orderID := utils.ConvertToUint(req)
	if orderID == 0 {
		InvalidParams(c)
		return
	}
	refund, err := h.service.GetRefundByOrderIDOne(orderID)
	if err != nil {
		Error(c, 26003, "退款记录不存在")
		return
	}

	Success(c, gin.H{
		"refund_no":     refund.RefundNo,
		"status":        refund.Status,
		"refund_time":   refund.RefundTime,
		"reason":        refund.Reason,
		"images":        refund.Images,
		"refund_amount": refund.RefundAmount,
		"created_time":  refund.CreatedTime,
	})
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

func (h *SpOrderRefundHandler) ListSpOrderRefund(c *gin.Context) {
	var req ListSpOrderRefundRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}
	orderParam := sp.ListOrdersQueryParam{
		Code:     req.OrderCode,
		NikeName: req.NikeName,
		Email:    req.Email,
	}
	orders, _, err := h.orderService.List(orderParam)
	if err != nil {
		Error(c, 26007, err.Error())
		return
	}

	// 提取订单ID列表，并构建ID到订单的映射，避免重复查询
	orderIDs := make([]uint, 0, len(orders))
	orderMap := make(map[uint]sp.SpOrder, len(orders)) // 假设订单实体为sp.SpOrder
	for _, order := range orders {
		orderIDs = append(orderIDs, order.ID)
		orderMap[order.ID] = order // 存储ID与订单的映射关系
	}

	refunds, total, err_1 := h.service.ListWithPagination(orderIDs, req.RefundNo, req.Status)
	if err_1 != nil {
		Error(c, 26008, err_1.Error())
		return
	}

	// 转换为VO列表
	refundVoList := make([]SpOrderRefundListVo, 0, len(refunds))
	for _, refund := range refunds {
		// 基础属性复制
		vo := SpOrderRefundListVo{
			ID:           refund.ID,
			RefundNo:     refund.RefundNo,
			Status:       refund.Status,
			RefundAmount: float64(refund.RefundAmount),
			RefundTime:   refund.RefundTime.Format("2006-01-02 15:04:05"),
			Reason:       refund.Reason,
		}

		// 直接从已查询的订单映射中获取信息，无需再次查询数据库
		if order, exists := orderMap[refund.OrderID]; exists {
			vo.OrderCode = order.Code // 直接访问订单的Code字段
			vo.Name = order.Nickname  // 直接访问订单的NickName字段
			vo.Email = order.Email    // 直接访问订单的Email字段
		}

		refundVoList = append(refundVoList, vo)
	}
	// 返回结果
	Success(c, gin.H{
		"total": total,
		"list":  refundVoList,
	})
}
