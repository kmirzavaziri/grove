import Box from '@mui/material/Box';
import React from 'react';

import type {ComponentProps} from '../../grove/src/Component';
import {Component} from '../../grove/src/Component';

export interface LBox {
    children: ComponentProps[];
}

export const LBox: React.FC<LBox> = (props) => (
    <Box>
        {props.children.map((child) => (
            <Component key={child.key} props={child} />
        ))}
    </Box>
);
