<template>
	<el-form :model="signInForm" ref="signInForm" status-icon :rules="rules">
		<el-form-item prop="email">
			<el-input placeholder="Email o nombre de usuario" v-model="signInForm.email" autofocus clearable autocomplete="off">
				<template slot="prepend">
					<font-awesome-icon icon="user" />
				</template>
			</el-input>
		</el-form-item>

		<el-form-item prop="password">
			<el-input placeholder="Password" v-model="signInForm.password" show-password clearable autocomplete="off">
				<template slot="prepend">
					<font-awesome-icon icon="key" />
				</template>
			</el-input>
		</el-form-item>

		<el-link type="primary" class="" style="padding-top: 5px">Restablecer password</el-link>

		<el-form-item class="content-center" style="padding: 12px 20px">
			<el-button type="primary" round @click="submitForm('signInForm')"> Iniciar Sesión </el-button>
		</el-form-item>
	</el-form>
</template>

<script>
import AuthServices from '../modules/auth/services'

export default {
	name: 'LogIn',
	data() {
		return {
			signInForm: {
				email: '',
				password: '',
			},
			rules: {
				email: [{ validator: AuthServices.validateEmail, trigger: 'blur' }],
				password: [{ required: true, message: 'Por favor ingresa el password', trigger: 'blur' }],
			},
			sigInData: {
				email: '',
				passwordHash: '',
			},
			errorMessage: undefined,
		}
	},
	methods: {
		submitForm(formName) {
			this.$refs[formName].validate((valid) => {
				if (valid) {
					this.signIn(formName)
				} else {
					console.log('Error al enviar el formulario')
					return false
				}
			})
		},
		async signIn() {
			this.sigInData.email = this.signInForm.email
			this.sigInData.passwordHash = this.signInForm.password
			const loading = this.$loading({
				lock: true,
				text: 'Enviando...',
				spinner: 'el-icon-loading',
				background: 'rgba(0, 0, 0, 0.7)',
			})
			try {
				await this.$store.dispatch('auth/signIn', this.sigInData)
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
