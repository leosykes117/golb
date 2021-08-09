import { httpClient } from '@/plugins/axios'

const auth = {
	login: ({ email, passwordHash }) => {
		const headers = {
			'Content-Type': 'application/json',
		}
		const payload = {
			email,
			password: passwordHash,
		}

		return httpClient.post(`users/auth`, payload, { headers })
	},
}

export default auth
