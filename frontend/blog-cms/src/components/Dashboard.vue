<template>
	<div>
		<nav class="navbar navbar-expand-lg navbar-dark nav-bg">
			<div class="container-fluid">
				<a class="navbar-brand nav-title" href="#">Golb</a>
				<button
					class="navbar-toggler"
					type="button"
					data-bs-toggle="collapse"
					data-bs-target="#navbarSupportedContent"
					aria-controls="navbarSupportedContent"
					aria-expanded="false"
					aria-label="Toggle navigation"
				>
					<span class="navbar-toggler-icon"></span>
				</button>
				<div class="collapse navbar-collapse" id="navbarSupportedContent">
					<form class="d-flex">
						<el-input
							class="me-1"
							size="mini"
							placeholder="Search"
							v-model="navForm.search"
							clearable
							autocomplete="off"
						></el-input>
					</form>
					<ul class="navbar-nav me-auto mb-2 mb-lg-0">
						<li class="nav-item">
							<a class="nav-link" aria-current="page" href="#">Item 1</a>
						</li>
						<li class="nav-item">
							<a class="nav-link" href="#">Item 2</a>
						</li>
						<li class="nav-item"><a href="" class="nav-link">Item 3</a></li>
					</ul>
					<ul v-if="isLoggedIn" class="navbar-nav mb-2 mb-lg-0">
						<li class="nav-item dropdown">
							<a
								class="nav-link dropdown-toggle"
								data-bs-toggle="dropdown"
								href="#"
								role="button"
								aria-expanded="false"
								>Dropdown</a
							>
							<ul class="dropdown-menu dropdown-menu-end">
								<li><a class="dropdown-item" href="#">Action</a></li>
								<li><a class="dropdown-item" href="#">Another action</a></li>
								<li><a class="dropdown-item" href="#">Something else here</a></li>
								<li><hr class="dropdown-divider" /></li>
								<li>
									<a class="dropdown-item" @click="logout">Sign out</a>
								</li>
							</ul>
						</li>
					</ul>
					<ul v-else class="navbar-nav mb-2 mb-lg-0">
						<li class="nav-item">
							<a class="nav-link" aria-current="page" href="#">Sign in</a>
						</li>
						<li class="nav-item"><a href="" class="nav-link">Sign up</a></li>
					</ul>
				</div>
			</div>
		</nav>
		<div class="vh-100">
			<el-form :inline="true" class="demo-form-inline">
				<el-form-item>
					<el-button type="primary" @click="onSubmit">Query</el-button>
				</el-form-item>
			</el-form>
			<p class="container border border-dark border-3 rounded">
				{{ JSON.stringify(response, undefined, '\t') }}
			</p>
		</div>
	</div>
</template>

<script>
import { mapGetters } from 'vuex'

export default {
	name: 'Dashboard',
	data() {
		return {
			navForm: {
				search: '',
			},
			errorMessage: {
				message: 'Ocurrio un error inesperado',
				type: 'error',
			},
			response: {},
		}
	},
	methods: {
		async logout() {
			const loading = this.$loading({
				lock: true,
				text: 'Enviando...',
				spinner: 'el-icon-loading',
				background: 'rgba(0, 0, 0, 0.7)',
			})
			try {
				await this.$store.dispatch('auth/logOut')
				console.log('Logout SUCCESS!')
				this.$router.push('/signin')
			} catch (err) {
				console.error('Logout FAIL!')
				console.error(err)
				this.errorMessage = {
					message: err.message || 'Ocurrio un error al cerrar sesión',
					type: 'error',
				}
			}
			loading.close()
			if (this.errorMessage) {
				this.$message(this.errorMessage)
				this.errorMessage = undefined
			}
		},
		async onSubmit() {
			console.log('submit!')
			try {
				console.log('token =>', this.axios.defaults.headers.common['Authorization'])
				let response = await this.axios.post('users/test', {})
				console.log('RESPONSE =>', { response })
				this.response = response.data.result
			} catch (err) {
				console.error(err)
				this.errorMessage = {
					message: err.message || 'Ocurrio un error al cerrar sesión',
					type: 'error',
				}
			}
			if (this.errorMessage) {
				this.$message(this.errorMessage)
				this.errorMessage = undefined
			}
		},
	},
	computed: {
		...mapGetters('auth', ['isLoggedIn']),
	},
}
</script>

<style lang="scss" scoped>
.nav-bg {
	background-color: #ff7a4f;
}

.navbar-dark .navbar-nav .nav-link,
.navbar-dark .navbar-nav .show > .nav-link {
	color: #085464;
}

.navbar-dark .navbar-nav .nav-link:focus,
.navbar-dark .navbar-nav .nav-link:hover {
	color: #e8e8e6;
}

.nav-title {
	color: #e8e8e6 !important;
}
</style>
