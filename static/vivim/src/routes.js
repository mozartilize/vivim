import VueRouter from 'vue-router'
import HelloWorld from './components/HelloWorld'
import Me from './components/Me'

export default new VueRouter({
  mode: process.env.NODE_ENV === 'production' ? 'history' : 'hash',
  routes: [
    { path: '/hello', component: HelloWorld },
    { path: '/me', component: Me }
  ]
})
