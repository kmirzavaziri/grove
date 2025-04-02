import React from "react";
import {Component, ComponentProps} from "../grove/Component";
import {ActionProps, perform} from "../grove/action";
import {useAppContext} from "../grove/app-state";
import {Box} from "@mui/material";

export interface XClickable {
    action: ActionProps;
    children: ComponentProps[];
}

export const XClickable: React.FC<XClickable> = (props) => {
    const appContextValue = useAppContext();
    return <Box onClick={() => perform(appContextValue, props.action)} sx={{ cursor: 'pointer' }}>
        {props.children.map(child => (<Component key={child.key} props={child}/>))}
    </Box>
};
