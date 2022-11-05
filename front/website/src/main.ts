import {
    createApp
} from 'vue'
import App from './App.vue'
// 通用字体
import 'vfonts/Lato.css'
// 等宽字体
import 'vfonts/FiraCode.css'
// 路由
import router from '@/router/index'

const app = createApp(App)
app.use(router)
app.mount('#app')