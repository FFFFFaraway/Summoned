import Vue from 'vue';
import VueRouter from 'vue-router';
Vue.use(VueRouter);

export default new VueRouter({
    routes: [
        { path: '/', component: () => import('./pages/Home') },
        { path: '/home', component: () => import('./pages/Home') },
        { path: '/login', component: () => import('./pages/Login') },
        { path: '/signup', component: () => import('./pages/Signup') },
        { path: '/profile/:id', component: () => import('./pages/Profile') },
        { path: '/summoned', component: () => import('./pages/Summoned') },
        { path: '/mysummoned', component: () => import('./pages/MySummoned') },
        { path: '/othersummoned', component: () => import('./pages/OtherSummoned') },
        { path: '/summoned/:id', component: () => import('./pages/OneSummoned') }
    ]
});