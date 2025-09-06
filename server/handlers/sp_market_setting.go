package handlers

import (
	"fmt"
	"net/http"
	"server/service"
	"server/utils"

	"github.com/gin-gonic/gin"
)

type SpMarketSettingHandler struct {
	categoryService *service.SpCategoryService
	productService  *service.SpProductService
}

type BreadCrumbRequest struct {
	Mode         interface{} `json:"mode"`
	ProductID    interface{}  `json:"product_id"`
	CategoryCode string `json:"category_code"`
}

type BreadcrumbVo struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}

func NewSpMarketSettingHandler(
	categoryService *service.SpCategoryService,
	productService *service.SpProductService,
) *SpMarketSettingHandler {
	return &SpMarketSettingHandler{
		categoryService: categoryService,
		productService:  productService,
	}
}

// GetBreadcrumb 获取面包屑导航
func (h *SpMarketSettingHandler) GetBreadcrumb(c *gin.Context) {
	var params BreadCrumbRequest
	if err := c.ShouldBind(&params); err != nil {
		InvalidParams(c)
		return
	}

	var out []BreadcrumbVo

	// 模式1：根据分类代码获取面包屑
	if utils.ConvertToUint(params.Mode) == 1 {
		if params.CategoryCode == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "分类代码不能为空"})
			return
		}

		parents, err := h.categoryService.GetParents(params.CategoryCode)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "获取分类信息失败"})
			return
		}

		for _, category := range parents {
			out = append(out, BreadcrumbVo{
				Title: category.Title,
				Link:  fmt.Sprintf("/collections/%s", category.Code),
			})
		}
	} else if utils.ConvertToUint(params.Mode) == 2 { // 模式2：根据产品ID获取面包屑
		if params.ProductID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "产品ID不能为空"})
			return
		}

		// 获取产品信息
		product, err := h.productService.GetProductByID(uint(utils.ConvertToUint(params.ProductID)))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "获取产品信息失败"})
			return
		}

		// 获取产品分类
		category, err := h.categoryService.GetCategoryByID(product.CategoryID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "获取分类信息失败"})
			return
		}

		if category != nil {
			parents, err := h.categoryService.GetParents(category.Code)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "获取父级分类失败"})
				return
			}

			for _, cat := range parents {
				out = append(out, BreadcrumbVo{
					Title: cat.Title,
					Link:  fmt.Sprintf("/collections/%s", cat.Code),
				})
			}
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "不支持的模式类型"})
		return
	}

	c.JSON(http.StatusOK, out)
}