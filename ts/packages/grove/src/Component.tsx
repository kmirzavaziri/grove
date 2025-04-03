import React from 'react';

// TODO do not use grovex from grove // TODO add allowed import DAG as a lint rule to both ts & go
import {DTypography} from '../../grovex/src/DTypography';

import type {Input} from './input';

export const componentRegistry = new Map<string, React.ComponentType<any>>();

export interface ComponentProps {
    type?: string;
    key?: string;
    path?: string[];
    role?: string;
    props?: {[key: string]: any};
    input?: Input;
    children?: ComponentProps[];
}

export const Component: React.FC<{props: ComponentProps}> = ({props}) => {
    if (props.type === undefined) {
        props.type = '';
    }

    if (props.key === undefined) {
        props.key = '';
    }

    if (props.path === undefined) {
        props.path = [];
    }

    const ComponentFn = componentRegistry.get(props.type);
    if (ComponentFn === undefined) {
        return <DTypography text={`Component ${props.type} not found`} color="error" />;
    }

    return (
        <ComponentFn key={props.key} path={props.path} children={props.children} input={props.input} {...props.props} />
    );
};

export const patchComponentProps = (dest: ComponentProps, src: ComponentProps) => {
    if (src.type) {
        dest.type = src.type;
    }
    if (src.key) {
        dest.key = src.key;
    }
    if (src.role) {
        dest.role = src.role;
    }
    dest.props = {...dest.props, ...src.props};
    if (src.children?.length) {
        dest.children = src.children;
    }
};
