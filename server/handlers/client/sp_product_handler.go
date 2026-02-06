package client

import (
	"encoding/json"
	"server/models/common"
	"server/models/shop"
	"server/models/sp"
	"server/service"
	"server/utils"

	"github.com/gin-gonic/gin"
)

type ListProductsRequest struct {
	CategoryID interface{}    `json:"category_id"` // 使用interface{}接收任何类型
	State      common.MyState `json:"state"`
	Title      string         `json:"title"`
	Page       int            `json:"page_no"`
	PageSize   int            `json:"page_size"`
}
type CreateProductPropertyRequest struct {
	Title    string `json:"title"`
	Value    string `json:"value"`
	Sort_Num int    `json:"sort_num"`
	Key      string `json:"_key"`
}

// CreateProductRequest 创建商品的请求结构体
// Request
type CreateProductRequest struct {
	CategoryID     int64            `json:"category_id"`
	Content        string           `json:"content"`
	CostPrice      interface{}      `json:"cost_price"`
	Description    string           `json:"description"`
	Hot            int64            `json:"hot"`
	OpenSku        common.MyState   `json:"open_sku"`
	OriginalPrice  interface{}      `json:"original_price"`
	Picture        string           `json:"picture"`
	PictureGallery []string         `json:"picture_gallery"`
	Price          interface{}      `json:"price"`
	PropertyList   []PropertyList   `json:"property_list"`
	PutawayTime    string           `json:"putaway_time"`
	SEODescription string           `json:"seo_description"`
	SEOKeyword     string           `json:"seo_keyword"`
	SEOTitle       string           `json:"seo_title"`
	SkuList        []SkuList        `json:"sku_list"`
	SoldNum        common.MySoldNum `json:"sold_num"`
	SortNum        common.MySortNum `json:"sort_num"`
	State          common.MyState   `json:"state"`
	Stock          common.MyNumber  `json:"stock"`
	Tags           []common.MyID    `json:"tags"`
	Title          string           `json:"title"`
}

type UpdateProductRequest struct {
	ProductID      interface{}      `json:"id"`
	CategoryID     common.MyID      `json:"category_id"`
	Content        string           `json:"content"`
	CostPrice      interface{}      `json:"cost_price"`
	Description    string           `json:"description"`
	Hot            int64            `json:"hot"`
	OpenSku        common.MyState   `json:"open_sku"`
	OriginalPrice  interface{}      `json:"original_price"`
	Picture        string           `json:"picture"`
	PictureGallery []string         `json:"picture_gallery"`
	Price          interface{}      `json:"price"`
	PropertyList   []PropertyList   `json:"property_list"`
	PutawayTime    string           `json:"putaway_time"`
	SEODescription string           `json:"seo_description"`
	SEOKeyword     string           `json:"seo_keyword"`
	SEOTitle       string           `json:"seo_title"`
	SkuList        []SkuList        `json:"sku_list"`
	SoldNum        common.MySoldNum `json:"sold_num"`
	SortNum        common.MySortNum `json:"sort_num"`
	State          common.MyState   `json:"state"`
	Stock          common.MyNumber  `json:"stock"`
	Tags           []common.MyID    `json:"tags"`
	Title          string           `json:"title"`
}

type PropertyList struct {
	Key     string           `json:"_key"`
	Name    string           `json:"name"`
	SortNum common.MySortNum `json:"sort_num"`
	Title   string           `json:"title"`
	Value   string           `json:"value"`
}

type SkuList struct {
	Pord          []Pord          `json:"_pord"`
	CostPrice     int64           `json:"cost_price"`
	DefaultShow   common.MyState  `json:"default_show"`
	ID            common.MyID     `json:"id"`
	OriginalPrice int64           `json:"original_price"`
	Price         int64           `json:"price"`
	SkuCode       string          `json:"sku_code"`
	State         common.MyState  `json:"state"`
	Stock         common.MyNumber `json:"stock"`
	Title         string          `json:"title"`
}

type Pord struct {
	Label string `json:"label"`
	Value int64  `json:"value"`
}

