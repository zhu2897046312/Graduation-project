package handlers

import (
	"github.com/gin-gonic/gin"
	"server/models/mypaypal"
	"server/service"
	"strconv"
)

type PaypalWebhookLogsHandler struct {
	service *service.PaypalWebhookLogsService
}

func NewPaypalWebhookLogsHandler(service *service.PaypalWebhookLogsService) *PaypalWebhookLogsHandler {
	return &PaypalWebhookLogsHandler{service: service}
}

// 创建Webhook日志
func (h *PaypalWebhookLogsHandler) CreateWebhookLog(c *gin.Context) {
	var log mypaypal.PaypalWebhookLogs
	if err := c.ShouldBindJSON(&log); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.CreateWebhookLog(&log); err != nil {
		Error(c, 18001, err.Error())
		return
	}

	Success(c, log)
}

// 根据事件ID获取日志
func (h *PaypalWebhookLogsHandler) GetLogByEventID(c *gin.Context) {
	eventID := c.Param("event_id")
	if eventID == "" {
		InvalidParams(c)
		return
	}

	log, err := h.service.GetLogByEventID(eventID)
	if err != nil {
		Error(c, 18002, "日志不存在")
		return
	}

	Success(c, log)
}

// 根据本地订单ID获取日志
func (h *PaypalWebhookLogsHandler) GetLogsByLocalOrder(c *gin.Context) {
	localOrderID := c.Param("local_order_id")
	if localOrderID == "" {
		InvalidParams(c)
		return
	}

	logs, err := h.service.GetLogsByLocalOrderID(localOrderID)
	if err != nil {
		Error(c, 18003, "获取日志失败")
		return
	}

	Success(c, logs)
}

// 根据PayPal订单ID获取日志
func (h *PaypalWebhookLogsHandler) GetLogsByPaypalOrder(c *gin.Context) {
	paypalOrderID := c.Param("paypal_order_id")
	if paypalOrderID == "" {
		InvalidParams(c)
		return
	}

	logs, err := h.service.GetLogsByPaypalOrderID(paypalOrderID)
	if err != nil {
		Error(c, 18004, "获取日志失败")
		return
	}

	Success(c, logs)
}

// 根据事件类型获取日志
func (h *PaypalWebhookLogsHandler) GetLogsByEventType(c *gin.Context) {
	eventType := c.Param("event_type")
	if eventType == "" {
		InvalidParams(c)
		return
	}

	logs, err := h.service.GetLogsByEventType(eventType)
	if err != nil {
		Error(c, 18005, "获取日志失败")
		return
	}

	Success(c, logs)
}

// 更新处理结果
func (h *PaypalWebhookLogsHandler) UpdateProcessResult(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil || id == 0 {
		InvalidParams(c)
		return
	}

	var req struct {
		Result string `json:"result"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.UpdateProcessResult(uint(id), req.Result); err != nil {
		Error(c, 18006, err.Error())
		return
	}

	Success(c, nil)
}