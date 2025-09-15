package handlers

import (
	"github.com/gin-gonic/gin"
	"server/service"
	"strconv"
)

type CoreRequestLogHandler struct {
	service *service.CoreRequestLogService
}

func NewCoreRequestLogHandler(service *service.CoreRequestLogService) *CoreRequestLogHandler {
	return &CoreRequestLogHandler{service: service}
}

// 分页查询请求日志
func (h *CoreRequestLogHandler) ListRequestLogs(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	tag := c.Query("tag")
	method := c.Query("method")

	logs, total, err := h.service.ListRequestLogs(page, pageSize, tag, method)
	if err != nil {
		Error(c, 9001, "获取日志失败")
		return
	}

	Success(c, gin.H{
		"list":  logs,
		"total": total,
	})
}

// 根据IP查询日志
func (h *CoreRequestLogHandler) GetLogsByIP(c *gin.Context) {
	ip := c.Param("ip")
	if ip == "" {
		InvalidParams(c)
		return
	}

	logs, err := h.service.GetLogsByIP(ip)
	if err != nil {
		Error(c, 9002, "获取日志失败")
		return
	}

	Success(c, logs)
}

// 删除过期日志
func (h *CoreRequestLogHandler) CleanupOldLogs(c *gin.Context) {
	days, err := strconv.Atoi(c.Query("days"))
	if err != nil || days <= 0 {
		InvalidParams(c)
		return
	}

	if err := h.service.CleanupOldLogs(days); err != nil {
		Error(c, 9003, err.Error())
		return
	}

	Success(c, nil)
}