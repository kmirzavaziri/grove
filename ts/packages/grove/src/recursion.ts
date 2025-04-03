import {ComponentProps} from './Component';


export function flatten(node: ComponentProps): ComponentProps[] {
//     TODO implement recursive
}

export function getNodeAt(node: ComponentProps, path: string[]): ComponentProps {
    // TODO maybe keep a flat map that has node path to node, and access easier
    if (path.length === 0) {
        return node;
    }

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

export function updatePaths(node: ComponentProps, parentPath: string[] = []): void {
    if (node.key === undefined) {
        node.key = '';
    }

    if (node.children === undefined) {
        node.children = [];
    }

    node.path = [...parentPath, node.key];

    for (const child of node.children) {
        updatePaths(child, node.path);
    }
}