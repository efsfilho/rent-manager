// Composables
import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    // component: () => import('@/views/Panel.vue'),
    component: () => import('@/layouts/Home.vue'),
    // component: () => import('@/App.vue'),
    children: [
      {
        path: '/rents',
        component: () => import('@/views/Rents.vue')
      },
      {
        path: '/tenants',
        component: () => import('@/views/Tenants.vue')
      },
      {
        path: '/properties',
        component: () => import('@/views/Properties.vue')
      }
    ]
  },
  {
      path: '/login',
      name: 'Login',
      component: () => import(/* webpackChunkName: "home" */ '@/views/Login.vue'),
    }
    // children: [
    //   {
    //     path: '',
    //     name: 'Home',
    //     // route level code-splitting
    //     // this generates a separate chunk (about.[hash].js) for this route
    //     // which is lazy-loaded when the route is visited.
    //     component: () => import(/* webpackChunkName: "home" */ '@/views/Home.vue'),
    //   },
    // ],
  // {
  //   path: '/teste',
  //   name: 'Teste',
  //   component: () => import(/* webpackChunkName: "home" */ '@/views/Tenant.vue'),
  // },
  // {
  //   path: '/panel',
  //   name: 'Panel',
  //   component: () => import(/* webpackChunkName: "home" */ '@/views/Panel.vue'),
  // }, 
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
})

export default router
