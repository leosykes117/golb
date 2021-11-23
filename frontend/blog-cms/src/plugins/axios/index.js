import Vue from 'vue'
import { golbClient } from './axiosClients'
import VueAxios from 'vue-axios'

Vue.use(VueAxios, golbClient)