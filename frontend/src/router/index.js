import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import Login from '../views/Login.vue'
import Register from '../views/Register.vue'
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
      footer: Footer,
    }
  },
  {
    path: "/login",
    name: "Login",
    components: {
      header: Header,
      default: Login,
      footer: Footer,
    }
  },
  {
    path: "/register",
    name: "Register",
    components: {
      header: Header,
      default: Register,
      footer: Footer,
    }
  }
]

const router = new VueRouter({
  routes
})

export default router
