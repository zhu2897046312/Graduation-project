package handlers

import (
	"github.com/gin-gonic/gin"
	"server/models/sp"
	"server/service"
	"strconv"
)

type SpUserAddressHandler struct {
	service *service.SpUserAddressService
}

func NewSpUserAddressHandler(service *service.SpUserAddressService) *SpUserAddressHandler {
	return &SpUserAddressHandler{service: service}
}

// CreateAddress 创建地址
func (h *SpUserAddressHandler) CreateAddress(c *gin.Context) {
	var address sp.SpUserAddress
	if err := c.ShouldBindJSON(&address); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.CreateAddress(&address); err != nil {
		Error(c, 3301, err.Error())
		return
	}

	Success(c, address)
}

// UpdateAddress 更新地址
func (h *SpUserAddressHandler) UpdateAddress(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		InvalidParams(c)
		return
	}

	var address sp.SpUserAddress
	if err := c.ShouldBindJSON(&address); err != nil {
		InvalidParams(c)
		return
	}
	address.ID = uint(id)

	if err := h.service.UpdateAddress(&address); err != nil {
		Error(c, 3302, err.Error())
		return
	}

	Success(c, address)
}

// GetAddresses 获取用户地址列表
func (h *SpUserAddressHandler) GetAddresses(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil || userID == 0 {
		InvalidParams(c)
		return
	}

	addresses, err := h.service.GetAddressesByUserID(uint(userID))
	if err != nil {
		Error(c, 3303, err.Error())
		return
	}

	Success(c, addresses)
}

// SetDefaultAddress 设置默认地址
func (h *SpUserAddressHandler) SetDefaultAddress(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		InvalidParams(c)
		return
	}

	userID, err := strconv.ParseUint(c.Query("user_id"), 10, 64)
	if err != nil || userID == 0 {
		InvalidParams(c)
		return
	}

	if err := h.service.SetDefaultAddress(uint(id), uint(userID)); err != nil {
		Error(c, 3304, err.Error())
		return
	}

	Success(c, nil)
}