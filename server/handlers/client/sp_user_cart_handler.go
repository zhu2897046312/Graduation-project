package client

import (
	"errors"
	"server/middleware"
	"server/models/sp"
	"server/models/common"
	"server/service"
	"server/utils"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ClientSpUserCartHandler struct {
	service        *service.SpUserCartService
	productService *service.SpProductService
	skuService     *service.SpSkuService
}

func NewClientSpUserCartHandler(
	service *service.SpUserCartService,
	productService *service.SpProductService,
	skuService     *service.SpSkuService,
	) *ClientSpUserCartHandler {
	return &ClientSpUserCartHandler{
		service: service,
		productService: productService,
		skuService:     skuService,
	}
}

func (h *ClientSpUserCartHandler) List(c *gin.Context) {
	fingerprint := middleware.GetDeviceFingerprintFromContext(c)
	// userID := middleware.GetUserIDFromContext(c)

	if fingerprint == "" {
		utils.InvalidParams(c)
		return
	}

	cartItems, total, err := h.service.List(0, fingerprint, 0, 0)
	if err != nil {
		utils.Error(c, 3406, err.Error())
		return
	}

	utils.Success(c, gin.H{
		"list":  cartItems,
		"total": total,
	})
}

func (h *ClientSpUserCartHandler) CarAction(c *gin.Context) {
	type SpUserCartActRequest struct {
		ProductID common.MyID `json:"product_id"` // 商品ID
		SkuID     common.MyID `json:"sku_id"`                        // SKU ID
		Quantity  uint `json:"quantity"`  // 数量
		Add       bool `json:"add"`                          // 操作类型：true=添加，false=减少
	}

	var req SpUserCartActRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.InvalidParams(c)
		return
	}

	userID := middleware.GetUserIDFromContext(c)
	fingerprint := middleware.GetDeviceFingerprintFromContext(c)
	if userID == 0 && fingerprint == "" {
		utils.Error(c, 3407, "用户未登录且缺少设备指纹")
		return
	}
	
	if userID !=0 {
		h.service.MergeGuestCart(common.MyID(userID), fingerprint)
	}

	currentCarts, err := h.service.GetCartItemByProduct(common.MyID(userID), fingerprint, req.ProductID, req.SkuID)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		utils.Error(c, 3408, err.Error())
		return
	}

	if req.Add {
		if currentCarts.ID == 0 {
			newCart := &sp.SpUserCart{
				UserID:      common.MyID(userID),
				Fingerprint: fingerprint,
				ProductID:   req.ProductID,
				SkuID:       req.SkuID,
				Quantity:    req.Quantity,
				CreatedTime: time.Now(),
				UpdatedTime: time.Now(),
			}
			// 同步商品信息
			if err_1 := h.syncCartProductInfo(newCart); err_1 != nil {
				utils.Error(c, 3409, err_1.Error())
				return
			}

			h.service.AddToCart(newCart)
			utils.Success(c, newCart)
		}else{
			currentCarts.Quantity += req.Quantity
			// 重新计算金额
			if err := h.syncCartProductInfo(currentCarts); err != nil {
				return 
			}
			if err_2 := h.service.UpdateCartItem(currentCarts); err_2 != nil {
				utils.Error(c, 3410, err_2.Error())
				return
			}
			utils.Success(c, currentCarts)
		}
	} else {
		if currentCarts == nil {
			utils.Error(c, 3411, "购物车中无此商品")
			return
		}
		if req.Quantity >= currentCarts.Quantity {
			h.service.DeleteCartItem(currentCarts.ID)
			utils.Success(c, nil)
		} else {
			currentCarts.Quantity -= req.Quantity
			// 重新计算金额
			if err := h.syncCartProductInfo(currentCarts); err != nil {
				return 
			}
			if err_3 := h.service.UpdateCartItem(currentCarts); err_3 != nil {
				utils.Error(c, 3412, err_3.Error())
				return
			}
			utils.Success(c, currentCarts) 
		}
	}
}

func (h *ClientSpUserCartHandler) syncCartProductInfo(cart *sp.SpUserCart) error {

	product, err := h.productService.GetProductByID(cart.ProductID)
	if err != nil {
		return err
	}
	if product == nil {
		return errors.New("商品不存在")
	}

	// 检查商品状态
	if product.State != 1 { // 假设1表示上架状态
		return errors.New("商品已下架")
	}

	if product.OpenSku == 1 {
		sku, err_1 := h.skuService.GetSkuByID(cart.SkuID)
		if err_1 != nil {
			return err_1
		}
		if sku == nil {
			return errors.New("SKU不存在")
		}
		if sku.ProductID != cart.ProductID {
			return errors.New("SKU不属于该商品")
		}

		cart.Price = sku.Price
		cart.OriginalPrice = sku.OriginalPrice
		cart.SkuCode = sku.SkuCode
		cart.SkuTitle = sku.Title
	} else {
		cart.Price = product.Price
		cart.OriginalPrice = product.OriginalPrice
		cart.SkuID = 0 // 无SKU时设置为0
	}

	cart.Thumb = product.Picture
	cart.Title = product.Title

	// 计算金额
	cart.TotalAmount = cart.OriginalPrice * float64(cart.Quantity)
	cart.PayAmount = cart.Price * float64(cart.Quantity)

	return nil
}
