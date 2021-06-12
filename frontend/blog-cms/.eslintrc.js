module.exports = {
	root: true,
	env: {
		node: true,
	},
	extends: ['plugin:vue/essential', 'eslint:recommended', '@vue/prettier'],
	parserOptions: {
		parser: 'babel-eslint',
	},
	rules: {
		'no-console': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
		'no-debugger': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
		'prettier/prettier': [
			'warn',
			{
				semi: false,
				singleQuote: true,
				useTabs: true,
				tabWidth: 4,
				printWidth: 130,
				arrowParens: 'always',
				bracketSpacing: true,
				trailingComma: 'es5',
			},
		],
	},
}
