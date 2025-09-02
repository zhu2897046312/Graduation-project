package handlers

import (
	"server/models/sp"
	"server/service"
	"server/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type SpOrderListVo struct {
	ID              uint             `json:"id"`
	Code            string           `json:"code"`
	UserID          uint             `json:"user_id"`
	Nickname        string           `json:"nickname"`
	Email           string           `json:"email"`
	TotalAmount     float64          `json:"total_amount"`
	PayAmount       float64          `json:"pay_amount"`
	State           uint8            `json:"state"`
	PaymentTime     *time.Time       `json:"payment_time"`
	DeliveryTime    *time.Time       `json:"delivery_time"`
	ReceiveTime     *time.Time       `json:"receive_time"`
	DeliveryCompany string           `json:"delivery_company"`
	DeliverySn      string           `json:"delivery_sn"`
	Remark          string           `json:"remark"`
	Freight         float64          `json:"freight"`
	CreatedTime     time.Time        `json:"created_time"`
	Items           []sp.SpOrderItem `json:"items"`
}

type OrderListResponse struct {
	List  []SpOrderListVo `json:"list"`
	Total int64           `json:"total"`
}

type ListOrdersRequest struct {
	NikeName string `json:"nickname"`
	Email    string `json:"email"`
	Code     string `json:"code"`
	State    uint8  `json:"state"`
	Page     int    `json:"page_no"`
	PageSize int    `json:"page_size"`
}
type SpOrderHandler struct {
	service *service.SpOrderService
	orderItemService *service.SpOrderItemService
	orederReceiveService *service.SpOrderReceiveAddressService
}

func NewSpOrderHandler(service *service.SpOrderService,orderItemService *service.SpOrderItemService,orederReceiveService *service.SpOrderReceiveAddressService) *SpOrderHandler {
	return &SpOrderHandler{
		service: service,
		orderItemService: orderItemService,
		orederReceiveService: orederReceiveService,
	}
}

// 创建订单
func (h *SpOrderHandler) CreateOrder(c *gin.Context) {
	var order sp.SpOrder
	if err := c.ShouldBindJSON(&order); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.CreateOrder(&order); err != nil {
		Error(c, 27001, err.Error())
		return
	}

	Success(c, order)
}

// 更新订单
func (h *SpOrderHandler) UpdateOrder(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil || id == 0 {
		InvalidParams(c)
		return
	}

	var order sp.SpOrder
	if err := c.ShouldBindJSON(&order); err != nil {
		InvalidParams(c)
		return
	}
	order.ID = uint(id)

	if err := h.service.UpdateOrder(&order); err != nil {
		Error(c, 27002, err.Error())
		return
	}

	Success(c, order)
}

// 获取订单详情
func (h *SpOrderHandler) GetOrder(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		InvalidParams(c)
		return
	}
	uid := utils.ConvertToUint(id)
	order, err := h.service.GetOrderByID(uid)
	if err != nil {
		Error(c, 27003, "订单不存在")
		return
	}
	items, err := h.orderItemService.GetItemsByOrderID(order.ID)
	if err != nil {
		Error(c, 27003, "订单不存在")
		return
	}
	receiveAddress, err := h.orederReceiveService.GetAddressByOrderID(order.ID)
	if err != nil {
		Error(c, 27003, "订单不存在")
		return
	}
	Success(c, gin.H{
		"order": order,
		"items": items,
		"receive_address": receiveAddress,
	})
}


// 根据订单号获取订单
func (h *SpOrderHandler) GetOrderByCode(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		InvalidParams(c)
		return
	}

	order, err := h.service.GetOrderByCode(code)
	if err != nil {
		Error(c, 27004, "订单不存在")
		return
	}

	Success(c, order)
}

// 根据用户ID获取订单列表
func (h *SpOrderHandler) GetOrdersByUser(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 32)
	if err != nil || userID == 0 {
		InvalidParams(c)
		return
	}

	orders, err := h.service.GetOrdersByUserID(uint(userID))
	if err != nil {
		Error(c, 27005, "获取订单列表失败")
		return
	}

	Success(c, orders)
}

// 根据状态获取订单列表
func (h *SpOrderHandler) GetOrdersByState(c *gin.Context) {
	state, err := strconv.ParseUint(c.Query("state"), 10, 8)
	if err != nil {
		InvalidParams(c)
		return
	}

	orders, err := h.service.GetOrdersByState(uint8(state))
	if err != nil {
		Error(c, 27006, "获取订单列表失败")
		return
	}

	Success(c, orders)
}

// 更新订单状态
func (h *SpOrderHandler) UpdateOrderState(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil || id == 0 {
		InvalidParams(c)
		return
	}

	var req struct {
		State uint8 `json:"state"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.UpdateOrderState(uint(id), req.State); err != nil {
		Error(c, 27007, err.Error())
		return
	}

	Success(c, nil)
}

// 更新物流信息
func (h *SpOrderHandler) UpdateDeliveryInfo(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil || id == 0 {
		InvalidParams(c)
		return
	}

	var req struct {
		Company string `json:"company"`
		SN      string `json:"sn"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.UpdateDeliveryInfo(uint(id), req.Company, req.SN); err != nil {
		Error(c, 27008, err.Error())
		return
	}

	Success(c, nil)
}

func (h *SpOrderHandler) ListOrders(c *gin.Context) {
	var req ListOrdersRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}
	val := sp.ListOrdersQueryParam{
		NikeName: req.NikeName,
		Email:    req.Email,
		Code:     req.Code,
		State:    req.State,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	orders, total, err := h.service.List(val)
	if err != nil {
		Error(c, 27009, err.Error())
		return
	}
	var orderList []SpOrderListVo
    for _, order := range orders {
        orderVo := SpOrderListVo{
            ID:              order.ID,
            Code:            order.Code,
            UserID:          order.UserID,
            Nickname:        order.Nickname,
            Email:           order.Email,
            TotalAmount:     order.TotalAmount,
            PayAmount:       order.PayAmount,
            State:           order.State,
            PaymentTime:     order.PaymentTime,
            DeliveryTime:    order.DeliveryTime,
            ReceiveTime:     order.ReceiveTime,
            DeliveryCompany: order.DeliveryCompany,
            DeliverySn:      order.DeliverySn,
            Remark:          order.Remark,
            Freight:         order.Freight,
            CreatedTime:     order.CreatedTime,
        }
		items, err := h.orderItemService.GetItemsByOrderID(order.ID)
		if err != nil {	
			Error(c, 27010, err.Error())
			return
		}
		orderVo.Items = items

        orderList = append(orderList, orderVo)
    }
	
	Success(c, gin.H{	
		"total": total,
		"list": orderList,
	})
}
