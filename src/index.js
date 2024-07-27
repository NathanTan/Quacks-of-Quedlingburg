import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import { Provider } from 'mobx-react';
import reportWebVitals from './reportWebVitals';
import CounterStore from './store/CounterStore'; // import the CounterStore component

const root = ReactDOM.createRoot(document.getElementById('root') ?? document.createElement('div'));


const counterStore = new CounterStore(); // remove the '()' when creating an instance of CounterStore

root.render(
    <Provider CounterStore={counterStore}>
        <App />
    </Provider>,
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
