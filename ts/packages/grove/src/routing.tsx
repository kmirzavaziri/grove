import type {MatchFunction} from 'path-to-regexp';
import {match, compile} from 'path-to-regexp';

import type {Struct} from './value';

let matchers: [MatchFunction<Struct>, string][] = [];
let nodePathToPattern: Map<string, string> = new Map();

export function initRouting(routes: Map<string, string>): void {
    matchers = Array.from(routes.entries()).map(([pattern, nodePath]) => [
        match(pattern, {decode: decodeURIComponent}),
        nodePath,
    ]);
    nodePathToPattern = new Map(Array.from(routes.entries()).map(([pattern, nodePath]) => [nodePath, pattern]));
}

export function resolvePath(path: string): {nodePath?: string[]; params: Struct} {
    for (const [matcher, nodePath] of matchers) {
        const result = matcher(path);
        if (result) {
            return {nodePath: nodePath.split('/'), params: result.params as Struct};
        }
    }
    return {params: {}};
}

export function reverseRoute(nodePath: string[], params: Struct): string {
    const pattern = nodePathToPattern.get(nodePath.join('/'));
    if (!pattern) {
        throw new Error(`No pattern found for nodePath: ${nodePath}`);
    }
    const toPath = compile(pattern);
    return toPath(params);
}
