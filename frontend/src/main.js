import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import naive from 'naive-ui'
import axios from 'axios'
import VueAxios from 'vue-axios'

createApp(App).use(router).use(naive).use(VueAxios, axios).mount('#app')
