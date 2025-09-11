// handlers/payment_handler.go
package handlers

import (
	"server/models/sp"
	"server/service"
	"server/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// PaymentSimulateReq 模拟支付请求
type PaymentSimulateReq struct {
	OrderID   string   `json:"order_id" form:"order_id" binding:"required"`
	PayType   uint8  `json:"pay_type" form:"pay_type" binding:"required"`
	Success   bool   `json:"success" form:"success"` // 是否模拟成功支付，默认true
	ReturnURL string `json:"return_url" form:"return_url"`
}

// PaymentSimulateResp 模拟支付响应
type PaymentSimulateResp struct {
	ApproveURL string `json:"approve_url"`
	Status     string `json:"status"`
	OrderID    string `json:"order_id"`
	Message    string `json:"message"`
}
type PaymentHandler struct {
	orderService *service.SpOrderService
}

func NewPaymentHandler(orderService *service.SpOrderService) *PaymentHandler {
	return &PaymentHandler{
		orderService: orderService,
	}
}

// SimulatePayment 模拟支付
func (h *PaymentHandler) SimulatePayment(c *gin.Context) {
	var req PaymentSimulateReq

	// 支持JSON和表单两种方式
	if c.ContentType() == "application/json" {
		if err := c.ShouldBindJSON(&req); err != nil {
			InvalidParams(c)
			return
		}
	} else {
		if err := c.ShouldBind(&req); err != nil {
			InvalidParams(c)
			return
		}
	}

	// 设置默认值
	if !req.Success {
		req.Success = true // 默认模拟成功支付
	}

	// 根据订单ID类型查询订单
	var order *sp.SpOrder
	var err error

	// 尝试将订单ID解析为uint（数字ID）
	order, err = h.orderService.GetByVisitorQueryCode(req.OrderID)

	if err != nil {
		Error(c, 17001, "订单不存在")
		return
	}

	// 检查订单状态
	if order.State != 2 {
		Error(c, 17002, "订单状态不支持支付")
		return
	}

	var resp PaymentSimulateResp

	if req.Success {
		// 模拟成功支付
		resp = h.simulateSuccessPayment(c, order, req)
	} else {
		// 模拟支付失败
		resp = h.simulateFailedPayment(c, order, req)
	}

	Success(c, resp)
}

// simulateSuccessPayment 模拟成功支付
func (h *PaymentHandler) simulateSuccessPayment(c *gin.Context, order *sp.SpOrder, req PaymentSimulateReq) PaymentSimulateResp {
	// 更新订单状态为已支付
	err := h.orderService.UpdateOrderState(order.ID, 1, "模拟支付成功")
	if err != nil {
		Error(c, 17003, "更新订单状态失败")
	}

	// 构建返回URL
	returnURL := req.ReturnURL
	if returnURL == "" {
		// 默认的成功页面URL
		returnURL = "/payment/success?order_id=" + strconv.FormatUint(uint64(order.ID), 10) +
			"&code=" + order.VisitorQueryCode +
			"&amount=" + utils.FormatPrice(order.PayAmount) +
			"&timestamp=" + strconv.FormatInt(time.Now().Unix(), 10)
	}

	return PaymentSimulateResp{
		ApproveURL: returnURL,
		Status:     "success",
		OrderID:    req.OrderID, // 返回客户端传入的原始订单ID
		Message:    "支付成功",
	}
}

// simulateFailedPayment 模拟支付失败
func (h *PaymentHandler) simulateFailedPayment(c *gin.Context, order *sp.SpOrder, req PaymentSimulateReq) PaymentSimulateResp {
	// 构建失败返回URL
	returnURL := req.ReturnURL
	if returnURL == "" {
		returnURL = "/payment/failed?order_id=" + strconv.FormatUint(uint64(order.ID), 10) +
			"&code=" + order.VisitorQueryCode +
			"&reason=simulated_failure"
	}

	return PaymentSimulateResp{
		ApproveURL: returnURL,
		Status:     "failed",
		OrderID:    req.OrderID, // 返回客户端传入的原始订单ID
		Message:    "模拟支付失败",
	}
}

// PaymentCallback 支付回调（用于前端重定向后的处理）
func (h *PaymentHandler) PaymentCallback(c *gin.Context) {
	orderIDStr := c.Query("order_id")
	code := c.Query("code")
	status := c.Query("status")

	orderID, err := strconv.ParseUint(orderIDStr, 10, 32)
	if err != nil {
		c.HTML(400, "payment_error.html", gin.H{
			"message": "无效的订单ID",
		})
		return
	}

	// 验证订单和查询码
	order, err := h.orderService.GetOrderByID(uint(orderID))
	if err != nil || order.VisitorQueryCode != code {
		c.HTML(400, "payment_error.html", gin.H{
			"message": "订单验证失败",
		})
		return
	}

	// 根据状态显示不同的页面
	if status == "success" {
		c.HTML(200, "payment_success.html", gin.H{
			"order":     order,
			"amount":    utils.FormatPrice(order.PayAmount),
			"orderCode": order.Code,
		})
	} else {
		c.HTML(200, "payment_failed.html", gin.H{
			"order":   order,
			"message": "支付失败，请重试",
		})
	}
}

// GetPaymentStatus 获取支付状态
func (h *PaymentHandler) GetPaymentStatus(c *gin.Context) {
	orderIDStr := c.Param("id")
	orderID, err := strconv.ParseUint(orderIDStr, 10, 32)
	if err != nil {
		Error(c, 17003, "无效的订单ID")
		return
	}

	order, err := h.orderService.GetOrderByID(uint(orderID))
	if err != nil {
		Error(c, 17001, "订单不存在")
		return
	}

	Success(c, gin.H{
		"order_id": order.ID,
		"status":   order.State,
		"paid":     order.State == 1,
		"amount":   utils.FormatPrice(order.PayAmount),
	})
}
