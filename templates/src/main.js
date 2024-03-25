import { createApp } from 'vue'
import App from './App.vue'

import { createRouter, createWebHistory } from 'vue-router'

// Vuetify
import 'vuetify/styles'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import DayJsAdapter from '@date-io/dayjs'


import HomePage from './components/HomePage.vue'
import RegisterPage from './components/RegisterPage.vue'
import LoginPage from './components/LoginPage.vue'
import TicketList from './components/TicketList.vue'

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
        path: '/ticket/list',
        component:  TicketList,
        name: 'ticketList'
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes
})
const vuetify = createVuetify({
    components,
    directives,
    date: {
        adapter: new DayJsAdapter(),
    },
})

createApp(App).use(vuetify).use(router).mount('#app')

