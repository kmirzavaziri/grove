import {ComponentProps} from "../grove/Component";
import {useAppContext} from "../grove/app-state";
import {reverseRoute} from "./routing";

export type Action =
    | {
    type: "grovex.Replace";
    payload: { nodePath: string[]; request?: any; node?: ComponentProps; updateHistory?: boolean }
}
    | { type: "grovex.Submit"; payload: { nodePath: string[] } };

export function useActions(groveServer: string) {
    const {dispatch} = useAppContext();

    function perform(action: Action) {
        switch (action.type) {
            case "grovex.Replace": {
                const {nodePath, request, node, updateHistory} = action.payload;
                if (updateHistory) window.history.pushState({}, "", reverseRoute(nodePath.join("/"), request));
                if (node) return dispatch({nodePath, node});

                fetch(`${groveServer}/fetch/${nodePath.join("/")}`, {
                    method: "POST",
                    headers: {"Content-Type": "application/json"},
                    body: JSON.stringify(request),
                })
                    .then((resp) => {
                        if (!resp.ok) throw new Error("Grove API returned error");
                        return resp.json();
                    })
                    .then((node) => dispatch({nodePath, node}));
                break;
            }
            case "grovex.Submit": {
                // TODO implement
                break;
            }
        }
    }

    return perform;
}