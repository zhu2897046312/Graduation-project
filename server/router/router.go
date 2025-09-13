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
	r.Static("/api/oss", "./oss")
	api := r.Group("/api")
	{
		ossHandler := handlers.NewOssHandler()
		// SP分类路由组
		spCategoryHandler := handlers.NewSpCategoryHandler(factory.GetSpCategoryService())
		// SP商品属性路由组
		spAttrHandler := handlers.NewSpProdAttributesHandler(
			factory.GetSpProdAttributesService(),
			factory.GetSpProdAttributesValueService(),
		)
		// 商品标签路由组
		tagHandler := handlers.NewShopTagHandler(factory.GetShopTagService(), factory.GetShopTagMateService())
		// SP商品属性值路由组
		spAttrValueHandler := handlers.NewSpProdAttributesValueHandler(factory.GetSpProdAttributesValueService())
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
			factory.GetSpProdAttributesValueService(),
		)
		// 系统配置路由组
		configHandler := handlers.NewCoreConfigHandler(factory.GetCoreConfigService())

		// 文档路由组
		documentHandler := handlers.NewCmsDocumentHandler(factory.GetCmsDocumentService(), factory.GetCmsDocumentArchiveService())

		// 推荐路由组
		recommendHandler := handlers.NewCmsRecommendHandler(factory.GetCmsRecommendService(), factory.GetCmsRecommendIndexService())

		// 推荐索引路由组
		recIndexHandler := handlers.NewCmsRecommendIndexHandler(factory.GetCmsRecommendIndexService())

		// SP订单路由组
		spOrderHandler := handlers.NewSpOrderHandler(
			factory.GetSpOrderService(),
			factory.GetSpOrderItemService(),
			factory.GetSpOrderReceiveAddressService(),
			factory.GetSpOrderRefundService(),
			factory.GetSpOrderReceiveAddressService(),
			factory.GetSpProductService(),
			factory.GetSpUserCartService(),
		)

		// SP订单退款路由组
		spRefundHandler := handlers.NewSpOrderRefundHandler(
			factory.GetSpOrderRefundService(),
			factory.GetSpOrderService(),
		)

		// 部门路由组
		deptHandler := handlers.NewCoreDeptHandler(factory.GetCoreDeptService())

		// 获取当前管理员信息
		adminHandler := handlers.NewCoreAdminHandler(
			factory.GetCoreAdminService(),
			factory.GetCoreDeptService(),
			factory.GetCoreRoleService(),
			factory.GetCoreAdminRoleIndexService(),
			rdb,
		)

		// 管理员角色路由组
		// adminRoleHandler := handlers.NewCoreAdminRoleIndexHandler(factory.GetCoreAdminRoleIndexService())

		// 角色路由组
		roleHandler := handlers.NewCoreRoleHandler(factory.GetCoreRoleService())

		// 权限路由组
		permissionHandler := handlers.NewCorePermissionHandler(factory.GetCorePermissionService())

		marketSettingHandler := handlers.NewSpMarketSettingHandler(
			factory.GetSpCategoryService(),
			factory.GetSpProductService(),
		)

		// 用户购物车路由组
		cartHandler := handlers.NewSpUserCartHandler(
			factory.GetSpUserCartService(),
			factory.GetSpProductService(),
			factory.GetSpSkuService(),
		)

		// 用户路由组
		mpUserHandler := handlers.NewMpUserHandler(factory.GetMpUserService(), factory.GetMpUserTokenService())

		payHandler := handlers.NewPaymentHandler(factory.GetSpOrderService())

		// 公开路由（不需要认证）
		public := api.Group("")
		{
			// // 管理员登录
			// adminHandler := handlers.NewCoreAdminHandler(factory.GetCoreAdminService(), rdb)
			public.POST("/manage/core/auth/login", adminHandler.LoginAdmin)
		}

		// 管理员认证路由
		adminAuth := api.Group("/manage")
		adminAuth.Use(middleware.AuthMiddleware(rdb)) // 添加管理员认证中间件
		{

			adminAuth.GET("/core/auth/info", adminHandler.GetAdminInfo)
			adminAuth.GET("/core/auth/enumDict", adminHandler.GetEnumDict)
			adminOssGroup := adminAuth.Group("/core/oss")
			{
				adminOssGroup.POST("/uploadFile", ossHandler.UploadFile)
				adminOssGroup.POST("/uploadFiles", ossHandler.UploadMultipleFiles)
				adminOssGroup.DELETE("/deleteFile", ossHandler.DeleteFile)
				adminOssGroup.GET("/fileInfo", ossHandler.GetFileInfo)
			}
			adminGroup := adminAuth.Group("/core/admin")
			{
				adminGroup.POST("/list", adminHandler.List)
				adminGroup.POST("/create", adminHandler.CreateAdmin)
				adminGroup.GET("/info", adminHandler.GetAdmin)
				adminGroup.POST("/update", adminHandler.UpdateAdmin)
				adminGroup.PATCH("/:id/password", adminHandler.UpdateAdminPassword)
			}

			productGroup := adminAuth.Group("/shop/product")
			{
				productGroup.POST("/create", productHandler.CreateProduct)
				productGroup.PUT("/:id", productHandler.UpdateProduct)
				productGroup.GET("/info", productHandler.GetProduct)
				productGroup.POST("/list", productHandler.ListProducts)
				productGroup.POST("/modify", productHandler.UpdateProduct)
				productGroup.GET("/del", productHandler.SoftDeleteProduct)
			}
			spCategoryGroup := adminAuth.Group("/shop/category")
			{
				spCategoryGroup.GET("/tree", spCategoryHandler.GetCategoryTree)
				spCategoryGroup.POST("/create", spCategoryHandler.CreateCategory)
				spCategoryGroup.GET("/info", spCategoryHandler.GetCategory)
				spCategoryGroup.POST("/modify", spCategoryHandler.UpdateCategory)
			}
			spAttrGroup := adminAuth.Group("/shop/prodAttributes")
			{
				spAttrGroup.POST("/list", spAttrHandler.GetAttributesByPage)
				spAttrGroup.POST("/create", spAttrHandler.CreateAttribute)
				spAttrGroup.GET("/info", spAttrHandler.GetAllAttributes)
				spAttrGroup.POST("/modify", spAttrHandler.UpdateAttribute)
				spAttrGroup.GET("/del", spAttrHandler.DeleteAttribute)
			}
			spAttrValueGroup := adminAuth.Group("/shop/prodAttributesValue")
			{
				spAttrValueGroup.POST("/list", spAttrValueHandler.List)
				spAttrValueGroup.POST("/create", spAttrValueHandler.CreateAttributeValue)
			}
			tagGroup := adminAuth.Group("/shop/tag")
			{
				tagGroup.POST("/list", tagHandler.ListTags)
				tagGroup.GET("/info", tagHandler.GetTag)
				tagGroup.POST("/create", tagHandler.CreateTag)
				tagGroup.POST("/modify", tagHandler.UpdateTag)
				tagGroup.GET("/delete", tagHandler.DeleteTag)
			}
			configGroup := adminAuth.Group("/shop/marketSetting")
			{
				configGroup.POST("/siteInfo", configHandler.GetSiteInfo)
				configGroup.POST("/saveSiteInfo", configHandler.SaveSiteInfo)
				configGroup.POST("/saveMarketSetting", configHandler.SaveMarketSetting)
				configGroup.POST("/info", configHandler.GetMarketInfo)
			}
			documentGroup := adminAuth.Group("/shop/document")
			{
				documentGroup.POST("/create", documentHandler.SaveDocument)
				documentGroup.POST("/update", documentHandler.SaveDocument)
				documentGroup.POST("/list", documentHandler.ListDocuments)
				documentGroup.GET("/delete", documentHandler.DeleteDocument)
			}

			recommendGroup := adminAuth.Group("/shop/recommend")
			{
				recommendGroup.POST("/list", recommendHandler.ListRecommends)
				recommendGroup.POST("/modify", recommendHandler.UpdateRecommend)
				recommendGroup.POST("/create", recommendHandler.CreateRecommend)
				recommendGroup.GET("/delete", recommendHandler.DeleteRecommendByID)
				recommendGroup.GET("/info", recommendHandler.GetRecommendByID)
			}

			recIndexGroup := adminAuth.Group("/shop/recommendIndex")
			{
				recIndexGroup.POST("/list", recIndexHandler.ListRecommendsIndex)
				recIndexGroup.POST("/modify", recIndexHandler.UpdateIndex)
				recIndexGroup.POST("/create", recIndexHandler.CreateIndex)
				recIndexGroup.GET("/info", recIndexHandler.GetRecommendIndexByID)
				recIndexGroup.GET("/delete", recIndexHandler.DeleteRecommendIndexByID)
			}

			spOrderGroup := adminAuth.Group("/shop/order")
			{
				spOrderGroup.POST("", spOrderHandler.CreateOrder)
				spOrderGroup.PUT("/:id", spOrderHandler.UpdateOrder)
				spOrderGroup.GET("/info", spOrderHandler.GetOrder)
				spOrderGroup.GET("/code/:code", spOrderHandler.GetOrderByCode)
				spOrderGroup.POST("/updateState", spOrderHandler.UpdateOrderState)
				spOrderGroup.GET("/infoByCode", spRefundHandler.GetOrderInfoByOrderCode)
				spOrderGroup.POST("/delivery", spOrderHandler.UpdateDeliveryInfo)
				spOrderGroup.POST("/list", spOrderHandler.ListOrders)
			}

			refundGroup := adminAuth.Group("/payment/paypal")
			{
				refundGroup.POST("/refund", spOrderHandler.OrderRefund)

			}

			spRefundGroup := adminAuth.Group("/shop/refund")
			{
				spRefundGroup.POST("/list", spRefundHandler.ListSpOrderRefund)
				spRefundGroup.PUT("/:id", spRefundHandler.UpdateRefund)
				spRefundGroup.GET("/info", spRefundHandler.GetRefundByOrder)
				spRefundGroup.GET("/:refund_no", spRefundHandler.GetRefundByRefundNo)
				spRefundGroup.PATCH("/:id/status", spRefundHandler.UpdateRefundStatus)
				// spRefundGroup.PATCH("/:id/amount", spRefundHandler.UpdateRefundAmount)
			}

			deptGroup := adminAuth.Group("/core/dept")
			{
				deptGroup.GET("/tree", deptHandler.Tree)
				deptGroup.POST("/update", deptHandler.UpdateDept)
				deptGroup.GET("/info", deptHandler.GetDept)
				deptGroup.POST("/create", deptHandler.CreateDept)
				deptGroup.GET("/del", deptHandler.DeleteDept)
			}

			roleGroup := adminAuth.Group("/core/role")
			{
				roleGroup.POST("/list", roleHandler.List)
				roleGroup.GET("/delete", roleHandler.DeleteRole)
				roleGroup.GET("/info", roleHandler.GetRole)
				roleGroup.GET("", roleHandler.GetAllRoles)
				roleGroup.POST("/create", roleHandler.CreateRole)
				roleGroup.POST("/update", roleHandler.UpdateRole)
			}

			permissionGroup := adminAuth.Group("/core/permission")
			{
				permissionGroup.POST("/create", permissionHandler.CreatePermission)
				permissionGroup.POST("/update", permissionHandler.UpdatePermission)
				permissionGroup.GET("/list", permissionHandler.List)
				permissionGroup.GET("/info", permissionHandler.GetPermission)
				permissionGroup.GET("/topList", permissionHandler.List)
			}
		}

		clientAuth := api.Group("/client")
		clientAuth.Use(middleware.DeviceFingerprintMiddleware())
		clientAuth.Use(middleware.OptionalClientAuthMiddleware())
		{
			shopGroup := clientAuth.Group("/shop")
			{
				documentGroup := shopGroup.Group("/document")
				{
					documentGroup.GET("/list", documentHandler.ListDocuments)
					documentGroup.GET("/info", documentHandler.GetDocumentByCode)
				}

				productGroup := shopGroup.Group("/product")
				{
					productGroup.POST("/list", productHandler.ListProducts)
					productGroup.GET("/info", productHandler.GetProductFrontInfo)
				}

				categoryGroup := shopGroup.Group("/category")
				{
					categoryGroup.GET("/tree", spCategoryHandler.GetCategoryTree)
					categoryGroup.GET("/info", spCategoryHandler.GetCategory)
					categoryGroup.GET("/getInfoByCode", spCategoryHandler.GetCategoryByCode)
					categoryGroup.GET("/getParents", spCategoryHandler.GetCategoryParents)
				}

				marketGroup := shopGroup.Group("/market")
				{
					marketGroup.GET("/siteInfo", configHandler.GetMarketInfo)
					marketGroup.POST("/breadcrumb", marketSettingHandler.GetBreadcrumb)
				}

				tagGroup := shopGroup.Group("/tag")
				{
					tagGroup.GET("/info", tagHandler.GetTagByCode)
					tagGroup.POST("/list", tagHandler.ListTags)
				}

				recommendIndexGrop := shopGroup.Group("/recommendIndex")
				{
					recommendIndexGrop.GET("/list", recommendHandler.ListRecommendsIndex)
				}

				cartGroup := shopGroup.Group("/userCart")
				{
					cartGroup.POST("/list", cartHandler.List)
					cartGroup.POST("/act", cartHandler.CarAction)
				}

				mpUserGroup := shopGroup.Group("/userAuth")
				{
					mpUserGroup.POST("/login", mpUserHandler.Login)
					mpUserGroup.POST("/register", mpUserHandler.Register)
				}

				marketClientGroup := shopGroup.Group("/market")
				{
					marketClientGroup.GET("/freight", marketSettingHandler.GetFreight)
				}

				orderGrop := shopGroup.Group("/order")
				{
					orderGrop.POST("/create", spOrderHandler.CreateOrder)
					orderGrop.GET("/query-code", spOrderHandler.GetOrderByQueryCode)
				}

				paymentGroup := clientAuth.Group("/payment")
				// 模拟支付接口（开发环境使用）
				paymentGroup.POST("/paypal/create-order", payHandler.SimulatePayment)
				paymentGroup.GET("/simulate", payHandler.SimulatePayment) // 支持GET请求

				// 支付回调页面
				paymentGroup.GET("/callback", payHandler.PaymentCallback)

				// 支付状态查询
				paymentGroup.GET("/status/:id", payHandler.GetPaymentStatus)

			}
		}

		// spCategoryGroup := api.Group("/sp/categories")
		// {
		// 	spCategoryGroup.PUT("/:id", spCategoryHandler.UpdateCategory)
		// 	spCategoryGroup.GET("/:id", spCategoryHandler.GetCategory)
		// 	spCategoryGroup.GET("", spCategoryHandler.GetSubCategories)
		// 	spCategoryGroup.PATCH("/:id/state", spCategoryHandler.UpdateCategoryState)
		// 	spCategoryGroup.PATCH("/:id/sort", spCategoryHandler.UpdateCategorySortNum)
		// }
		// // 用户认证路由（需要登录用户）
		// userAuth := api.Group("")
		// userAuth.Use(middleware.AuthMiddleware(rdb)) // 添加用户认证中间件
		// {

		// }
		// // 可选认证路由（游客可访问，但如果有有效token会设置用户信息）
		// optionalAuth := api.Group("")
		// optionalAuth.Use(middleware.OptionalAuthMiddleware(rdb)) // 添加可选认证中间件
		// {

		// }
		// // 地点路由组
		// placeHandler := handlers.NewCmsPlaceHandler(factory.GetCmsAssociatedPlaceService())
		// placeGroup := api.Group("/cms/places")
		// {
		// 	placeGroup.POST("", placeHandler.CreatePlace)
		// 	placeGroup.PUT("/:id", placeHandler.UpdatePlace)
		// 	placeGroup.GET("/:id", placeHandler.GetPlace)
		// 	placeGroup.GET("", placeHandler.ListPlaces)
		// 	placeGroup.GET("/search", placeHandler.SearchPlaces)
		// 	placeGroup.PATCH("/:id/state", placeHandler.UpdatePlaceState)
		// 	placeGroup.DELETE("/:id", placeHandler.DeletePlace)
		// }

		// // 分类路由组
		// categoryHandler := handlers.NewCmsCategoryHandler(factory.GetCmsCategoryService())
		// categoryGroup := api.Group("/cms/categories")
		// {
		// 	categoryGroup.POST("", categoryHandler.CreateCategory)
		// 	categoryGroup.PUT("/:id", categoryHandler.UpdateCategory)
		// 	categoryGroup.GET("/:id", categoryHandler.GetCategory)
		// 	categoryGroup.GET("/sub", categoryHandler.GetSubCategories)
		// 	categoryGroup.GET("", categoryHandler.ListCategories)
		// 	categoryGroup.PATCH("/:id/sort", categoryHandler.UpdateCategorySort)
		// 	categoryGroup.DELETE("/:id", categoryHandler.DeleteCategory)
		// }

		// // 评论路由组
		// commentHandler := handlers.NewCmsCommentHandler(factory.GetCmsCommentService())
		// commentGroup := api.Group("/cms/comments")
		// {
		// 	commentGroup.GET("/document/:document_id", commentHandler.GetByDocumentID)
		// 	commentGroup.GET("/user/:user_id", commentHandler.GetByUserID)
		// 	commentGroup.GET("/top", commentHandler.GetTopLevel)
		// 	commentGroup.GET("/replies/:comment_id", commentHandler.GetReplies)
		// 	commentGroup.GET("", commentHandler.ListComments)
		// }

		// // 文档存档路由组
		// archiveHandler := handlers.NewCmsDocumentArchiveHandler(factory.GetCmsDocumentArchiveService())
		// archiveGroup := api.Group("/cms/archives")
		// {
		// 	archiveGroup.POST("/:document_id", archiveHandler.CreateArchive)
		// 	archiveGroup.PUT("/:document_id", archiveHandler.UpdateArchive)
		// 	archiveGroup.GET("/:document_id", archiveHandler.GetArchive)
		// }

		// // 文档标签路由组
		// docTagHandler := handlers.NewCmsDocumentTagHandler(factory.GetCmsDocumentTagService())
		// docTagGroup := api.Group("/cms/document-tags")
		// {
		// 	docTagGroup.POST("", docTagHandler.CreateDocumentTag)
		// 	docTagGroup.DELETE("/:document_id/:tag_id", docTagHandler.DeleteDocumentTag)
		// 	docTagGroup.GET("/document/:document_id", docTagHandler.GetTagsByDocument)
		// 	docTagGroup.GET("/tag/:tag_id", docTagHandler.GetDocumentsByTag)
		// }

		// videoHandler := handlers.NewCmsDocumentVideoHandler(factory.GetCmsDocumentVideoService())
		// videoGroup := api.Group("/cms/videos")
		// {
		// 	videoGroup.POST("/:document_id", videoHandler.CreateVideo)
		// 	videoGroup.PUT("/:document_id", videoHandler.UpdateVideo)
		// 	videoGroup.GET("/:document_id", videoHandler.GetVideo)
		// }

		// // 文件路由组
		// fileHandler := handlers.NewCmsFileHandler(factory.GetCmsFileService())
		// fileGroup := api.Group("/cms/files")
		// {
		// 	fileGroup.POST("", fileHandler.CreateFile)
		// 	fileGroup.PUT("/:id", fileHandler.UpdateFile)
		// 	fileGroup.GET("/:id", fileHandler.GetFile)
		// 	fileGroup.GET("/md5", fileHandler.GetFileByMD5)
		// }

		// 景点路由组
		// spotHandler := handlers.NewCmsScenicSpotHandler(factory.GetCmsScenicSpotService())
		// spotGroup := api.Group("/cms/scenic-spots")
		// {
		// 	spotGroup.POST("", spotHandler.CreateSpot)
		// 	spotGroup.PUT("/:id", spotHandler.UpdateSpot)
		// 	spotGroup.GET("/:id", spotHandler.GetSpot)
		// 	spotGroup.GET("/place/:place_id", spotHandler.GetByPlace)
		// 	spotGroup.PATCH("/:id/read", spotHandler.IncrementReadNum)
		// 	spotGroup.GET("", spotHandler.ListSpots)
		// }

		// 标签路由组
		// cmsTagHandler := handlers.NewCmsTagHandler(factory.GetCmsTagService())
		// tagGroup := api.Group("/cms/tags")
		// {
		// 	tagGroup.POST("", cmsTagHandler.CreateTag)
		// 	tagGroup.PUT("/:id", cmsTagHandler.UpdateTag)
		// 	tagGroup.GET("/:id", cmsTagHandler.GetTag)
		// 	tagGroup.GET("", cmsTagHandler.GetByState)
		// 	tagGroup.PATCH("/:id/read", cmsTagHandler.IncrementReadNum)
		// 	tagGroup.GET("/search", cmsTagHandler.SearchTags)
		// }

		// 用户点赞路由组
		// likeHistoryHandler := handlers.NewCmsUserLikeHistoryHandler(factory.GetCmsUserLikeHistoryService())
		// likeHistoryGroup := api.Group("/cms/user-like-histories")
		// {
		// 	likeHistoryGroup.POST("", likeHistoryHandler.CreateLikeHistory)
		// 	likeHistoryGroup.PUT("/:id", likeHistoryHandler.UpdateLikeHistory)
		// 	likeHistoryGroup.GET("/user/:user_id", likeHistoryHandler.GetLikeHistoryByUser)
		// 	likeHistoryGroup.GET("/check", likeHistoryHandler.CheckUserLiked)
		// 	likeHistoryGroup.GET("/count/document/:document_id", likeHistoryHandler.GetLikeCount)
		// }

		// 请求日志路由组
		// requestLogHandler := handlers.NewCoreRequestLogHandler(factory.GetCoreRequestLogService())
		// logGroup := api.Group("/core/logs")
		// {
		// 	logGroup.GET("", requestLogHandler.ListRequestLogs)
		// 	logGroup.GET("/ip/:ip", requestLogHandler.GetLogsByIP)
		// 	logGroup.DELETE("/cleanup", requestLogHandler.CleanupOldLogs)
		// }
		// adminRoleGroup := api.Group("/cores/roles")
		// {
		// 	adminRoleGroup.POST("/list", adminRoleHandler.CreateAdminRole)
		// 	adminRoleGroup.DELETE("", adminRoleHandler.DeleteAdminRole)
		// 	adminRoleGroup.GET("/admin/:admin_id", adminRoleHandler.GetAdminRoles)
		// 	adminRoleGroup.DELETE("/admin/:admin_id/all", adminRoleHandler.DeleteAllAdminRoles)
		// }

		// 订单路由组
		// orderHandler := handlers.NewMpOrderHandler(factory.GetMpOrderService())
		// orderGroup := api.Group("/mp/orders")
		// {
		// 	orderGroup.POST("", orderHandler.CreateOrder)
		// 	orderGroup.PUT("/:id", orderHandler.UpdateOrder)
		// 	orderGroup.GET("/:id", orderHandler.GetOrder)
		// 	orderGroup.GET("/user/:user_id", orderHandler.GetOrdersByUser)
		// 	orderGroup.GET("", orderHandler.GetOrdersByState)
		// 	orderGroup.PATCH("/:id/state", orderHandler.UpdateOrderState)
		// 	orderGroup.GET("/third/:third_id", orderHandler.GetOrderByThirdID)
		// }

		// 支付配置路由组
		// payConfigHandler := handlers.NewMpPayConfigHandler(factory.GetMpPayConfigService())
		// payConfigGroup := api.Group("/mp/pay-configs")
		// {
		// 	payConfigGroup.POST("", payConfigHandler.CreatePayConfig)
		// 	payConfigGroup.PUT("/:id", payConfigHandler.UpdatePayConfig)
		// 	payConfigGroup.GET("/:id", payConfigHandler.GetPayConfig)
		// 	payConfigGroup.GET("/active", payConfigHandler.GetActivePayConfigs)
		// 	payConfigGroup.GET("/code", payConfigHandler.GetPayConfigByCode)
		// 	payConfigGroup.PATCH("/:id/state", payConfigHandler.UpdatePayConfigState)
		// }

		// 产品路由组
		// mpProductHandler := handlers.NewMpProductHandler(factory.GetMpProductService())
		// mpProductGroup := api.Group("/mp/products")
		// {
		// 	mpProductGroup.POST("", mpProductHandler.CreateProduct)
		// 	mpProductGroup.PUT("/:id", mpProductHandler.UpdateProduct)
		// 	mpProductGroup.GET("/:id", mpProductHandler.GetProduct)
		// 	mpProductGroup.GET("/type", mpProductHandler.GetProductsByType)
		// 	mpProductGroup.GET("/terminal", mpProductHandler.GetProductsByTerminal)
		// 	mpProductGroup.GET("/code", mpProductHandler.GetProductByCode)
		// 	mpProductGroup.PATCH("/:id/state", mpProductHandler.UpdateProductState)
		// }

		// 密码重置令牌路由组
		// resetTokenHandler := handlers.NewMpResetPwdTokensHandler(factory.GetMpResetPwdTokensService())
		// resetTokenGroup := api.Group("/mp/reset-tokens")
		// {
		// 	resetTokenGroup.POST("", resetTokenHandler.CreateResetToken)
		// 	resetTokenGroup.GET("/:token", resetTokenHandler.GetTokenRecord)
		// 	resetTokenGroup.GET("/email/:email", resetTokenHandler.GetTokenByEmail)
		// 	resetTokenGroup.PATCH("/increment/:email", resetTokenHandler.IncrementTokenCount)
		// 	resetTokenGroup.DELETE("/cleanup", resetTokenHandler.DeleteExpiredTokens)
		// 	resetTokenGroup.DELETE("/email/:email", resetTokenHandler.DeleteTokenByEmail)
		// }

		// 用户令牌路由组
		// userTokenHandler := handlers.NewMpUserTokenHandler(factory.GetMpUserTokenService())
		// userTokenGroup := api.Group("/mp/user-tokens")
		// {
		// 	userTokenGroup.POST("", userTokenHandler.CreateUserToken)
		// 	userTokenGroup.GET("/:token", userTokenHandler.GetToken)
		// 	userTokenGroup.GET("/user/:user_id", userTokenHandler.GetUserTokens)
		// 	userTokenGroup.DELETE("/:id", userTokenHandler.DeleteToken)
		// 	userTokenGroup.DELETE("/user/:user_id/all", userTokenHandler.DeleteUserTokens)
		// 	userTokenGroup.DELETE("/cleanup", userTokenHandler.CleanupExpiredTokens)
		// }

		// PayPal订单日志路由组
		// paypalOrderLogsHandler := handlers.NewPaypalOrderLogsHandler(factory.GetPaypalOrderLogsService())
		// paypalOrderLogsGroup := api.Group("/paypal/order-logs")
		// {
		// 	paypalOrderLogsGroup.POST("", paypalOrderLogsHandler.CreateOrderLog)
		// 	paypalOrderLogsGroup.GET("/local/:local_order_id", paypalOrderLogsHandler.GetLogsByLocalOrder)
		// 	paypalOrderLogsGroup.GET("/paypal/:paypal_order_id", paypalOrderLogsHandler.GetLogByPaypalOrder)
		// 	paypalOrderLogsGroup.GET("", paypalOrderLogsHandler.GetAllOrderLogs)
		// }

		// PayPal Webhook日志路由组
		// paypalWebhookLogsHandler := handlers.NewPaypalWebhookLogsHandler(factory.GetPaypalWebhookLogsService())
		// paypalWebhookLogsGroup := api.Group("/paypal/webhook-logs")
		// {
		// 	paypalWebhookLogsGroup.POST("", paypalWebhookLogsHandler.CreateWebhookLog)
		// 	paypalWebhookLogsGroup.GET("/event/:event_id", paypalWebhookLogsHandler.GetLogByEventID)
		// 	paypalWebhookLogsGroup.GET("/local/:local_order_id", paypalWebhookLogsHandler.GetLogsByLocalOrder)
		// 	paypalWebhookLogsGroup.GET("/paypal/:paypal_order_id", paypalWebhookLogsHandler.GetLogsByPaypalOrder)
		// 	paypalWebhookLogsGroup.GET("/event-type/:event_type", paypalWebhookLogsHandler.GetLogsByEventType)
		// 	paypalWebhookLogsGroup.PATCH("/:id/result", paypalWebhookLogsHandler.UpdateProcessResult)
		// }

		// 商品标签关联路由组
		// tagIndexHandler := handlers.NewShopTagIndexHandler(factory.GetShopTagIndexService())
		// tagIndexGroup := api.Group("/shop/tag-indices")
		// {
		// 	tagIndexGroup.POST("", tagIndexHandler.CreateTagIndex)
		// 	tagIndexGroup.DELETE("", tagIndexHandler.DeleteTagIndex)
		// 	tagIndexGroup.GET("/product/:product_id", tagIndexHandler.GetTagIndicesByProduct)
		// 	tagIndexGroup.GET("/tag/:tag_id", tagIndexHandler.GetTagIndicesByTag)
		// 	tagIndexGroup.PATCH("/:id/sort", tagIndexHandler.UpdateTagSortNum)
		// 	tagIndexGroup.DELETE("/product/:product_id/all", tagIndexHandler.DeleteAllTagsByProduct)
		// }

		// 标签元数据路由组
		// tagMateHandler := handlers.NewShopTagMateHandler(factory.GetShopTagMateService())
		// tagMateGroup := api.Group("/shop/tag-mates")
		// {
		// 	tagMateGroup.POST("", tagMateHandler.CreateTagMate)
		// 	tagMateGroup.PUT("/:id", tagMateHandler.UpdateTagMate)
		// 	tagMateGroup.GET("/:id", tagMateHandler.GetTagMate)
		// 	tagMateGroup.PATCH("/:id/seo", tagMateHandler.UpdateTagSEO)
		// 	tagMateGroup.PATCH("/:id/content", tagMateHandler.UpdateTagContent)
		// }

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
		// spOrderItemHandler := handlers.NewSpOrderItemHandler(factory.GetSpOrderItemService())
		// spOrderItemGroup := api.Group("/sp/order-items")
		// {
		// 	spOrderItemGroup.POST("", spOrderItemHandler.CreateOrderItem)
		// 	spOrderItemGroup.POST("/batch", spOrderItemHandler.BatchCreateOrderItems)
		// 	spOrderItemGroup.GET("/order/:order_id", spOrderItemHandler.GetItemsByOrder)
		// 	spOrderItemGroup.GET("/product/:product_id", spOrderItemHandler.GetItemsByProduct)
		// 	spOrderItemGroup.GET("/sku/:sku_id", spOrderItemHandler.GetItemsBySku)
		// 	spOrderItemGroup.GET("/product/:product_id/sales", spOrderItemHandler.CalculateProductSales)
		// }

		// // SP订单操作历史路由组
		// spOrderHistoryHandler := handlers.NewSpOrderOperateHistoryHandler(factory.GetSpOrderOperateHistoryService())
		// spOrderHistoryGroup := api.Group("/sp/order-histories")
		// {
		// 	spOrderHistoryGroup.POST("", spOrderHistoryHandler.CreateHistory)
		// 	spOrderHistoryGroup.GET("/order/:order_id", spOrderHistoryHandler.GetHistoriesByOrder)
		// 	spOrderHistoryGroup.GET("/user/:user", spOrderHistoryHandler.GetHistoriesByUser)
		// }

		// // SP订单收货地址路由组
		// spAddressHandler := handlers.NewSpOrderReceiveAddressHandler(factory.GetSpOrderReceiveAddressService())
		// spAddressGroup := api.Group("/sp/addresses")
		// {
		// 	spAddressGroup.POST("", spAddressHandler.CreateAddress)
		// 	spAddressGroup.PUT("/:id", spAddressHandler.UpdateAddress)
		// 	spAddressGroup.GET("/order/:order_id", spAddressHandler.GetAddressByOrder)
		// 	spAddressGroup.GET("/email/:email", spAddressHandler.GetAddressesByEmail)
		// }

		// spAttrGroup := api.Group("/sp/attributes")
		// {
		// 	spAttrGroup.POST("", spAttrHandler.CreateAttribute)
		// 	spAttrGroup.PUT("/:id", spAttrHandler.UpdateAttribute)
		// 	spAttrGroup.GET("/:id", spAttrHandler.GetAttribute)
		// 	spAttrGroup.GET("", spAttrHandler.GetAllAttributes)
		// 	spAttrGroup.PATCH("/:id/sort", spAttrHandler.UpdateAttributeSortNum)
		// 	spAttrGroup.DELETE("/:id", spAttrHandler.DeleteAttribute)
		// }

		// spAttrValueGroup := api.Group("/sp/attribute-values")
		// {
		// 	spAttrValueGroup.POST("", spAttrValueHandler.CreateAttributeValue)
		// 	spAttrValueGroup.PUT("/:id", spAttrValueHandler.UpdateAttributeValue)
		// 	// spAttrValueGroup.GET("/attribute/:attr_id", spAttrValueHandler.GetValuesByAttribute)
		// 	spAttrValueGroup.GET("/:id", spAttrValueHandler.GetValue)
		// 	spAttrValueGroup.POST("/batch", spAttrValueHandler.BatchCreateAttributeValues)
		// 	spAttrValueGroup.DELETE("/attribute/:attr_id/all", spAttrValueHandler.DeleteValuesByAttribute)
		// }

		// // SP商品内容路由组
		// spContentHandler := handlers.NewSpProductContentHandler(factory.GetSpProductContentService())
		// spContentGroup := api.Group("/sp/product-contents")
		// {
		// 	spContentGroup.POST("", spContentHandler.CreateContent)
		// 	spContentGroup.PUT("/:id", spContentHandler.UpdateContent)
		// 	spContentGroup.GET("/product/:product_id", spContentHandler.GetContentByProduct)
		// 	spContentGroup.PATCH("/:product_id/seo", spContentHandler.UpdateSEO)
		// 	spContentGroup.PATCH("/:product_id/content", spContentHandler.UpdateContentText)
		// }

		// productPropertyHandler := handlers.NewSpProductPropertyHandler(factory.GetSpProductPropertyService())
		// productPropertyGroup := api.Group("/sp/product-properties")
		// {
		// 	productPropertyGroup.POST("", productPropertyHandler.CreateProperty)
		// 	productPropertyGroup.PUT("/:id", productPropertyHandler.UpdateProperty)
		// 	productPropertyGroup.GET("/product/:product_id", productPropertyHandler.GetPropertiesByProduct)
		// 	productPropertyGroup.DELETE("/product/:product_id", productPropertyHandler.DeletePropertiesByProduct)
		// }

		// // SKU路由组
		// skuHandler := handlers.NewSpSkuHandler(factory.GetSpSkuService())
		// skuGroup := api.Group("/sp/skus")
		// {
		// 	skuGroup.POST("", skuHandler.CreateSku)
		// 	skuGroup.PUT("/:id", skuHandler.UpdateSku)
		// 	skuGroup.GET("/product/:product_id", skuHandler.GetSkusByProduct)
		// 	skuGroup.PATCH("/:id/stock", skuHandler.UpdateSkuStock)
		// }

		// // 用户地址路由组
		// addressHandler := handlers.NewSpUserAddressHandler(factory.GetSpUserAddressService())
		// addressGroup := api.Group("/sp/user-addresses")
		// {
		// 	addressGroup.POST("", addressHandler.CreateAddress)
		// 	addressGroup.PUT("/:id", addressHandler.UpdateAddress)
		// 	addressGroup.GET("/user/:user_id", addressHandler.GetAddresses)
		// 	addressGroup.PATCH("/:id/default", addressHandler.SetDefaultAddress)
		// }

		// cartGroup := api.Group("/sp/user-carts")
		// {
		// 	cartGroup.POST("", cartHandler.AddToCart)
		// 	cartGroup.PUT("/:id", cartHandler.UpdateCartItem)
		// 	cartGroup.GET("/user/:user_id", cartHandler.GetCartItems)
		// 	cartGroup.DELETE("/:id", cartHandler.DeleteCartItem)
		// 	cartGroup.DELETE("/user/:user_id/clear", cartHandler.ClearCart)
		// }
	}
	
	return r
}
