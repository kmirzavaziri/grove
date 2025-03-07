import React from "react";
import Button from "@mui/material/Button";
import ListItem from "@mui/material/ListItem";
import ListItemButton from "@mui/material/ListItemButton";
import ListItemText from "@mui/material/ListItemText";
import ListItemIcon from "@mui/material/ListItemIcon";
import {ButtonVariant, Color, Size} from "./enums";
import {DynamicIcon} from "./DynamicIcon";

export interface DButton {
    text: string

    color?: Color
    variant?: ButtonVariant
    full_width: boolean
    disabled: boolean
    selected: boolean

    start_icon?: string
    end_icon?: string
    size?: Size

}

export const DButton: React.FC<DButton> = (props) => (
    <>
        {(props.variant === 'item' && <ListItem disablePadding sx={{display: 'block'}}>
            <ListItemButton selected={props.selected}>
                {(props.start_icon && <ListItemIcon><DynamicIcon iconName={props.start_icon}/></ListItemIcon>)}
                <ListItemText primary={props.text}/>
                {(props.end_icon && <ListItemIcon><DynamicIcon iconName={props.end_icon}/></ListItemIcon>)}
            </ListItemButton>
        </ListItem>)}
        {(props.variant !== 'item' && <Button
            variant={props.variant}
            fullWidth={props.full_width}
            startIcon={props.start_icon ? <DynamicIcon iconName={props.start_icon}/> : null}
            endIcon={props.end_icon ? <DynamicIcon iconName={props.end_icon}/> : null}
        >
            {props.text}
        </Button>)}
    </>
)

