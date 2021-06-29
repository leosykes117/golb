import axios from 'axios'

axios.defaults.baseURL = 'http://localhost:3000/api/'
axios.defaults.headers.common['Access-Control-Allow-Origin'] = '*'
axios.defaults.headers.common['X-Requested-With'] = 'XMLHttpRequest'
