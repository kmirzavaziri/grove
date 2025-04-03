import {ComponentProps, AppContextValue, perform, Input, flatten} from '@grove/grove';
import {reverseRoute} from '@grove/grove';
import type {Struct} from '@grove/grove/src/value';

export interface ASubmitProps {
    node_path: string[];
    request?: Struct;
}

export function ASubmit(appContextValue: AppContextValue, props: ASubmitProps): void {
    console.log("here")

    // TODO state.tree might be null
    const inputs = extractInputsFromNode(getNodeAt(appContextValue.state.tree!, props.node_path));

    const payload = {}; // TODO gather user input values

    // TODO run validations

    appContextValue.apiHandlers
        .submit(props.node_path, {...payload, ...props.request})
        .then((action) => perform(appContextValue, action));
}



function extractInputsFromNode(node: ComponentProps): Input[] {
    flatten(node);
    // TODO get inputs and return
    return [];
}