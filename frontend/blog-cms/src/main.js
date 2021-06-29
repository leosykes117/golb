import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import './plugins/element.js'
import './axios'

import axios from 'axios'
import vueAxios from 'vue-axios'

import './plugins/font-awesome'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

Vue.use(vueAxios, axios)

Vue.component('font-awesome-icon', FontAwesomeIcon)

Vue.config.productionTip = false

new Vue({
	router,
	store,
	render: (h) => h(App),
}).$mount('#app')
