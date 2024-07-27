import React from 'react';
import ReactDOM from 'react-dom';
import MyComponent from './MyComponent'; // adjust the path as necessary
import './store'; // Ensure the store is imported if it needs to be initialized

ReactDOM.render(
    <React.StrictMode>
        <MyComponent message="Hello, World!"/>
    </React.StrictMode>,
    document.getElementById('root')
);