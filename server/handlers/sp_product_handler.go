package handlers

import (
	"encoding/json"
	"fmt"
	"server/models/shop"
	"server/models/sp"
	"server/service"
	"server/utils"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type ListProductsRequest struct {
	CategoryID interface{} `json:"category_id"` // 使用interface{}接收任何类型
	State      int         `json:"state"`
	Title      string      `json:"title"`
	Page       int         `json:"page_no"`
	PageSize   int         `json:"page_size"`
}
type CreateProductPropertyRequest struct {
	Title     string `json:"title"`
	Value    string `json:"value"`
	Sort_Num int    `json:"sort_num"`
	Key string `json:"_key"`
}

// CreateProductRequest 创建商品的请求结构体
type CreateProductRequest struct {
	CategoryID     uint                           `json:"category_id"`
	Title          string                         `json:"title"`
	State          int                            `json:"state"`
	Price          float64                        `json:"price"`
	OriginalPrice  float64                         `json:"original_price"`
	CostPrice      float64                         `json:"cost_price"`
	Stock          int                            `json:"stock"`
	Picture        string                         `json:"picture"`
	PictureGallery []string                       `json:"picture_gallery"`
	Description    string                         `json:"description"`
	SoldNum        int                            `json:"sold_num"`
	SortNum        int                            `json:"sort_num"`
	PutawayTime    string                         `json:"putaway_time"`
	Content        string                         `json:"content"`
	SeoTitle       string                         `json:"seo_title"`
	SeoKeyword     string                         `json:"seo_keyword"`
	SeoDescription string                         `json:"seo_description"`
	PropertyList   []CreateProductPropertyRequest `json:"property_list"`
	OpenSku        int                            `json:"open_sku"`
	SkuList        []sp.SpSku                     `json:"sku_list"`
	Tags           []int                          `json:"tags"`
	Hot            int                            `json:"hot"`
}

type SpProductHandler struct {
	service              *service.SpProductService
	categoryService      *service.SpCategoryService
	contentService       *service.SpProductContentService
	propertyService      *service.SpProductPropertyService
	skuService           *service.SpSkuService
	skuIndexService      *service.SpSkuIndexService
	tagIndexService      *service.ShopTagIndexService
	tagService           *service.ShopTagService
	ProdAttributesValues *service.SpProdAttributesValueService
}

func NewSpProductHandler(
	service *service.SpProductService,
	categoryService *service.SpCategoryService,
	contentService *service.SpProductContentService,
	propertyService *service.SpProductPropertyService,
	skuService *service.SpSkuService,
	skuIndexService *service.SpSkuIndexService,
	tagIndexService *service.ShopTagIndexService,
	tagService *service.ShopTagService,
	prodAttributes *service.SpProdAttributesValueService,

) *SpProductHandler {
	return &SpProductHandler{
		service:              service,
		categoryService:      categoryService,
		contentService:       contentService,
		propertyService:      propertyService,
		skuService:           skuService,
		skuIndexService:      skuIndexService,
		tagIndexService:      tagIndexService,
		tagService:           tagService,
		ProdAttributesValues: prodAttributes,
	}
}

// CreateProduct 创建商品
func (h *SpProductHandler) CreateProduct(c *gin.Context) {
	var req CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}

	// 检查分类是否存在
	if req.CategoryID != 0 {
		category, err := h.categoryService.GetCategoryByID(req.CategoryID)
		if err != nil || category == nil {
			Error(c, 3105, "分类不存在")
			return
		}
	}

	// 解析上架时间
	var putawayTime time.Time
	if req.PutawayTime != "" {
		var err error
		putawayTime, err = time.Parse(time.RFC3339, req.PutawayTime)
		if err != nil {
			// 如果解析失败，使用当前时间
			putawayTime = time.Now()
		}
	} else {
		putawayTime = time.Now()
	}

	// price, _ := strconv.ParseFloat(req.Price, 64)
	// originalPrice, _ := strconv.ParseFloat(req.OriginalPrice, 64)
	// costPrice, _ := strconv.ParseFloat(req.CostPrice, 64)

	rawJSON, _ := json.Marshal(req.PictureGallery)
	PictureGallery := json.RawMessage(rawJSON)
	// 创建商品基本信息
	product := sp.SpProduct{
		CategoryID:     req.CategoryID,
		Title:          req.Title,
		State:          uint8(req.State),
		Price:          req.Price,
		OriginalPrice:  req.OriginalPrice,
		CostPrice:      req.CostPrice,
		Stock:          uint(req.Stock),
		OpenSku:        uint8(req.OpenSku),
		Picture:        req.Picture,
		PictureGallery: PictureGallery,
		Description:    req.Description,
		SoldNum:        uint16(req.SoldNum),
		Version:        0,
		SortNum:        uint16(req.SortNum),
		PutawayTime:    &putawayTime,
		Hot:            uint8(req.Hot),
	}

	pro,err := h.service.CreateProduct(&product)
	// 创建商品基本信息
	if err != nil {
		Error(c, 3101, err.Error())
		return
	} 

	content := sp.SpProductContent{
		ProductID:      pro.ID,
		Content:        req.Content,
		SeoTitle:       req.SeoTitle,
		SeoKeyword:     req.SeoKeyword,
		SeoDescription: req.SeoDescription,
	}

	if err := h.contentService.CreateContent(&content); err != nil {
		Error(c, 3101, err.Error())
		return
	}
	fmt.Println(pro)
	// 保存商品属性
	if len(req.PropertyList) > 0 {
		if err := h.saveProperties(pro.ID, req.PropertyList); err != nil {
			Error(c, 3101, "保存商品属性失败: "+err.Error())
			return
		}
	}

	// 保存SKU
	if len(req.SkuList) > 0 {
		if err := h.saveSkus(pro.ID, req.SkuList); err != nil {
			Error(c, 3101, "保存SKU失败: "+err.Error())
			return
		}

		// 同步SKU配置
		if err := h.syncProductSkuConfig(pro.ID); err != nil {
			Error(c, 3101, "同步SKU配置失败: "+err.Error())
			return
		}
	}

	// 保存标签
	if len(req.Tags) > 0 {
		if err := h.saveTags(pro.ID, req.Tags); err != nil {
			Error(c, 3101, "保存标签失败: "+err.Error())
			return
		}
	}

	// 获取完整的商品信息并返回
	fullProductInfo, err := h.getFullProductInfo(pro.ID)
	if err != nil {
		Error(c, 3101, "获取商品详情失败: "+err.Error())
		return
	}

	Success(c, fullProductInfo)

}

