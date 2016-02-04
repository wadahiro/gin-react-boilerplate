var path = require('path')
var webpack = require('webpack')

var config = require('./webpack.base.config.js')

config.debug = true
config.profile = false
config.devtool = 'inline-source-map'

// config.entry.unshift('webpack/hot/only-dev-server')
// config.entry.unshift('webpack-dev-server/client?http://localhost:9000')

// config.module.loaders = config.module.loaders.map(function (loader) {
//   loader.loader = 'react-hot!' + loader.loader
//   return loader
// })

config.plugins = config.plugins.concat([
//   new webpack.HotModuleReplacementPlugin(),
  new webpack.NoErrorsPlugin()
])

module.exports = config
