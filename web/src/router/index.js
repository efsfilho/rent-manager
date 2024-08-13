import { createRouter, createWebHistory } from 'vue-router'
import App from '../App.vue'
import NewBlock from '../views/NewBlock.vue'
import EditBlock from '../components/EditBlock.vue'
import Teste1 from '../views/Teste1.vue'
import Teste2 from '../views/Teste2.vue'
import Dashboard from '../views/Dashboard.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    // {
    //   path: '',
    //   name: 'home',
    //   component: App,
    //   // children: [
    //   //   // UserHome will be rendered inside User's <router-view>
    //   //   // when /user/:id is matched
    //   //   {
    //   //       path: 'dashboard',
    //   //       component: Dashboard
    //   //   },
  
    //   //   // ...other sub routes
    //   // ],
    // },
    {
      path: '',
      name: 'home',
      component: Dashboard,
      // props: { default: true }
    },
    {
      name: 'new-block',
      component: NewBlock,
      meta: { menuTitle: 'New Block'}
      // props: { newsletterPopup: false }
      // props:  true,
      // props: {
      //   // You must set `true` to the default view, which uses ProductDetail
      //   default: true
      // }
    },
    {
      name: 'teste1',
      component: Teste1,
      meta: { menuTitle: 'Teste1111'}
    },
    {
      name: 'teste2',
      component: Teste2,
      meta: { menuTitle: 'T 22 2'}
    }
    // {
    //   path: '/about',
    //   name: 'about',
    //   // route level code-splitting
    //   // this generates a separate chunk (About.[hash].js) for this route
    //   // which is lazy-loaded when the route is visited.
    //   component: () => import('../views/AboutView.vue')
    // }
  ]
})

export default router
