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

    const modify = (newNode: ComponentProps) => (node: ComponentProps) => {
        modifyComponentProps(node, newNode, props.patch);
    };

    if (props.node) {
        return appContextValue.dispatch({path: props.node_path, node: modify(props.node)});
    }

    appContextValue.apiHandlers
        .fetch(props.node_path, props.request)
        .then((node) => appContextValue.dispatch({path: props.node_path, node: modify(node)}));
}
