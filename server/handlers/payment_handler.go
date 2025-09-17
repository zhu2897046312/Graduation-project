// handlers/payment_handler.go
package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"server/config"
	"server/models/mypaypal"
	"server/models/sp"
	"server/service"
	"server/utils"

	"github.com/gin-gonic/gin"
	"github.com/plutov/paypal/v4"
)

// PaymentRequest 支付请求
type PaymentRequest struct {
	OrderCode string `json:"order_id"`
	PayType   uint8  `json:"pay_type"`
	ReturnURL string `json:"return_url"`
	CancelURL string `json:"cancel_url"`
}

// PaymentResponse 支付响应
type PaymentResponse struct {
	ApproveURL string `json:"approveUrl"`
	Status     string `json:"status"`
	OrderID    string `json:"order_id"`
	PaymentID  string `json:"payment_id"`
	Message    string `json:"message"`
}

type PaymentHandler struct {
	orderService    *service.SpOrderService
	paypalOrderLogs *service.PaypalOrderLogsService
	paypalClient    *paypal.Client
}

func NewPaymentHandler(orderService *service.SpOrderService, paypalOrderLogs *service.PaypalOrderLogsService) *PaymentHandler {
	cfg := config.GlobalConfig.Payment

	// 创建 PayPal 客户端
	client, err := paypal.NewClient(cfg.PayPal.ClientID, cfg.PayPal.Secret, getAPIBase(cfg.PayPal.APIBase))
	if err != nil {
		panic(fmt.Sprintf("Failed to create PayPal client: %v", err))
	}

	// 设置 PayPal API 基础 URL
	client.SetLog(os.Stdout) // 可选：启用日志记录

	return &PaymentHandler{
		orderService:    orderService,
		paypalOrderLogs: paypalOrderLogs,
		paypalClient:    client,
	}
}

// getAPIBase 根据配置返回 PayPal API 基础 URL
func getAPIBase(apiBase string) string {
	// 如果配置中已经是完整的 URL，直接返回
	if apiBase == "https://api.sandbox.paypal.com" || apiBase == "https://api.paypal.com" {
		return apiBase
	}

	return paypal.APIBaseSandBox
}

// CreatePayment 创建支付订单
func (h *PaymentHandler) CreatePayment(c *gin.Context) {
	var req PaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		Error(c, 17000, "无效的请求参数")
		return
	}

	// 根据订单ID类型查询订单
	order, err := h.orderService.GetByVisitorQueryCode(req.OrderCode)
	if err != nil {
		Error(c, 17001, "订单不存在")
		return
	}

	// 检查订单状态
	if order.State != 2 {
		Error(c, 17002, "订单状态不支持支付")
		return
	}

	// 设置返回URL

	req.CancelURL = "http://localhost:3000/order-detail/" + order.VisitorQueryCode

	// 根据支付类型处理
	switch req.PayType {
	case 1: // PayPal支付
		resp, err := h.createPayPalPayment(order, req)
		if err != nil {
			Error(c, 17004, "创建支付订单失败: "+err.Error())
			return
		}
		Success(c, resp)
	case 2: // 信用卡支付
		Error(c, 17005, "暂不支持该支付方式")
	default:
		Error(c, 17006, "不支持的支付类型")
	}
}

