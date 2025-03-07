import React from "react";
// TODO do not use grovex from grove // TODO add allowed import DAG as a lint rule to both ts & go
import {DTypography} from "../grovex/DTypography";

export const componentRegistry = new Map<string, React.ComponentType<any>>();

export interface ComponentProps {
    type: string;
    key: string;
    role?: string;
    props?: { [key: string]: any };
    children?: ComponentProps[];
}

export const Component: React.FC<{ props: ComponentProps }> = ({props}) => {
    const ComponentFn = componentRegistry.get(props.type);
    if (ComponentFn === undefined) {
        return <DTypography text={`Component ${props.type} not found`} color="error"/>;
    }

    return <ComponentFn key={props.key} children={props.children} {...props.props}/>;
};

