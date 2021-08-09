module.exports = {
	lintOnSave: false,
	devServer: {
		disableHostCheck: true,
		proxy: {
			'^/api': {
				target: 'http://gocms:3000/api/',
				changeOrigin: true,
			},
		},
	},
}
