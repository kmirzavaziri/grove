import {AppContextValue, ComponentProps, modifyComponentProps, reverseRoute} from '@grove/grove';
import type {Struct} from '@grove/grove/src/value';

export interface ARenderProps {
    node_path: string[];
    request?: Struct;
    node?: ComponentProps;
    patch: boolean;
    update_history?: boolean;
}

export function ARender(appContextValue: AppContextValue, props: ARenderProps): void {
    if (props.update_history) {
        window.history.pushState({}, '', reverseRoute(props.node_path, props.request || {}));
    }

    if (props.node) {
        return appContextValue.dispatch({path: props.node_path, modify: modifyComponentProps(props.node, props.patch)});
    }

    appContextValue.apiHandlers
        .fetch(props.node_path, props.request)
        .then((node) => appContextValue.dispatch({
            path: props.node_path,
            node: modifyComponentProps(node, props.patch),
        }));
}
