import { createRouter, createWebHistory } from 'vue-router'
import SearchView from '@/views/SearchView.vue'

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            redirect: 'search',
        },
        {
            path: '/search',
            name: 'search',
            component: SearchView
        },
        {
            path: '/result',
            name: 'result',
            component: () => import('@/views/ResultView.vue')
        }
    ]
})

export default router