type ClientSpProductHandler struct {
	service              *service.SpProductService
	categoryService      *service.SpCategoryService
	contentService       *service.SpProductContentService
	propertyService      *service.SpProductPropertyService
	skuService           *service.SpSkuService
	skuIndexService      *service.SpSkuIndexService
	tagIndexService      *service.ShopTagIndexService
	tagService           *service.ShopTagService
	ProdAttributes       *service.SpProdAttributesService
	ProdAttributesValues *service.SpProdAttributesValueService
}

func NewClientSpProductHandler(
	service *service.SpProductService,
	categoryService *service.SpCategoryService,
	contentService *service.SpProductContentService,
	propertyService *service.SpProductPropertyService,
	skuService *service.SpSkuService,
	skuIndexService *service.SpSkuIndexService,
	tagIndexService *service.ShopTagIndexService,
	tagService *service.ShopTagService,
	ProdAttributes *service.SpProdAttributesService,
	prodAttributes *service.SpProdAttributesValueService,

) *ClientSpProductHandler {
	return &ClientSpProductHandler{
		service:              service,
		categoryService:      categoryService,
		contentService:       contentService,
		propertyService:      propertyService,
		skuService:           skuService,
		skuIndexService:      skuIndexService,
		tagIndexService:      tagIndexService,
		tagService:           tagService,
		ProdAttributes:       ProdAttributes,
		ProdAttributesValues: prodAttributes,
	}
}

// ListProducts 分页获取商品
func (h *ClientSpProductHandler) ListProducts(c *gin.Context) {
	var req ListProductsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.InvalidParams(c)
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
		utils.ServerError(c, err)
		return
	}

	// 为每个商品查询并添加分类信息
	productsWithCategory, err := h.enrichProductsWithCategory(products)
	if err != nil {
		utils.ServerError(c, err)
		return
	}

	utils.Success(c, gin.H{
		"list":  productsWithCategory,
		"total": total,
	})
}

