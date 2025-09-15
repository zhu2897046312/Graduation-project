package handlers

import (
	"github.com/gin-gonic/gin"
	"server/service"
	"server/models/cms"
	"strconv"
)

type CmsScenicSpotHandler struct {
	service *service.CmsScenicSpotService
}

func NewCmsScenicSpotHandler(service *service.CmsScenicSpotService) *CmsScenicSpotHandler {
	return &CmsScenicSpotHandler{service: service}
}

// 创建景点
func (h *CmsScenicSpotHandler) CreateSpot(c *gin.Context) {
	var spot cms.CmsScenicSpot
	if err := c.ShouldBindJSON(&spot); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.CreateScenicSpot(&spot); err != nil {
		Error(c, 9001, err.Error())
		return
	}

	Success(c, spot)
}

// 更新景点
func (h *CmsScenicSpotHandler) UpdateSpot(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		InvalidParams(c)
		return
	}

	var spot cms.CmsScenicSpot
	if err := c.ShouldBindJSON(&spot); err != nil {
		InvalidParams(c)
		return
	}
	spot.ID = id

	if err := h.service.UpdateScenicSpot(&spot); err != nil {
		Error(c, 9002, err.Error())
		return
	}

	Success(c, spot)
}

// 获取景点详情
func (h *CmsScenicSpotHandler) GetSpot(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		InvalidParams(c)
		return
	}

	spot, err := h.service.GetScenicSpotByID(id)
	if err != nil {
		Error(c, 9003, "景点不存在")
		return
	}

	Success(c, spot)
}

// 根据地点获取景点
func (h *CmsScenicSpotHandler) GetByPlace(c *gin.Context) {
	placeID, err := strconv.ParseInt(c.Param("place_id"), 10, 64)
	if err != nil || placeID <= 0 {
		InvalidParams(c)
		return
	}

	spots, err := h.service.GetScenicSpotsByPlaceID(placeID)
	if err != nil {
		ServerError(c, err)
		return
	}

	Success(c, spots)
}

// 增加阅读量
func (h *CmsScenicSpotHandler) IncrementReadNum(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		InvalidParams(c)
		return
	}

	if err := h.service.IncrementReadNum(id); err != nil {
		Error(c, 9004, err.Error())
		return
	}

	Success(c, nil)
}

// 分页获取景点
func (h *CmsScenicSpotHandler) ListSpots(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	spots, total, err := h.service.ListScenicSpots(page, pageSize)
	if err != nil {
		ServerError(c, err)
		return
	}

	Success(c, gin.H{
		"list":  spots,
		"total": total,
	})
}