import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import Header from '../layout/Header.vue'
import Footer from '../layout/Footer.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    components: {
      header: Header,
      default: Home,
      footer: Footer
    }
  }
]

const router = new VueRouter({
  routes
})

export default router
