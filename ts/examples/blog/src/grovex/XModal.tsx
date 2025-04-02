import React, {useEffect, useState} from "react";
import {Component, ComponentProps} from "../grove/Component";
import {Box, Modal} from "@mui/material";
import {useAppContext} from "../grove/app-state";
import {ARender} from "./ARender";

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

    const handleClose = () => {
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
        >
            <Box
                sx={{
                    position: "absolute",
                    top: "50%",
                    left: "50%",
                    transform: "translate(-50%, -50%)",
                    width: 500,
                    bgcolor: "background.paper",
                    border: "2px solid #000",
                    boxShadow: 24,
                    p: 4,
                }}
            >
                {props.children.map((child) => (
                    <Component key={child.key} props={child}/>
                ))}
            </Box>
        </Modal>
    );
};