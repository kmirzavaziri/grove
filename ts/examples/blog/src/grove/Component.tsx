import React from "react";

export const componentRegistry = new Map<string, React.ComponentType<any>>();

export interface ComponentProps {
    type: string;
    key: string;
    role: string;
    data: { [key: string]: any };
    children: ComponentProps[];
}

export const Component: React.FC<{ props: ComponentProps }> = ({props}) => {
    const ComponentFn = componentRegistry.get(props.type);
    if (ComponentFn === undefined) {
        return null;
    }

    return <ComponentFn children={props.children} {...props.data}/>;
};

