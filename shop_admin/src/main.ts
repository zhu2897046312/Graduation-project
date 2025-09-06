import { createApp } from 'vue'
import App from './App.vue'
import { createPinia } from 'pinia'
import router from './router'
import permission from "/@/directive/permission.ts";

import "quill/dist/quill.core.css";
import 'quill/dist/quill.snow.css';

import './assets/main.css'
import './style.css'

const pinia = createPinia()

createApp(App)
  .use(pinia)
  .use(router)
  .directive('permission', permission)
  .mount('#app')
