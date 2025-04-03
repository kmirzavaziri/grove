import type {Dispatch} from 'react';
import {createContext, useContext} from 'react';

import type {ActionProps} from './action';
import type {ComponentProps} from './Component';
import type {Struct} from './value';
import {flatten, updatePaths} from './recursion';

export interface AppState {
    node: ComponentProps | null;
    flatNodes: Map<string, ComponentProps>;
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
    if (!state.node) {
        state.node = {};
    }

    // TODO don't modify, return a new state
    // TODO does react handle partial updates or re-renders everything?
    const node = state.flatNodes.get(action.path.join('/'));
    if (node) {
        action.modify(node);
    }

    updatePaths(state.node);

    return {
        node: state.node,
        flatNodes: flatten(state.node),
    };
}

export {AppContext};
