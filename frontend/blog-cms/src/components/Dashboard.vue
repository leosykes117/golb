<template>
	<div>
		<el-header>
			<el-menu mode="horizontal" collapse="true" background-color="#FF7A4F" text-color="#085464" active-text-color="#085464">
				<el-menu-item index="1" class="menu-title">Golb</el-menu-item>
				<el-menu-item index="2">
					<font-awesome-icon icon="sign-in-alt" />Iniciar Sesión
				</el-menu-item>
				<el-menu-item index="3">
					<font-awesome-icon icon="user-plus" />Registrate
				</el-menu-item>
			</el-menu>
		</el-header>
	</div>
</template>

<script>
export default {
	name: 'Dashboard',
	data() {
		return {}
	},
}
</script>

<style lang="scss" scoped>
.el-header {
	background-color: #ff7a4f;
	color: #085464;
}
.menu-title {
	color: #e8e8e6 !important;
}
</style>
