import type {AppContextValue} from './app-state';
import type {Struct} from './value';

type Action<T> = (appContextValue: AppContextValue, payload: T) => void;

export const actionRegistry = new Map<string, Action<Struct>>();

export interface ActionProps {
    type: string;
    payload: Struct;
}

export function perform(appContextValue: AppContextValue, props: ActionProps): void {
    const action = actionRegistry.get(props?.type);
    if (action) {
        action(appContextValue, props.payload);
    }
}
