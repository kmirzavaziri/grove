import React from "react";
import {Align, align2MUI, Color, TypographyVariant} from "./enums";
import {Typography} from "@mui/material";


export interface DTypographyProps {
    text: string
    align?: Align
    color?: Color
    variant?: TypographyVariant
}

export const DTypography: React.FC<DTypographyProps> = (props) => (
    <Typography
        align={align2MUI(props.align)}
        color={props.color}
        variant={props.variant}
        sx={{py: 1}}
    >
        {props.text}
    </Typography>
);
