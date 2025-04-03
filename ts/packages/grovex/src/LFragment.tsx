import React from 'react';

import type {ComponentProps} from '../../grove/src/Component';
import {Component} from '../../grove/src/Component';

export interface LFragment {
    children: ComponentProps[];
}

export const LFragment: React.FC<LFragment> = (props) => (
    <>
        {props.children.map((child) => (
            <Component key={child.key} props={child} />
        ))}
    </>
);
