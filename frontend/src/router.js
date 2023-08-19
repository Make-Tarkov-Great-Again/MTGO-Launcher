import App from './App.vue'
import ModPageComponent from './components/ModPageComponent.vue'
import { createRouter, createWebHistory, createWebHashHistory } from "vue-router"


const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/', component: App },
    { path: '/mod/:id', component: ModPageComponent, props: true },
  ],
})

export default router