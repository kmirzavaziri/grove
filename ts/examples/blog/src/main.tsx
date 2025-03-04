import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './grove/App';
import {base_url} from './const';
import "./grovex/setup";

ReactDOM.createRoot(document.getElementById('root')!).render(
    <React.StrictMode>
        <App grove_server={base_url}/>
    </React.StrictMode>
);
