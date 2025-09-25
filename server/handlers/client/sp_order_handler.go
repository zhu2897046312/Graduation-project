package client

import (
	"encoding/json"
	"fmt"
	"server/middleware"
	"server/models/common"
	"server/models/sp"
	"server/service"
	"server/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type SpOrderListVo struct {
	ID               common.MyID      `json:"id"`
	Code             string           `json:"code"`
	UserID           common.MyID      `json:"user_id"`
	Nickname         string           `json:"nickname"`
	Email            string           `json:"email"`
	TotalAmount      float64          `json:"total_amount"`
	PayAmount        float64          `json:"pay_amount"`
	State            common.MyState   `json:"state"`
	PaymentTime      *time.Time       `json:"payment_time"`
	DeliveryTime     *time.Time       `json:"delivery_time"`
	ReceiveTime      *time.Time       `json:"receive_time"`
	DeliveryCompany  string           `json:"delivery_company"`
	DeliverySn       string           `json:"delivery_sn"`
	Remark           string           `json:"remark"`
	Freight          float64          `json:"freight"`
	CreatedTime      time.Time        `json:"created_time"`
	Items            []sp.SpOrderItem `json:"items"`
	VisitorQueryCode string           `json:"visitor_query_code"`
}

type OrderListResponse struct {
	List  []SpOrderListVo `json:"list"`
	Total int64           `json:"total"`
}

type ListOrdersRequest struct {
	NikeName string         `json:"nickname"`
	Email    string         `json:"email"`
	Code     string         `json:"code"`
	State    common.MyState `json:"state"`
	Page     int            `json:"page_no"`
	PageSize int            `json:"page_size"`
}

type ProductItemRequest struct {
	ProductID common.MyID `json:"product_id"`
	SkuID     common.MyID `json:"sku_id"`
	Quantity  uint        `json:"quantity"`
}

type OrderCreateRequest struct {
	ProductItem []ProductItemRequest `json:"product_items"`
	PayType     interface{}          `json:"pay_type"`
	FirstName   string               `json:"first_name"`
	LastName    string               `json:"last_name"`
	Email       string               `json:"email"`
	Phone       string               `json:"phone"`
	PostCode    string               `json:"postal_code"`
	Country     string               `json:"country"`
	Province    string               `json:"province"`
	City        string               `json:"city"`
	Region      string               `json:"region"`
	Detail      string               `json:"detail_address"`
}

// SpOrderCreateResp 创建订单响应
type SpOrderCreateResp struct {
	OrderID          uint   `json:"order_id"`
	OrderCode        string `json:"order_code"`
	VisitorQueryCode string `json:"visitor_query_code"`
	TotalAmount      string `json:"total_amount"`
	PayAmount        string `json:"pay_amount"`
	Freight          string `json:"freight"`
}

// SpOrderFrontInfoVo 前端订单信息视图对象
type SpOrderFrontInfoVo struct {
	Order   SpOrderFrontQueryVo      `json:"order"`
	Address sp.SpOrderReceiveAddress `json:"address"`
	Items   []SpOrderItemFrontVo     `json:"items"`
}

// SpOrderFrontQueryVo 前端订单查询视图对象
type SpOrderFrontQueryVo struct {
	ID               common.MyID          `json:"id" description:"主键"`
	Code             string               `json:"code" description:"订单号"`
	UserID           common.MyID          `json:"user_id" description:"用户id"`
	VisitorQueryCode string               `json:"visitor_query_code" description:"访客查询码"`
	Nickname         string               `json:"nickname" description:"昵称"`
	Email            string               `json:"email" description:"邮箱"`
	TotalAmount      float64              `json:"total_amount" description:"订单总金额"`
	PayAmount        float64              `json:"pay_amount" description:"实际支付总金额"`
	PayType          uint8                `json:"pay_type" description:"支付方式:1=货到付款"`
	State            common.MyState       `json:"state" description:"订单状态:1=待付款;2=待发货;3=已发货;4=已完成;5=已关闭;6=无效订单"`
	PaymentTime      time.Time            `json:"payment_time" description:"支付时间"`
	DeliveryTime     time.Time            `json:"delivery_time" description:"发货时间"`
	ReceiveTime      time.Time            `json:"receive_time" description:"确认收货时间"`
	DeliveryCompany  string               `json:"delivery_company" description:"物流公司(配送方式)"`
	DeliverySn       string               `json:"delivery_sn" description:"物流单号"`
	Items            []SpOrderItemFrontVo `json:"items" description:"商品信息"`
	Freight          float64              `json:"freight" description:"运费"`
}

