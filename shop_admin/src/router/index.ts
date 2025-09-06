import { createRouter, createWebHashHistory } from 'vue-router';
import type { RouteRecordRaw } from 'vue-router';
import { useAuthStore } from '/@/store/authStore';
import { apiGetAuthInfo, apiGetEnumDict } from '/@/api/auth';
import { message } from 'ant-design-vue';
import enumDict from '/@/utils/enum-dict';

import core from './modules/core';
// import cms from './modules/cms';
import shop from './modules/shop';

const routes: Array<RouteRecordRaw> = [
    {
        path: '/',
        name: 'Base',
        component: () => import('../components/Kernel/Layout/Base.vue'),
        // redirect:'/dashboard',
        children: [
            ...shop,
            ...core,
        ],
    },
    {
        path: '/login',
        name: 'Login',
        component: () => import('../views/login/index.vue'),
    },
];

const router = createRouter({
    history: createWebHashHistory(),
    routes,
})

router.beforeEach(async (to, _, next) => {
    const authStroe = useAuthStore();
    if (to.name != 'Login') {
        if (!authStroe.currentToken || authStroe.currentToken.length == 0) {
            next({ name: 'Login' });
        }
        if (!authStroe.user) {
            const hide = message.loading('请稍后...', 0)
            try {
                const user = await apiGetAuthInfo();
                authStroe.setUser(user);
                const dicts: any = await apiGetEnumDict()
                enumDict.initDict(dicts);
                next();
            } catch (e: any) {
                message.warn(e.toString());
                console.warn(e);
                next({ name: 'Login' })
            } finally {
                hide();
            }
        }
    }
    next();
})

export default router