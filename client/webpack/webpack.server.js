var webpack = require('webpack')
var WebpackDevServer = require('webpack-dev-server')
var config = require('./webpack.config')

new WebpackDevServer(webpack(config), {
  publicPath: 'http://localhost:9000' + config.output.publicPath,
  contentBase: __dirname + '/public',
  hot: true,
  historyApiFallback: true,
  stats: {colors: true},
  proxy: {
    '/example/*': 'http://localhost:3000'
  }
}).listen(9000, 'localhost', function (err, result) {
  if (err) {
    console.log(err)
  }

  console.log('Listening at localhost:9000')
})
