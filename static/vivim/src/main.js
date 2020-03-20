import Vue from 'vue'
import VueRouter from 'vue-router'
import App from './App.vue'
import router from './routes'

Vue.config.productionTip = false
Vue.use(VueRouter)

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
