import {ComponentProps} from './Component';

export function flatten(node: ComponentProps): Map<string, ComponentProps> {
    const map = new Map<string, ComponentProps>();

    function traverse(n: ComponentProps) {
        map.set(n.path?.join('/') || '', n);
        n.children?.forEach(traverse);
    }

    traverse(node);

    return map;
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