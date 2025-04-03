import type {Dispatch} from 'react';
import {createContext, useContext} from 'react';

import type {ActionProps} from './action';
import type {ComponentProps} from './Component';
import {getNodeAt, updatePaths} from './recursion';
import type {Struct} from './value';

export interface AppState {
    tree: ComponentProps | null;
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
    if (!state.tree) {
        state.tree = {};
    }

    // TODO don't modify, return a new state
    // TODO does react handle partial updates or re-renders everything?
    action.modify(getNodeAt(state.tree, action.path.slice(1)));

    updatePaths(state.tree);

    return {tree: state.tree};
}

export {AppContext};
