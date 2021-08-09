import AuthServices from './services'

export const SET_TOKEN = 'auth/SET_TOKEN'

export const AuthStore = {
	namespaced: true,
	state: {
		token: '',
	},
	actions: {
		signIn: ({ commit }, payload) => {
			return new Promise((resolve, reject) => {
				AuthServices.login(payload)
					.then(res => {
						console.log({ res })
						const token = res.data.result.token
						commit(SET_TOKEN, token)
						resolve()
					})
					.catch(err => {
						console.error({err})
						reject(err.response.data)
					})
			})
		},
	},
	mutations: {
		[SET_TOKEN](state, payload) {
			state.token = payload
		},
	},
}
