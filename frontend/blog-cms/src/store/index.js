import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

import { AuthStore } from '../modules/auth/store'

export default new Vuex.Store({
	modules: {
		auth: { ...AuthStore },
	},
})