// saveProperties 保存商品属性
func (h *SpProductHandler) saveProperties(productID uint, properties []CreateProductPropertyRequest) error {
	for i := range properties {
		property := sp.SpProductProperty{
			ProductID: productID,
			Title:   properties[i].Title,
			Value:   properties[i].Value,
			SortNum: uint16(properties[i].Sort_Num),
		}
		if err := h.propertyService.CreateProperty(&property); err != nil {
			return err
		}
	}
	return nil
}

// saveSkus 保存SKU
func (h *SpProductHandler) saveSkus(productID uint, skus []sp.SpSku) error {
	for i := range skus {
		skus[i].ProductID = productID
		if err := h.skuService.CreateSku(&skus[i]); err != nil {
			return err
		}
	}
	return nil
}

// syncProductSkuConfig 同步商品SKU配置
func (h *SpProductHandler) syncProductSkuConfig(productID uint) error {
	// 1. 获取商品SKU列表
	skus, err := h.skuService.GetSkusByProductID(productID)
	if err != nil {
		return fmt.Errorf("获取SKU列表失败: %v", err)
	}

	// 2. 删除旧的SKU索引
	if err := h.skuIndexService.DeleteByProductID(productID); err != nil {
		return fmt.Errorf("删除旧SKU索引失败: %v", err)
	}

	// 3. 如果SKU列表不为空，处理新的索引
	if len(skus) > 0 {
		// 收集所有属性值ID
		prodValueIds := make(map[string]struct{})
		for _, sku := range skus {
			// 分割SKU编码（Java中是分号分隔）
			splitCodes := strings.Split(sku.SkuCode, ";")
			for _, code := range splitCodes {
				if code != "" {
					prodValueIds[code] = struct{}{}
				}
			}
		}

		// 逐个查询属性值信息
		for valueID := range prodValueIds {
			parsedValue, err := strconv.ParseUint(valueID, 10, 32)
			if err != nil {
				return nil
			}
			attrValue, err := h.ProdAttributesValues.GetValuesByAttributeID(uint(parsedValue))
			if err != nil {
				return fmt.Errorf("查询属性值失败(ID=%s): %v", valueID, err)
			}

			// 创建SKU索引
			index := &sp.SpSkuIndex{
				ProductID:             productID,
				ProdAttributesID:      attrValue[0].ProdAttributesID,
				ProdAttributesValueID: attrValue[0].ID,
			}
			if err := h.skuIndexService.CreateIndex(index); err != nil {
				return fmt.Errorf("创建SKU索引失败: %v", err)
			}
		}
	}

	return nil
}

