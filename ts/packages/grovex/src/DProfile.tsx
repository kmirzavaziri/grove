import Stack from "@mui/material/Stack";
import Avatar from "@mui/material/Avatar";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import React from "react";

export interface DProfileProps {
    image: string;
    title: string;
    subtitle: string;
}

export const DProfile: React.FC<DProfileProps> = (props) => {
    return (
        <>
            <Stack
                direction="row"
                sx={{
                    display: {xs: 'none', md: 'flex'},
                    py: 2,
                    gap: 1,
                    alignItems: 'center',
                    mt: 'calc(var(--template-frame-height, 0px) + 4px)',
                }}
            >
                <Avatar sizes="small" alt={props.title} src={props.image} sx={{width: 36, height: 36}} />
                <Box sx={{mr: 'auto'}}>
                    <Typography variant="body2" sx={{fontWeight: 500, lineHeight: '16px'}}>
                        {props.title}
                    </Typography>
                    <Typography variant="caption" sx={{color: 'text.secondary'}}>
                        {props.subtitle}
                    </Typography>
                </Box>
            </Stack>
            <Stack
                direction="row"
                sx={{
                    display: {xs: 'auto', md: 'none'},
                    py: 1,
                    gap: 1,
                }}
            >
                <Stack direction="row" sx={{gap: 1, alignItems: 'center', flexGrow: 1, p: 1}}>
                    <Avatar sizes="small" alt={props.title} src={props.image} sx={{width: 24, height: 24}} />
                    <Typography component="p" variant="h6">
                        {props.title}
                    </Typography>
                </Stack>
            </Stack>
        </>;
    );
};
