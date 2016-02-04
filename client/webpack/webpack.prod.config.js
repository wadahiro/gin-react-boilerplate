var path = require('path')
var webpack = require('webpack')

var config = require('./webpack.base.config.js')

config.bail = true
config.debug = false
config.profile = false
// config.devtool = '#source-map'

config.plugins = config.plugins.concat([
    new webpack.optimize.OccurenceOrderPlugin(true),
    new webpack.optimize.DedupePlugin(),
//   new webpack.optimize.UglifyJsPlugin({ output: {comments: false} }),
])

module.exports = config