// saveTags 保存标签关联
func (h *SpProductHandler) saveTags(productID uint, tagIDs []int) error {
	// 这里实现标签关联的保存逻辑
	// 需要调用 tagIndexService 来创建标签索引
	for _, tagID := range tagIDs {
		tagIndex := shop.ShopTagIndex{
			ProductID: int(productID),
			TagID:     tagID,
		}
		if err := h.tagIndexService.CreateTagIndex(&tagIndex); err != nil {
			return err
		}
	}
	return nil
}

// getFullProductInfo 获取完整的商品信息
func (h *SpProductHandler) getFullProductInfo(productID uint) (gin.H, error) {
	// 获取商品基本信息
	product, err := h.service.GetProductByID(productID)
	if err != nil {
		return nil, err
	}

	// 获取商品内容
	content, err := h.contentService.GetContentByProductID(productID)
	if err != nil {
		content = &sp.SpProductContent{ProductID: productID}
	}

	// 获取商品属性
	properties, err := h.propertyService.GetPropertiesByProductID(productID)
	if err != nil {
		properties = []sp.SpProductProperty{}
	}

	// 获取SKU列表
	skus, err := h.skuService.GetSkusByProductID(productID)
	if err != nil {
		skus = []sp.SpSku{}
	}

	// 获取SKU配置
	skuConfigList, err := h.skuIndexService.GetIndicesByProductID(productID)
	if err != nil {
		skuConfigList = []sp.SpSkuIndex{}
	}

	// 获取标签
	tagIds, err := h.tagIndexService.GetTagIndicesByProductID(int(productID))
	var tags []shop.ShopTag
	if err == nil && len(tagIds) > 0 {
		for _, tagId := range tagIds {
			tag, err := h.tagService.GetTagByID(int(tagId.ID))
			if err == nil {
				tags = append(tags, *tag)
			}
		}
	}

	return gin.H{
		"product":         product,
		"content":         content,
		"property_list":   properties,
		"sku_list":        skus,
		"sku_config_list": skuConfigList,
		"tags":            tags,
	}, nil
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

// SoftDeleteProduct 软删除商品
func (h *SpProductHandler) SoftDeleteProduct(c *gin.Context) {
	id := c.Query("id")
	uintId := utils.ConvertToUint(id)
	if uintId == 0 {
		InvalidParams_1(c, uintId)
		return
	}

	if err := h.service.SoftDeleteProduct(uintId); err != nil {
		Error(c, 3105, err.Error())
		return
	}

	Success(c, nil)
}