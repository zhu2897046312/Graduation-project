package handlers

import (
	"fmt"
	"server/models/common"
	"server/models/core"
	"server/service"
	"server/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type CoreConfigHandler struct {
	service *service.CoreConfigService
}
type Request struct {
	Logo           string `json:"logo"`
	SEODescription string `json:"seo_description"`
	SEOKeyword     string `json:"seo_keyword"`
	SEOTitle       string `json:"seo_title"`
	Title          string `json:"title"`
}

type MarketRequest struct {
	Exchange interface{} `json:"exchange"`
	Freight  interface{} `json:"freight"`
	Original interface{} `json:"original"`
}

func NewCoreConfigHandler(service *service.CoreConfigService) *CoreConfigHandler {
	return &CoreConfigHandler{service: service}
}

// 保存站点信息
// 保存站点信息 - 根据配置键创建或更新配置
func (h *CoreConfigHandler) SaveSiteInfo(c *gin.Context) {
	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}

	if req.Logo != "" {
		result, _ := h.service.GetConfigByKey("logo")
		config := core.CoreConfig{
			ConfigKey:   "logo",
			ConfigValue: req.Logo,
			UpdatedTime: time.Now(),
		}
		if result.ConfigValue == "" {
			config.CreatedTime = time.Now()
			h.service.CreateConfig(&config)
		} else {
			h.service.UpdateConfig(&config)
		}
	}

	if req.SEOKeyword != "" {
		result, _ := h.service.GetConfigByKey("seo_keyword")
		config := core.CoreConfig{
			ConfigKey:   "seo_keyword",
			ConfigValue: req.SEOKeyword,
			UpdatedTime: time.Now(),
		}
		if result.ConfigValue == "" {
			config.CreatedTime = time.Now()
			err := h.service.CreateConfig(&config)
			fmt.Println(err)
		} else {
			h.service.UpdateConfig(&config)
		}
	}

	if req.SEODescription != "" {
		result, _ := h.service.GetConfigByKey("seo_description")
		config := core.CoreConfig{
			ConfigKey:   "seo_description",
			ConfigValue: req.SEODescription,
			UpdatedTime: time.Now(),
		}
		if result.ConfigValue == "" {
			config.CreatedTime = time.Now()
			h.service.CreateConfig(&config)
		} else {
			h.service.UpdateConfig(&config)
		}
	}

	if req.SEOTitle != "" {
		result, _ := h.service.GetConfigByKey("seo_title")
		config := core.CoreConfig{
			ConfigKey:   "seo_title",
			ConfigValue: req.SEOTitle,
			UpdatedTime: time.Now(),
		}
		if result.ConfigValue == "" {
			config.CreatedTime = time.Now()
			h.service.CreateConfig(&config)
		} else {
			h.service.UpdateConfig(&config)
		}
	}

	if req.Title != "" {
		result, _ := h.service.GetConfigByKey("title")
		config := core.CoreConfig{
			ConfigKey:   "title",
			ConfigValue: req.Title,
			UpdatedTime: time.Now(),
		}
		if result.ConfigValue == "" {
			config.CreatedTime = time.Now()
			h.service.CreateConfig(&config)
		} else {
			h.service.UpdateConfig(&config)
		}
	}
	Success(c, "配置保存成功")
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
	config.ID = common.MyID(id)

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

type SiteInfoResponse struct {
	Title          string `json:"title"`
	Logo           string `json:"logo"`
	SeoTitle       string `json:"seo_title"`
	SeoKeyword     string `json:"seo_keyword"`
	SeoDescription string `json:"seo_description"`
}

type MarketInfoResponse struct {
	Exchange       string `json:"exchange"`
	Freight        string `json:"freight"`
	Original       string `json:"original"`
	SEOTitle       string `json:"seo_title"`
	SEODescription string `json:"seo_description"`
	SEOKeyword     string `json:"seo_keyword"`
}

// GetAllConfigs 获取所有配置项
func (h *CoreConfigHandler) GetSiteInfo(c *gin.Context) {
	configs, err := h.service.GetAllConfigs()
	if err != nil {
		Error(c, 6004, "获取配置失败")
		return
	}
	fmt.Println(configs)
	// 初始化返回对象
	response := &SiteInfoResponse{}

	// 遍历配置项，提取需要的键值
	for i := range configs {
		switch configs[i].ConfigKey {
		case "title":
			response.Title = configs[i].ConfigValue
		case "logo":
			response.Logo = configs[i].ConfigValue
		case "seo_title":
			response.SeoTitle = configs[i].ConfigValue
		case "seo_keyword":
			response.SeoKeyword = configs[i].ConfigValue
		case "seo_description":
			response.SeoDescription = configs[i].ConfigValue
		}
	}
	fmt.Println(response)
	Success(c, response)
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

func (h *CoreConfigHandler) SaveMarketSetting(c *gin.Context) {
	var req MarketRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}

	if req.Freight != "" {
		result, _ := h.service.GetConfigByKey("Freight")
		config := core.CoreConfig{
			ConfigKey:   "Freight",
			ConfigValue: fmt.Sprintf("%v", (utils.ConvertToFloat64(req.Freight))),
			UpdatedTime: time.Now(),
		}
		if result.ConfigValue == "" {
			config.CreatedTime = time.Now()
			h.service.CreateConfig(&config)
		} else {
			h.service.UpdateConfig(&config)
		}
	}

	if req.Original != "" {
		result, _ := h.service.GetConfigByKey("Original")
		config := core.CoreConfig{
			ConfigKey:   "Original",
			ConfigValue: fmt.Sprintf("%v", (utils.ConvertToFloat64(req.Original))),
			UpdatedTime: time.Now(),
		}
		if result.ConfigValue == "" {
			config.CreatedTime = time.Now()
			err := h.service.CreateConfig(&config)
			fmt.Println(err)
		} else {
			h.service.UpdateConfig(&config)
		}
	}

	if req.Exchange != "" {
		result, _ := h.service.GetConfigByKey("Exchange")
		config := core.CoreConfig{
			ConfigKey:   "Exchange",
			ConfigValue: fmt.Sprintf("%v", (utils.ConvertToFloat64(req.Exchange))),
			UpdatedTime: time.Now(),
		}
		if result.ConfigValue == "" {
			config.CreatedTime = time.Now()
			h.service.CreateConfig(&config)
		} else {
			h.service.UpdateConfig(&config)
		}
	}
	Success(c, "配置保存成功")
}

func (h *CoreConfigHandler) GetMarketInfo(c *gin.Context) {
	configs, err := h.service.GetAllConfigs()
	if err != nil {
		Error(c, 6004, "获取配置失败")
		return
	}
	fmt.Println(configs)
	// 初始化返回对象
	response := &MarketInfoResponse{}

	// 遍历配置项，提取需要的键值
	for i := range configs {
		switch configs[i].ConfigKey {
		case "Freight":
			response.Freight = configs[i].ConfigValue
		case "Exchange":
			response.Exchange = configs[i].ConfigValue
		case "Original":
			response.Original = configs[i].ConfigValue
		case "seo_title":
			response.SEOTitle = configs[i].ConfigValue
		case "seo_description":
			response.SEODescription = configs[i].ConfigValue
		case "seo_keyword":
			response.SEOKeyword = configs[i].ConfigValue
		}
	}
	fmt.Println(response)
	Success(c, response)
}
