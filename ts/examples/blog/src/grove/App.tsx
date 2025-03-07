import React, {useEffect, useReducer} from "react";
import {Component, ComponentProps} from "./Component";
import {resolvePath} from "../grovex/routing";
import {AppContext, appReducer} from "./app-state";

export interface AppProps {
    grove_server: string;
    err_404: ComponentProps;
    err_500: ComponentProps;
    loading: ComponentProps;
}

export const App: React.FC<AppProps> = (props) => {
    const [state, dispatch] = useReducer(appReducer, {tree: null});

    useEffect(() => {
        const {nodePath, params} = resolvePath(window.location.pathname);

        if (!nodePath) {
            dispatch({nodePath: [], node: props.err_404});
            return;
        }

        fetch(`${props.grove_server}/fetch/${nodePath}`, {
            method: "POST",
            headers: {"Content-Type": "application/json"},
            body: JSON.stringify(params),
        })
            .then((resp) => {
                if (!resp.ok) throw new Error("Grove API returned error");
                return resp.json();
            })
            .then((node) => dispatch({nodePath: [], node}))
            .catch(() => dispatch({nodePath: [], node: props.err_500}));
    }, [props.grove_server]);

    if (!state.tree) return <Component props={props.loading}/>;

    return (
        <AppContext.Provider value={{state, dispatch}}>
            <Component props={state.tree}/>
        </AppContext.Provider>
    );
};

export default App;