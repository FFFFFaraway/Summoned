import Vue from 'vue';
import VueRouter from 'vue-router';
import Home from './pages/Home';
import Login from './pages/Login';
import Signup from './pages/Signup';
import Profile from './pages/Profile';
import Summoned from './pages/Summoned';
import MySummoned from './pages/MySummoned';
import OtherSummoned from './pages/OtherSummoned';
import OneSummoned from './pages/OneSummoned'


Vue.use(VueRouter);

export default new VueRouter({
    routes: [
        { path: '/', component: Home },
        { path: '/home', component: Home },
        { path: '/login', component: Login },
        { path: '/signup', component: Signup },
        { path: '/profile', component: Profile },
        { path: '/summoned', component: Summoned },
        { path: '/mysummoned', component: MySummoned },
        { path: '/othersummoned', component: OtherSummoned },
        { path: '/summoned/:id', component: OneSummoned }
    ]
});