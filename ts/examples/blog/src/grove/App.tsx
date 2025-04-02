import React, {useEffect, useReducer} from "react";
import {Component, ComponentProps} from "./Component";
import {resolvePath} from "../grovex/routing";
import {ApiHandlers, AppContext, appReducer} from "./app-state";

export interface AppProps {
    apiHandlers: ApiHandlers;
    err404: ComponentProps;
    err500: ComponentProps;
    loading: ComponentProps;
}

export const App: React.FC<AppProps> = (props) => {
    const [state, render] = useReducer(appReducer, {tree: null});

    useEffect(() => {
        const {nodePath, params} = resolvePath(window.location.pathname);

        if (!nodePath) {
            render({path: [], node: props.err404});
            return;
        }

        props.apiHandlers.fetch(nodePath, params)
            .then((node) => render({path: [], node}))
            .catch(() => render({path: [], node: props.err500}));
    }, [props.apiHandlers, props.err404, props.err500]);

    if (!state.tree) return <Component props={props.loading}/>;

    return (
        <AppContext.Provider value={{state, render, apiHandlers: props.apiHandlers}}>
            <Component props={state.tree}/>
        </AppContext.Provider>
    );
};

export default App;

export const groveFetch = (groveServer: string) => (nodePath: string[], request?: any) =>
    fetch(`${groveServer}/fetch/${nodePath.join("/")}`, {
        method: "POST",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify(request),
    })
        .then((resp) => {
            if (!resp.ok) throw new Error("Grove API returned error");
            return resp.json();
        })

export const groveSubmit = (groveServer: string) => (nodePath: string[], request?: any) =>
    fetch(`${groveServer}/submit/${nodePath.join("/")}`, {
        method: "POST",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify(request),
    })
        .then((resp) => {
            if (!resp.ok) throw new Error("Grove API returned error");
            return resp.json();
        })

