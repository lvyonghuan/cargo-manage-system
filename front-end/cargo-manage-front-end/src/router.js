import { createRouter, createWebHistory } from 'vue-router'
import Dashboard from './dashboard.vue'
import Login from './login.vue'
import Add from './child/add.vue'
import Restock from './child/restock.vue'
import Products from './child/products.vue'

const routes = [
    { path: '/', component: Login },
    { path: '/login', component: Login },
    {
        path: '/dashboard',
        component: Dashboard,
        children: [
            { path: 'add', component: Add },
            {path: 'restock',component: Restock},
            {path:'products',component: Products},
        ]
    },
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router