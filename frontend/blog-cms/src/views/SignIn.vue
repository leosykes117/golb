<template>
	<div class="container-login">
		<el-card shadow="always">
			<LogIn />
			<div class="content-center">
				<span>¿Aún no tienes cuenta?</span>
				<router-link to="/signup" v-slot="{ href, navigate }">
					<el-link type="primary" :href="href" @click="navigate">Registrate</el-link>
				</router-link>
			</div>
		</el-card>
	</div>
</template>

<script>
import LogIn from '../components/LogIn'

export default {
	name: 'SignIn',
	components: {
		LogIn,
	},
}
</script>

<style lang="scss" scoped>
.container-login {
	max-width: 450px;
	padding: 0 10px;
	margin: 170px auto;
}

.el-card {
	width: 100%;
}

.container-login .el-card .content-center a.el-link {
	margin-left: 3px;
	font-size: 16px;
	height: 16px;
}
</style>