// createPayPalPayment 创建PayPal支付订单
func (h *PaymentHandler) createPayPalPayment(order *sp.SpOrder, req PaymentRequest) (*PaymentResponse, error) {
	purchaseUnits := []paypal.PurchaseUnitRequest{
		{
			ReferenceID: order.VisitorQueryCode,
			Amount: &paypal.PurchaseUnitAmount{
				Currency: "USD",
				Value:    utils.FormatPrice(order.PayAmount),
			},
			Description: "Order #" + order.VisitorQueryCode,
		},
	}

	// ⚡ ReturnURL 直接写后端 capture 接口
	returnURL := fmt.Sprintf(
		"http://localhost:8080/api/client/payment/capture-order?redirect=%s",
		"http://localhost:3000/order-detail/"+order.VisitorQueryCode,
	)
	// 创建 PayPal 订单
	createdOrder, err := h.paypalClient.CreateOrder(
		context.Background(),
		paypal.OrderIntentCapture,
		purchaseUnits,
		nil,
		&paypal.ApplicationContext{
			ReturnURL:          returnURL,
			CancelURL:          "http://localhost:3000/order-detail/" + order.VisitorQueryCode,
			BrandName:          "Your Store Name",
			UserAction:         "PAY_NOW",
			ShippingPreference: paypal.ShippingPreferenceNoShipping,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("创建PayPal订单失败: %v", err)
	}

	// 查找批准URL
	var approveURL string
	for _, link := range createdOrder.Links {
		if link.Rel == "approve" {
			approveURL = link.Href
			break
		}
	}
	if approveURL == "" {
		return nil, fmt.Errorf("未找到PayPal批准URL")
	}

	// 保存本地订单日志
	if err := h.paypalOrderLogs.CreateOrderLog(&mypaypal.PaypalOrderLogs{
		LocalOrderID:  req.OrderCode,
		PaypalOrderID: createdOrder.ID,
	}); err != nil {
		return nil, fmt.Errorf("更新订单支付信息失败: %v", err)
	}

	return &PaymentResponse{
		ApproveURL: approveURL,
		Status:     string(createdOrder.Status),
		OrderID:    req.OrderCode,
		PaymentID:  createdOrder.ID,
		Message:    "PayPal支付订单创建成功",
	}, nil
}


// CapturePayment 捕获支付（确认支付）
func (h *PaymentHandler) CapturePayment(c *gin.Context) {
    paymentID := c.Query("token")   // PayPal 会带 token=订单ID
    redirectURL := c.Query("redirect")

    if paymentID == "" {
        Error(c, 17007, "支付ID不能为空")
        return
    }

    captureResult, err := h.paypalClient.CaptureOrder(context.Background(), paymentID, paypal.CaptureOrderRequest{})
    if err != nil || captureResult.Status != "COMPLETED" {
        if redirectURL != "" {
            c.Redirect(302, redirectURL+"?status=failed")
            return
        }
        Error(c, 17009, "捕获支付失败: "+err.Error())
        return
    }

    // 更新本地订单状态
    order, _ := h.paypalOrderLogs.GetLogByPaypalOrderID(paymentID)
    if order != nil {
        h.orderService.UpdateOrderState(order.LocalOrderID, 1, "PayPal支付成功")
    }

    if redirectURL != "" {
        c.Redirect(302, redirectURL+"?status=success")
        return
    }

    Success(c, gin.H{"status": "success"})
}


// PaymentWebhook PayPal支付webhook处理
type PaypalWebhookPayload struct {
	paypal.Event
	Resource json.RawMessage `json:"resource"`
}

func (h *PaymentHandler) PaymentWebhook(c *gin.Context) {
	var payload PaypalWebhookPayload
	if err := c.BindJSON(&payload); err != nil {
		InvalidParams(c)
		return
	}

	switch payload.EventType {
	case "PAYMENT.CAPTURE.COMPLETED":
		var capture paypal.Capture
		if err := json.Unmarshal(payload.Resource, &capture); err == nil {
			if order, err := h.paypalOrderLogs.GetLogByPaypalOrderID(capture.ID); err == nil {
				h.orderService.UpdateOrderState(order.LocalOrderID, 1, "PayPal支付成功")
			}
		}

	case "PAYMENT.CAPTURE.DENIED", "PAYMENT.CAPTURE.FAILED":
		var capture paypal.Capture
		if err := json.Unmarshal(payload.Resource, &capture); err == nil {
			if order, err := h.paypalOrderLogs.GetLogByPaypalOrderID(capture.ID); err == nil {
				h.orderService.UpdateOrderState(order.LocalOrderID, 4, "支付失败")
			}
		}
	}

	Success(c, gin.H{"status": "success"})
}

// GetPaymentStatus 获取支付状态
func (h *PaymentHandler) GetPaymentStatus(c *gin.Context) {
	orderIDStr := c.Param("id")
	if orderIDStr == "" {
		Error(c, 17002, "订单ID不能为空")
		return
	}

	// 使用 PayPal API 获取订单详情
	orderDetails, err := h.paypalClient.GetOrder(context.Background(), orderIDStr)
	if err != nil {
		Error(c, 17013, "获取支付状态失败: "+err.Error())
		return
	}

	// 根据 PayPal 订单 ID 查找本地订单记录
	orderLog, err := h.paypalOrderLogs.GetLogByPaypalOrderID(orderIDStr)
	if err != nil {
		Error(c, 17001, "订单不存在")
		return
	}

	// 获取本地订单信息
	localOrderInfo, err := h.orderService.GetByVisitorQueryCode(orderLog.LocalOrderID)
	if err != nil {
		Error(c, 17001, "订单不存在")
		return
	}

	Success(c, gin.H{
		"order_id":      localOrderInfo.ID,
		"status":        localOrderInfo.State,
		"paid":          localOrderInfo.State == 1,
		"amount":        utils.FormatPrice(localOrderInfo.PayAmount),
		"payment_id":    orderLog.PaypalOrderID,
		"paypal_status": string(orderDetails.Status),
	})
}
