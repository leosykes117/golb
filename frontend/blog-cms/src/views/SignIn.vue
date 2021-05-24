<template>
    <div class="signin">
		<div class="container-form">
			<div class="form-vertical-align">
				<h1>Inicia Sesión</h1>
				<LogIn />
			</div>
		</div>
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

<style>

</style>