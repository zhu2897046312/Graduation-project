package client

import (
	"server/middleware"
	"server/models/common"
	"server/models/sp"
	"server/service"
	"server/utils"

	"github.com/gin-gonic/gin"
)

type SpUserAddressHandler struct {
	service *service.SpUserAddressService
}

func NewSpUserAddressHandler(service *service.SpUserAddressService) *SpUserAddressHandler {
	return &SpUserAddressHandler{service: service}
}

type ListAddressRequest struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

func (h *SpUserAddressHandler) ListAddress(c *gin.Context) {
	var req ListAddressRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.InvalidParams(c)
	}
	userId := middleware.GetUserIDFromContext(c)

	addresses, total, err := h.service.ListAddress(&sp.SpUserAddressListParam{
		UserID:   common.MyID(userId),
		Page:     req.Page,
		PageSize: req.PageSize,
	})

	if err != nil {
		utils.ServerError(c, err)
		return
	}

	utils.Success(c, gin.H{
		"list":  addresses,
		"total": total,
	})
}

func (h *SpUserAddressHandler) CreateAddress(c *gin.Context) {
	var address sp.SpUserAddress
	if err := c.ShouldBindJSON(&address); err != nil {
		utils.InvalidParams(c)
	}
	userId := middleware.GetUserIDFromContext(c)
	address.UserID = common.MyID(userId)

	if err := h.service.Create(&address); err != nil {
		utils.ServerError(c, err)
		return
	}

	utils.Success(c, gin.H{
		"id": address.ID,
	})
}

func (h *SpUserAddressHandler) UpdateAddress(c *gin.Context) {
	var address sp.SpUserAddress
	if err := c.ShouldBindJSON(&address); err != nil {
		utils.InvalidParams(c)
	}

	if err := h.service.Update(&address); err != nil {
		utils.ServerError(c, err)
		return
	}

	utils.Success(c, gin.H{
		"id": address.ID,
	})
}

func (h *SpUserAddressHandler) DeleteAddress(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		utils.InvalidParams(c)
		return
	}

	myID := common.MyID(utils.ConvertToUint(id))

	if err := h.service.Delete(myID); err != nil {
		utils.ServerError(c, err)
		return
	}

	utils.Success(c, gin.H{
		"id": myID,
	})
}

func (h *SpUserAddressHandler) GetAddress(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		utils.InvalidParams(c)
		return
	}

	myID := common.MyID(utils.ConvertToUint(id))
	address, err := h.service.GetAddressByID(myID)

	if err != nil {
		utils.ServerError(c, err)
		return
	}

	utils.Success(c, address)
}
