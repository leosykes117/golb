<template>
	<div class="login">
		<el-form :model="signInForm" ref="signInForm" status-icon :rules="rules" class="demo-ruleForm">
			<el-form-item prop="email">
				<el-input placeholder="Email" v-model="signInForm.email" autofocus clearable autocomplete="off"></el-input>
			</el-form-item>

			<el-form-item prop="password">
				<el-input
					placeholder="Password"
					v-model="signInForm.password"
					show-password
					clearable
					autocomplete="off"
				></el-input>
			</el-form-item>

			<el-form-item>
				<el-button type="primary" round @click="submitForm('signInForm')"> Iniciar Sesión </el-button>
			</el-form-item>
		</el-form>
	</div>
</template>

<script>
export default {
	name: 'LogIn',
	data() {
		var validateEmail = (rule, value, callback) => {
			const re =
				/^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
			if (typeof value !== 'string' || value.length < 1) {
				callback(new Error('Por favor ingresa un email'))
			} else if (!re.test(value)) {
				callback(new Error('El email no es válido'))
			} else {
				callback()
			}
		}
		return {
			signInForm: {
				email: '',
				password: '',
			},
			rules: {
				email: [{ validator: validateEmail, trigger: 'blur' }],
				password: [{ required: true, message: 'Por favor ingresa el password', trigger: 'blur' }],
			},
			sigInData: {
				email: '',
				passwordHash: '',
			},
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
			this.sigInData.password = this.signInForm.password
			const loading = this.$loading({
				lock: true,
				text: 'Enviando...',
				spinner: 'el-icon-loading',
				background: 'rgba(0, 0, 0, 0.7)',
			})
			const headers = {
				'Content-Type': 'application/json',
			}
			try {
				let response = await this.axios.post(`users/auth`, this.sigInData, { headers })
				loading.close()
				console.log(response)
			} catch (err) {
				loading.close()
				console.error('Ocurrio un error al iniciar sesión')
				let data = err.response.data
				console.error({ data })
				this.$message({
					message: data.message,
					type: 'error',
				})
				this.signInForm.password = ''
			}
		},
	},
}
</script>

<style></style>
