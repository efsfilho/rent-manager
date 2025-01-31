import { createRouter, createWebHistory } from 'vue-router'
import Rents from '../views/Rents.vue'
import Settings from '../views/Settings.vue'
import Dashboard from '../views/Dashboard.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '',
      name: 'home',
      component: Dashboard,
    },
    {
      name: 'rents',
      component: Rents,
      meta: {
        title: 'Rents'
      }
    },
    {
      name: 'settings',
      component: Settings,
      meta: {
        title: 'Settings'
      }
    },
  ]
})

export default router
