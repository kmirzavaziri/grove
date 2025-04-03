import {Typography} from '@mui/material';
import React from 'react';

import type {Align, Color, DTypographyVariant} from './variants';
import {align2MUI} from './variants';

export interface DTypographyProps {
    text: string;
    align?: Align;
    color?: Color;
    variant?: DTypographyVariant;
}

export const DTypography: React.FC<DTypographyProps> = (props) => (
    <Typography align={align2MUI(props.align)} color={props.color} variant={props.variant} sx={{py: 1}}>
        {props.text}
    </Typography>
);
