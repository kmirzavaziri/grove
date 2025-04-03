import Button from '@mui/material/Button';
import ListItem from '@mui/material/ListItem';
import ListItemButton from '@mui/material/ListItemButton';
import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';
import React from 'react';

import {DynamicIcon} from './DynamicIcon';
import type {DButtonVariant, Color, Size} from './variants';

export interface DButton {
    text: string;

    color?: Color;
    variant?: DButtonVariant;
    full_width: boolean;
    disabled: boolean;
    selected: boolean;

    start_icon?: string;
    end_icon?: string;
    size?: Size;
}

export const DButton: React.FC<DButton> = (props) => (
    <>
        {props.variant === 'item' && (
            <ListItem disablePadding sx={{display: 'block'}}>
                <ListItemButton selected={props.selected}>
                    {props.start_icon && (
                        <ListItemIcon>
                            <DynamicIcon iconName={props.start_icon} />
                        </ListItemIcon>
                    )}
                    <ListItemText primary={props.text} />
                    {props.end_icon && (
                        <ListItemIcon>
                            <DynamicIcon iconName={props.end_icon} />
                        </ListItemIcon>
                    )}
                </ListItemButton>
            </ListItem>
        )}
        {props.variant !== 'item' && (
            <Button
                variant={props.variant}
                fullWidth={props.full_width}
                startIcon={props.start_icon ? <DynamicIcon iconName={props.start_icon} /> : null}
                endIcon={props.end_icon ? <DynamicIcon iconName={props.end_icon} /> : null}
                sx={{my: 1}}
            >
                {props.text}
            </Button>
        )}
    </>
);
