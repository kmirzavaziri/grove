import {Box, Modal} from '@mui/material';
import React, {useEffect, useState} from 'react';

import {useAppContext} from '../../grove/src/app-state';
import type {ComponentProps} from '../../grove/src/Component';
import {Component} from '../../grove/src/Component';

import {ARender} from './ARender';

export interface XModalProps {
    path: string[];
    title: string;
    description: string;
    open: boolean;
    children: ComponentProps[];
}

export const XModal: React.FC<XModalProps> = (props) => {
    const {state, render, apiHandlers} = useAppContext();
    const [open, setOpen] = useState(props.open ?? false);

    useEffect(() => {
        setOpen(props.open ?? false);
    }, [props.open]);

    const handleClose = (): void => {
        if (props.path) {
            ARender({state, render, apiHandlers}, {node_path: props.path, patch: true, node: {props: {open: false}}});
        }
    };

    return (
        <Modal
            open={open}
            onClose={handleClose}
            aria-labelledby={props.title}
            aria-describedby={props.description}
            keepMounted
            sx={{
                display: open ? 'block' : 'none',
            }}
        >
            <Box
                sx={{
                    position: 'absolute',
                    top: '50%',
                    left: '50%',
                    transform: 'translate(-50%, -50%)',
                    width: 400,
                    bgcolor: 'background.paper',
                    border: '1px solid',
                    borderColor: 'divider',
                    borderRadius: 1,
                    boxShadow: 24,
                    p: 3,
                }}
            >
                {props.children.map((child) => (
                    <Component key={child.key} props={child} />
                ))}
            </Box>
        </Modal>
    );
};
