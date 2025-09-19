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

type SpUserAddressCreateRequest struct {
	ID            common.MyID `json:"id"`
	UserID        common.MyID `json:"user_id"`
	Title         string      `json:"title"`
	DefaultStatus int16       `json:"default_status"`
	FirstName     string      `json:"first_name"`
	LastName      string      `json:"last_name"`
	Email         string      `json:"email"`
	Phone         string      `json:"phone"`
	Province      string      `json:"province"`
	City          string      `json:"city"`
	Region        string      `json:"region"`
	DetailAddress string      `json:"detail_address"`
	Country       string      `json:"country"`
	PostalCode    string      `json:"postal_code"`
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
	var req SpUserAddressCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.InvalidParams(c)
	}
	userId := middleware.GetUserIDFromContext(c)

	address := sp.SpUserAddress{
		UserID:        common.MyID(userId),
		Title:         req.Title,
		DefaultStatus: req.DefaultStatus,
		FirstName:     req.FirstName,
		LastName:      req.LastName,
		Email:         req.Email,
		Phone:         req.Phone,
		Province:      req.Province,
		City:          req.City,
		Region:        req.Region,
		DetailAddress: req.DetailAddress,
		Country:       req.Country,
		PostalCode:    req.PostalCode,
	}

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
