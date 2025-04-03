import React from "react";
import {FormLabel, Grid, OutlinedInput} from "@mui/material";
import {Input} from "../grove/input";
import {ITextVariant} from "./variants";

export interface ITextProps {
    input: Input;
    label: string;
    variant: ITextVariant;
    placeholder: string;
    auto_complete: string;
}

export const IText: React.FC<ITextProps> = (props) => {
    return (<Grid
        sx={{
            display: 'flex',
            flexDirection: 'column',
            my: 1,
        }}
        size={{xs: 12}}
    >
        <FormLabel htmlFor={props.input.key}>
            {props.label}
        </FormLabel>
        <OutlinedInput
            id={props.input.key}
            name={props.input.key}
            type={props.variant}  // TODO add switch for other types
            placeholder={props.placeholder}
            autoComplete={props.auto_complete}
        />
    </Grid>);
};