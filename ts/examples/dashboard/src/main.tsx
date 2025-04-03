import {App, groveFetch, groveSubmit, initRouting} from '@grove/grove';
import {CssBaseline} from '@mui/material';
import React from 'react';
import ReactDOM from 'react-dom/client';

import {base_url} from './const';
import AppTheme from './theme/AppTheme';
import '@grove/grovex';
import './setup';

initRouting(
    new Map<string, string>([
        ['/', 'home'],
        ['/users', 'users'],
        ['/permissions', 'permissions'],
        ['/events', 'events'],
    ]),
);

ReactDOM.createRoot(document.getElementById('root')!).render(
    <React.StrictMode>
        <AppTheme>
            <CssBaseline />
            <App
                apiHandlers={{
                    fetch: groveFetch(base_url),
                    submit: groveSubmit(base_url),
                }}
                err404={{type: 'admin.SimplePage', key: '404', props: {code: '404', message: 'Not Found'}}}
                err500={{type: 'admin.SimplePage', key: '500', props: {code: '500', message: 'Internal Error'}}}
                loading={{type: 'admin.SimplePage', key: 'loading', props: {code: 'Loading', message: 'Fetching'}}}
            />
        </AppTheme>
    </React.StrictMode>,
);
