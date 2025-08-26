package router

import (
	"server/handlers"
	"server/middleware"
	"server/service"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func SetupRouter(r *gin.Engine, factory *service.ServiceFactory, rdb *redis.Client) *gin.Engine {
	r.Use(
		middleware.Cors(), // CORS中间件         // 服务注入中间件
	)
	api := r.Group("/api")
	{
		ossHandler := handlers.NewOssHandler()
		// SP分类路由组
		spCategoryHandler := handlers.NewSpCategoryHandler(factory.GetSpCategoryService())
		// SP商品属性路由组
		spAttrHandler := handlers.NewSpProdAttributesHandler(factory.GetSpProdAttributesService())
		// 商品标签路由组
		tagHandler := handlers.NewShopTagHandler(factory.GetShopTagService())
		// 公开路由（不需要认证）
		public := api.Group("")
		{
			// 管理员登录
			adminHandler := handlers.NewCoreAdminHandler(factory.GetCoreAdminService(), rdb)
			public.POST("/manage/core/auth/login", adminHandler.LoginAdmin)
		}
		// 管理员认证路由
		adminAuth := api.Group("/manage")
		adminAuth.Use(middleware.AuthMiddleware(rdb)) // 添加管理员认证中间件
		{
			// 获取当前管理员信息
			adminHandler := handlers.NewCoreAdminHandler(factory.GetCoreAdminService(), rdb)
			adminAuth.GET("/core/auth/info", adminHandler.GetAdminInfo)
			adminAuth.GET("/core/auth/enumDict", adminHandler.GetEnumDict)
			adminOssGroup := adminAuth.Group("/core/oss")
			{
				adminOssGroup.POST("/uploadFile", ossHandler.UploadFile)
				adminOssGroup.POST("/uploadFiles", ossHandler.UploadMultipleFiles)
				adminOssGroup.DELETE("/deleteFile", ossHandler.DeleteFile)
				adminOssGroup.GET("/fileInfo", ossHandler.GetFileInfo)
			}
			adminGroup := adminAuth.Group("/core/admins")
			{
				adminGroup.POST("", adminHandler.CreateAdmin)
				adminGroup.PUT("/:id", adminHandler.UpdateAdmin)
				adminGroup.GET("/:id", adminHandler.GetAdmin)
				adminGroup.PATCH("/:id/status", adminHandler.UpdateAdminStatus)
				adminGroup.PATCH("/:id/password", adminHandler.UpdateAdminPassword)
			}
			//商品路由组
			productHandler := handlers.NewSpProductHandler(
				factory.GetSpProductService(),
				factory.GetSpCategoryService(),
				factory.GetSpProductContentService(),
				factory.GetSpProductPropertyService(),
				factory.GetSpSkuService(),
				factory.GetSpSkuIndexService(),
				factory.GetShopTagIndexService(),
				factory.GetShopTagService(),
			)
			productGroup := adminAuth.Group("/shop/product")
			{
				productGroup.POST("", productHandler.CreateProduct)
				productGroup.PUT("/:id", productHandler.UpdateProduct)
				productGroup.GET("/info", productHandler.GetProduct)
				productGroup.POST("/list", productHandler.ListProducts)
				productGroup.PATCH("/:id/stock", productHandler.UpdateStock)
			}
			spCategoryGroup := adminAuth.Group("/shop/category")
			{
				spCategoryGroup.GET("/tree", spCategoryHandler.GetCategoryTree)
			}
			spAttrGroup := adminAuth.Group("/shop/prodAttributes")
			{
				spAttrGroup.POST("/list", spAttrHandler.GetAttributesByPage)
			}
			tagGroup := adminAuth.Group("/shop/tag")
			{
				tagGroup.POST("/list", tagHandler.ListTags)
			}
		}

		spCategoryGroup := api.Group("/sp/categories")
		{
			spCategoryGroup.POST("", spCategoryHandler.CreateCategory)
			spCategoryGroup.PUT("/:id", spCategoryHandler.UpdateCategory)
			spCategoryGroup.GET("/:id", spCategoryHandler.GetCategory)
			spCategoryGroup.GET("", spCategoryHandler.GetSubCategories)
			spCategoryGroup.PATCH("/:id/state", spCategoryHandler.UpdateCategoryState)
			spCategoryGroup.PATCH("/:id/sort", spCategoryHandler.UpdateCategorySortNum)
		}
		// 用户认证路由（需要登录用户）
		userAuth := api.Group("")
		userAuth.Use(middleware.AuthMiddleware(rdb)) // 添加用户认证中间件
		{
			// 用户相关操作
			// 用户路由组
			userHandler := handlers.NewMpUserHandler(factory.GetMpUserService())
			userGroup := api.Group("/mp/users")
			{
				userGroup.POST("", userHandler.CreateUser)
				userGroup.PUT("/:id", userHandler.UpdateUser)
				userGroup.GET("/:id", userHandler.GetUser)
				userGroup.GET("/email", userHandler.GetUserByEmail)
				userGroup.PATCH("/:id/status", userHandler.UpdateUserStatus)
				userGroup.PATCH("/:id/password", userHandler.UpdateUserPassword)
				userGroup.PATCH("/:id/verify-email", userHandler.VerifyUserEmail)
				userGroup.PATCH("/:id/token", userHandler.UpdateUserToken)
			}
		}
		// 可选认证路由（游客可访问，但如果有有效token会设置用户信息）
		optionalAuth := api.Group("")
		optionalAuth.Use(middleware.OptionalAuthMiddleware(rdb)) // 添加可选认证中间件
		{

		}
		// 地点路由组
		placeHandler := handlers.NewCmsPlaceHandler(factory.GetCmsAssociatedPlaceService())
		placeGroup := api.Group("/cms/places")
		{
			placeGroup.POST("", placeHandler.CreatePlace)
			placeGroup.PUT("/:id", placeHandler.UpdatePlace)
			placeGroup.GET("/:id", placeHandler.GetPlace)
			placeGroup.GET("", placeHandler.ListPlaces)
			placeGroup.GET("/search", placeHandler.SearchPlaces)
			placeGroup.PATCH("/:id/state", placeHandler.UpdatePlaceState)
			placeGroup.DELETE("/:id", placeHandler.DeletePlace)
		}

		// 分类路由组
		categoryHandler := handlers.NewCmsCategoryHandler(factory.GetCmsCategoryService())
		categoryGroup := api.Group("/cms/categories")
		{
			categoryGroup.POST("", categoryHandler.CreateCategory)
			categoryGroup.PUT("/:id", categoryHandler.UpdateCategory)
			categoryGroup.GET("/:id", categoryHandler.GetCategory)
			categoryGroup.GET("/sub", categoryHandler.GetSubCategories)
			categoryGroup.GET("", categoryHandler.ListCategories)
			categoryGroup.PATCH("/:id/sort", categoryHandler.UpdateCategorySort)
			categoryGroup.DELETE("/:id", categoryHandler.DeleteCategory)
		}

		// 评论路由组
		commentHandler := handlers.NewCmsCommentHandler(factory.GetCmsCommentService())
		commentGroup := api.Group("/cms/comments")
		{
			commentGroup.GET("/document/:document_id", commentHandler.GetByDocumentID)
			commentGroup.GET("/user/:user_id", commentHandler.GetByUserID)
			commentGroup.GET("/top", commentHandler.GetTopLevel)
			commentGroup.GET("/replies/:comment_id", commentHandler.GetReplies)
			commentGroup.GET("", commentHandler.ListComments)
		}

		// 文档存档路由组
		archiveHandler := handlers.NewCmsDocumentArchiveHandler(factory.GetCmsDocumentArchiveService())
		archiveGroup := api.Group("/cms/archives")
		{
			archiveGroup.POST("/:document_id", archiveHandler.CreateArchive)
			archiveGroup.PUT("/:document_id", archiveHandler.UpdateArchive)
			archiveGroup.GET("/:document_id", archiveHandler.GetArchive)
		}

		// 文档路由组
		documentHandler := handlers.NewCmsDocumentHandler(factory.GetCmsDocumentService())
		documentGroup := api.Group("/cms/documents")
		{
			documentGroup.GET("/category/:category_id", documentHandler.GetByCategoryID)
			documentGroup.GET("/popular", documentHandler.GetPopular)
			documentGroup.GET("", documentHandler.ListDocuments)
		}

		// 文档标签路由组
		docTagHandler := handlers.NewCmsDocumentTagHandler(factory.GetCmsDocumentTagService())
		docTagGroup := api.Group("/cms/document-tags")
		{
			docTagGroup.POST("", docTagHandler.CreateDocumentTag)
			docTagGroup.DELETE("/:document_id/:tag_id", docTagHandler.DeleteDocumentTag)
			docTagGroup.GET("/document/:document_id", docTagHandler.GetTagsByDocument)
			docTagGroup.GET("/tag/:tag_id", docTagHandler.GetDocumentsByTag)
		}

		videoHandler := handlers.NewCmsDocumentVideoHandler(factory.GetCmsDocumentVideoService())
		videoGroup := api.Group("/cms/videos")
		{
			videoGroup.POST("/:document_id", videoHandler.CreateVideo)
			videoGroup.PUT("/:document_id", videoHandler.UpdateVideo)
			videoGroup.GET("/:document_id", videoHandler.GetVideo)
		}

		// 文件路由组
		fileHandler := handlers.NewCmsFileHandler(factory.GetCmsFileService())
		fileGroup := api.Group("/cms/files")
		{
			fileGroup.POST("", fileHandler.CreateFile)
			fileGroup.PUT("/:id", fileHandler.UpdateFile)
			fileGroup.GET("/:id", fileHandler.GetFile)
			fileGroup.GET("/md5", fileHandler.GetFileByMD5)
		}

		// 推荐索引路由组
		recIndexHandler := handlers.NewCmsRecommendIndexHandler(factory.GetCmsRecommendIndexService())
		recIndexGroup := api.Group("/cms/recommend-indices")
		{
			recIndexGroup.POST("", recIndexHandler.CreateIndex)
			recIndexGroup.PUT("/:id", recIndexHandler.UpdateIndex)
			recIndexGroup.GET("/recommend/:recommend_id", recIndexHandler.GetByRecommendID)
			recIndexGroup.GET("/state", recIndexHandler.GetByState)
			recIndexGroup.DELETE("/:id", recIndexHandler.DeleteIndex)
		}

		// 推荐路由组
		recommendHandler := handlers.NewCmsRecommendHandler(factory.GetCmsRecommendService())
		recommendGroup := api.Group("/cms/recommends")
		{
			recommendGroup.POST("", recommendHandler.CreateRecommend)
			recommendGroup.PUT("/:id", recommendHandler.UpdateRecommend)
			recommendGroup.GET("/active", recommendHandler.GetActiveRecommends)
			recommendGroup.GET("", recommendHandler.GetRecommendsByState)
		}

		// 景点路由组
		spotHandler := handlers.NewCmsScenicSpotHandler(factory.GetCmsScenicSpotService())
		spotGroup := api.Group("/cms/scenic-spots")
		{
			spotGroup.POST("", spotHandler.CreateSpot)
			spotGroup.PUT("/:id", spotHandler.UpdateSpot)
			spotGroup.GET("/:id", spotHandler.GetSpot)
			spotGroup.GET("/place/:place_id", spotHandler.GetByPlace)
			spotGroup.PATCH("/:id/read", spotHandler.IncrementReadNum)
			spotGroup.GET("", spotHandler.ListSpots)
		}

		// 标签路由组
		cmsTagHandler := handlers.NewCmsTagHandler(factory.GetCmsTagService())
		tagGroup := api.Group("/cms/tags")
		{
			tagGroup.POST("", cmsTagHandler.CreateTag)
			tagGroup.PUT("/:id", cmsTagHandler.UpdateTag)
			tagGroup.GET("/:id", cmsTagHandler.GetTag)
			tagGroup.GET("", cmsTagHandler.GetByState)
			tagGroup.PATCH("/:id/read", cmsTagHandler.IncrementReadNum)
			tagGroup.GET("/search", cmsTagHandler.SearchTags)
		}

		// 用户点赞路由组
		likeHistoryHandler := handlers.NewCmsUserLikeHistoryHandler(factory.GetCmsUserLikeHistoryService())
		likeHistoryGroup := api.Group("/cms/user-like-histories")
		{
			likeHistoryGroup.POST("", likeHistoryHandler.CreateLikeHistory)
			likeHistoryGroup.PUT("/:id", likeHistoryHandler.UpdateLikeHistory)
			likeHistoryGroup.GET("/user/:user_id", likeHistoryHandler.GetLikeHistoryByUser)
			likeHistoryGroup.GET("/check", likeHistoryHandler.CheckUserLiked)
			likeHistoryGroup.GET("/count/document/:document_id", likeHistoryHandler.GetLikeCount)
		}

		// 管理员角色路由组
		adminRoleHandler := handlers.NewCoreAdminRoleIndexHandler(factory.GetCoreAdminRoleIndexService())
		adminRoleGroup := api.Group("/core/admin-roles")
		{
			adminRoleGroup.POST("", adminRoleHandler.CreateAdminRole)
			adminRoleGroup.DELETE("", adminRoleHandler.DeleteAdminRole)
			adminRoleGroup.GET("/admin/:admin_id", adminRoleHandler.GetAdminRoles)
			adminRoleGroup.DELETE("/admin/:admin_id/all", adminRoleHandler.DeleteAllAdminRoles)
		}

		// 系统配置路由组
		configHandler := handlers.NewCoreConfigHandler(factory.GetCoreConfigService())
		configGroup := api.Group("/core/configs")
		{
			configGroup.POST("", configHandler.CreateConfig)
			configGroup.PUT("/:id", configHandler.UpdateConfig)
			configGroup.GET("", configHandler.GetConfigByKey)
			configGroup.GET("/all", configHandler.GetAllConfigs)
			configGroup.PUT("/batch", configHandler.BatchUpdateConfigs)
		}

		// 部门路由组
		deptHandler := handlers.NewCoreDeptHandler(factory.GetCoreDeptService())
		deptGroup := api.Group("/core/depts")
		{
			deptGroup.POST("", deptHandler.CreateDept)
			deptGroup.PUT("/:id", deptHandler.UpdateDept)
			deptGroup.GET("/:id", deptHandler.GetDept)
			deptGroup.GET("", deptHandler.GetSubDepts)
			deptGroup.DELETE("/:id", deptHandler.DeleteDept)
		}

		// 权限路由组
		permissionHandler := handlers.NewCorePermissionHandler(factory.GetCorePermissionService())
		permissionGroup := api.Group("/core/permissions")
		{
			permissionGroup.POST("", permissionHandler.CreatePermission)
			permissionGroup.PUT("/:id", permissionHandler.UpdatePermission)
			permissionGroup.GET("/:id", permissionHandler.GetPermission)
			permissionGroup.GET("/code", permissionHandler.GetPermissionByCode)
		}

		// 请求日志路由组
		requestLogHandler := handlers.NewCoreRequestLogHandler(factory.GetCoreRequestLogService())
		logGroup := api.Group("/core/logs")
		{
			logGroup.GET("", requestLogHandler.ListRequestLogs)
			logGroup.GET("/ip/:ip", requestLogHandler.GetLogsByIP)
			logGroup.DELETE("/cleanup", requestLogHandler.CleanupOldLogs)
		}

		// 角色路由组
		roleHandler := handlers.NewCoreRoleHandler(factory.GetCoreRoleService())
		roleGroup := api.Group("/core/roles")
		{
			roleGroup.POST("", roleHandler.CreateRole)
			roleGroup.PUT("/:id", roleHandler.UpdateRole)
			roleGroup.GET("/:id", roleHandler.GetRole)
			roleGroup.GET("", roleHandler.GetAllRoles)
			roleGroup.PATCH("/:id/status", roleHandler.UpdateRoleStatus)
			roleGroup.PATCH("/:id/permissions", roleHandler.UpdateRolePermissions)
		}

		// 订单路由组
		orderHandler := handlers.NewMpOrderHandler(factory.GetMpOrderService())
		orderGroup := api.Group("/mp/orders")
		{
			orderGroup.POST("", orderHandler.CreateOrder)
			orderGroup.PUT("/:id", orderHandler.UpdateOrder)
			orderGroup.GET("/:id", orderHandler.GetOrder)
			orderGroup.GET("/user/:user_id", orderHandler.GetOrdersByUser)
			orderGroup.GET("", orderHandler.GetOrdersByState)
			orderGroup.PATCH("/:id/state", orderHandler.UpdateOrderState)
			orderGroup.GET("/third/:third_id", orderHandler.GetOrderByThirdID)
		}

		// 支付配置路由组
		payConfigHandler := handlers.NewMpPayConfigHandler(factory.GetMpPayConfigService())
		payConfigGroup := api.Group("/mp/pay-configs")
		{
			payConfigGroup.POST("", payConfigHandler.CreatePayConfig)
			payConfigGroup.PUT("/:id", payConfigHandler.UpdatePayConfig)
			payConfigGroup.GET("/:id", payConfigHandler.GetPayConfig)
			payConfigGroup.GET("/active", payConfigHandler.GetActivePayConfigs)
			payConfigGroup.GET("/code", payConfigHandler.GetPayConfigByCode)
			payConfigGroup.PATCH("/:id/state", payConfigHandler.UpdatePayConfigState)
		}

		// 产品路由组
		productHandler := handlers.NewMpProductHandler(factory.GetMpProductService())
		productGroup := api.Group("/mp/products")
		{
			productGroup.POST("", productHandler.CreateProduct)
			productGroup.PUT("/:id", productHandler.UpdateProduct)
			productGroup.GET("/:id", productHandler.GetProduct)
			productGroup.GET("/type", productHandler.GetProductsByType)
			productGroup.GET("/terminal", productHandler.GetProductsByTerminal)
			productGroup.GET("/code", productHandler.GetProductByCode)
			productGroup.PATCH("/:id/state", productHandler.UpdateProductState)
		}

		// 密码重置令牌路由组
		resetTokenHandler := handlers.NewMpResetPwdTokensHandler(factory.GetMpResetPwdTokensService())
		resetTokenGroup := api.Group("/mp/reset-tokens")
		{
			resetTokenGroup.POST("", resetTokenHandler.CreateResetToken)
			resetTokenGroup.GET("/:token", resetTokenHandler.GetTokenRecord)
			resetTokenGroup.GET("/email/:email", resetTokenHandler.GetTokenByEmail)
			resetTokenGroup.PATCH("/increment/:email", resetTokenHandler.IncrementTokenCount)
			resetTokenGroup.DELETE("/cleanup", resetTokenHandler.DeleteExpiredTokens)
			resetTokenGroup.DELETE("/email/:email", resetTokenHandler.DeleteTokenByEmail)
		}

		// 用户令牌路由组
		userTokenHandler := handlers.NewMpUserTokenHandler(factory.GetMpUserTokenService())
		userTokenGroup := api.Group("/mp/user-tokens")
		{
			userTokenGroup.POST("", userTokenHandler.CreateUserToken)
			userTokenGroup.GET("/:token", userTokenHandler.GetToken)
			userTokenGroup.GET("/user/:user_id", userTokenHandler.GetUserTokens)
			userTokenGroup.DELETE("/:id", userTokenHandler.DeleteToken)
			userTokenGroup.DELETE("/user/:user_id/all", userTokenHandler.DeleteUserTokens)
			userTokenGroup.DELETE("/cleanup", userTokenHandler.CleanupExpiredTokens)
		}

		// PayPal订单日志路由组
		paypalOrderLogsHandler := handlers.NewPaypalOrderLogsHandler(factory.GetPaypalOrderLogsService())
		paypalOrderLogsGroup := api.Group("/paypal/order-logs")
		{
			paypalOrderLogsGroup.POST("", paypalOrderLogsHandler.CreateOrderLog)
			paypalOrderLogsGroup.GET("/local/:local_order_id", paypalOrderLogsHandler.GetLogsByLocalOrder)
			paypalOrderLogsGroup.GET("/paypal/:paypal_order_id", paypalOrderLogsHandler.GetLogByPaypalOrder)
			paypalOrderLogsGroup.GET("", paypalOrderLogsHandler.GetAllOrderLogs)
		}

		// PayPal Webhook日志路由组
		paypalWebhookLogsHandler := handlers.NewPaypalWebhookLogsHandler(factory.GetPaypalWebhookLogsService())
		paypalWebhookLogsGroup := api.Group("/paypal/webhook-logs")
		{
			paypalWebhookLogsGroup.POST("", paypalWebhookLogsHandler.CreateWebhookLog)
			paypalWebhookLogsGroup.GET("/event/:event_id", paypalWebhookLogsHandler.GetLogByEventID)
			paypalWebhookLogsGroup.GET("/local/:local_order_id", paypalWebhookLogsHandler.GetLogsByLocalOrder)
			paypalWebhookLogsGroup.GET("/paypal/:paypal_order_id", paypalWebhookLogsHandler.GetLogsByPaypalOrder)
			paypalWebhookLogsGroup.GET("/event-type/:event_type", paypalWebhookLogsHandler.GetLogsByEventType)
			paypalWebhookLogsGroup.PATCH("/:id/result", paypalWebhookLogsHandler.UpdateProcessResult)
		}

		// 商品标签关联路由组
		tagIndexHandler := handlers.NewShopTagIndexHandler(factory.GetShopTagIndexService())
		tagIndexGroup := api.Group("/shop/tag-indices")
		{
			tagIndexGroup.POST("", tagIndexHandler.CreateTagIndex)
			tagIndexGroup.DELETE("", tagIndexHandler.DeleteTagIndex)
			tagIndexGroup.GET("/product/:product_id", tagIndexHandler.GetTagIndicesByProduct)
			tagIndexGroup.GET("/tag/:tag_id", tagIndexHandler.GetTagIndicesByTag)
			tagIndexGroup.PATCH("/:id/sort", tagIndexHandler.UpdateTagSortNum)
			tagIndexGroup.DELETE("/product/:product_id/all", tagIndexHandler.DeleteAllTagsByProduct)
		}

		// 标签元数据路由组
		tagMateHandler := handlers.NewShopTagMateHandler(factory.GetShopTagMateService())
		tagMateGroup := api.Group("/shop/tag-mates")
		{
			tagMateGroup.POST("", tagMateHandler.CreateTagMate)
			tagMateGroup.PUT("/:id", tagMateHandler.UpdateTagMate)
			tagMateGroup.GET("/:id", tagMateHandler.GetTagMate)
			tagMateGroup.PATCH("/:id/seo", tagMateHandler.UpdateTagSEO)
			tagMateGroup.PATCH("/:id/content", tagMateHandler.UpdateTagContent)
		}

		// 商品标签路由组
		// tagHandler := handlers.NewShopTagHandler(factory.GetShopTagService())
		// tagGroup := api.Group("/shop/tags")
		// {
		// 	tagGroup.POST("", tagHandler.CreateTag)
		// 	tagGroup.PUT("/:id", tagHandler.UpdateTag)
		// 	tagGroup.GET("/:id", tagHandler.GetTag)
		// 	tagGroup.GET("", tagHandler.GetTagsByState)
		// 	tagGroup.GET("/search", tagHandler.SearchTags)
		// 	tagGroup.PATCH("/:id/read", tagHandler.IncrementTagReadNum)
		// 	tagGroup.GET("/list", tagHandler.ListTags)
		// }

		// SP订单项路由组
		spOrderItemHandler := handlers.NewSpOrderItemHandler(factory.GetSpOrderItemService())
		spOrderItemGroup := api.Group("/sp/order-items")
		{
			spOrderItemGroup.POST("", spOrderItemHandler.CreateOrderItem)
			spOrderItemGroup.POST("/batch", spOrderItemHandler.BatchCreateOrderItems)
			spOrderItemGroup.GET("/order/:order_id", spOrderItemHandler.GetItemsByOrder)
			spOrderItemGroup.GET("/product/:product_id", spOrderItemHandler.GetItemsByProduct)
			spOrderItemGroup.GET("/sku/:sku_id", spOrderItemHandler.GetItemsBySku)
			spOrderItemGroup.GET("/product/:product_id/sales", spOrderItemHandler.CalculateProductSales)
		}

		// SP订单操作历史路由组
		spOrderHistoryHandler := handlers.NewSpOrderOperateHistoryHandler(factory.GetSpOrderOperateHistoryService())
		spOrderHistoryGroup := api.Group("/sp/order-histories")
		{
			spOrderHistoryGroup.POST("", spOrderHistoryHandler.CreateHistory)
			spOrderHistoryGroup.GET("/order/:order_id", spOrderHistoryHandler.GetHistoriesByOrder)
			spOrderHistoryGroup.GET("/user/:user", spOrderHistoryHandler.GetHistoriesByUser)
		}

		// SP订单收货地址路由组
		spAddressHandler := handlers.NewSpOrderReceiveAddressHandler(factory.GetSpOrderReceiveAddressService())
		spAddressGroup := api.Group("/sp/addresses")
		{
			spAddressGroup.POST("", spAddressHandler.CreateAddress)
			spAddressGroup.PUT("/:id", spAddressHandler.UpdateAddress)
			spAddressGroup.GET("/order/:order_id", spAddressHandler.GetAddressByOrder)
			spAddressGroup.GET("/email/:email", spAddressHandler.GetAddressesByEmail)
		}

		// SP订单退款路由组
		spRefundHandler := handlers.NewSpOrderRefundHandler(factory.GetSpOrderRefundService())
		spRefundGroup := api.Group("/sp/refunds")
		{
			spRefundGroup.POST("", spRefundHandler.CreateRefund)
			spRefundGroup.PUT("/:id", spRefundHandler.UpdateRefund)
			spRefundGroup.GET("/order/:order_id", spRefundHandler.GetRefundByOrder)
			spRefundGroup.GET("/:refund_no", spRefundHandler.GetRefundByRefundNo)
			spRefundGroup.PATCH("/:id/status", spRefundHandler.UpdateRefundStatus)
			spRefundGroup.PATCH("/:id/amount", spRefundHandler.UpdateRefundAmount)
		}

		// SP订单路由组
		spOrderHandler := handlers.NewSpOrderHandler(factory.GetSpOrderService())
		spOrderGroup := api.Group("/sp/orders")
		{
			spOrderGroup.POST("", spOrderHandler.CreateOrder)
			spOrderGroup.PUT("/:id", spOrderHandler.UpdateOrder)
			spOrderGroup.GET("/:id", spOrderHandler.GetOrder)
			spOrderGroup.GET("/code/:code", spOrderHandler.GetOrderByCode)
			spOrderGroup.GET("/user/:user_id", spOrderHandler.GetOrdersByUser)
			spOrderGroup.GET("", spOrderHandler.GetOrdersByState)
			spOrderGroup.PATCH("/:id/state", spOrderHandler.UpdateOrderState)
			spOrderGroup.PATCH("/:id/delivery", spOrderHandler.UpdateDeliveryInfo)
		}

		spAttrGroup := api.Group("/sp/attributes")
		{
			spAttrGroup.POST("", spAttrHandler.CreateAttribute)
			spAttrGroup.PUT("/:id", spAttrHandler.UpdateAttribute)
			spAttrGroup.GET("/:id", spAttrHandler.GetAttribute)
			spAttrGroup.GET("", spAttrHandler.GetAllAttributes)
			spAttrGroup.PATCH("/:id/sort", spAttrHandler.UpdateAttributeSortNum)
			spAttrGroup.DELETE("/:id", spAttrHandler.DeleteAttribute)
		}

		// SP商品属性值路由组
		spAttrValueHandler := handlers.NewSpProdAttributesValueHandler(factory.GetSpProdAttributesValueService())
		spAttrValueGroup := api.Group("/sp/attribute-values")
		{
			spAttrValueGroup.POST("", spAttrValueHandler.CreateAttributeValue)
			spAttrValueGroup.PUT("/:id", spAttrValueHandler.UpdateAttributeValue)
			spAttrValueGroup.GET("/attribute/:attr_id", spAttrValueHandler.GetValuesByAttribute)
			spAttrValueGroup.GET("/:id", spAttrValueHandler.GetValue)
			spAttrValueGroup.POST("/batch", spAttrValueHandler.BatchCreateAttributeValues)
			spAttrValueGroup.DELETE("/attribute/:attr_id/all", spAttrValueHandler.DeleteValuesByAttribute)
		}

		// SP商品内容路由组
		spContentHandler := handlers.NewSpProductContentHandler(factory.GetSpProductContentService())
		spContentGroup := api.Group("/sp/product-contents")
		{
			spContentGroup.POST("", spContentHandler.CreateContent)
			spContentGroup.PUT("/:id", spContentHandler.UpdateContent)
			spContentGroup.GET("/product/:product_id", spContentHandler.GetContentByProduct)
			spContentGroup.PATCH("/:product_id/seo", spContentHandler.UpdateSEO)
			spContentGroup.PATCH("/:product_id/content", spContentHandler.UpdateContentText)
		}

		productPropertyHandler := handlers.NewSpProductPropertyHandler(factory.GetSpProductPropertyService())
		productPropertyGroup := api.Group("/sp/product-properties")
		{
			productPropertyGroup.POST("", productPropertyHandler.CreateProperty)
			productPropertyGroup.PUT("/:id", productPropertyHandler.UpdateProperty)
			productPropertyGroup.GET("/product/:product_id", productPropertyHandler.GetPropertiesByProduct)
			productPropertyGroup.DELETE("/product/:product_id", productPropertyHandler.DeletePropertiesByProduct)
		}

		// SKU路由组
		skuHandler := handlers.NewSpSkuHandler(factory.GetSpSkuService())
		skuGroup := api.Group("/sp/skus")
		{
			skuGroup.POST("", skuHandler.CreateSku)
			skuGroup.PUT("/:id", skuHandler.UpdateSku)
			skuGroup.GET("/product/:product_id", skuHandler.GetSkusByProduct)
			skuGroup.PATCH("/:id/stock", skuHandler.UpdateSkuStock)
		}

		// 用户地址路由组
		addressHandler := handlers.NewSpUserAddressHandler(factory.GetSpUserAddressService())
		addressGroup := api.Group("/sp/user-addresses")
		{
			addressGroup.POST("", addressHandler.CreateAddress)
			addressGroup.PUT("/:id", addressHandler.UpdateAddress)
			addressGroup.GET("/user/:user_id", addressHandler.GetAddresses)
			addressGroup.PATCH("/:id/default", addressHandler.SetDefaultAddress)
		}

		// 用户购物车路由组
		cartHandler := handlers.NewSpUserCartHandler(factory.GetSpUserCartService())
		cartGroup := api.Group("/sp/user-carts")
		{
			cartGroup.POST("", cartHandler.AddToCart)
			cartGroup.PUT("/:id", cartHandler.UpdateCartItem)
			cartGroup.GET("/user/:user_id", cartHandler.GetCartItems)
			cartGroup.DELETE("/:id", cartHandler.DeleteCartItem)
			cartGroup.DELETE("/user/:user_id/clear", cartHandler.ClearCart)
		}
	}
	r.Static("/api/oss", "./oss")
	return r
}
