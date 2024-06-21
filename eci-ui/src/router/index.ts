import {createRouter, createWebHistory, RouteRecordRaw} from 'vue-router';

const routes: Array<RouteRecordRaw> = [
    {
        path: '/login',
        name: 'Login',
        component: () => import('@/components/Login.vue'),
    },
    {
        path: '/',
        name: 'Index',
        component: () => import('@/components/Index.vue'),
        children: [
            {
                path: '',
                name: 'Home',
                component: () => import('@/components/Home.vue'),
            },
            {
                path: '/eciList',
                name: 'EciList',
                component: () => import('@/components/EciList.vue'),
            },
            {
                path: '/containerList',
                name: 'ContainerList',
                component: () => import('@/components/ContainerList.vue'),
                /* 允许通过路由传递 props 数据 */
                props: true,
            },
        ]
    },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

/* beforeEach 中用不到的参数不建议写上去 */
router.beforeEach((to) => {
    if (to.path !== '/login') {
        let authToken = localStorage.getItem('authToken');
        if (!authToken) {
            // noinspection JSIgnoredPromiseFromCall
            router.push('/login');
            return false;
        }
    }

});

export default router;