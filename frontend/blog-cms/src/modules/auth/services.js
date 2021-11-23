import { golbClient } from '@/plugins/axios/axiosClients'

const auth = {
	login({ email, passwordHash }) {
		const headers = {
			'Content-Type': 'application/json',
		}
		const payload = {
			email,
			password: passwordHash,
		}
		return golbClient.post('users/auth', payload, { headers })
	},
	signUp({ name, surname, email, phone, gender, password }) {
		const headers = {
			'Content-Type': 'application/json',
		}
		const payload = {
			name,
			surname,
			email,
			phone,
			gender,
			password,
		}
		return golbClient.post('users', payload, { headers })
	},
	validateEmail(rule, value, callback) {
		const re =
			/^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
		if (typeof value !== 'string' || value.length < 1) {
			callback(new Error('Por favor ingresa un email'))
		} else if (!re.test(value)) {
			callback(new Error('El email no es válido'))
		} else {
			callback()
		}
	},
}

export default auth
