import { createApp } from 'vue'
import App from './App.vue'

import { createRouter, createWebHistory } from 'vue-router'

// Vuetify
import 'vuetify/styles'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'


import HomePage from './components/HomePage.vue'
import RegisterPage from './components/RegisterPage.vue'
import LoginPage from './components/LoginPage.vue'
import OrderList from './components/OrderList.vue'

const routes = [
    { 
        path: '/', 
        component: HomePage , 
        name: 'home'
    },
    { 
        path: '/register', 
        component: RegisterPage , 
        name: 'register'
    },
    { 
        path: '/login', 
        component: LoginPage , 
        name: 'login'
    },
    {
        path: '/order/list',
        component:  OrderList,
        name: 'orderList'
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes
})
const vuetify = createVuetify({
    components,
    directives,
})

createApp(App).use(vuetify).use(router).mount('#app')

