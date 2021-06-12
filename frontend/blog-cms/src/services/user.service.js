export const userService = {
	login,
	logout,
}

const login = async (username, password) => {
	const headers = {
		'Content-Type': 'application/json',
	}
	const payload = {
		username,
		password,
	}
	this.axios
		.post(`users/auth`, payload, { headers })
		.then(handleResponse)
		.then((user) => {
			const isObject = user instanceof Object,
				hasProperty = 'token' in user,
				tokenIsString = typeof user.token === 'string',
				tokenIsNotEmpty = user.token.length > 0
			if (isObject && hasProperty && tokenIsString && tokenIsNotEmpty) {
				localStorage.setItem('IDToken', user.token)
				delete user.token
				return user
			} else {
				return Promise.reject('La respuesta no tiene la estrucutura correcta')
			}
		})
}

const logout = () => {
	localStorage.removeItem('IDToken')
}

const handleResponse = (response) => {
	let info = response.data
	if (!info.success) {
		return Promise.reject(info.message)
	}
	return info.result
}
