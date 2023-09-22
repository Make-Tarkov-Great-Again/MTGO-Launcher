import { createApp } from 'vue';
import App from './App.vue';
import { createRouter, createWebHistory } from 'vue-router'
import FloatingVue from 'floating-vue'
import 'floating-vue/dist/style.css'

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
});

const app = createApp(App);

const vueApp = app
  .use(router)
  .mount('#app');



  import {
    // Directives
    VTooltip,
    VClosePopper,
    // Components
    Dropdown,
    Tooltip,
    Menu
  } from 'floating-vue'

  app.directive('tooltip', VTooltip)
  app.directive('close-popper', VClosePopper)
  app.component('VDropdown', Dropdown)
  app.component('VTooltip', Tooltip)
  app.component('VMenu', Menu)

