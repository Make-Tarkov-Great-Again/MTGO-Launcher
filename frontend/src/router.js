import App from './App.vue'
import ModPageComponent from './components/ModPageComponent.vue'
import { createRouter, createWebHistory, createWebHashHistory } from "vue-router"


const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/', name: "home", component: App },
    { path: '/mod/:id', name: "modpage", component: ModPageComponent, props: true },
  ],
})

export default router