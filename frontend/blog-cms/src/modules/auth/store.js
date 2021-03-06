import AuthServices from './services'
import { golbClient } from '@/plugins/axios/axiosClients'

export const SET_TOKEN = 'auth/SET_TOKEN'
export const SET_NAME = 'auth/SET_NAME'
export const SET_GENDER = 'auth/SET_GENDER'
export const SET_LOG_OUT = 'auth/SET_LOG_OUT'

export const AuthStore = {
	namespaced: true,
	state: {
		token: localStorage.getItem('user-token') || '',
		name: '',
		gender: false,
	},
	actions: {
		signIn({ dispatch }, payload) {
			return new Promise((resolve, reject) => {
				AuthServices.login(payload)
					.then((response) => {
						console.log({ response })
						const token = response.data.result.token
						localStorage.setItem('user-token', token)
						golbClient.defaults.headers.common['Authorization'] = `Bearer ${token}`
						dispatch('populateAuthData', response.data.result)
						resolve()
					})
					.catch((err) => {
						console.error({ err })
						let existsResponse = typeof err.response === 'object' && err.response !== null
						if (existsResponse && 'data' in err.response) {
							reject(err.response.data)
						} else {
							reject(err)
						}
					})
			})
		},
		signUp({ dispatch }, payload) {
			return new Promise((resolve, reject) => {
				AuthServices.signUp(payload)
					.then((response) => {
						console.log({ response })
						const token = response.data.result.token
						localStorage.setItem('user-token', token)
						golbClient.defaults.headers.common['Authorization'] = `Bearer ${token}`
						dispatch('populateAuthData', response.data.result)
						resolve()
					})
					.catch((err) => {
						console.error({ err })
						let existsResponse = typeof err.response === 'object' && err.response !== null
						if (existsResponse && 'data' in err.response) {
							reject(err.response.data)
						} else {
							reject(err)
						}
					})
			})
		},
		populateAuthData({ commit }, payload) {
			commit(SET_TOKEN, payload.token)
			commit(SET_NAME, payload.name)
			commit(SET_GENDER, payload.gender)
		},
		logOut({ commit }) {
			return new Promise((resolve) => {
				commit(SET_LOG_OUT)
				localStorage.removeItem('user-token')
				delete golbClient.defaults.headers.common['Authorization']
				resolve()
			})
		},
	},
	mutations: {
		[SET_TOKEN](state, payload) {
			state.token = payload
		},
		[SET_NAME](state, payload) {
			state.name = payload
		},
		[SET_GENDER](state, payload) {
			state.gender = payload
		},
		[SET_LOG_OUT](state) {
			state.token = ''
			state.name = ''
			state.gender = false
		},
	},
	getters: {
		isLoggedIn: (state) => !!state.token,
	},
}
