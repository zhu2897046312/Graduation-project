import type { RouteRecordRaw } from 'vue-router';

const routes: Array<RouteRecordRaw> = [
  
  {
    path: '/shop/core',
    name: 'ShopCore',
    redirect: '/shop/core/category',
    meta: {
      icon: 'shop-twotone',
      title: '商品管理',
      // permission: ['CoreDept', 'CoreRole', 'CoreAdmin', 'CorePermission', 'Core'],
    },
    children: [
      {
        path: 'product',
        name: 'ShopCoreProduct',
        meta: {
          icon: 'shopping-outlined',
          title: '商品列表',
          permission: ['CoreAdmin'],
        },
        component: () => import('/@/views/shop/product/index.vue'),
      },
      {
        path: 'create-product',
        name: 'ShopCoreProductCreate',
        meta: {
          icon: 'shopping-outlined',
          title: '录入商品',
          hideMenu: true,
          permission: ['CoreAdmin'],
        },
        component: () => import('/@/views/shop/product/add.vue'),
      },
      {
        path: 'modify-product',
        name: 'ShopCoreProductModify',
        meta: {
          icon: 'shopping-outlined',
          title: '商品信息',
          permission: ['CoreAdmin'],
          hideMenu: true,
        },
        component: () => import('/@/views/shop/product/modify.vue'),
      },
      {
        path: 'category',
        name: 'ShopCoreCategory',
        meta: {
          icon: 'right-circle-outlined',
          title: '类目管理',
          permission: ['CoreAdmin'],
        },
        component: () => import('/@/views/shop/category/index.vue'),
      },
      {
        path: 'prod-attributes',
        name: 'ShopCoreProdAttributes',
        meta: {
          icon: 'right-circle-outlined',
          title: '商品属性',
          permission: ['CoreAdmin'],
        },
        component: () => import('/@/views/shop/prod_attributes/index.vue'),
      },
      {
        path: 'prod-attributes-value',
        name: 'ShopCoreProdAttributesValue',
        meta: {
          icon: 'user-outlined',
          title: '商品属性值',
          permission: ['CoreAdmin'],
          hideMenu: true,
        },
        component: () => import('/@/views/shop/prod_attributes_value/index.vue'),
      },
      // {
      //   path: 'task',
      //   name: 'ShopTask',
      //   meta: {
      //     icon: 'right-circle-outlined',
      //     title: '采集任务',
      //     permission: ['CoreAdmin'],
      //   },
      //   component: () => import('/@/views/shop/task/index.vue'),
      // },
    ],
  },
  {
    path: '/shop/order',
    name: 'ShopOrders',
    redirect: '/shop/order/list',
    meta: {
      icon: 'code-sandbox-outlined',
      title: '订单管理',
      permission: ['CoreAdmin'],

    },
    children: [
      {
        path: 'list',
        name: 'OrdersList',
        meta: {
          icon: 'book-outlined',
          title: '订单列表',
          permission: ['CoreAdmin'],
        },
        component: () => import('/@/views/shop/order/index.vue'),
      },
      {
        path: 'refund',
        name: 'RefundList',
        meta: {
          icon: 'book-outlined',
          title: '退款列表',
          permission: ['CoreAdmin'],
        },
        component: () => import('/@/views/shop/refund/index.vue'),
      },
      
    ]
  },
  {
    path: '/shop/cms',
    name: 'ShopCms',
    redirect: '/shop/cms/blogs',
    meta: {
      icon: 'code-sandbox-outlined',
      title: '内容管理',
      // permission: ['CoreDept', 'CoreRole', 'CoreAdmin', 'CorePermission', 'Core'],
    },
    children: [
      {
        path: 'blogs',
        name: 'ShopBlogs',
        meta: {
          icon: 'book-outlined',
          title: '文档管理',
          permission: ['CoreAdmin'],
        },
        component: () => import('/@/views/cms/archive/index.vue'),
      },
      {
        path: 'blogs/add',
        name: 'ShopBlogsAdd',
        meta: {
          icon: 'book-outlined',
          title: '添加文档',
          hideMenu: true,
          permission: ['CoreAdmin'],
        },
        component: () => import('/@/views/cms/archive/add.vue'),
      },
      {
        path: 'blogs/modify',
        name: 'ShopBlogsModify',
        meta: {
          icon: 'user-outlined',
          title: '编辑文档',
          hideMenu: true,
          permission: ['CoreAdmin'],
        },
        component: () => import('/@/views/cms/archive/modify.vue'),
      },
      {
        path: 'tag',
        name: 'ShopTag',
        meta: {
          icon: 'tags-outlined',
          title: '标签管理',
          permission: ['CoreAdmin'],
        },
        component: () => import('/@/views/cms/tag/index.vue'),
      },
      {
        path: 'recommend',
        name: 'ShopRecommend',
        meta: {
          icon: 'pushpin-outlined',
          title: '推荐位管理',
          permission: ['CoreAdmin'],
        },
        component: () => import('/@/views/cms/recommend/index.vue'),
      },
      {
        path: 'recommend-index',
        name: 'ShopRecommendIndex',
        meta: {
          icon: 'pushpin-outlined',
          title: '推荐位内容管理',
          hideMenu: true,
          permission: ['CoreAdmin'],
        },
        component: () => import('/@/views/cms/recommend_index/index.vue'),
      },
    ]
  },
  {
    path: '/shop/setting',
    name: 'ShopSetting',
    meta: {
      icon: 'code-sandbox-outlined',
      title: '站点管理',
      // permission: ['CoreDept', 'CoreRole', 'CoreAdmin', 'CorePermission', 'Core'],
    },
    children: [
      {
        path: 'market-setting',
        name: 'ShopSettingMarketSetting',
        meta: {
          icon: 'calendar-filled',
          title: '商店设置',
          permission: ['CoreAdmin'],
        },
        component: () => import('/@/views/shop/market/setting.vue'),
      },
      {
        path: 'site-setting',
        name: 'ShopSettingSiteSetting',
        meta: {
          icon: 'right-circle-outlined',
          title: '站点设置',
          permission: ['CoreAdmin'],
        },
        component: () => import('/@/views/shop/market/siteinfo.vue'),
      },
    ],
  }
];

export default routes;