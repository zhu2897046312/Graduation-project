package handlers

import (
	"encoding/json"
	"fmt"
	"server/models/common"
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
	Title    string `json:"title"`
	Value    string `json:"value"`
	Sort_Num int    `json:"sort_num"`
	Key      string `json:"_key"`
}

// CreateProductRequest 创建商品的请求结构体
// Request
type CreateProductRequest struct {
	CategoryID     int64          `json:"category_id"`
	Content        string         `json:"content"`
	CostPrice      interface{}    `json:"cost_price"`
	Description    string         `json:"description"`
	Hot            int64          `json:"hot"`
	OpenSku        int64          `json:"open_sku"`
	OriginalPrice  interface{}    `json:"original_price"`
	Picture        string         `json:"picture"`
	PictureGallery []string       `json:"picture_gallery"`
	Price          interface{}    `json:"price"`
	PropertyList   []PropertyList `json:"property_list"`
	PutawayTime    string         `json:"putaway_time"`
	SEODescription string         `json:"seo_description"`
	SEOKeyword     string         `json:"seo_keyword"`
	SEOTitle       string         `json:"seo_title"`
	SkuList        []SkuList      `json:"sku_list"`
	SoldNum        int64          `json:"sold_num"`
	SortNum        int64          `json:"sort_num"`
	State          int64          `json:"state"`
	Stock          int64          `json:"stock"`
	Tags           []common.MyID  `json:"tags"`
	Title          string         `json:"title"`
}

type UpdateProductRequest struct {
	ProductID      interface{}    `json:"id"`
	CategoryID     int64          `json:"category_id"`
	Content        string         `json:"content"`
	CostPrice      interface{}    `json:"cost_price"`
	Description    string         `json:"description"`
	Hot            int64          `json:"hot"`
	OpenSku        int64          `json:"open_sku"`
	OriginalPrice  interface{}    `json:"original_price"`
	Picture        string         `json:"picture"`
	PictureGallery []string       `json:"picture_gallery"`
	Price          interface{}    `json:"price"`
	PropertyList   []PropertyList `json:"property_list"`
	PutawayTime    string         `json:"putaway_time"`
	SEODescription string         `json:"seo_description"`
	SEOKeyword     string         `json:"seo_keyword"`
	SEOTitle       string         `json:"seo_title"`
	SkuList        []SkuList      `json:"sku_list"`
	SoldNum        int64          `json:"sold_num"`
	SortNum        int64          `json:"sort_num"`
	State          int64          `json:"state"`
	Stock          int64          `json:"stock"`
	Tags           []common.MyID  `json:"tags"`
	Title          string         `json:"title"`
}

type PropertyList struct {
	Key     string `json:"_key"`
	Name    string `json:"name"`
	SortNum int64  `json:"sort_num"`
	Title   string `json:"title"`
	Value   string `json:"value"`
}

type SkuList struct {
	Pord          []Pord `json:"_pord"`
	CostPrice     int64  `json:"cost_price"`
	DefaultShow   int64  `json:"default_show"`
	ID            int64  `json:"id"`
	OriginalPrice int64  `json:"original_price"`
	Price         int64  `json:"price"`
	SkuCode       string `json:"sku_code"`
	State         int64  `json:"state"`
	Stock         int64  `json:"stock"`
	Title         string `json:"title"`
}

