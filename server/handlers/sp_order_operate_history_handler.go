package handlers

import (
	"github.com/gin-gonic/gin"
	"server/models/sp"
	"server/service"
	"strconv"
)

type SpOrderOperateHistoryHandler struct {
	service *service.SpOrderOperateHistoryService
}

func NewSpOrderOperateHistoryHandler(service *service.SpOrderOperateHistoryService) *SpOrderOperateHistoryHandler {
	return &SpOrderOperateHistoryHandler{service: service}
}

// 创建操作历史
func (h *SpOrderOperateHistoryHandler) CreateHistory(c *gin.Context) {
	var history sp.SpOrderOperateHistory
	if err := c.ShouldBindJSON(&history); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.CreateHistory(&history); err != nil {
		Error(c, 24001, err.Error())
		return
	}

	Success(c, history)
}

// 根据订单ID获取操作历史
func (h *SpOrderOperateHistoryHandler) GetHistoriesByOrder(c *gin.Context) {
	orderID, err := strconv.ParseUint(c.Param("order_id"), 10, 32)
	if err != nil || orderID == 0 {
		InvalidParams(c)
		return
	}

	histories, err := h.service.GetHistoriesByOrderID(uint(orderID))
	if err != nil {
		Error(c, 24002, "获取操作历史失败")
		return
	}

	Success(c, histories)
}

// 根据操作人获取操作历史
func (h *SpOrderOperateHistoryHandler) GetHistoriesByUser(c *gin.Context) {
	user := c.Param("user")
	if user == "" {
		InvalidParams(c)
		return
	}

	histories, err := h.service.GetHistoriesByOperateUser(user)
	if err != nil {
		Error(c, 24003, "获取操作历史失败")
		return
	}

	Success(c, histories)
}