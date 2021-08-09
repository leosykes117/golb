'use strict'

import Vue from 'vue'
import axios from 'axios'

axios.defaults.headers.common['Access-Control-Allow-Origin'] = '*'
axios.defaults.headers.common['X-Requested-With'] = 'XMLHttpRequest'

let config = {
	baseURL: process.env.VUE_APP_APIURL || 'http://localhost:3000/api/',
}

export const httpClient = axios.create(config)

Plugin.install = function (Vue) {
	Vue.axios = httpClient
	window.axios = httpClient
	Object.defineProperties(Vue.prototype, {
		axios: {
			get() {
				return httpClient
			},
		},
		$axios: {
			get() {
				return httpClient
			},
		},
	})
}

Vue.use(Plugin)

export default Plugin
