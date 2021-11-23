'use strict'

import axios from 'axios'
import router from '@/router'
import store from '@/store'

const token = localStorage.getItem('token')
let config = {
	baseURL: process.env.VUE_APP_APIURL,
}

const golbClient = axios.create(config)
golbClient.defaults.headers.common['Access-Control-Allow-Origin'] = '*'
golbClient.defaults.headers.common['X-Requested-With'] = 'XMLHttpRequest'
if (token) {
	console.log('Sí existe token =>', token)
	golbClient.defaults.headers.common['Authorization'] = `Bearer ${token}`
}

console.log("interceptor creado")
golbClient.interceptors.response.use(
    (response) => {
        return response
    },
    (error) => {
        // Return any error which is not due to authentication back to the calling service
        console.log('axios.interceptors.error =>', { error })
        if (error.response.status !== 401) {
            return new Promise((resolve, reject) => {
                reject(error)
            })
        } else {
            return new Promise((resolve, reject) => {
                store
                    .dispatch('auth/logOut')
                    .then((response) => {
                        console.log('axios.interceptors.store.dispatch.then: response =>', response)
                        console.log('axios.interceptors.store.dispatch.then: Logout SUCCESS!')
                        router.push({name: 'SignIn'})
						error.response.data.message = "La sesión ha expirado"
                        reject(error.response.data)
                    })
                    .catch((err) => {
                        onsole.error('axios.interceptors.store.dispatch.catch: err =>', err)
                        reject(error)
                    })
            })
        }
    }
)

export { golbClient }