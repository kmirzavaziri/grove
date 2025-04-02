import { match, MatchFunction, compile } from 'path-to-regexp';

export type Params = { [key: string]: string };

let matchers: [MatchFunction<any>, string][] = [];
let nodePathToPattern: Map<string, string> = new Map();

export function initRouting(routes: Map<string, string>) {
    matchers = Array.from(routes.entries()).map(
        ([pattern, nodePath]) => [match(pattern, { decode: decodeURIComponent }), nodePath]
    );
    nodePathToPattern = new Map(Array.from(routes.entries()).map(([pattern, nodePath]) => [nodePath, pattern]));
}

export function resolvePath(path: string): { nodePath?: string[]; params: Params } {
    for (const [matcher, nodePath] of matchers) {
        const result = matcher(path);
        if (result) {
            return { nodePath: nodePath.split("/"), params: result.params as Params };
        }
    }
    return { params: {} };
}

export function reverseRoute(nodePath: string[], params: Params): string {
    const pattern = nodePathToPattern.get(nodePath.join("/"));
    if (!pattern) {
        throw new Error(`No pattern found for nodePath: ${nodePath}`);
    }
    const toPath = compile(pattern);
    return toPath(params);
}