// SpOrderItemFrontVo 前端订单项视图对象
type SpOrderItemFrontVo struct {
	ID            common.MyID `json:"id" description:"主键"`
	Title         string      `json:"title" description:"商品标题"`
	SkuTitle      string      `json:"sku_title" description:"商品SKU内容"`
	Thumb         string      `json:"thumb" description:"商品图片"`
	OrderID       common.MyID `json:"order_id" description:"订单id"`
	ProductID     common.MyID `json:"product_id" description:"商品id"`
	SkuID         common.MyID `json:"sku_id" description:"商品SKUid"`
	TotalAmount   float64     `json:"total_amount" description:"总金额"`
	PayAmount     float64     `json:"pay_amount" description:"实际支付金额"`
	Quantity      uint        `json:"quantity" description:"购买数量"`
	Price         float64     `json:"price" description:"单价"`
	OriginalPrice float64     `json:"original_price" description:"原价单价"`
}

type ClientSpOrderHandler struct {
	service              *service.SpOrderService
	orderItemService     *service.SpOrderItemService
	orederReceiveService *service.SpOrderReceiveAddressService
	orderRefundService   *service.SpOrderRefundService
	addressService       *service.SpOrderReceiveAddressService
	productService       *service.SpProductService
	cartService          *service.SpUserCartService
}

func NewClientSpOrderHandler(
	service *service.SpOrderService,
	orderItemService *service.SpOrderItemService,
	orederReceiveService *service.SpOrderReceiveAddressService,
	orderRefundService *service.SpOrderRefundService,
	addressService *service.SpOrderReceiveAddressService,
	productService *service.SpProductService,
	cartService *service.SpUserCartService,
) *ClientSpOrderHandler {
	return &ClientSpOrderHandler{
		service:              service,
		orderItemService:     orderItemService,
		orederReceiveService: orederReceiveService,
		orderRefundService:   orderRefundService,
		addressService:       addressService,
		productService:       productService,
		cartService:          cartService,
	}
}

// 创建订单
func (h *ClientSpOrderHandler) CreateOrder(c *gin.Context) {
	var req OrderCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.InvalidParams(c)
		return
	}
	// 获取用户ID和设备指纹
	userID := middleware.GetUserIDFromContext(c)
	fingerprint := middleware.GetDeviceFingerprintFromContext(c)

	// 创建订单项并计算金额
	orderItems, totalAmount, payAmount, err := h.createOrderItems(req.ProductItem)
	if err != nil {
		utils.Error(c, 16001, err.Error())
		return
	}

	// 获取运费
	freight, err := h.getFreight()
	if err != nil {
		utils.Error(c, 16002, "获取运费失败")
		return
	}

	// 计算总金额（含运费）
	totalAmountWithFreight := totalAmount + freight
	payAmountWithFreight := payAmount + freight

	// 生成订单号
	orderCode := utils.GenerateOrderSn()

	// 生成访客查询码
	visitorQueryCode := utils.GenerateUUID()

	// 创建订单主记录
	order := &sp.SpOrder{
		Code:             orderCode,
		UserID:           common.MyID(userID),
		Nickname:         req.FirstName + " " + req.LastName,
		Email:            req.Email,
		TotalAmount:      totalAmountWithFreight,
		PayAmount:        payAmountWithFreight,
		PayType:          common.MyType(utils.ConvertToUint(req.PayType)),
		State:            2, // 待支付
		Freight:          freight,
		VisitorQueryCode: visitorQueryCode,
	}

	// 保存订单
	if err := h.service.CreateOrder(order); err != nil {
		utils.Error(c, 16003, "创建订单失败: "+err.Error())
		return
	}

	// 保存订单项
	if err := h.saveOrderItems(order.ID, orderItems); err != nil {
		// 如果保存订单项失败，删除订单
		h.service.DeleteOrder(order.ID)
		utils.Error(c, 16004, "保存订单项失败")
		return
	}

	// 保存收货地址
	if err := h.saveOrderAddress(order.ID, req); err != nil {
		// 如果保存地址失败，删除订单和订单项
		h.service.DeleteOrder(order.ID)
		utils.Error(c, 16005, "保存收货地址失败")
		return
	}

	// 清空购物车
	if err := h.clearCart(common.MyID(userID), fingerprint); err != nil {
		utils.Error(c, 16006, "清空购物车失败")
	}

	// 返回响应
	utils.Success(c, visitorQueryCode)
}