type Pord struct {
	Label string `json:"label"`
	Value int64  `json:"value"`
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
	ProdAttributes       *service.SpProdAttributesService
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
	ProdAttributes *service.SpProdAttributesService,
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
		ProdAttributes:       ProdAttributes,
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
		category, err := h.categoryService.GetCategoryByID(common.MyID(req.CategoryID))
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
	Price := float64(utils.ConvertToUint(req.Price))
	OriginalPrice := float64(utils.ConvertToUint(req.OriginalPrice))
	CostPrice := float64(utils.ConvertToUint(req.CostPrice))
	if Price == 0 || OriginalPrice == 0 || CostPrice == 0 {
		for i := range req.SkuList {
			if utils.ConvertToUint(req.SkuList[i].State) == 1 {
				Price = float64(utils.ConvertToUint(req.SkuList[i].Price))
				OriginalPrice = float64(utils.ConvertToUint(req.SkuList[i].OriginalPrice))
				CostPrice = float64(utils.ConvertToUint(req.SkuList[i].CostPrice))
				break
			}
		}
	}
	// 创建商品基本信息
	product := sp.SpProduct{
		CategoryID:     common.MyID(req.CategoryID),
		Title:          req.Title,
		State:          uint8(req.State),
		Price:          Price,
		OriginalPrice:  OriginalPrice,
		CostPrice:      CostPrice,
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

	pro, err := h.service.CreateProduct(&product)
	// 创建商品基本信息
	if err != nil {
		Error(c, 3101, err.Error())
		return
	}

	content := sp.SpProductContent{
		ProductID:      pro.ID,
		Content:        req.Content,
		SeoTitle:       req.SEOTitle,
		SeoKeyword:     req.SEOKeyword,
		SeoDescription: req.SEODescription,
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
func (h *SpProductHandler) saveProperties(productID common.MyID, properties []PropertyList) error {
	for i := range properties {
		property := sp.SpProductProperty{
			ProductID: productID,
			Title:     properties[i].Title,
			Value:     properties[i].Value,
			SortNum:   uint16(properties[i].SortNum),
		}
		if err := h.propertyService.CreateProperty(&property); err != nil {
			return err
		}
	}
	return nil
}

// saveSkus 保存SKU
func (h *SpProductHandler) saveSkus(productID common.MyID, skus []SkuList) error {
	if err := h.skuService.DeleteByProductID(productID); err != nil {
		return err
	}
	for i := range skus {
		sku := sp.SpSku{
			ProductID:     productID,
			SkuCode:       skus[i].SkuCode,
			Title:         skus[i].Title,
			Price:         float64(utils.ConvertToUint(skus[i].Price)),
			OriginalPrice: float64(utils.ConvertToUint(skus[i].OriginalPrice)),
			CostPrice:     float64(utils.ConvertToUint(skus[i].CostPrice)),
			Stock:         uint(skus[i].Stock),
			DefaultShow:   uint8(skus[i].DefaultShow),
			State:         uint8(skus[i].State),
			Version:       0,
		}
		if err := h.skuService.CreateSku(&sku); err != nil {
			return err
		}
	}
	return nil
}

// syncProductSkuConfig 同步商品SKU配置
func (h *SpProductHandler) syncProductSkuConfig(productID common.MyID) error {
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
		// 遍历所有SKU
		for _, sku := range skus {
			// 分割SKU编码（分隔）
			splitCodes := strings.Split(sku.SkuCode, ";")

			// 为当前SKU的每个属性值创建索引
			for _, code := range splitCodes {
				if code == "" {
					continue
				}

				// 解析属性值ID
				parsedValue, err := strconv.ParseUint(code, 10, 32)
				if err != nil {
					return fmt.Errorf("解析属性值ID失败(%s): %v", code, err)
				}

				// 获取属性值信息
				attrValue, err := h.ProdAttributesValues.GetValuesByAttributeID(common.MyID(parsedValue))
				if err != nil {
					return fmt.Errorf("查询属性值失败(ID=%s): %v", code, err)
				}
				if len(attrValue) == 0 {
					return fmt.Errorf("属性值不存在(ID=%s)", code)
				}

				// 创建SKU索引，使用当前SKU的ID作为SkuID
				index := &sp.SpSkuIndex{
					SkuID:                 sku.ID, // 这里填充当前SKU的ID
					ProductID:             productID,
					ProdAttributesID:      attrValue[0].ProdAttributesID,
					ProdAttributesValueID: attrValue[0].ID,
				}

				if err := h.skuIndexService.CreateIndex(index); err != nil {
					return fmt.Errorf("创建SKU索引失败: %v", err)
				}
			}
		}
	}

	return nil
}

// saveTags 保存标签关联
func (h *SpProductHandler) saveTags(productID common.MyID, tagIDs []common.MyID) error {
	// 这里实现标签关联的保存逻辑
	// 需要调用 tagIndexService 来创建标签索引
	for _, tagID := range tagIDs {
		tagIndex := shop.ShopTagIndex{
			ProductID: common.MyID(productID),
			TagID:     tagID,
		}
		if err := h.tagIndexService.CreateTagIndex(&tagIndex); err != nil {
			return err
		}
	}
	return nil
}

// getFullProductInfo 获取完整的商品信息
func (h *SpProductHandler) getFullProductInfo(productID common.MyID) (gin.H, error) {
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
	tagIds, err := h.tagIndexService.GetTagIndicesByProductID(common.MyID(productID))
	var tags []shop.ShopTag
	if err == nil && len(tagIds) > 0 {
		for _, tagId := range tagIds {
			tag, err := h.tagService.GetTagByID(common.MyID(tagId.ID))
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

	var req UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}
	fmt.Println(req)
	// 检查分类是否存在
	if req.CategoryID != 0 {
		category, err := h.categoryService.GetCategoryByID(common.MyID(req.CategoryID))
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
	Price := float64(utils.ConvertToUint(req.Price))
	OriginalPrice := float64(utils.ConvertToUint(req.OriginalPrice))
	CostPrice := float64(utils.ConvertToUint(req.CostPrice))
	if Price == 0 || OriginalPrice == 0 || CostPrice == 0 {
		for i := range req.SkuList {
			if utils.ConvertToUint(req.SkuList[i].State) == 1 {
				Price = float64(utils.ConvertToUint(req.SkuList[i].Price))
				OriginalPrice = float64(utils.ConvertToUint(req.SkuList[i].OriginalPrice))
				CostPrice = float64(utils.ConvertToUint(req.SkuList[i].CostPrice))
				break
			}
		}
	}
	productID := utils.ConvertToUint(req.ProductID)
	// 创建商品基本信息
	product := sp.SpProduct{
		ID:             common.MyID(productID),
		CategoryID:     common.MyID(req.CategoryID),
		Title:          req.Title,
		State:          uint8(req.State),
		Price:          Price,
		OriginalPrice:  OriginalPrice,
		CostPrice:      CostPrice,
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

	err := h.service.UpdateProduct(&product)
	// 创建商品基本信息
	if err != nil {
		Error(c, 3101, err.Error())
		return
	}

	content := sp.SpProductContent{
		ProductID:      common.MyID(productID),
		Content:        req.Content,
		SeoTitle:       req.SEOTitle,
		SeoKeyword:     req.SEOKeyword,
		SeoDescription: req.SEODescription,
	}

	if err := h.contentService.UpdateContent(&content); err != nil {
		Error(c, 3101, err.Error())
		return
	}
	// 保存商品属性
	if len(req.PropertyList) > 0 {
		if err := h.saveProperties(common.MyID(productID), req.PropertyList); err != nil {
			Error(c, 3101, "保存商品属性失败: "+err.Error())
			return
		}
	}

	// 保存SKU
	if len(req.SkuList) > 0 {
		if err := h.saveSkus(common.MyID(productID), req.SkuList); err != nil {
			Error(c, 3101, "保存SKU失败: "+err.Error())
			return
		}

		// 同步SKU配置
		if err := h.syncProductSkuConfig(common.MyID(productID)); err != nil {
			Error(c, 3101, "同步SKU配置失败: "+err.Error())
			return
		}
	}

	// 保存标签
	if len(req.Tags) > 0 {
		if err := h.saveTags(common.MyID(productID), req.Tags); err != nil {
			Error(c, 3101, "保存标签失败: "+err.Error())
			return
		}
	}
	Success(c, nil)
}

// GetProduct 获取商品详情
func (h *SpProductHandler) GetProduct(c *gin.Context) {
	id := c.Query("id")
	uintId := utils.ConvertToUint(id)
	if uintId == 0 {
		InvalidParams_1(c, uintId)
		return
	}

	product, err := h.service.GetProductByID(common.MyID(uintId))
	if err != nil {
		Error(c, 3103, "商品不存在")
		return
	}

	// 获取商品内容
	content, err := h.contentService.GetContentByProductID(product.ID)
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

	// 获取SKU配置列表
	skuConfigList, err := h.skuIndexService.GetIndicesByProductID(common.MyID(uintId))
	if err != nil {
		skuConfigList = []sp.SpSkuIndex{}
	}

	// 获取标签
	tagIds, err := h.tagIndexService.GetTagIndicesByProductID(common.MyID(uintId))
	var tags []shop.ShopTag
	if err == nil && len(tagIds) > 0 {
		// 使用循环逐个获取标签
		for _, tagId := range tagIds {
			tag, err := h.tagService.GetTagByID(common.MyID(tagId.ID))
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

	if err := h.service.UpdateStock(common.MyID(id), req.Stock); err != nil {
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

	if err := h.service.SoftDeleteProduct(common.MyID(uintId)); err != nil {
		Error(c, 3105, err.Error())
		return
	}

	Success(c, nil)
}

// GetProductFrontInfo 获取前端商品详情
func (h *SpProductHandler) GetProductFrontInfo(c *gin.Context) {
	id := c.Query("id")
	uintId := utils.ConvertToUint(id)
	if uintId == 0 {
		InvalidParams_1(c, uintId)
		return
	}

	// 获取完整的商品信息
	fullInfo, err := h.getFullProductInfo(common.MyID(uintId))
	if err != nil {
		Error(c, 3103, "商品不存在")
		return
	}

	// 转换为前端需要的格式
	frontInfo := h.convertToFrontInfo(fullInfo)
	Success(c, frontInfo)
}

// convertToFrontInfo 将内部数据结构转换为前端需要的格式
func (h *SpProductHandler) convertToFrontInfo(fullInfo gin.H) gin.H {
	product := fullInfo["product"].(*sp.SpProduct)
	content := fullInfo["content"].(*sp.SpProductContent)
	properties := fullInfo["property_list"].([]sp.SpProductProperty)
	skus := fullInfo["sku_list"].([]sp.SpSku)
	skuConfig := fullInfo["sku_config_list"].([]sp.SpSkuIndex)
	tags := fullInfo["tags"].([]shop.ShopTag)

	// 处理图片集
	var pictureGallery []string
	if product.PictureGallery != nil {
		json.Unmarshal(product.PictureGallery, &pictureGallery)
	}

	// 转换为前端需要的结构
	return gin.H{
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
		"sku_config":      skuConfig,
		"tags":            tags,
	}
}

// GetClientProduct 获取前端商品详情
func (h *SpProductHandler) GetClientProduct(c *gin.Context) {
	id := c.Query("id")
	uintId := utils.ConvertToUint(id)
	if uintId == 0 {
		InvalidParams_1(c, uintId)
		return
	}

	product, err := h.service.GetProductByID(common.MyID(uintId))
	if err != nil {
		Error(c, 3103, "商品不存在")
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

	Success(c, response)
}

type SpProductProdFrontVo struct {
	ID      common.MyID                 `json:"id"`
	Title   string                      `json:"title"` //属性名称
	SortNum uint16                      `json:"sort_num"`
	Value   []SpProductProdValueFrontVo `json:"value"` // 属性值
}
type SpProductProdValueFrontVo struct {
	ID      common.MyID `json:"id"`
	Title   string      `json:"title"` //值名称
	SortNum uint16      `json:"sort_num"`
}

func (h *SpProductHandler) getSkuConfig(productID common.MyID) ([]SpProductProdFrontVo, error) {
	// 获取 SKU 索引列表
	skuIndices, err := h.skuIndexService.GetIndicesByProductID(productID)
	if err != nil {
		return nil, err
	}
	if len(skuIndices) == 0 {
		return []SpProductProdFrontVo{}, nil
	}

	// 收集属性 ID 和属性值 ID
	var attributeIDs []common.MyID
	var attributeValueIDs []common.MyID
	for _, index := range skuIndices {
		attributeIDs = append(attributeIDs, index.ProdAttributesID)
		attributeValueIDs = append(attributeValueIDs, index.ProdAttributesValueID)
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
