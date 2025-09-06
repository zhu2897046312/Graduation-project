import type { RouteRecordRaw } from 'vue-router';

const routes: Array<RouteRecordRaw> = [
  {
    path: '/core',
    name: 'Core',
    redirect: '/core/dept',
    meta: {
      icon: 'setting-outlined',
      title: '系统管理',
      permission: ['CoreDept', 'CoreRole', 'CoreAdmin', 'CorePermission', 'Core'],
    },
    children: [
      {
        path: 'admin',
        name: 'CoreAdmin',
        meta: {
          icon: 'user-outlined',
          title: '管理员',
          permission: ['CoreAdmin'],
        },
        component: () => import('/@/views/core/admin/index.vue'),
      },
      {
        path: 'dept',
        name: 'CoreDept',
        meta: {
          icon: 'cluster-outlined',
          title: '部门管理',
          permission: ['CoreDept'],
        },
        component: () => import('/@/views/core/dept/index.vue'),
      },
      {
        path: 'role',
        name: 'CoreRole',
        meta: {
          icon: 'usergroup-add-outlined',
          title: '角色管理',
          permission: ['CoreRole'],
        },
        component: () => import('/@/views/core/role/index.vue'),
      },
      {
        path: 'permission',
        name: 'CorePermission',
        meta: {
          icon: 'safety-certificate-outlined',
          title: '权限管理',
          permission: ['CorePermission'],
        },
        component: () => import('/@/views/core/permission/index.vue'),
      },
      // {
      //   path: 'dict',
      //   name: 'CoreDict',
      //   meta: {
      //     icon: 'safety-certificate-outlined',
      //     title: '字典管理',
      //     permission: ['CoreDict'],
      //   },
      //   component: () => import('/@/views/core/dict/index.vue'),
      // },
      {
        path: 'dictItem',
        name: 'CoreDictItem',
        meta: {
          icon: 'safety-certificate-outlined',
          title: '字典属性管理',
          hideMenu: true,
        },
        component: () => import('/@/views/core/dict/info.vue'),
      },
      {
        path: 'iconDemo',
        name: 'CoreIconDemo',
        meta: {
          icon: 'safety-certificate-outlined',
          title: '字体查看',
        },
        component: () => import('/@/views/core/fonticon/index.vue'),
      },
    ],
  }
];

export default routes;