// clearCart 清空购物车
func (h *ClientSpOrderHandler) clearCart(userID common.MyID, fingerprint string) error {
	if userID == 0 {
		// 游客 - 根据设备指纹清空
		return h.cartService.ClearCartByFingerprint(fingerprint)
	} else {
		// 用户 - 根据用户ID清空
		return h.cartService.ClearCartByUserID(userID)
	}
}

// getFreight 获取运费
func (h *ClientSpOrderHandler) getFreight() (float64, error) {
	// 这里可以从配置或数据库中获取运费
	// 暂时返回固定值
	return 20, nil
}

// createOrderItems 创建订单项并计算金额
func (h *ClientSpOrderHandler) createOrderItems(productItems []ProductItemRequest) ([]*sp.SpOrderItem, float64, float64, error) {
	var orderItems []*sp.SpOrderItem
	var totalAmount float64
	var payAmount float64

	for _, item := range productItems {
		// 获取商品信息
		product, err := h.productService.GetProductByID(item.ProductID)
		if err != nil {
			return nil, 0, 0, err
		}

		// 计算商品金额
		itemTotalAmount := product.Price * float64(item.Quantity)
		itemPayAmount := product.Price * float64(item.Quantity) // 这里可以根据折扣策略调整
		pictureGallery, _ := json.Marshal(product.PictureGallery)
		thumb := string(pictureGallery)
		orderItem := &sp.SpOrderItem{
			ProductID:   common.MyID(item.ProductID),
			SkuID:       common.MyID(item.SkuID),
			Quantity:    item.Quantity,
			Price:       product.Price,
			TotalAmount: itemTotalAmount,
			PayAmount:   itemPayAmount,
			Title:       product.Title,
			Thumb:       thumb,
		}

		orderItems = append(orderItems, orderItem)
		totalAmount += itemTotalAmount
		payAmount += itemPayAmount
	}

	return orderItems, totalAmount, payAmount, nil
}

// saveOrderItems 保存订单项
func (h *ClientSpOrderHandler) saveOrderItems(orderID common.MyID, items []*sp.SpOrderItem) error {
	for _, item := range items {
		item.OrderID = orderID
		if err := h.orderItemService.CreateOrderItem(item); err != nil {
			return err
		}
	}
	return nil
}

// saveOrderAddress 保存收货地址
func (h *ClientSpOrderHandler) saveOrderAddress(orderID common.MyID, req OrderCreateRequest) error {
	address := &sp.SpOrderReceiveAddress{
		OrderID:       orderID,
		FirstName:     req.FirstName,
		LastName:      req.LastName,
		Email:         req.Email,
		Phone:         req.Phone,
		Country:       req.Country,
		Province:      req.Province,
		City:          req.City,
		PostalCode:    req.PostCode,
		DetailAddress: req.Detail,
	}

	return h.addressService.CreateAddress(address)
}

