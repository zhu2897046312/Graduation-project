package handlers

import (
	"github.com/gin-gonic/gin"
	"server/models/core"
	"server/service"
	"strconv"
)

type CoreConfigHandler struct {
	service *service.CoreConfigService
}

func NewCoreConfigHandler(service *service.CoreConfigService) *CoreConfigHandler {
	return &CoreConfigHandler{service: service}
}

// 创建配置
func (h *CoreConfigHandler) CreateConfig(c *gin.Context) {
	var config core.CoreConfig
	if err := c.ShouldBindJSON(&config); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.CreateConfig(&config); err != nil {
		Error(c, 6001, err.Error())
		return
	}

	Success(c, config)
}

// 更新配置
func (h *CoreConfigHandler) UpdateConfig(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		InvalidParams(c)
		return
	}

	var config core.CoreConfig
	if err := c.ShouldBindJSON(&config); err != nil {
		InvalidParams(c)
		return
	}
	config.ID = id

	if err := h.service.UpdateConfig(&config); err != nil {
		Error(c, 6002, err.Error())
		return
	}

	Success(c, config)
}

// 获取配置详情
func (h *CoreConfigHandler) GetConfigByKey(c *gin.Context) {
	key := c.Query("key")
	if key == "" {
		InvalidParams(c)
		return
	}

	config, err := h.service.GetConfigByKey(key)
	if err != nil {
		Error(c, 6003, "配置不存在")
		return
	}

	Success(c, config)
}

// 获取所有配置
func (h *CoreConfigHandler) GetAllConfigs(c *gin.Context) {
	configs, err := h.service.GetAllConfigs()
	if err != nil {
		Error(c, 6004, "获取配置失败")
		return
	}

	Success(c, configs)
}

// 批量更新配置
func (h *CoreConfigHandler) BatchUpdateConfigs(c *gin.Context) {
	var configs []core.CoreConfig
	if err := c.ShouldBindJSON(&configs); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.BatchUpdateConfigs(configs); err != nil {
		Error(c, 6005, err.Error())
		return
	}

	Success(c, nil)
}