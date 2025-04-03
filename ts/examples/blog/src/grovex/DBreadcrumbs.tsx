import * as React from 'react';
import { styled } from '@mui/material/styles';
import Typography from '@mui/material/Typography';
import MuiBreadcrumbs, { breadcrumbsClasses } from '@mui/material/Breadcrumbs';
import NavigateNextRoundedIcon from '@mui/icons-material/NavigateNextRounded';
import {XClickable} from "./XClickable";
import {useAppContext} from "../grove/app-state";
import {ActionProps, perform} from "../grove/action";

const Breadcrumbs = styled(MuiBreadcrumbs)(({ theme }) => ({
    margin: theme.spacing(1, 0),
    [`& .${breadcrumbsClasses.separator}`]: {
        color: (theme.vars || theme).palette.action.disabled,
        margin: 1,
    },
    [`& .${breadcrumbsClasses.ol}`]: {
        alignItems: 'center',
    },
}));

export interface DBreadcrumbsProps {
    items: DBreadcrumbsItem[];
}

export interface DBreadcrumbsItem {
    title: string;
    action: ActionProps;
}

export const DBreadcrumbs: React.FC<DBreadcrumbsProps> = (props) => {
    const appContextValue = useAppContext();
    return (
        <Breadcrumbs
            aria-label="breadcrumb"
            separator={<NavigateNextRoundedIcon fontSize="small" />}
        >
            {props.items.map((item, index) => (
                <Typography
                    key={index}
                    variant="body1"
                    sx={{
                        ...(index === props.items.length - 1 && {
                            color: 'text.primary',
                            fontWeight: 600
                        }),
                        cursor: 'pointer',
                    }}
                    onClick={() => perform(appContextValue, item.action)}
                >
                    {item.title}
                </Typography>
            ))}
        </Breadcrumbs>
    );
}
