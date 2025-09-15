package handlers

import (
	"github.com/gin-gonic/gin"
	"server/models/sp"
	"server/models/common"
	"server/service"
	"strconv"
)

type SpOrderReceiveAddressHandler struct {
	service *service.SpOrderReceiveAddressService
}

func NewSpOrderReceiveAddressHandler(service *service.SpOrderReceiveAddressService) *SpOrderReceiveAddressHandler {
	return &SpOrderReceiveAddressHandler{service: service}
}

// 创建收货地址
func (h *SpOrderReceiveAddressHandler) CreateAddress(c *gin.Context) {
	var address sp.SpOrderReceiveAddress
	if err := c.ShouldBindJSON(&address); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.CreateAddress(&address); err != nil {
		Error(c, 25001, err.Error())
		return
	}

	Success(c, address)
}

// 更新收货地址
func (h *SpOrderReceiveAddressHandler) UpdateAddress(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil || id == 0 {
		InvalidParams(c)
		return
	}

	var address sp.SpOrderReceiveAddress
	if err := c.ShouldBindJSON(&address); err != nil {
		InvalidParams(c)
		return
	}
	address.ID = common.MyID(id)

	if err := h.service.UpdateAddress(&address); err != nil {
		Error(c, 25002, err.Error())
		return
	}

	Success(c, address)
}

// 根据订单ID获取收货地址
func (h *SpOrderReceiveAddressHandler) GetAddressByOrder(c *gin.Context) {
	orderID, err := strconv.ParseUint(c.Param("order_id"), 10, 32)
	if err != nil || orderID == 0 {
		InvalidParams(c)
		return
	}

	address, err := h.service.GetAddressByOrderID(common.MyID(orderID))
	if err != nil {
		Error(c, 25003, "收货地址不存在")
		return
	}

	Success(c, address)
}

// 根据邮箱获取收货地址
func (h *SpOrderReceiveAddressHandler) GetAddressesByEmail(c *gin.Context) {
	email := c.Param("email")
	if email == "" {
		InvalidParams(c)
		return
	}

	addresses, err := h.service.GetAddressesByEmail(email)
	if err != nil {
		Error(c, 25004, "获取收货地址失败")
		return
	}

	Success(c, addresses)
}