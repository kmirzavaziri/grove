import {ComponentProps} from "../../grove/src/Component";

export function findComponentByRole(components: ComponentProps[], role: string): ComponentProps | undefined {
    return components.find(component => component.role === role);
}