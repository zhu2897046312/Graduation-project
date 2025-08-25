package handlers

import (
	"server/models/sp"
	"server/models/shop"
	"server/service"
	"server/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ListProductsRequest struct {
	CategoryID interface{} `json:"category_id"` // 使用interface{}接收任何类型
	State      int         `json:"state"`
	Title      string      `json:"title"`
	Page       int         `json:"page_no"`
	PageSize   int         `json:"page_size"`
}

type SpProductHandler struct {
	service         *service.SpProductService
	categoryService *service.SpCategoryService
	contentService  *service.SpProductContentService
	propertyService *service.SpProductPropertyService
	skuService      *service.SpSkuService 
	skuIndexService *service.SpSkuIndexService
	tagIndexService *service.ShopTagIndexService
	tagService      *service.ShopTagService
}

func NewSpProductHandler(
	service *service.SpProductService, 
	categoryService *service.SpCategoryService,
	contentService *service.SpProductContentService,
	propertyService *service.SpProductPropertyService,
	skuService *service.SpSkuService,
	skuIndexService *service.SpSkuIndexService,
	tagIndexService *service.ShopTagIndexService,
	tagService      *service.ShopTagService,
	) *SpProductHandler {
	return &SpProductHandler{
		service:         service,
		categoryService: categoryService,
		contentService:  contentService,
		propertyService: propertyService,
		skuService:      skuService,
		skuIndexService: skuIndexService,
		tagIndexService: tagIndexService,
		tagService:      tagService,
	}
}

// CreateProduct 创建商品
func (h *SpProductHandler) CreateProduct(c *gin.Context) {
	var product sp.SpProduct
	if err := c.ShouldBindJSON(&product); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.CreateProduct(&product); err != nil {
		Error(c, 3101, err.Error())
		return
	}

	Success(c, product)
}

// UpdateProduct 更新商品
func (h *SpProductHandler) UpdateProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		InvalidParams(c)
		return
	}

	var product sp.SpProduct
	if err := c.ShouldBindJSON(&product); err != nil {
		InvalidParams(c)
		return
	}
	product.ID = uint(id)

	if err := h.service.UpdateProduct(&product); err != nil {
		Error(c, 3102, err.Error())
		return
	}

	Success(c, product)
}

// GetProduct 获取商品详情
// GetProduct 获取商品详情
func (h *SpProductHandler) GetProduct(c *gin.Context) {
    id := c.Query("id")
    uintId := utils.ConvertToUint(id)
    if uintId == 0 {
        InvalidParams_1(c, uintId)
        return
    }

    product, err := h.service.GetProductByID(uintId)
    if err != nil {
        Error(c, 3103, "商品不存在")
        return
    }

    // 获取商品内容
    content, err := h.contentService.GetContentByProductID(product.ID)
    if err != nil {
        content = &sp.SpProductContent{
            ProductID: uintId,
        }
    }

    // 获取商品属性列表
    properties, err := h.propertyService.GetPropertiesByProductID(uintId)
    if err != nil {
        properties = []sp.SpProductProperty{}
    }

    // 获取SKU列表
    skus, err := h.skuService.GetSkusByProductID(uintId)
    if err != nil {
        skus = []sp.SpSku{}
    }

    // 获取SKU配置列表
    skuConfigList, err := h.skuIndexService.GetIndicesByProductID(uintId)
    if err != nil {
        skuConfigList = []sp.SpSkuIndex{}
    }

    // 获取标签
    tagIds, err := h.tagIndexService.GetTagIndicesByProductID(int(uintId))
    var tags []shop.ShopTag
    if err == nil && len(tagIds) > 0 {
        // 使用循环逐个获取标签
        for _, tagId := range tagIds {
            tag, err := h.tagService.GetTagByID(int(tagId.ID))
            if err == nil {
                tags = append(tags, *tag)
            }
            // 如果获取失败，可以选择记录日志但继续处理其他标签
        }
    }

    // 构建返回结构
    response := gin.H{
        "product":         product,
        "content":         content,
        "property_list":   properties,
        "sku_list":        skus,
        "sku_config_list": skuConfigList,
        "tags":            tags,
    }
    Success(c, response)
}

// ListProducts 分页获取商品
func (h *SpProductHandler) ListProducts(c *gin.Context) {
	var req ListProductsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}
	// 统一转换CategoryID为uint
	categoryID := utils.ConvertToUint(req.CategoryID)
	products, total, err := h.service.ListProducts(sp.ProductQueryParams{
		CategoryID: categoryID,
		State:      req.State,
		Title:      req.Title,
		Page:       req.Page,
		PageSize:   req.PageSize,
	})
	if err != nil {
		ServerError(c, err)
		return
	}

	// 为每个商品查询并添加分类信息
	productsWithCategory, err := h.enrichProductsWithCategory(products)
	if err != nil {
		ServerError(c, err)
		return
	}

	Success(c, gin.H{
		"list":  productsWithCategory,
		"total": total,
	})
}

// enrichProductsWithCategory 为商品列表添加分类信息
func (h *SpProductHandler) enrichProductsWithCategory(products []sp.SpProduct) ([]map[string]interface{}, error) {
	var result []map[string]interface{}

	for _, product := range products {
		// 将商品转换为map
		productMap := make(map[string]interface{})
		productMap["id"] = product.ID
		productMap["category_id"] = product.CategoryID
		productMap["title"] = product.Title
		productMap["state"] = product.State
		productMap["price"] = product.Price
		productMap["original_price"] = product.OriginalPrice
		productMap["cost_price"] = product.CostPrice
		productMap["stock"] = product.Stock
		productMap["picture"] = product.Picture
		productMap["sold_num"] = product.SoldNum
		productMap["sort_num"] = product.SortNum
		productMap["putaway_time"] = product.PutawayTime
		productMap["open_sku"] = product.OpenSku
		productMap["created_time"] = product.CreatedTime
		productMap["updated_time"] = product.UpdatedTime

		// 查询分类信息
		category, err := h.categoryService.GetCategoryByID(product.CategoryID)
		if err != nil {
			// 如果分类查询失败，可以记录日志但继续处理其他商品
			productMap["category"] = nil
		} else {
			productMap["category"] = category
		}

		result = append(result, productMap)
	}

	return result, nil
}

// UpdateStock 更新库存
func (h *SpProductHandler) UpdateStock(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		InvalidParams(c)
		return
	}

	var req struct {
		Stock int `json:"stock"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.UpdateStock(uint(id), req.Stock); err != nil {
		Error(c, 3104, err.Error())
		return
	}

	Success(c, nil)
}
