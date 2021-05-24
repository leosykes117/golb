module.exports = {
  lintOnSave: false,
  devServer: {
    proxy: {
      '^/api': {
        target: 'http://gocms:3000/api/',
        changeOrigin: true
      },
    }
  }
};
