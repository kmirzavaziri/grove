import React from "react";
import {Component, ComponentProps} from "../grove/Component";
import Box from "@mui/material/Box";

export interface LBox {
    children: ComponentProps[]
}

export const LBox: React.FC<LBox> = (props) => (
    <Box>
        {props.children.map(child => (<Component key={child.key} props={child}/>))}
    </Box>
);
