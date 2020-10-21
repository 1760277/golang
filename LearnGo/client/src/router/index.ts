import Vue from 'vue';
import VueRoute, { RouteConfig } from 'vue-router';
import auth from '../services/auth';

Vue.use(VueRoute);

const routes: RouteConfig[] = [
    {
        path: '',
        component: () => import('../components/User.vue'),
        children: [
            {
                path: '/',
                name: 'user',
                component: () => import('../components/User.vue'),
            },
        ],
    },
    {
        path: '/login',
        name: 'login',
        component: () => import('../components/Login.vue'),
    },
];

const router = new VueRoute({
    routes,
    mode: 'history',
});

router.beforeEach((to, from, next) => {
    if (to.name !== 'login') {
        if (!auth.isAuthenticated()) {
            next('/login');
        }
    } else {
        const authBool = auth.isAuthenticated();
        if (authBool) {
            next('/');
        }
    }
    next();
});

export default router;
