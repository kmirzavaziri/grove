import React from "react";
import {Box} from "@mui/material";
import {Component, ComponentProps} from "../../grove/src/Component";
import {ActionProps, perform} from "../../grove/src/action";
import {useAppContext} from "../../grove/src/app-state";

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
