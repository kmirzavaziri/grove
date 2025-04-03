import {AppContextValue} from "./app-state";

type Action<T> = (appContextValue: AppContextValue, payload: T) => void

export const actionRegistry = new Map<string, Action<any>>();

export interface ActionProps {
    type: string;
    payload: { [key: string]: any };
}

export function perform(appContextValue: AppContextValue, props: ActionProps) {
    const action = actionRegistry.get(props.type);
    if (action) {
        action(appContextValue, props.payload);
    }
}
