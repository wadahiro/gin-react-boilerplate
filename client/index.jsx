import * as ReactDOM from 'react-dom';
import ReactDOMServer from 'react-dom/server';
import * as React from 'react';
import App from './app/App';

require("babel-polyfill");
require('whatwg-fetch');



if (typeof window !== 'undefined') {
    ReactDOM.render(<App />, document.getElementById('app'));
} else {
    global.main = (options, callback) => {
        console.log('render server side', JSON.stringify(options))
        const s = ReactDOMServer.renderToString(React.createElement(App, {}));
        
        console.log(s)
        
        callback(JSON.stringify({
            uuid: options.uuid,
            app: s,
            title: null,
            meta: null,
            initial: null,
            error: null,
            redirect: null
        }));
    };
}