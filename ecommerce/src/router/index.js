import { createRouter, createWebHistory } from 'vue-router'
import Orders from "@/views/Orders.vue";

const routes = [
  {
    path: '/orders',
    name: 'orders',
    component: Orders
  },
  {
    path: '/',
    name: 'main',
    component: () => import(/* webpackChunkName: "about" */ '../views/Main.vue')
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
