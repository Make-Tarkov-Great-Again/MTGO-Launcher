import { createApp } from 'vue';
import App from './App.vue';
import { createRouter, createWebHistory } from 'vue-router';
import FloatingVue from 'floating-vue';
import 'floating-vue/dist/style.css';
import mitt from 'mitt'


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
const emitter = mitt()
const app = createApp(App);
app.config.globalProperties.emitter = emitter

const vueApp = app
  .use(router)
  .mount('#app');

app.use(FloatingVue, {
  container: document.getElementById("#systemApp")
});
