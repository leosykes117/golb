<template>
	<div>
		<el-row>
			<el-col :span="24">
				<h1>Registrate</h1>
			</el-col>
		</el-row>
		<div class="container-signup">
			<el-card shadow="always">
				<Register />
			</el-card>
		</div>
	</div>
</template>

<script>
import Register from '../components/Register'

export default {
	name: 'SignUp',
	components: {
		Register,
	},
}
</script>

<style lang="scss" scoped>
.container-signup {
	max-width: 450px;
	padding: 0 10px;
	margin: 0 auto;
}
</style>
