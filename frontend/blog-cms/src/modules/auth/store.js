import AuthServices from './services'

export const SET_TOKEN = 'auth/SET_TOKEN'
export const SET_NAME = 'auth/SET_NAME'
export const SET_GENDER = 'auth/SET_GENDER'

export const AuthStore = {
	namespaced: true,
	state: {
		token: '',
		name: '',
		gender: false,
	},
	actions: {
		signIn({ dispatch }, payload) {
			return new Promise((resolve, reject) => {
				AuthServices.login(payload)
					.then((response) => {
						console.log({ response })
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
	},
}
