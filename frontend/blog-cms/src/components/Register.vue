<template>
	<el-form :model="signUpForm" ref="signUpForm" status-icon>
		<el-form-item prop="name">
			<el-input placeholder="Nombre" v-model="signUpForm.name" autofocus clearable autocomplete="off"></el-input>
		</el-form-item>

		<el-form-item prop="surname">
			<el-input placeholder="Apellidos" v-model="signUpForm.surname" clearable autocomplete="off"></el-input>
		</el-form-item>

		<el-form-item prop="email">
			<el-input
				placeholder="Correo electrónico o teléfono"
				v-model="signUpForm.email"
				clearable
				autocomplete="off"
			></el-input>
		</el-form-item>

		<el-form-item prop="gender" label="Genero" class="content-center">
			<el-switch v-model="signUpForm.gender" active-value="M" active-text="Masculino" class="switch-element"></el-switch>
			<el-switch
				v-model="signUpForm.gender"
				active-value="F"
				active-text="Femenino"
				class="switch-element"
				tabindex="-1"
			></el-switch>
		</el-form-item>

		<el-form-item prop="password">
			<el-input placeholder="Password" v-model="signUpForm.password" show-password clearable autocomplete="off"></el-input>
		</el-form-item>

		<el-form-item prop="confirmPassword">
			<el-input
				placeholder="Confirma tu password"
				v-model="signUpForm.confirmPassword"
				show-password
				clearable
				autocomplete="off"
			></el-input>
		</el-form-item>

		<el-form-item class="content-center">
			<el-button type="primary" round @click="submitForm('signUpForm')">Registarse</el-button>
		</el-form-item>
	</el-form>
</template>

<script>
export default {
	name: 'Register',
	data() {
		return {
			signUpForm: {
				name: '',
				surname: '',
				email: '',
				phone: '',
				gender: '',
				password: '',
			},
		}
	},
	methods: {
		submitForm(formName) {
			this.$refs[formName].validate((valid) => {
				if (valid) {
					this.signUp(formName)
				} else {
					console.log('Error al enviar el formulario')
					return false
				}
			})
		},
		async signUp() {
			const loading = this.$loading({
				lock: true,
				text: 'Enviando...',
				spinner: 'el-icon-loading',
				background: 'rgba(0, 0, 0, 0.7)',
			})
			try {
				await this.$store.dispatch('auth/signUp', this.signUpForm)
				console.log('SUCCESS!')
				this.$router.push('/')
			} catch (err) {
				console.error('FAIL!')
				console.error(err)
				this.errorMessage = {
					message: err.message,
					type: 'error',
				}
			}
			loading.close()
			if (this.errorMessage) {
				this.$message(this.errorMessage)
				this.errorMessage = undefined
			}
		},
	},
}
</script>

<style lang="scss" scoped></style>
