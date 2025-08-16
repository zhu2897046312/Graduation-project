package handlers

import (
	"github.com/gin-gonic/gin"
	"server/models/mp"
	"server/service"
	"strconv"
)

type MpPayConfigHandler struct {
	service *service.MpPayConfigService
}

func NewMpPayConfigHandler(service *service.MpPayConfigService) *MpPayConfigHandler {
	return &MpPayConfigHandler{service: service}
}

// 创建支付配置
func (h *MpPayConfigHandler) CreatePayConfig(c *gin.Context) {
	var config mp.MpPayConfig
	if err := c.ShouldBindJSON(&config); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.CreatePayConfig(&config); err != nil {
		Error(c, 12001, err.Error())
		return
	}

	Success(c, config)
}

// 更新支付配置
func (h *MpPayConfigHandler) UpdatePayConfig(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		InvalidParams(c)
		return
	}

	var config mp.MpPayConfig
	if err := c.ShouldBindJSON(&config); err != nil {
		InvalidParams(c)
		return
	}
	config.ID = id

	if err := h.service.UpdatePayConfig(&config); err != nil {
		Error(c, 12002, err.Error())
		return
	}

	Success(c, config)
}

// 获取支付配置详情
func (h *MpPayConfigHandler) GetPayConfig(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		InvalidParams(c)
		return
	}

	config, err := h.service.GetPayConfigByID(id)
	if err != nil {
		Error(c, 12003, "支付配置不存在")
		return
	}

	Success(c, config)
}

// 获取所有启用的支付配置
func (h *MpPayConfigHandler) GetActivePayConfigs(c *gin.Context) {
	configs, err := h.service.GetActivePayConfigs()
	if err != nil {
		Error(c, 12004, "获取支付配置失败")
		return
	}

	Success(c, configs)
}

// 根据Code获取支付配置
func (h *MpPayConfigHandler) GetPayConfigByCode(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		InvalidParams(c)
		return
	}

	config, err := h.service.GetPayConfigByCode(code)
	if err != nil {
		Error(c, 12005, "支付配置不存在")
		return
	}

	Success(c, config)
}

// 更新支付配置状态
func (h *MpPayConfigHandler) UpdatePayConfigState(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
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

	if err := h.service.UpdatePayConfigState(id, req.State); err != nil {
		Error(c, 12006, err.Error())
		return
	}

	Success(c, nil)
}