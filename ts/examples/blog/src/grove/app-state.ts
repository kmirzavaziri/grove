import { createContext, useContext, Dispatch } from "react";
import { ComponentProps } from "./Component";

export interface AppState {
    tree: ComponentProps | null;
}

interface AppContextValue {
    state: AppState;
    dispatch: Dispatch<{ nodePath: string[]; node: ComponentProps }>;
}

const AppContext = createContext<AppContextValue | undefined>(undefined);

export function useAppContext() {
    const context = useContext(AppContext);
    if (!context) throw new Error("useAppContext must be used within an AppProvider");
    return context;
}

function updateTree(tree: ComponentProps, path: string[], node: ComponentProps): ComponentProps {
    if (path.length === 0) return node;
    const [head, ...tail] = path;
    const children = (tree.children || []).map(child =>
        child.key === head ? updateTree(child, tail, node) : { ...child }
    );
    return { ...tree, children };
}

export function appReducer(state: AppState, action: { nodePath: string[]; node: ComponentProps }) {
    if (!state.tree || action.nodePath.length === 0) return { tree: action.node };
    return { tree: updateTree(state.tree, action.nodePath, action.node) };
}

export { AppContext };
