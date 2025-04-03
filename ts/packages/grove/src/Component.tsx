import React from 'react';

import type {Input} from './input';
import type {Struct} from './value';

export const componentRegistry = new Map<string, React.ComponentType<Struct>>();

export interface ComponentProps {
    type?: string;
    key?: string;
    role?: string;
    props?: Struct;
    input?: Input;
    children?: ComponentProps[];

    path?: string[];
}

export const modifyComponentProps = (src: ComponentProps, patch: boolean = false) => (dest: ComponentProps): void => {
    if (patch) {
        patchComponentProps(dest, src);
    } else {
        overrideComponentProps(dest, src);
    }
};

export const patchComponentProps = (dest: ComponentProps, src: ComponentProps): void => {
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

    if (src.input) {
        dest.input = {...dest.input, ...src.input};
    }

    if (src.children?.length) {
        dest.children = src.children;
    }
};

export const overrideComponentProps = (dest: ComponentProps, src: ComponentProps): void => {
    dest.type = src.type;
    dest.key = src.key;
    dest.role = src.role;
    dest.props = src.props;
    dest.input = src.input;
    dest.children = src.children;
};

export const Component: React.FC<{props: ComponentProps}> = ({props}) => {
    const np = {
        type: props.type ?? '',
        key: props.key ?? '',
        path: props.path ?? [],
        role: props.role,
        props: props.props,
        input: props.input,
        children: props.children,
    };

    const ComponentFn = componentRegistry.get(np.type);
    if (ComponentFn === undefined) {
        return <div style="color: red;">Component ${np.type} not found</div>;
    }

    return (
        <ComponentFn key={np.key} path={np.path} children={np.children} input={np.input} {...np.props} />
    );
};
