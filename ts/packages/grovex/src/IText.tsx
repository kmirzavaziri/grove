import {FormLabel, Grid, OutlinedInput} from '@mui/material';
import React from 'react';

import type {Input} from '@grove/grove';

import type {ITextVariant} from './variants';
import {ComponentProps, useAppContext} from '@grove/grove';

export interface ITextProps {
    path: string;
    input: Input;
    label: string;
    variant: ITextVariant;
    placeholder: string;
    auto_complete: string;
}

export const IText: React.FC<ITextProps> = (props) => {
    const {dispatch} = useAppContext();

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        dispatch({
            path: props.path, modify: (node: ComponentProps) => {
                if (node.input) {
                    node.input.value = e.target.value;
                }
            },
        });
    };

    return (
        <Grid sx={{display: 'flex', flexDirection: 'column', my: 1}} size={{xs: 12}}>
            <FormLabel htmlFor={props.input.key}>{props.label}</FormLabel>
            <OutlinedInput
                id={props.input.key}
                name={props.input.key}
                type={props.variant}
                placeholder={props.placeholder}
                autoComplete={props.auto_complete}
                value={props.input.value || ''}
                onChange={handleChange}
            />
        </Grid>
    );
};
