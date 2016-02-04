import * as ReactDOM from 'react-dom';
import * as React from 'react';
import App from './App';

require("babel-polyfill");
require('whatwg-fetch');

ReactDOM.render(<App />, document.getElementById('app'));