import {ComponentProps} from './Component';

export function getNodeAt(node: ComponentProps, path: string[]): ComponentProps {
    if (path.length === 0) {
        return node;
    }

    const children = (node.children ?? []);

    const childIndex = children.findIndex((child) => child.key === path[0]);

    if (childIndex === -1) {
        throw new Error(
            `not found: node ${node.path} has no child with key ${path[0]}, available keys:` +
            children.map((child) => child.key).join(', '),
        );
    }

    return getNodeAt(children[childIndex], path.slice(1));
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



