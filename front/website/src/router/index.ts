import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router";

// 2. 配置路由
const routes: Array<RouteRecordRaw> = [
    {
        path: "/login",
        name: 'Login',
        component: () => import("@/components/login/index.vue"),
    },
    {
        path: "/register",
        name: 'Register',
        component: () => import("@/components/register/index.vue")
    },
    {
        path: "/calendar",
        name: 'Calendar',
        component: () => import("@/components/editCalendar/index.vue")
    },
    {
        path: "/clock",
        name: 'Clock',
        component: () => import("@/components/editClock/index.vue")
    },
    {
        path: "/",
        name: "Panel",
        component: () => import("@/components/panel/index.vue")
    },
];
// 1.返回一个 router 实列，为函数，里面有配置项（对象） history
const router = createRouter({
    history: createWebHashHistory(),
    routes,
});

// 3导出路由   然后去 main.ts 注册 router.ts
export default router