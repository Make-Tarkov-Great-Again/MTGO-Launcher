import { createApp } from 'vue';
import App from './App.vue';
import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: '/workshop/mod/:id',
            component: () => import('@/components/ModPageComponent.vue'),
            name: 'modpage',
            props: true
        }
    ]

})


createApp(App)
  .use(router)
  .mount('#app');