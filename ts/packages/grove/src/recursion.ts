import {ComponentProps, patchComponentProps} from './Component';


export function flatten(node: ComponentProps): ComponentProps[] {
//     TODO implement recursive
}

export function getNodeAt(node: ComponentProps, path: string[]): ComponentProps {
    // TODO maybe keep a flat map that has node path to node, and access easier

    if (!node.children) {
        node.children = [];
    }

    const childIndex = node.children.findIndex((child) => child.key === path[0]);

    if (childIndex === -1) {
        throw new Error(
            `not found: node ${node.path} has no child with key ${path[0]}, available keys:` +
            node.children.map((child) => child.key).join(', '),
        );
    }

    if (path.length === 1) {
        return node.children[childIndex];
    }

    return getNodeAt(node.children[childIndex], path.slice(1));
}

export function setNodeAt(node: ComponentProps, path: string[], newNode: ComponentProps, patch: boolean): void {
    // TODO maybe keep a flat map that has node path to node, and access easier
    // TODO can we somehow reuse getNodeAt, and set the value into it here?

    if (!node.children) {
        node.children = [];
    }

    const childIndex = node.children.findIndex((child) => child.key === path[0]);

    if (childIndex === -1) {
        throw new Error(
            `not found: node ${node.path} has no child with key ${path[0]}, available keys:` +
            node.children.map((child) => child.key).join(', '),
        );
    }

    if (path.length === 1) {
        if (patch) {
            patchComponentProps(node.children[childIndex], newNode);
        } else {
            node.children[childIndex] = {...newNode};
        }

        return;
    }

    setNodeAt(node.children[childIndex], path.slice(1), newNode, patch);
}