// enrichProductsWithCategory 为商品列表添加分类信息
func (h *ClientSpProductHandler) enrichProductsWithCategory(products []sp.SpProduct) ([]map[string]interface{}, error) {
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

// SoftDeleteProduct 软删除商品
func (h *ClientSpProductHandler) SoftDeleteProduct(c *gin.Context) {
	id := c.Query("id")
	uintId := utils.ConvertToUint(id)
	if uintId == 0 {
		utils.InvalidParams_1(c, uintId)
		return
	}

	if err := h.service.SoftDeleteProduct(common.MyID(uintId)); err != nil {
		utils.Error(c, 3105, err.Error())
		return
	}

	utils.Success(c, nil)
}

// GetClientProduct 获取前端商品详情
func (h *ClientSpProductHandler) GetClientProduct(c *gin.Context) {
	id := c.Query("id")
	uintId := utils.ConvertToUint(id)
	if uintId == 0 {
		utils.InvalidParams_1(c, uintId)
		return
	}

	product, err := h.service.GetProductByID(common.MyID(uintId))
	if err != nil {
		utils.Error(c, 3103, "商品不存在")
		return
	}

	// 获取商品内容
	content, err := h.contentService.GetContentByProductID(common.MyID(uintId))
	if err != nil {
		content = &sp.SpProductContent{
			ProductID: common.MyID(uintId),
		}
	}

	// 获取商品属性列表
	properties, err := h.propertyService.GetPropertiesByProductID(common.MyID(uintId))
	if err != nil {
		properties = []sp.SpProductProperty{}
	}

	// 获取SKU列表
	skus, err := h.skuService.GetSkusByProductID(common.MyID(uintId))
	if err != nil {
		skus = []sp.SpSku{}
	}

	// 转换SKU配置为前端需要的格式
	frontSkuConfig, err := h.getSkuConfig(common.MyID(uintId))
	if err != nil {
		frontSkuConfig = []SpProductProdFrontVo{}
	}

	// 获取标签
	tagIds, err := h.tagIndexService.GetTagIndicesByProductID(common.MyID(uintId))
	var tags []shop.ShopTag
	if err == nil && len(tagIds) > 0 {
		for _, tagId := range tagIds {
			tag, err := h.tagService.GetTagByID(common.MyID(tagId.ID))
			if err == nil {
				tags = append(tags, *tag)
			}
		}
	}

	// 处理图片集
	var pictureGallery []string
	if product.PictureGallery != nil {
		json.Unmarshal(product.PictureGallery, &pictureGallery)
	}

	// 构建前端需要的响应结构 - ProductInfo
	response := gin.H{
		"id":              product.ID,
		"category_id":     product.CategoryID,
		"title":           product.Title,
		"state":           product.State,
		"price":           product.Price,
		"original_price":  product.OriginalPrice,
		"stock":           product.Stock,
		"picture":         product.Picture,
		"picture_gallery": pictureGallery,
		"description":     product.Description,
		"sold_num":        product.SoldNum,
		"open_sku":        product.OpenSku,
		"sort_num":        product.SortNum,
		"putaway_time":    product.PutawayTime,
		"content":         content.Content,
		"seo_title":       content.SeoTitle,
		"seo_keyword":     content.SeoKeyword,
		"seo_description": content.SeoDescription,
		"property_list":   properties,
		"sku_list":        skus,
		"sku_config":      frontSkuConfig,
		"tags":            tags,
	}

	utils.Success(c, response)
}

type SpProductProdFrontVo struct {
	ID      common.MyID                 `json:"id"`
	Title   string                      `json:"title"` //属性名称
	SortNum common.MySortNum            `json:"sort_num"`
	Value   []SpProductProdValueFrontVo `json:"value"` // 属性值
}
type SpProductProdValueFrontVo struct {
	ID      common.MyID      `json:"id"`
	Title   string           `json:"title"` //值名称
	SortNum common.MySortNum `json:"sort_num"`
}

func (h *ClientSpProductHandler) getSkuConfig(productID common.MyID) ([]SpProductProdFrontVo, error) {
	// 获取 SKU 索引列表
	skuIndices, err := h.skuIndexService.GetIndicesByProductID(productID)
	if err != nil {
		return nil, err
	}
	if len(skuIndices) == 0 {
		return []SpProductProdFrontVo{}, nil
	}

	// 使用 map 去重收集属性 ID 和属性值 ID
	attributeIDMap := make(map[common.MyID]bool)
	attributeValueIDMap := make(map[common.MyID]bool)

	for _, index := range skuIndices {
		attributeIDMap[index.ProdAttributesID] = true
		attributeValueIDMap[index.ProdAttributesValueID] = true
	}

	// 将去重后的属性 ID 转换为切片
	var attributeIDs []common.MyID
	for id := range attributeIDMap {
		attributeIDs = append(attributeIDs, id)
	}

	// 将去重后的属性值 ID 转换为切片
	var attributeValueIDs []common.MyID
	for id := range attributeValueIDMap {
		attributeValueIDs = append(attributeValueIDs, id)
	}

	// 获取属性列表
	var attributes []sp.SpProdAttributes
	for _, id := range attributeIDs {
		attr, err := h.ProdAttributes.GetAttributeByID(id)
		if err != nil {
			return nil, err
		}
		attributes = append(attributes, *attr)
	}

	// 获取属性值列表
	var attributeValues []sp.SpProdAttributesValue
	for _, id := range attributeValueIDs {
		val, err := h.ProdAttributesValues.GetValueByID(id)
		if err != nil {
			return nil, err
		}
		attributeValues = append(attributeValues, *val)
	}

	// 构建返回结果
	var result []SpProductProdFrontVo
	for _, attr := range attributes {
		vo := SpProductProdFrontVo{
			ID:      attr.ID,
			Title:   attr.Title,
			SortNum: attr.SortNum,
		}

		// 过滤并匹配属性值
		var values []SpProductProdValueFrontVo
		for _, val := range attributeValues {
			if val.ProdAttributesID == attr.ID {
				values = append(values, SpProductProdValueFrontVo{
					ID:      val.ID,
					Title:   val.Title,
					SortNum: val.SortNum,
				})
			}
		}
		vo.Value = values
		result = append(result, vo)
	}

	return result, nil
}
