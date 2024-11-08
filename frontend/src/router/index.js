import { createMemoryHistory, createRouter } from 'vue-router'

const routes = [
    {
        path: '/',
        redirect: '/Database'
    },
    {
        path: '/Database',
        component: () => import('../view/Database.vue')
    }
]

const router = createRouter({
    history: createMemoryHistory(),
    routes
})

export default router