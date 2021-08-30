<template>
	<div class="home">
		<Dashboard v-if="isLoggedIn" />
		<router-link to="/signin" v-else>Iniciar Sesión</router-link>
	</div>
</template>

<script>
import Dashboard from '@/components/Dashboard.vue'
import { mapGetters } from 'vuex'

export default {
	name: 'Home',
	components: {
		Dashboard,
	},
	computed: {
		...mapGetters('auth', ['isLoggedIn']),
	},
}
</script>
