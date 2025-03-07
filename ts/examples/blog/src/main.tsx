import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './grove/App';
import {base_url} from './const';
import {CssBaseline} from "@mui/material";
import AppTheme from "./theme/AppTheme";
import {initRouting} from "./grovex/routing";
import "./grovex/setup";
import "./setup";

initRouting(new Map<string, string>([
    ['/', 'home'],
    ['/users', 'users'],
    ['/permissions', 'permissions'],
    ['/events', 'events'],
]));

ReactDOM.createRoot(document.getElementById('root')!).render(
    <React.StrictMode>
        <AppTheme>
            <CssBaseline/>
            <App
                grove_server={base_url}
                err_404={{type: "admin.SimplePage", key: "404", props: {code: "404", message: "Not Found"}}}
                err_500={{type: "admin.SimplePage", key: "500", props: {code: "500", message: "Internal Error"}}}
                loading={{type: "admin.SimplePage", key: "loading", props: {code: "Loading", message: "Fetching"}}}
            />
        </AppTheme>
    </React.StrictMode>
);

