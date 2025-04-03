import type {Dispatch} from 'react';
import {createContext, useContext} from 'react';

import type {ActionProps} from './action';
import type {ComponentProps} from './Component';
import {patchComponentProps} from './Component';

export interface AppState {
    tree: ComponentProps | null;
}

export interface ApiHandlers {
    fetch: (nodePath: string[], request?: any) => Promise<ComponentProps>;
    submit: (nodePath: string[], request?: any) => Promise<ActionProps>;
}

export interface AppContextValue {
    state: AppState;
    render: Dispatch<{path: string[]; node: ComponentProps; patch?: boolean}>;
    apiHandlers: ApiHandlers;
}

const AppContext = createContext<AppContextValue | undefined>(undefined);

export function useAppContext() {
    const context = useContext(AppContext);
    if (!context) {
        throw new Error('useAppContext must be used within an AppProvider');
    }
    return context;
}

function updateTree(tree: ComponentProps, path: string[], node: ComponentProps, patch: boolean) {
    if (!tree.children) {
        tree.children = [];
    }

    const childIndex = tree.children.findIndex((child) => child.key === path[1]);

    if (path.length === 2) {
        if (childIndex >= 0) {
            if (patch) {
                patchComponentProps(tree.children[childIndex], node);
            } else {
                tree.children[childIndex] = {...node};
            }
        } else {
            tree.children.push({...node});
        }

        return;
    }

    if (childIndex === -1) {
        throw new Error(`Path ${path.join('.')} not found: no child with key ${path[1]}`);
    }

    updateTree(tree.children[childIndex], path.slice(1), node, patch);
}

export function appReducer(state: AppState, action: {path: string[]; node: ComponentProps; patch?: boolean}) {
    if (!state.tree) {
        state.tree = action.node;
    }

    const path = action.path?.length ? action.path : [action.node.key || ''];

    if (path.length === 1) {
        state.tree = action.node;
    } else {
        updateTree(state.tree, path, action.node, action.patch || false);
    }

    function populatePath(node: ComponentProps, parentPath: string[] = []) {
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
