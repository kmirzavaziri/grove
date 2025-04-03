import type {Dispatch} from 'react';
import {createContext, useContext} from 'react';

import type {ActionProps} from './action';
import type {ComponentProps} from './Component';
import {setNodeAt} from './recursion';
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
    render: Dispatch<{path: string[]; node: ComponentProps; patch?: boolean}>;
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

export function appReducer(state: AppState, action: {path: string[]; node: ComponentProps; patch?: boolean}): AppState {
    if (!state.tree) {
        state.tree = action.node;
    }

    const path = action.path?.length ? action.path : [action.node.key || ''];

    if (path.length === 1) {
        state.tree = action.node;
    } else {
        if (state.tree.key !== path[0]) {
            throw new Error(`Path ${path.join('.')} not found: tree root is not ${path[0]}, it is ${state.tree.key}`);
        }
        setNodeAt(state.tree, path.slice(1), action.node, action.patch || false);
    }

    function populatePath(node: ComponentProps, parentPath: string[] = []): void {
        if (node.key === undefined) {
            node.key = '';
        }

        if (node.children === undefined) {
            node.children = [];
        }

        node.path = [...parentPath, node.key];

        for (const child of node.children) {
            populatePath(child, node.path);
        }
    }

    populatePath(state.tree);

    return {tree: state.tree};
}

export {AppContext};
