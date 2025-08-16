package handlers

import (
	"github.com/gin-gonic/gin"
	"server/models/paypal"
	"server/service"
)

type PaypalOrderLogsHandler struct {
	service *service.PaypalOrderLogsService
}

func NewPaypalOrderLogsHandler(service *service.PaypalOrderLogsService) *PaypalOrderLogsHandler {
	return &PaypalOrderLogsHandler{service: service}
}

// 创建订单日志
func (h *PaypalOrderLogsHandler) CreateOrderLog(c *gin.Context) {
	var log paypal.PaypalOrderLogs
	if err := c.ShouldBindJSON(&log); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.CreateOrderLog(&log); err != nil {
		Error(c, 17001, err.Error())
		return
	}

	Success(c, log)
}

// 根据本地订单ID获取日志
func (h *PaypalOrderLogsHandler) GetLogsByLocalOrder(c *gin.Context) {
	localOrderID := c.Param("local_order_id")
	if localOrderID == "" {
		InvalidParams(c)
		return
	}

	logs, err := h.service.GetLogsByLocalOrderID(localOrderID)
	if err != nil {
		Error(c, 17002, "获取日志失败")
		return
	}

	Success(c, logs)
}

// 根据PayPal订单ID获取日志
func (h *PaypalOrderLogsHandler) GetLogByPaypalOrder(c *gin.Context) {
	paypalOrderID := c.Param("paypal_order_id")
	if paypalOrderID == "" {
		InvalidParams(c)
		return
	}

	log, err := h.service.GetLogByPaypalOrderID(paypalOrderID)
	if err != nil {
		Error(c, 17003, "日志不存在")
		return
	}

	Success(c, log)
}

// 获取所有订单日志
func (h *PaypalOrderLogsHandler) GetAllOrderLogs(c *gin.Context) {
	logs, err := h.service.GetAllOrderLogs()
	if err != nil {
		Error(c, 17004, "获取日志失败")
		return
	}

	Success(c, logs)
}