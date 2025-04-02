import React, {useEffect, useState} from "react";
import {Component, ComponentProps} from "../grove/Component";
import {Box, FormLabel, Grid, Modal, OutlinedInput, TextField, Typography} from "@mui/material";
import {useAppContext} from "../grove/app-state";
import {ARender} from "./ARender";
import {DTypography} from "./DTypography";
import {styled} from "@mui/material/styles";

const FormGrid = styled(Grid)(() => ({
    display: 'flex',
    flexDirection: 'column',
}));

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
                    border: '1px solid',
                    borderColor: "divider",
                    borderRadius: 1,
                    boxShadow: 24,
                    p: 4,
                }}
            >
                <FormGrid size={{ xs: 12, md: 6 }}>

                <FormLabel htmlFor="first-name" required>
                    First name
                </FormLabel>
                <OutlinedInput
                    id="first-name"
                    name="first-name"
                    type="name"
                    placeholder="John"
                    autoComplete="first name"
                    required
                    size="small"
                />
                </FormGrid>

                {/*{props.children.map((child) => (*/}
                {/*    <Component key={child.key} props={child}/>*/}
                {/*))}*/}
            </Box>
        </Modal>
    );
};