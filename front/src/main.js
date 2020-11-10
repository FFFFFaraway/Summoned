import Vue from 'vue'
import App from './App.vue'
import router from './router'
import axios from 'axios'
import Multiselect from 'vue-multiselect'
import store from './store';
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';

Vue.config.productionTip = false
// 下面这行很重要，解决session不一致问题
axios.defaults.withCredentials = true
// 下面这行不能使用127.0.0.1, 防止cookie问题的出现
axios.defaults.baseURL = 'http://localhost:9999';
Vue.prototype.$axios = axios
Vue.component('multiselect', Multiselect)
Vue.use(ElementUI);


new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app')
