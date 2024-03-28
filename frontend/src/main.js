import { createMemoryHistory, createRouter } from 'vue-router'
import Vue, { createApp } from '@vue/compat';
import { BootstrapVue, IconsPlugin } from 'bootstrap-vue';

Vue.use(BootstrapVue)
Vue.use(IconsPlugin)

import './assets/base.css'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

import Calendar from './Calendar.vue'
import Register from './Register.vue'
import App from './App.vue'
import Main from './Main.vue'
import User from './User.vue'
import Conflict from './Conflict.vue'

const routes = [
    { path: '/', component: Main, name: 'Main' },
    { path: '/register', component: Register, name: 'Register' },
    { path: '/conflict', component: Conflict, name: 'Conflict' },
    { path: '/calendar', component: Calendar, name: 'Calendar' },
    { path: '/user', component: User, name: 'User' },
]
  
const router = createRouter({
    history: createMemoryHistory(),
    routes,
})

const app = createApp(App);
app.use(router);
app.mount("#app");