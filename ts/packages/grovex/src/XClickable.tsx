import {Box} from '@mui/material';
import React from 'react';

import type {ActionProps} from '../../grove/src/action';
import {perform} from '../../grove/src/action';
import {useAppContext} from '../../grove/src/app-state';
import {Component} from '../../grove/src/Component';
import type {ComponentProps} from '../../grove/src/Component';

export interface XClickable {
    action: ActionProps;
    children: ComponentProps[];
}

export const XClickable: React.FC<XClickable> = (props) => {
    const appContextValue = useAppContext();
    return (
        <Box onClick={() => perform(appContextValue, props.action)} sx={{cursor: 'pointer'}}>
            {props.children.map((child) => (
                <Component key={child.key} props={child} />
            ))}
        </Box>
    );
};
