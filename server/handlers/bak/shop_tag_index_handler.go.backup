package handlers

import (
	"github.com/gin-gonic/gin"
	"server/models/shop"
	"server/service"
	"strconv"
)

type ShopTagIndexHandler struct {
	service *service.ShopTagIndexService
}

func NewShopTagIndexHandler(service *service.ShopTagIndexService) *ShopTagIndexHandler {
	return &ShopTagIndexHandler{service: service}
}

// 创建标签关联
func (h *ShopTagIndexHandler) CreateTagIndex(c *gin.Context) {
	var index shop.ShopTagIndex
	if err := c.ShouldBindJSON(&index); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.CreateTagIndex(&index); err != nil {
		Error(c, 19001, err.Error())
		return
	}

	Success(c, index)
}

// 删除标签关联
// func (h *ShopTagIndexHandler) DeleteTagIndex(c *gin.Context) {
// 	productID, err := strconv.Atoi(c.Query("product_id"))
// 	if err != nil || productID == 0 {
// 		InvalidParams(c)
// 		return
// 	}

// 	tagID, err := strconv.Atoi(c.Query("tag_id"))
// 	if err != nil || tagID == 0 {
// 		InvalidParams(c)
// 		return
// 	}

// 	if err := h.service.DeleteTagIndex(productID, tagID); err != nil {
// 		Error(c, 19002, err.Error())
// 		return
// 	}

// 	Success(c, nil)
// }

// 根据产品ID获取标签关联
// func (h *ShopTagIndexHandler) GetTagIndicesByProduct(c *gin.Context) {
// 	productID, err := strconv.Atoi(c.Param("product_id"))
// 	if err != nil || productID == 0 {
// 		InvalidParams(c)
// 		return
// 	}

// 	indices, err := h.service.GetTagIndicesByProductID(productID)
// 	if err != nil {
// 		Error(c, 19003, "获取标签关联失败")
// 		return
// 	}

// 	Success(c, indices)
// }

// // 根据标签ID获取产品关联
// func (h *ShopTagIndexHandler) GetTagIndicesByTag(c *gin.Context) {
// 	tagID, err := strconv.Atoi(c.Param("tag_id"))
// 	if err != nil || tagID == 0 {
// 		InvalidParams(c)
// 		return
// 	}

// 	indices, err := h.service.GetTagIndicesByTagID(tagID)
// 	if err != nil {
// 		Error(c, 19004, "获取产品关联失败")
// 		return
// 	}

// 	Success(c, indices)
// }

// 更新标签关联排序
func (h *ShopTagIndexHandler) UpdateTagSortNum(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		InvalidParams(c)
		return
	}

	var req struct {
		SortNum int `json:"sort_num"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.UpdateTagSortNum(id, req.SortNum); err != nil {
		Error(c, 19005, err.Error())
		return
	}

	Success(c, nil)
}

// 删除产品的所有标签关联
// func (h *ShopTagIndexHandler) DeleteAllTagsByProduct(c *gin.Context) {
// 	productID, err := strconv.Atoi(c.Param("product_id"))
// 	if err != nil || productID == 0 {
// 		InvalidParams(c)
// 		return
// 	}

// 	if err := h.service.DeleteAllTagsByProductID(productID); err != nil {
// 		Error(c, 19006, err.Error())
// 		return
// 	}

// 	Success(c, nil)
// }