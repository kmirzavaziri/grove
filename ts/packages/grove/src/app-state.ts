import type {Dispatch} from 'react';
import {createContext, useContext} from 'react';

import type {ActionProps} from './action';
import type {ComponentProps} from './Component';
import type {Struct} from './value';
import {getNodeAt, updatePaths} from './recursion';

export interface AppState {
    node: ComponentProps | null;
}

export interface ApiHandlers {
    fetch: (nodePath: string[], request?: Struct) => Promise<ComponentProps>;
    submit: (nodePath: string[], request?: Struct) => Promise<ActionProps>;
}

export interface AppContextValue {
    state: AppState;
    dispatch: Dispatch<{path: string[]; modify: (node: ComponentProps) => void}>;
    apiHandlers: ApiHandlers;
}

const AppContext = createContext<AppContextValue | undefined>(undefined);

export function useAppContext(): AppContextValue {
    const context = useContext(AppContext);
    if (!context) {
        throw new Error('useAppContext must be used within an AppProvider');
    }
    return context;
}

export function appReducer(
    state: AppState,
    action: {
        path: string[];
        modify: (node: ComponentProps) => void;
    },
): AppState {
    let node = structuredClone(state.node);

    if (!node) {
        node = {};
    }

    const targetNode = getNodeAt(node, action.path.slice(1));

    action.modify(targetNode);

    updatePaths(node);

    return {node: node};
}

export {AppContext};
