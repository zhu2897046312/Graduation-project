package handlers

import (
	"github.com/gin-gonic/gin"
	"server/service"
	"server/models/cms"
	"strconv"
)

type CmsPlaceHandler struct {
	service *service.CmsAssociatedPlaceService
}

func NewCmsPlaceHandler(service *service.CmsAssociatedPlaceService) *CmsPlaceHandler {
	return &CmsPlaceHandler{service: service}
}

// 创建地点
func (h *CmsPlaceHandler) CreatePlace(c *gin.Context) {
	var place cms.CmsAssociatedPlace
	if err := c.ShouldBindJSON(&place); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.CreatePlace(&place); err != nil {
		Error(c, 1002, err.Error())
		return
	}

	Success(c, place)
}

// 更新地点
func (h *CmsPlaceHandler) UpdatePlace(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		InvalidParams(c)
		return
	}

	var place cms.CmsAssociatedPlace
	if err := c.ShouldBindJSON(&place); err != nil {
		InvalidParams(c)
		return
	}
	place.ID = id

	if err := h.service.UpdatePlace(&place); err != nil {
		Error(c, 1003, err.Error())
		return
	}

	Success(c, place)
}

// 获取地点详情
func (h *CmsPlaceHandler) GetPlace(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		InvalidParams(c)
		return
	}

	place, err := h.service.GetPlaceByID(id)
	if err != nil {
		Error(c, 1004, "地点不存在")
		return
	}

	Success(c, place)
}

// 分页获取地点列表
func (h *CmsPlaceHandler) ListPlaces(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	places, total, err := h.service.ListPlaces(page, pageSize)
	if err != nil {
		ServerError(c, err)
		return
	}

	Success(c, gin.H{
		"list":  places,
		"total": total,
	})
}

// 根据首字母搜索地点
func (h *CmsPlaceHandler) SearchPlaces(c *gin.Context) {
	initial := c.Query("initial")
	if initial == "" {
		InvalidParams(c)
		return
	}

	places, err := h.service.SearchPlacesByInitial(initial)
	if err != nil {
		ServerError(c, err)
		return
	}

	Success(c, places)
}

// 更新地点状态
func (h *CmsPlaceHandler) UpdatePlaceState(c *gin.Context) {
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

	if err := h.service.UpdatePlaceState(id, req.State); err != nil {
		Error(c, 1005, err.Error())
		return
	}

	Success(c, nil)
}

// 删除地点
func (h *CmsPlaceHandler) DeletePlace(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		InvalidParams(c)
		return
	}

	if err := h.service.DeletePlace(id); err != nil {
		Error(c, 1006, err.Error())
		return
	}

	Success(c, nil)
}