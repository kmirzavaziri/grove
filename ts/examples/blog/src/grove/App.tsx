import React, {useEffect, useState} from 'react';
import {Component, ComponentProps} from "./Component";


export interface AppProps {
    grove_server: string;
}

export const App: React.FC<AppProps> = (props) => {
    const [root, setRoot] = useState<ComponentProps | null>(null);

    useEffect(() => {
        fetch(`${props.grove_server}/home`)
            .then(resp => resp.json())
            .then(node => setRoot(node))
            .catch(err => console.error('failed to fetch:', err));
    }, []);

    if (!root) return <div>Loading...</div>;

    return <Component props={root}/>;
};

export default App;
