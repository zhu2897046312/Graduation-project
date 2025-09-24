package router

import (
	"server/handlers/admin"
	"server/handlers/client"
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
		ossHandler := admin.NewOssHandler()
		// SP分类路由组
		spCategoryHandler := admin.NewSpCategoryHandler(factory.GetSpCategoryService())

		// SP商品属性路由组
		spAttrHandler := admin.NewSpProdAttributesHandler(
			factory.GetSpProdAttributesService(),
			factory.GetSpProdAttributesValueService(),
		)
		// 商品标签路由组
		tagHandler := admin.NewShopTagHandler(factory.GetShopTagService(), factory.GetShopTagMateService())

		// SP商品属性值路由组
		spAttrValueHandler := admin.NewSpProdAttributesValueHandler(factory.GetSpProdAttributesValueService())
		//商品路由组
		productHandler := admin.NewSpProductHandler(
			factory.GetSpProductService(),
			factory.GetSpCategoryService(),
			factory.GetSpProductContentService(),
			factory.GetSpProductPropertyService(),
			factory.GetSpSkuService(),
			factory.GetSpSkuIndexService(),
			factory.GetShopTagIndexService(),
			factory.GetShopTagService(),
			factory.GetSpProdAttributesService(),
			factory.GetSpProdAttributesValueService(),
		)

		// 系统配置路由组
		configHandler := admin.NewCoreConfigHandler(factory.GetCoreConfigService())

		// 文档路由组
		documentHandler := admin.NewCmsDocumentHandler(factory.GetCmsDocumentService(), factory.GetCmsDocumentArchiveService())

		// 推荐路由组
		recommendHandler := admin.NewCmsRecommendHandler(factory.GetCmsRecommendService(), factory.GetCmsRecommendIndexService())

		recIndexHandler := admin.NewCmsRecommendIndexHandler(factory.GetCmsRecommendIndexService())

		// SP订单路由组
		spOrderHandler := admin.NewSpOrderHandler(
			factory.GetSpOrderService(),
			factory.GetSpOrderItemService(),
			factory.GetSpOrderReceiveAddressService(),
			factory.GetSpOrderRefundService(),
			factory.GetSpOrderReceiveAddressService(),
			factory.GetSpProductService(),
			factory.GetSpUserCartService(),
		)

		// SP订单退款路由组
		spRefundHandler := admin.NewSpOrderRefundHandler(
			factory.GetSpOrderRefundService(),
			factory.GetSpOrderService(),
		)

		// 部门路由组
		deptHandler := admin.NewCoreDeptHandler(factory.GetCoreDeptService())

		// 获取当前管理员信息
		adminHandler := admin.NewCoreAdminHandler(
			factory.GetCoreAdminService(),
			factory.GetCoreDeptService(),
			factory.GetCoreRoleService(),
			factory.GetCoreAdminRoleIndexService(),
			rdb,
		)

		// 管理员角色路由组
		// adminRoleHandler := admin.NewCoreAdminRoleIndexHandler(factory.GetCoreAdminRoleIndexService())

		// 角色路由组
		roleHandler := admin.NewCoreRoleHandler(factory.GetCoreRoleService())

		// 权限路由组
		permissionHandler := admin.NewCorePermissionHandler(factory.GetCorePermissionService())

		mpUseHandler := admin.NewMpUserHandler(factory.GetMpUserService())

		// 用户路由组

		// 公开路由（不需要认证）
		public := api.Group("")
		{
			// // 管理员登录
			// adminHandler := admin.NewCoreAdminHandler(factory.GetCoreAdminService(), rdb)
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
			}

			productGroup := adminAuth.Group("/shop/product")
			{
				productGroup.POST("/create", productHandler.CreateProduct)
				productGroup.GET("/info", productHandler.GetProduct)
				productGroup.POST("/list", productHandler.ListProducts)
				productGroup.POST("/modify", productHandler.UpdateProduct)
				productGroup.GET("/del", productHandler.SoftDeleteProduct)
			}
			spCategoryGroup := adminAuth.Group("/shop/category")
			{
				spCategoryGroup.GET("/tree", spCategoryHandler.Tree)
				spCategoryGroup.POST("/create", spCategoryHandler.CreateCategory)
				spCategoryGroup.GET("/info", spCategoryHandler.GetCategory)
				spCategoryGroup.POST("/modify", spCategoryHandler.UpdateCategory)
			}
			spAttrGroup := adminAuth.Group("/shop/prodAttributes")
			{
				spAttrGroup.POST("/list", spAttrHandler.List)
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
				spOrderGroup.GET("/info", spOrderHandler.GetOrder)
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
				spRefundGroup.GET("/info", spRefundHandler.GetRefundByOrder)
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

			mpUserGroup := adminAuth.Group("/mp/user")
			{
				mpUserGroup.POST("/list", mpUseHandler.List)
				mpUserGroup.GET("/delete", mpUseHandler.DeleteMpUser)
			}
		}

		clientSpUserAddressHandler := client.NewSpUserAddressHandler(factory.GetSpUserAddressService())
		clientConfigHandler := client.NewClientCoreConfigHandler(factory.GetCoreConfigService())
		clientSpCategoryHandler := client.NewClientSpCategoryHandler(factory.GetSpCategoryService())
		clientTagHandler := client.NewClientShopTagHandler(
			factory.GetShopTagService(),
			factory.GetShopTagMateService(),
		)
		clientDocumentHandler := client.NewClientCmsDocumentHandler(
			factory.GetCmsDocumentService(),
			factory.GetCmsDocumentArchiveService(),
		)
		clientMpUserHandler := client.NewClientMpUserHandler(
			factory.GetMpUserService(),
			factory.GetMpUserTokenService(),
		)

		clientPayHandler := client.NewClientPaymentHandler(
			factory.GetSpOrderService(),
			factory.GetPaypalOrderLogsService(),
		)
		clientMarketSettingHandler := client.NewClientSpMarketSettingHandler(
			factory.GetSpCategoryService(),
			factory.GetSpProductService(),
		)
		clientRecommendHandler := client.NewClientCmsRecommendHandler(
			factory.GetCmsRecommendService(),
			factory.GetCmsRecommendIndexService(),
		)
		clientCartHandler := client.NewClientSpUserCartHandler(
			factory.GetSpUserCartService(),
			factory.GetSpProductService(),
			factory.GetSpSkuService(),
		)
		clientSpOrderHandler := client.NewClientSpOrderHandler(
			factory.GetSpOrderService(),
			factory.GetSpOrderItemService(),
			factory.GetSpOrderReceiveAddressService(),
			factory.GetSpOrderRefundService(),
			factory.GetSpOrderReceiveAddressService(),
			factory.GetSpProductService(),
			factory.GetSpUserCartService(),
		)
		// 推荐索引路由组
		clientProductHandler := client.NewClientSpProductHandler(
			factory.GetSpProductService(),
			factory.GetSpCategoryService(),
			factory.GetSpProductContentService(),
			factory.GetSpProductPropertyService(),
			factory.GetSpSkuService(),
			factory.GetSpSkuIndexService(),
			factory.GetShopTagIndexService(),
			factory.GetShopTagService(),
			factory.GetSpProdAttributesService(),
			factory.GetSpProdAttributesValueService(),
		)
		clientAuth := api.Group("/client")
		clientAuth.Use(middleware.DeviceFingerprintMiddleware())
		clientAuth.Use(middleware.OptionalClientAuthMiddleware())
		{
			shopGroup := clientAuth.Group("/shop")
			{
				documentGroup := shopGroup.Group("/document")
				{
					documentGroup.GET("/list", clientDocumentHandler.GetAll)
					documentGroup.GET("/info", clientDocumentHandler.GetDocumentByCode)
				}

				productGroup := shopGroup.Group("/product")
				{
					productGroup.POST("/list", clientProductHandler.ListProducts)
					productGroup.GET("/info", clientProductHandler.GetClientProduct)
				}

				categoryGroup := shopGroup.Group("/category")
				{
					categoryGroup.GET("/tree", clientSpCategoryHandler.GetCategoryTree)
					categoryGroup.GET("/info", clientSpCategoryHandler.GetCategory)
					categoryGroup.GET("/getInfoByCode", clientSpCategoryHandler.GetCategoryByCode)
					categoryGroup.GET("/getParents", clientSpCategoryHandler.GetCategoryParents)
				}

				marketGroup := shopGroup.Group("/market")
				{
					marketGroup.GET("/siteInfo", clientConfigHandler.GetMarketInfo)
					marketGroup.POST("/breadcrumb", clientMarketSettingHandler.GetBreadcrumb)
					marketGroup.GET("/freight", clientMarketSettingHandler.GetFreight)
				}

				tagGroup := shopGroup.Group("/tag")
				{
					tagGroup.GET("/info", clientTagHandler.GetTagByCode)
					tagGroup.POST("/list", clientTagHandler.ListTags)
				}

				recommendIndexGrop := shopGroup.Group("/recommendIndex")
				{
					recommendIndexGrop.GET("/list", clientRecommendHandler.ListRecommendsIndex)
				}

				cartGroup := shopGroup.Group("/userCart")
				{
					cartGroup.POST("/list", clientCartHandler.List)
					cartGroup.POST("/act", clientCartHandler.CarAction)
				}

				mpUserGroup := shopGroup.Group("/userAuth")
				{
					mpUserGroup.POST("/login", clientMpUserHandler.Login)
					mpUserGroup.POST("/register", clientMpUserHandler.Register)
				}

				orderGrop := shopGroup.Group("/order")
				{
					orderGrop.POST("/create", clientSpOrderHandler.CreateOrder)
					orderGrop.GET("/query-code", clientSpOrderHandler.GetOrderByQueryCode)
					orderGrop.POST("/list", clientSpOrderHandler.ListOrders)
				}

				paymentGroup := clientAuth.Group("/payment")
				{
					paymentGroup.POST("/paypal/create-order", clientPayHandler.CreatePayment)
					paymentGroup.GET("/capture-order", clientPayHandler.CapturePayment)
					paymentGroup.POST("/webhook", clientPayHandler.PaymentWebhook)
					paymentGroup.GET("/status/:id", clientPayHandler.GetPaymentStatus)
				}

				SpUerAddressGroup := clientAuth.Group("/shop/userAddress")
				{
					SpUerAddressGroup.POST("/list", clientSpUserAddressHandler.ListAddress)
					SpUerAddressGroup.POST("/create", clientSpUserAddressHandler.CreateAddress)
					SpUerAddressGroup.POST("/modify", clientSpUserAddressHandler.UpdateAddress)
					SpUerAddressGroup.GET("/info", clientSpUserAddressHandler.GetAddress)
					SpUerAddressGroup.GET("/del", clientSpUserAddressHandler.DeleteAddress)
				}
			}
		}

	}
	return r
}
