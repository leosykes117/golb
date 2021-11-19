'use strict'

import Vue from 'vue'
import axios from 'axios'

let config = {
	baseURL: process.env.VUE_APP_APIURL
}
const token = localStorage.getItem('token')

export const httpClient = axios.create(config)

httpClient.defaults.headers.common['Access-Control-Allow-Origin'] = '*'
httpClient.defaults.headers.common['X-Requested-With'] = 'XMLHttpRequest'

if (token) {
	httpClient.defaults.headers.common['Authorization'] = `Bearer ${token}`
}

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
