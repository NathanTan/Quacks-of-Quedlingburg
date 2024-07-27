import React from 'react';
import ReactDOM from 'react-dom';
import App from './App'; // adjust the path as necessary
import { Provider } from 'mobx-react';
import CounterStore from './store/CounterStore'; // import the CounterStore component


const counterStore = new CounterStore(); // remove the '()' when creating an instance of CounterStore

ReactDOM.render(
    <Provider CounterStore={counterStore}>
        <App />
    </Provider>,
    document.getElementById('root')
);