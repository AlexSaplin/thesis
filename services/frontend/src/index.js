import React from 'react';
import ReactDOM from 'react-dom';

import {Provider} from "react-redux";
import {applyMiddleware, createStore} from 'redux';
import thunkMiddleware from 'redux-thunk';
import {createLogger} from 'redux-logger';
import rootReducer from './reducers';
import {loadState, saveState} from './persist_store'

import axios from "axios";

import App from './pages/app';

import {BrowserRouter as Router} from 'react-router-dom'

import moment from 'moment';
import 'moment/locale/ru';

moment.locale('ru');

const loggerMiddleware = createLogger();
const persistedState = loadState();
const store = createStore(
    rootReducer,
    persistedState,
    applyMiddleware(
        thunkMiddleware,
        loggerMiddleware
    )
);

store.subscribe(() => {
    saveState({
        authentication: {...store.getState().authentication, errorMessage: ''}
    });
});

const {authentication = {}} = store.getState();
const {token = ''} = authentication;
if (token !== '') {
    axios.defaults.headers.common['Authorization'] = 'Token ' + token;
    axios.defaults.headers.common['X-Token'] = token;
}

(function (b, o, n, g, s, r, c) {
    if (b[s]) return;
    b[s] = {};
    b[s].scriptToken = "XzEzMzIzNjgxODQ";
    b[s].callsQueue = [];
    b[s].api = function () {
        b[s].callsQueue.push(arguments);
    };
    r = o.createElement(n);
    c = o.getElementsByTagName(n)[0];
    r.async = 1;
    r.src = g;
    r.id = s + n;
    c.parentNode.insertBefore(r, c);
})(window, document, "script", "https://cdn.oribi.io/XzEzMzIzNjgxODQ/oribi.js", "ORIBI");


ReactDOM.render(
    <Provider store={store}>
        <Router>
            <App/>
        </Router>
    </Provider>,
    document.getElementById('root')
);