func (h *ClientSpOrderHandler) ListOrders(c *gin.Context) {
	var req ListOrdersRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.InvalidParams(c)
		return
	}
	userID := middleware.GetUserIDFromContext(c)
	val := sp.ListOrdersQueryParam{
		UserID: common.MyID(userID),
	}
	orders, total, err := h.service.List(val)
	if err != nil {
		utils.Error(c, 27009, err.Error())
		return
	}
	var orderList []SpOrderListVo
	for _, order := range orders {
		orderVo := SpOrderListVo{
			ID:               order.ID,
			Code:             order.Code,
			Nickname:         order.Nickname,
			Email:            order.Email,
			TotalAmount:      order.TotalAmount,
			PayAmount:        order.PayAmount,
			State:            order.State,
			PaymentTime:      order.PaymentTime,
			DeliveryTime:     order.DeliveryTime,
			ReceiveTime:      order.ReceiveTime,
			DeliveryCompany:  order.DeliveryCompany,
			DeliverySn:       order.DeliverySn,
			Remark:           order.Remark,
			Freight:          order.Freight,
			CreatedTime:      order.CreatedTime,
			VisitorQueryCode: order.VisitorQueryCode,
		}
		items, err := h.orderItemService.GetItemsByOrderID(order.ID)
		if err != nil {
			utils.Error(c, 27010, err.Error())
			return
		}
		orderVo.Items = items

		orderList = append(orderList, orderVo)
	}

	utils.Success(c, gin.H{
		"total": total,
		"list":  orderList,
	})
}

// GetOrderByQueryCode 根据访客查询码获取订单详情
func (h *ClientSpOrderHandler) GetOrderByQueryCode(c *gin.Context) {
	queryCode := c.Query("queryCode")
	if queryCode == "" {
		utils.Error(c, 27016, "查询码不能为空")
		return
	}

	// 1. 使用查询码查找订单
	order, err := h.service.GetByVisitorQueryCode(queryCode)
	if err != nil {
		utils.Error(c, 27017, "订单不存在")
		return
	}

	// 2. 获取订单的收货地址信息
	address, err := h.orederReceiveService.GetAddressByOrderID(order.ID)
	if err != nil {
		utils.Error(c, 27018, "获取收货地址失败")
		return
	}

	// 3. 获取订单的商品项信息
	items, err := h.orderItemService.GetItemsByOrderID(order.ID)
	if err != nil {
		utils.Error(c, 27019, "获取订单商品项失败")
		return
	}

	// 4. 转换为前端需要的VO对象
	frontItems := make([]SpOrderItemFrontVo, 0)
	for _, item := range items {
		frontItems = append(frontItems, SpOrderItemFrontVo{
			ID:            item.ID,
			Title:         item.Title,
			SkuTitle:      item.SkuTitle,
			Thumb:         item.Thumb,
			OrderID:       item.OrderID,
			ProductID:     item.ProductID,
			SkuID:         item.SkuID,
			TotalAmount:   item.TotalAmount,
			PayAmount:     item.PayAmount,
			Quantity:      item.Quantity,
			Price:         item.Price,
			OriginalPrice: item.OriginalPrice,
		})
	}

	// 处理时间字段，将指针转换为值
	var paymentTime, deliveryTime, receiveTime time.Time
	if order.PaymentTime != nil {
		paymentTime = *order.PaymentTime
	}
	if order.DeliveryTime != nil {
		deliveryTime = *order.DeliveryTime
	}
	if order.ReceiveTime != nil {
		receiveTime = *order.ReceiveTime
	}

	// 5. 组装响应数据
	orderVo := SpOrderFrontQueryVo{
		ID:               order.ID,
		Code:             order.Code,
		UserID:           order.UserID,
		VisitorQueryCode: order.VisitorQueryCode,
		Nickname:         order.Nickname,
		Email:            order.Email,
		TotalAmount:      order.TotalAmount,
		PayAmount:        order.PayAmount,
		PayType:          uint8(order.PayType),
		State:            order.State,
		PaymentTime:      paymentTime,
		DeliveryTime:     deliveryTime,
		ReceiveTime:      receiveTime,
		DeliveryCompany:  order.DeliveryCompany,
		DeliverySn:       order.DeliverySn,
		Items:            frontItems,
		Freight:          order.Freight,
	}

	response := SpOrderFrontInfoVo{
		Order:   orderVo,
		Address: *address,
		Items:   frontItems,
	}
	fmt.Println(response)
	// 6. 返回响应
	utils.Success(c, response)
}
