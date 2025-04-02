import {ComponentProps} from "../grove/Component";
import {reverseRoute} from "./routing";
import {AppContextValue} from "../grove/app-state";

export interface ARenderProps {
    node_path: string[];
    request?: any;
    node?: ComponentProps;
    patch: boolean;
    update_history?: boolean;
}

export function ARender(appContextValue: AppContextValue, props: ARenderProps) {
    if (props.update_history) {
        window.history.pushState({}, "", reverseRoute(props.node_path, props.request));
    }

    if (props.node) {
        return appContextValue.render({path: props.node_path, node: props.node, patch: props.patch});
    }

    appContextValue.apiHandlers.fetch(props.node_path, props.request)
        .then((node) => appContextValue.render({path: props.node_path, node, patch: props.patch}));
}