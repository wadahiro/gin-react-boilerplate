var NODE_ENV = process.env.NODE_ENV;
var isProduction = NODE_ENV && NODE_ENV.trim() === 'production';
console.log('NODE_ENV: ' + NODE_ENV);
var config = isProduction ? require('./webpack.prod.config.js') : require('./webpack.dev.config.js');

module.exports = config;