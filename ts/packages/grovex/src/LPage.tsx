import MenuRoundedIcon from '@mui/icons-material/MenuRounded';
import AppBar from '@mui/material/AppBar';
import Badge, {badgeClasses} from '@mui/material/Badge';
import Box from '@mui/material/Box';
import Divider from '@mui/material/Divider';
import Drawer, {drawerClasses} from '@mui/material/Drawer';
import IconButton from '@mui/material/IconButton';
import Stack from '@mui/material/Stack';
import {tabsClasses} from '@mui/material/Tabs';
import Toolbar from '@mui/material/Toolbar';
import React, {useEffect} from 'react';

import type {ComponentProps} from '../../grove/src/Component';
import {Component} from '../../grove/src/Component';

import {DDivider} from './DDivider';
import {findComponentByRole} from './role';

export interface LPageProps {
    title: string;
    children: ComponentProps[];
}

export const LPage: React.FC<LPageProps> = (props) => {
    useEffect(() => {
        document.title = props.title;
    }, []);

    const [open, setOpen] = React.useState(false);

    const toggleDrawer = (newOpen: boolean) => () => {
        setOpen(newOpen);
    };

    const main = findComponentByRole(props.children, 'main');
    const sidebarStart = findComponentByRole(props.children, 'sidebar_start');
    const sidebarEnd = findComponentByRole(props.children, 'sidebar_end');
    const sidebarStartXS = findComponentByRole(props.children, 'sidebar_start_xs');
    const sidebarEndXS = findComponentByRole(props.children, 'sidebar_end_xs');
    const navbarStart = findComponentByRole(props.children, 'navbar_start');
    const navbarEnd = findComponentByRole(props.children, 'navbar_end');
    const navbarStartXS = findComponentByRole(props.children, 'navbar_start_xs');
    const navbarEndXS = findComponentByRole(props.children, 'navbar_end_xs');

    return (
        <Box sx={{display: 'flex'}}>
            <Drawer
                variant="permanent"
                sx={{
                    width: 240,
                    flexShrink: 0,
                    boxSizing: 'border-box',
                    mt: 10,
                    display: {xs: 'none', md: 'block'},
                    [`& .${drawerClasses.paper}`]: {
                        px: 2,
                        width: 240,
                        boxSizing: 'border-box',
                        backgroundColor: 'background.paper',
                    },
                }}
            >
                <Box
                    sx={{
                        height: 'calc(var(--template-frame-height, 0px) + 4px)',
                    }}
                />
                {sidebarStart && (
                    <Box
                        sx={{
                            overflow: 'auto',
                            height: '100%',
                            display: 'flex',
                            flexDirection: 'column',
                        }}
                    >
                        <Component props={sidebarStart} />
                    </Box>
                )}
                {sidebarEnd && (
                    <>
                        <DDivider />
                        <Component props={sidebarEnd} />
                    </>
                )}
            </Drawer>
            <AppBar
                position="fixed"
                sx={{
                    display: {xs: 'auto', md: 'none'},
                    boxShadow: 0,
                    bgcolor: 'background.paper',
                    backgroundImage: 'none',
                    borderBottom: '1px solid',
                    borderColor: 'divider',
                    top: 'var(--template-frame-height, 0px)',
                }}
            >
                <Toolbar
                    sx={{
                        width: '100%',
                        padding: '12px',
                        display: 'flex',
                        flexDirection: 'column',
                        alignItems: 'start',
                        justifyContent: 'center',
                        gap: '12px',
                        flexShrink: 0,
                        [`& ${tabsClasses.flexContainer}`]: {
                            gap: '8px',
                            p: '8px',
                            pb: 0,
                        },
                    }}
                    variant="regular"
                >
                    <Stack
                        direction="row"
                        sx={{
                            alignItems: 'center',
                            flexGrow: 1,
                            width: '100%',
                            gap: 2,
                        }}
                    >
                        <Stack direction="row" spacing={1} sx={{justifyContent: 'center', mr: 'auto'}}>
                            {navbarStartXS && <Component props={navbarStartXS} />}
                        </Stack>
                        {navbarEndXS && <Component props={navbarEndXS} />}
                        <Badge
                            color="error"
                            variant="dot"
                            invisible={true}
                            sx={{[`& .${badgeClasses.badge}`]: {right: 2, top: 2}}}
                        >
                            <IconButton size="small" aria-label="menu" onClick={toggleDrawer(true)}>
                                <MenuRoundedIcon />
                            </IconButton>
                        </Badge>
                        <Drawer
                            anchor="right"
                            open={open}
                            onClose={toggleDrawer(false)}
                            sx={{
                                zIndex: (theme) => theme.zIndex.drawer + 1,
                                [`& .${drawerClasses.paper}`]: {
                                    backgroundImage: 'none',
                                    backgroundColor: 'background.paper',
                                    px: 2,
                                },
                            }}
                        >
                            <Stack
                                sx={{
                                    maxWidth: '70dvw',
                                    height: '100%',
                                }}
                            >
                                {sidebarStartXS && (
                                    <Box
                                        sx={{
                                            overflow: 'auto',
                                            height: '100%',
                                            display: 'flex',
                                            flexDirection: 'column',
                                        }}
                                    >
                                        <Component props={sidebarStartXS} />
                                    </Box>
                                )}
                                {sidebarEndXS && (
                                    <>
                                        <Divider />
                                        <Component props={sidebarEndXS} />
                                    </>
                                )}
                            </Stack>
                        </Drawer>
                    </Stack>
                </Toolbar>
            </AppBar>
            <Box component="main" sx={{flexGrow: 1, overflow: 'auto'}}>
                <Stack
                    spacing={2}
                    sx={{
                        alignItems: 'center',
                        mx: 3,
                        pb: 5,
                        mt: {xs: 8, md: 0},
                    }}
                >
                    <Stack
                        direction="row"
                        sx={{
                            display: {xs: 'none', md: 'flex'},
                            width: '100%',
                            alignItems: {xs: 'flex-start', md: 'center'},
                            justifyContent: 'space-between',
                            maxWidth: {sm: '100%', md: '1700px'},
                            pt: 1.5,
                        }}
                        spacing={2}
                    >
                        {navbarStart && <Component props={navbarStart} />}
                        <Stack direction="row" sx={{gap: 1}}>
                            {navbarEnd && <Component props={navbarEnd} />}
                        </Stack>
                    </Stack>
                    <Box sx={{width: '100%', maxWidth: {sm: '100%', md: '1700px'}}}>
                        {main && <Component props={main} />}
                    </Box>
                </Stack>
            </Box>
        </Box>
    );
};
