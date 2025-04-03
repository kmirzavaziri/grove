import CheckBoxOutlineBlankRoundedIcon from '@mui/icons-material/CheckBoxOutlineBlankRounded';
import CheckRoundedIcon from '@mui/icons-material/CheckRounded';
import RemoveRoundedIcon from '@mui/icons-material/RemoveRounded';
import UnfoldMoreRoundedIcon from '@mui/icons-material/UnfoldMoreRounded';
import {buttonBaseClasses} from '@mui/material/ButtonBase';
import {checkboxClasses} from '@mui/material/Checkbox';
import {chipClasses} from '@mui/material/Chip';
import {dividerClasses} from '@mui/material/Divider';
import {iconButtonClasses} from '@mui/material/IconButton';
import {listClasses} from '@mui/material/List';
import {listItemIconClasses} from '@mui/material/ListItemIcon';
import {menuItemClasses} from '@mui/material/MenuItem';
import {outlinedInputClasses} from '@mui/material/OutlinedInput';
import {paperClasses} from '@mui/material/Paper';
import {selectClasses} from '@mui/material/Select';
import {alpha, createTheme, ThemeProvider} from '@mui/material/styles';
import type {SvgIconProps} from '@mui/material/SvgIcon';
import {svgIconClasses} from '@mui/material/SvgIcon';
import {tabClasses} from '@mui/material/Tab';
import {tablePaginationClasses} from '@mui/material/TablePagination';
import {toggleButtonClasses} from '@mui/material/ToggleButton';
import {toggleButtonGroupClasses} from '@mui/material/ToggleButtonGroup';
import {typographyClasses} from '@mui/material/Typography';
import {axisClasses, chartsGridClasses, legendClasses} from '@mui/x-charts';
import {gridClasses} from '@mui/x-data-grid';
import {pickersDayClasses, pickersMonthClasses, pickersYearClasses} from '@mui/x-date-pickers';
import * as React from 'react';

declare module '@mui/material/Paper' {
    interface PaperPropsVariantOverrides {
        highlighted: true;
    }
}

declare module '@mui/material/styles' {
    interface ColorRange {
        50: string;
        100: string;
        200: string;
        300: string;
        400: string;
        500: string;
        600: string;
        700: string;
        800: string;
        900: string;
    }

    interface Palette {
        baseShadow: string;
    }
}

const defaultTheme = createTheme();

const brand = {
    50: 'hsl(210, 100%, 95%)',
    100: 'hsl(210, 100%, 92%)',
    200: 'hsl(210, 100%, 80%)',
    300: 'hsl(210, 100%, 65%)',
    400: 'hsl(210, 98%, 48%)',
    500: 'hsl(210, 98%, 42%)',
    600: 'hsl(210, 98%, 55%)',
    700: 'hsl(210, 100%, 35%)',
    800: 'hsl(210, 100%, 16%)',
    900: 'hsl(210, 100%, 21%)',
};

const gray = {
    50: 'hsl(220, 35%, 97%)',
    100: 'hsl(220, 30%, 94%)',
    200: 'hsl(220, 20%, 88%)',
    300: 'hsl(220, 20%, 80%)',
    400: 'hsl(220, 20%, 65%)',
    500: 'hsl(220, 20%, 42%)',
    600: 'hsl(220, 20%, 35%)',
    700: 'hsl(220, 20%, 25%)',
    800: 'hsl(220, 30%, 6%)',
    900: 'hsl(220, 35%, 3%)',
};

const green = {
    50: 'hsl(120, 80%, 98%)',
    100: 'hsl(120, 75%, 94%)',
    200: 'hsl(120, 75%, 87%)',
    300: 'hsl(120, 61%, 77%)',
    400: 'hsl(120, 44%, 53%)',
    500: 'hsl(120, 59%, 30%)',
    600: 'hsl(120, 70%, 25%)',
    700: 'hsl(120, 75%, 16%)',
    800: 'hsl(120, 84%, 10%)',
    900: 'hsl(120, 87%, 6%)',
};

const orange = {
    50: 'hsl(45, 100%, 97%)',
    100: 'hsl(45, 92%, 90%)',
    200: 'hsl(45, 94%, 80%)',
    300: 'hsl(45, 90%, 65%)',
    400: 'hsl(45, 90%, 40%)',
    500: 'hsl(45, 90%, 35%)',
    600: 'hsl(45, 91%, 25%)',
    700: 'hsl(45, 94%, 20%)',
    800: 'hsl(45, 95%, 16%)',
    900: 'hsl(45, 93%, 12%)',
};

const red = {
    50: 'hsl(0, 100%, 97%)',
    100: 'hsl(0, 92%, 90%)',
    200: 'hsl(0, 94%, 80%)',
    300: 'hsl(0, 90%, 65%)',
    400: 'hsl(0, 90%, 40%)',
    500: 'hsl(0, 90%, 30%)',
    600: 'hsl(0, 91%, 25%)',
    700: 'hsl(0, 94%, 18%)',
    800: 'hsl(0, 95%, 12%)',
    900: 'hsl(0, 93%, 6%)',
};

const colorSchemes = {
    light: {
        palette: {
            primary: {
                light: brand[200],
                main: brand[400],
                dark: brand[700],
                contrastText: brand[50],
            },
            info: {
                light: brand[100],
                main: brand[300],
                dark: brand[600],
                contrastText: gray[50],
            },
            warning: {
                light: orange[300],
                main: orange[400],
                dark: orange[800],
            },
            error: {
                light: red[300],
                main: red[400],
                dark: red[800],
            },
            success: {
                light: green[300],
                main: green[400],
                dark: green[800],
            },
            grey: {
                ...gray,
            },
            divider: alpha(gray[300], 0.4),
            background: {
                default: 'hsl(0, 0%, 99%)',
                paper: 'hsl(220, 35%, 97%)',
            },
            text: {
                primary: gray[800],
                secondary: gray[600],
                warning: orange[400],
            },
            action: {
                hover: alpha(gray[200], 0.2),
                selected: `${alpha(gray[200], 0.3)}`,
            },
            baseShadow: 'hsla(220, 30%, 5%, 0.07) 0px 4px 16px 0px, hsla(220, 25%, 10%, 0.07) 0px 8px 16px -5px',
        },
    },
    dark: {
        palette: {
            primary: {
                contrastText: brand[50],
                light: brand[300],
                main: brand[400],
                dark: brand[700],
            },
            info: {
                contrastText: brand[300],
                light: brand[500],
                main: brand[700],
                dark: brand[900],
            },
            warning: {
                light: orange[400],
                main: orange[500],
                dark: orange[700],
            },
            error: {
                light: red[400],
                main: red[500],
                dark: red[700],
            },
            success: {
                light: green[400],
                main: green[500],
                dark: green[700],
            },
            grey: {
                ...gray,
            },
            divider: alpha(gray[700], 0.6),
            background: {
                default: gray[900],
                paper: 'hsl(220, 30%, 7%)',
            },
            text: {
                primary: 'hsl(0, 0%, 100%)',
                secondary: gray[400],
            },
            action: {
                hover: alpha(gray[600], 0.2),
                selected: alpha(gray[600], 0.3),
            },
            baseShadow: 'hsla(220, 30%, 5%, 0.7) 0px 4px 16px 0px, hsla(220, 25%, 10%, 0.8) 0px 8px 16px -5px',
        },
    },
};

const typography = {
    fontFamily: 'Inter, sans-serif',
    h1: {
        fontSize: defaultTheme.typography.pxToRem(48),
        fontWeight: 600,
        lineHeight: 1.2,
        letterSpacing: -0.5,
    },
    h2: {
        fontSize: defaultTheme.typography.pxToRem(36),
        fontWeight: 600,
        lineHeight: 1.2,
    },
    h3: {
        fontSize: defaultTheme.typography.pxToRem(30),
        lineHeight: 1.2,
    },
    h4: {
        fontSize: defaultTheme.typography.pxToRem(24),
        fontWeight: 600,
        lineHeight: 1.5,
    },
    h5: {
        fontSize: defaultTheme.typography.pxToRem(20),
        fontWeight: 600,
    },
    h6: {
        fontSize: defaultTheme.typography.pxToRem(18),
        fontWeight: 600,
    },
    subtitle1: {
        fontSize: defaultTheme.typography.pxToRem(18),
    },
    subtitle2: {
        fontSize: defaultTheme.typography.pxToRem(14),
        fontWeight: 500,
    },
    body1: {
        fontSize: defaultTheme.typography.pxToRem(14),
    },
    body2: {
        fontSize: defaultTheme.typography.pxToRem(14),
        fontWeight: 400,
    },
    caption: {
        fontSize: defaultTheme.typography.pxToRem(12),
        fontWeight: 400,
    },
};

const shape = {
    borderRadius: 8,
};

const shadows = ['none', 'var(--template-palette-baseShadow)', ...defaultTheme.shadows.slice(2)];

interface AppThemeProps {
    children: React.ReactNode;
}

export default function AppTheme(props: AppThemeProps): JSX.Element {
    const theme = createTheme({
        cssVariables: {
            colorSchemeSelector: 'data-mui-color-scheme',
            cssVarPrefix: 'template',
        },
        colorSchemes,
        typography,
        shadows,
        shape,
        components: {
            MuiButtonBase: {
                defaultProps: {
                    disableTouchRipple: true,
                    disableRipple: true,
                },
                styleOverrides: {
                    root: ({theme}) => ({
                        boxSizing: 'border-box',
                        transition: 'all 100ms ease-in',
                        '&:focus-visible': {
                            outline: `3px solid ${alpha(theme.palette.primary.main, 0.5)}`,
                            outlineOffset: '2px',
                        },
                    }),
                },
            },
            MuiButton: {
                styleOverrides: {
                    root: ({theme}) => ({
                        boxShadow: 'none',
                        borderRadius: (theme.vars || theme).shape.borderRadius,
                        textTransform: 'none',
                        variants: [
                            {
                                props: {
                                    size: 'small',
                                },
                                style: {
                                    height: '2.25rem',
                                    padding: '8px 12px',
                                },
                            },
                            {
                                props: {
                                    size: 'medium',
                                },
                                style: {
                                    height: '2.5rem', // 40px
                                },
                            },
                            {
                                props: {
                                    color: 'primary',
                                    variant: 'contained',
                                },
                                style: {
                                    color: 'white',
                                    backgroundColor: gray[900],
                                    backgroundImage: `linear-gradient(to bottom, ${gray[700]}, ${gray[800]})`,
                                    boxShadow: `inset 0 1px 0 ${gray[600]}, inset 0 -1px 0 1px hsl(220, 0%, 0%)`,
                                    border: `1px solid ${gray[700]}`,
                                    '&:hover': {
                                        backgroundImage: 'none',
                                        backgroundColor: gray[700],
                                        boxShadow: 'none',
                                    },
                                    '&:active': {
                                        backgroundColor: gray[800],
                                    },
                                    ...theme.applyStyles('dark', {
                                        color: 'black',
                                        backgroundColor: gray[50],
                                        backgroundImage: `linear-gradient(to bottom, ${gray[100]}, ${gray[50]})`,
                                        boxShadow: 'inset 0 -1px 0  hsl(220, 30%, 80%)',
                                        border: `1px solid ${gray[50]}`,
                                        '&:hover': {
                                            backgroundImage: 'none',
                                            backgroundColor: gray[300],
                                            boxShadow: 'none',
                                        },
                                        '&:active': {
                                            backgroundColor: gray[400],
                                        },
                                    }),
                                },
                            },
                            {
                                props: {
                                    color: 'secondary',
                                    variant: 'contained',
                                },
                                style: {
                                    color: 'white',
                                    backgroundColor: brand[300],
                                    backgroundImage:
                                        'linear-gradient(to bottom, ' + `${alpha(brand[400], 0.8)}, ${brand[500]})`,
                                    boxShadow:
                                        `inset 0 2px 0 ${alpha(brand[200], 0.2)}, ` +
                                        `inset 0 -2px 0 ${alpha(brand[700], 0.4)}`,
                                    border: `1px solid ${brand[500]}`,
                                    '&:hover': {
                                        backgroundColor: brand[700],
                                        boxShadow: 'none',
                                    },
                                    '&:active': {
                                        backgroundColor: brand[700],
                                        backgroundImage: 'none',
                                    },
                                },
                            },
                            {
                                props: {
                                    variant: 'outlined',
                                },
                                style: {
                                    color: (theme.vars || theme).palette.text.primary,
                                    border: '1px solid',
                                    borderColor: gray[200],
                                    backgroundColor: alpha(gray[50], 0.3),
                                    '&:hover': {
                                        backgroundColor: gray[100],
                                        borderColor: gray[300],
                                    },
                                    '&:active': {
                                        backgroundColor: gray[200],
                                    },
                                    ...theme.applyStyles('dark', {
                                        backgroundColor: gray[800],
                                        borderColor: gray[700],

                                        '&:hover': {
                                            backgroundColor: gray[900],
                                            borderColor: gray[600],
                                        },
                                        '&:active': {
                                            backgroundColor: gray[900],
                                        },
                                    }),
                                },
                            },
                            {
                                props: {
                                    color: 'secondary',
                                    variant: 'outlined',
                                },
                                style: {
                                    color: brand[700],
                                    border: '1px solid',
                                    borderColor: brand[200],
                                    backgroundColor: brand[50],
                                    '&:hover': {
                                        backgroundColor: brand[100],
                                        borderColor: brand[400],
                                    },
                                    '&:active': {
                                        backgroundColor: alpha(brand[200], 0.7),
                                    },
                                    ...theme.applyStyles('dark', {
                                        color: brand[50],
                                        border: '1px solid',
                                        borderColor: brand[900],
                                        backgroundColor: alpha(brand[900], 0.3),
                                        '&:hover': {
                                            borderColor: brand[700],
                                            backgroundColor: alpha(brand[900], 0.6),
                                        },
                                        '&:active': {
                                            backgroundColor: alpha(brand[900], 0.5),
                                        },
                                    }),
                                },
                            },
                            {
                                props: {
                                    variant: 'text',
                                },
                                style: {
                                    color: gray[600],
                                    '&:hover': {
                                        backgroundColor: gray[100],
                                    },
                                    '&:active': {
                                        backgroundColor: gray[200],
                                    },
                                    ...theme.applyStyles('dark', {
                                        color: gray[50],
                                        '&:hover': {
                                            backgroundColor: gray[700],
                                        },
                                        '&:active': {
                                            backgroundColor: alpha(gray[700], 0.7),
                                        },
                                    }),
                                },
                            },
                            {
                                props: {
                                    color: 'secondary',
                                    variant: 'text',
                                },
                                style: {
                                    color: brand[700],
                                    '&:hover': {
                                        backgroundColor: alpha(brand[100], 0.5),
                                    },
                                    '&:active': {
                                        backgroundColor: alpha(brand[200], 0.7),
                                    },
                                    ...theme.applyStyles('dark', {
                                        color: brand[100],
                                        '&:hover': {
                                            backgroundColor: alpha(brand[900], 0.5),
                                        },
                                        '&:active': {
                                            backgroundColor: alpha(brand[900], 0.3),
                                        },
                                    }),
                                },
                            },
                        ],
                    }),
                },
            },
            MuiIconButton: {
                styleOverrides: {
                    root: ({theme}) => ({
                        boxShadow: 'none',
                        borderRadius: (theme.vars || theme).shape.borderRadius,
                        textTransform: 'none',
                        fontWeight: theme.typography.fontWeightMedium,
                        letterSpacing: 0,
                        color: (theme.vars || theme).palette.text.primary,
                        border: '1px solid ',
                        borderColor: gray[200],
                        backgroundColor: alpha(gray[50], 0.3),
                        '&:hover': {
                            backgroundColor: gray[100],
                            borderColor: gray[300],
                        },
                        '&:active': {
                            backgroundColor: gray[200],
                        },
                        ...theme.applyStyles('dark', {
                            backgroundColor: gray[800],
                            borderColor: gray[700],
                            '&:hover': {
                                backgroundColor: gray[900],
                                borderColor: gray[600],
                            },
                            '&:active': {
                                backgroundColor: gray[900],
                            },
                        }),
                        variants: [
                            {
                                props: {
                                    size: 'small',
                                },
                                style: {
                                    width: '2.25rem',
                                    height: '2.25rem',
                                    padding: '0.25rem',
                                    [`& .${svgIconClasses.root}`]: {fontSize: '1rem'},
                                },
                            },
                            {
                                props: {
                                    size: 'medium',
                                },
                                style: {
                                    width: '2.5rem',
                                    height: '2.5rem',
                                },
                            },
                        ],
                    }),
                },
            },
            MuiToggleButtonGroup: {
                styleOverrides: {
                    root: ({theme}) => ({
                        borderRadius: '10px',
                        boxShadow: `0 4px 16px ${alpha(gray[400], 0.2)}`,
                        [`& .${toggleButtonGroupClasses.selected}`]: {
                            color: brand[500],
                        },
                        ...theme.applyStyles('dark', {
                            [`& .${toggleButtonGroupClasses.selected}`]: {
                                color: '#fff',
                            },
                            boxShadow: `0 4px 16px ${alpha(brand[700], 0.5)}`,
                        }),
                    }),
                },
            },
            MuiToggleButton: {
                styleOverrides: {
                    root: ({theme}) => ({
                        padding: '12px 16px',
                        textTransform: 'none',
                        borderRadius: '10px',
                        fontWeight: 500,
                        ...theme.applyStyles('dark', {
                            color: gray[400],
                            boxShadow: '0 4px 16px rgba(0, 0, 0, 0.5)',
                            [`&.${toggleButtonClasses.selected}`]: {
                                color: brand[300],
                            },
                        }),
                    }),
                },
            },
            MuiCheckbox: {
                defaultProps: {
                    disableRipple: true,
                    icon: <CheckBoxOutlineBlankRoundedIcon sx={{color: 'hsla(210, 0%, 0%, 0.0)'}} />,
                    checkedIcon: <CheckRoundedIcon sx={{height: 14, width: 14}} />,
                    indeterminateIcon: <RemoveRoundedIcon sx={{height: 14, width: 14}} />,
                },
                styleOverrides: {
                    root: ({theme}) => ({
                        margin: 10,
                        height: 16,
                        width: 16,
                        borderRadius: 5,
                        border: '1px solid ',
                        borderColor: alpha(gray[300], 0.8),
                        boxShadow: '0 0 0 1.5px hsla(210, 0%, 0%, 0.04) inset',
                        backgroundColor: alpha(gray[100], 0.4),
                        transition: 'border-color, background-color, 120ms ease-in',
                        '&:hover': {
                            borderColor: brand[300],
                        },
                        '&.Mui-focusVisible': {
                            outline: `3px solid ${alpha(brand[500], 0.5)}`,
                            outlineOffset: '2px',
                            borderColor: brand[400],
                        },
                        '&.Mui-checked': {
                            color: 'white',
                            backgroundColor: brand[500],
                            borderColor: brand[500],
                            boxShadow: 'none',
                            '&:hover': {
                                backgroundColor: brand[600],
                            },
                        },
                        ...theme.applyStyles('dark', {
                            borderColor: alpha(gray[700], 0.8),
                            boxShadow: '0 0 0 1.5px hsl(210, 0%, 0%) inset',
                            backgroundColor: alpha(gray[900], 0.8),
                            '&:hover': {
                                borderColor: brand[300],
                            },
                            '&.Mui-focusVisible': {
                                borderColor: brand[400],
                                outline: `3px solid ${alpha(brand[500], 0.5)}`,
                                outlineOffset: '2px',
                            },
                        }),
                    }),
                },
            },
            MuiInputBase: {
                styleOverrides: {
                    root: {
                        border: 'none',
                    },
                    input: {
                        '&::placeholder': {
                            opacity: 0.7,
                            color: gray[500],
                        },
                    },
                },
            },
            MuiOutlinedInput: {
                styleOverrides: {
                    input: {
                        padding: 0,
                    },
                    root: ({theme}) => ({
                        padding: '8px 12px',
                        color: (theme.vars || theme).palette.text.primary,
                        borderRadius: (theme.vars || theme).shape.borderRadius,
                        border: `1px solid ${(theme.vars || theme).palette.divider}`,
                        backgroundColor: (theme.vars || theme).palette.background.default,
                        transition: 'border 120ms ease-in',
                        '&:hover': {
                            borderColor: gray[400],
                        },
                        [`&.${outlinedInputClasses.focused}`]: {
                            outline: `3px solid ${alpha(brand[500], 0.5)}`,
                            borderColor: brand[400],
                        },
                        ...theme.applyStyles('dark', {
                            '&:hover': {
                                borderColor: gray[500],
                            },
                        }),
                        variants: [
                            {
                                props: {
                                    size: 'small',
                                },
                                style: {
                                    height: '2.25rem',
                                },
                            },
                            {
                                props: {
                                    size: 'medium',
                                },
                                style: {
                                    height: '2.5rem',
                                },
                            },
                        ],
                    }),
                    notchedOutline: {
                        border: 'none',
                    },
                },
            },
            MuiInputAdornment: {
                styleOverrides: {
                    root: ({theme}) => ({
                        color: (theme.vars || theme).palette.grey[500],
                        ...theme.applyStyles('dark', {
                            color: (theme.vars || theme).palette.grey[400],
                        }),
                    }),
                },
            },
            MuiFormLabel: {
                styleOverrides: {
                    root: ({theme}) => ({
                        typography: theme.typography.caption,
                        marginBottom: 8,
                    }),
                },
            },
            MuiList: {
                styleOverrides: {
                    root: {
                        padding: '8px',
                        display: 'flex',
                        flexDirection: 'column',
                        gap: 0,
                    },
                },
            },
            MuiListItem: {
                styleOverrides: {
                    root: ({theme}) => ({
                        [`& .${svgIconClasses.root}`]: {
                            width: '1rem',
                            height: '1rem',
                            color: (theme.vars || theme).palette.text.secondary,
                        },
                        [`& .${typographyClasses.root}`]: {
                            fontWeight: 500,
                        },
                        [`& .${buttonBaseClasses.root}`]: {
                            display: 'flex',
                            gap: 8,
                            padding: '2px 8px',
                            borderRadius: (theme.vars || theme).shape.borderRadius,
                            opacity: 0.7,
                            '&.Mui-selected': {
                                opacity: 1,
                                backgroundColor: alpha(theme.palette.action.selected, 0.3),
                                [`& .${svgIconClasses.root}`]: {
                                    color: (theme.vars || theme).palette.text.primary,
                                },
                                '&:focus-visible': {
                                    backgroundColor: alpha(theme.palette.action.selected, 0.3),
                                },
                                '&:hover': {
                                    backgroundColor: alpha(theme.palette.action.selected, 0.5),
                                },
                            },
                            '&:focus-visible': {
                                backgroundColor: 'transparent',
                            },
                        },
                    }),
                },
            },
            MuiListItemText: {
                styleOverrides: {
                    primary: ({theme}) => ({
                        fontSize: theme.typography.body2.fontSize,
                        fontWeight: 500,
                        lineHeight: theme.typography.body2.lineHeight,
                    }),
                    secondary: ({theme}) => ({
                        fontSize: theme.typography.caption.fontSize,
                        lineHeight: theme.typography.caption.lineHeight,
                    }),
                },
            },
            MuiListSubheader: {
                styleOverrides: {
                    root: ({theme}) => ({
                        backgroundColor: 'transparent',
                        padding: '4px 8px',
                        fontSize: theme.typography.caption.fontSize,
                        fontWeight: 500,
                        lineHeight: theme.typography.caption.lineHeight,
                    }),
                },
            },
            MuiListItemIcon: {
                styleOverrides: {
                    root: {
                        minWidth: 0,
                    },
                },
            },
            MuiChip: {
                defaultProps: {
                    size: 'small',
                },
                styleOverrides: {
                    root: ({theme}) => ({
                        border: '1px solid',
                        borderRadius: '999px',
                        [`& .${chipClasses.label}`]: {
                            fontWeight: 600,
                        },
                        variants: [
                            {
                                props: {
                                    color: 'default',
                                },
                                style: {
                                    borderColor: gray[200],
                                    backgroundColor: gray[100],
                                    [`& .${chipClasses.label}`]: {
                                        color: gray[500],
                                    },
                                    [`& .${chipClasses.icon}`]: {
                                        color: gray[500],
                                    },
                                    ...theme.applyStyles('dark', {
                                        borderColor: gray[700],
                                        backgroundColor: gray[800],
                                        [`& .${chipClasses.label}`]: {
                                            color: gray[300],
                                        },
                                        [`& .${chipClasses.icon}`]: {
                                            color: gray[300],
                                        },
                                    }),
                                },
                            },
                            {
                                props: {
                                    color: 'success',
                                },
                                style: {
                                    borderColor: green[200],
                                    backgroundColor: green[50],
                                    [`& .${chipClasses.label}`]: {
                                        color: green[500],
                                    },
                                    [`& .${chipClasses.icon}`]: {
                                        color: green[500],
                                    },
                                    ...theme.applyStyles('dark', {
                                        borderColor: green[800],
                                        backgroundColor: green[900],
                                        [`& .${chipClasses.label}`]: {
                                            color: green[300],
                                        },
                                        [`& .${chipClasses.icon}`]: {
                                            color: green[300],
                                        },
                                    }),
                                },
                            },
                            {
                                props: {
                                    color: 'error',
                                },
                                style: {
                                    borderColor: red[100],
                                    backgroundColor: red[50],
                                    [`& .${chipClasses.label}`]: {
                                        color: red[500],
                                    },
                                    [`& .${chipClasses.icon}`]: {
                                        color: red[500],
                                    },
                                    ...theme.applyStyles('dark', {
                                        borderColor: red[800],
                                        backgroundColor: red[900],
                                        [`& .${chipClasses.label}`]: {
                                            color: red[200],
                                        },
                                        [`& .${chipClasses.icon}`]: {
                                            color: red[300],
                                        },
                                    }),
                                },
                            },
                            {
                                props: {size: 'small'},
                                style: {
                                    maxHeight: 20,
                                    [`& .${chipClasses.label}`]: {
                                        fontSize: theme.typography.caption.fontSize,
                                    },
                                    [`& .${svgIconClasses.root}`]: {
                                        fontSize: theme.typography.caption.fontSize,
                                    },
                                },
                            },
                            {
                                props: {size: 'medium'},
                                style: {
                                    [`& .${chipClasses.label}`]: {
                                        fontSize: theme.typography.caption.fontSize,
                                    },
                                },
                            },
                        ],
                    }),
                },
            },
            MuiTablePagination: {
                styleOverrides: {
                    actions: {
                        display: 'flex',
                        gap: 8,
                        marginRight: 6,
                        [`& .${iconButtonClasses.root}`]: {
                            minWidth: 0,
                            width: 36,
                            height: 36,
                        },
                    },
                },
            },
            MuiIcon: {
                defaultProps: {
                    fontSize: 'small',
                },
                styleOverrides: {
                    root: {
                        variants: [
                            {
                                props: {
                                    fontSize: 'small',
                                },
                                style: {
                                    fontSize: '1rem',
                                },
                            },
                        ],
                    },
                },
            },
            MuiAlert: {
                styleOverrides: {
                    root: ({theme}) => ({
                        borderRadius: 10,
                        backgroundColor: orange[100],
                        color: (theme.vars || theme).palette.text.primary,
                        border: `1px solid ${alpha(orange[300], 0.5)}`,
                        '& .MuiAlert-icon': {
                            color: orange[500],
                        },
                        ...theme.applyStyles('dark', {
                            backgroundColor: `${alpha(orange[900], 0.5)}`,
                            border: `1px solid ${alpha(orange[800], 0.5)}`,
                        }),
                    }),
                },
            },
            MuiDialog: {
                styleOverrides: {
                    root: ({theme}) => ({
                        '& .MuiDialog-paper': {
                            borderRadius: '10px',
                            border: '1px solid',
                            borderColor: (theme.vars || theme).palette.divider,
                        },
                    }),
                },
            },
            MuiLinearProgress: {
                styleOverrides: {
                    root: ({theme}) => ({
                        height: 8,
                        borderRadius: 8,
                        backgroundColor: gray[200],
                        ...theme.applyStyles('dark', {
                            backgroundColor: gray[800],
                        }),
                    }),
                },
            },
            MuiMenuItem: {
                styleOverrides: {
                    root: ({theme}) => ({
                        borderRadius: (theme.vars || theme).shape.borderRadius,
                        padding: '6px 8px',
                        [`&.${menuItemClasses.focusVisible}`]: {
                            backgroundColor: 'transparent',
                        },
                        [`&.${menuItemClasses.selected}`]: {
                            [`&.${menuItemClasses.focusVisible}`]: {
                                backgroundColor: alpha(theme.palette.action.selected, 0.3),
                            },
                        },
                    }),
                },
            },
            MuiMenu: {
                styleOverrides: {
                    list: {
                        gap: '0px',
                        [`&.${dividerClasses.root}`]: {
                            margin: '0 -8px',
                        },
                    },
                    paper: ({theme}) => ({
                        marginTop: '4px',
                        borderRadius: (theme.vars || theme).shape.borderRadius,
                        border: `1px solid ${(theme.vars || theme).palette.divider}`,
                        backgroundImage: 'none',
                        background: 'hsl(0, 0%, 100%)',
                        boxShadow:
                            'hsla(220, 30%, 5%, 0.07) 0px 4px 16px 0px, hsla(220, 25%, 10%, 0.07) 0px 8px 16px -5px',
                        [`& .${buttonBaseClasses.root}`]: {
                            '&.Mui-selected': {
                                backgroundColor: alpha(theme.palette.action.selected, 0.3),
                            },
                        },
                        ...theme.applyStyles('dark', {
                            background: gray[900],
                            boxShadow:
                                'hsla(220, 30%, 5%, 0.7) 0px 4px 16px 0px, hsla(220, 25%, 10%, 0.8) 0px 8px 16px -5px',
                        }),
                    }),
                },
            },
            MuiSelect: {
                defaultProps: {
                    IconComponent: React.forwardRef<SVGSVGElement, SvgIconProps>((props, ref) => (
                        <UnfoldMoreRoundedIcon fontSize="small" {...props} ref={ref} />
                    )),
                },
                styleOverrides: {
                    root: ({theme}) => ({
                        borderRadius: (theme.vars || theme).shape.borderRadius,
                        border: '1px solid',
                        borderColor: gray[200],
                        backgroundColor: (theme.vars || theme).palette.background.paper,
                        boxShadow:
                            'inset 0 1px 0 1px hsla(220, 0%, 100%, 0.6), inset 0 -1px 0 1px hsla(220, 35%, 90%, 0.5)',
                        '&:hover': {
                            borderColor: gray[300],
                            backgroundColor: (theme.vars || theme).palette.background.paper,
                            boxShadow: 'none',
                        },
                        [`&.${selectClasses.focused}`]: {
                            outlineOffset: 0,
                            borderColor: gray[400],
                        },
                        '&:before, &:after': {
                            display: 'none',
                        },

                        ...theme.applyStyles('dark', {
                            borderRadius: (theme.vars || theme).shape.borderRadius,
                            borderColor: gray[700],
                            backgroundColor: (theme.vars || theme).palette.background.paper,
                            boxShadow:
                                `inset 0 1px 0 1px ${alpha(gray[700], 0.15)}, inset 0 ` +
                                '-1px 0 1px hsla(220, 0%, 0%, 0.7)',
                            '&:hover': {
                                borderColor: alpha(gray[700], 0.7),
                                backgroundColor: (theme.vars || theme).palette.background.paper,
                                boxShadow: 'none',
                            },
                            [`&.${selectClasses.focused}`]: {
                                outlineOffset: 0,
                                borderColor: gray[900],
                            },
                            '&:before, &:after': {
                                display: 'none',
                            },
                        }),
                    }),
                    select: ({theme}) => ({
                        display: 'flex',
                        alignItems: 'center',
                        ...theme.applyStyles('dark', {
                            display: 'flex',
                            alignItems: 'center',
                            '&:focus-visible': {
                                backgroundColor: gray[900],
                            },
                        }),
                    }),
                },
            },
            MuiLink: {
                defaultProps: {
                    underline: 'none',
                },
                styleOverrides: {
                    root: ({theme}) => ({
                        color: (theme.vars || theme).palette.text.primary,
                        fontWeight: 500,
                        position: 'relative',
                        textDecoration: 'none',
                        width: 'fit-content',
                        '&::before': {
                            content: '""',
                            position: 'absolute',
                            width: '100%',
                            height: '1px',
                            bottom: 0,
                            left: 0,
                            backgroundColor: (theme.vars || theme).palette.text.secondary,
                            opacity: 0.3,
                            transition: 'width 0.3s ease, opacity 0.3s ease',
                        },
                        '&:hover::before': {
                            width: 0,
                        },
                        '&:focus-visible': {
                            outline: `3px solid ${alpha(brand[500], 0.5)}`,
                            outlineOffset: '4px',
                            borderRadius: '2px',
                        },
                    }),
                },
            },
            MuiDrawer: {
                styleOverrides: {
                    paper: ({theme}) => ({
                        backgroundColor: (theme.vars || theme).palette.background.default,
                    }),
                },
            },
            MuiPaginationItem: {
                styleOverrides: {
                    root: ({theme}) => ({
                        '&.Mui-selected': {
                            color: 'white',
                            backgroundColor: (theme.vars || theme).palette.grey[900],
                        },
                        ...theme.applyStyles('dark', {
                            '&.Mui-selected': {
                                color: 'black',
                                backgroundColor: (theme.vars || theme).palette.grey[50],
                            },
                        }),
                    }),
                },
            },
            MuiTabs: {
                styleOverrides: {
                    root: {minHeight: 'fit-content'},
                    indicator: ({theme}) => ({
                        backgroundColor: (theme.vars || theme).palette.grey[800],
                        ...theme.applyStyles('dark', {
                            backgroundColor: (theme.vars || theme).palette.grey[200],
                        }),
                    }),
                },
            },
            MuiTab: {
                styleOverrides: {
                    root: ({theme}) => ({
                        padding: '6px 8px',
                        marginBottom: '8px',
                        textTransform: 'none',
                        minWidth: 'fit-content',
                        minHeight: 'fit-content',
                        color: (theme.vars || theme).palette.text.secondary,
                        borderRadius: (theme.vars || theme).shape.borderRadius,
                        border: '1px solid',
                        borderColor: 'transparent',
                        ':hover': {
                            color: (theme.vars || theme).palette.text.primary,
                            backgroundColor: gray[100],
                            borderColor: gray[200],
                        },
                        [`&.${tabClasses.selected}`]: {
                            color: gray[900],
                        },
                        ...theme.applyStyles('dark', {
                            ':hover': {
                                color: (theme.vars || theme).palette.text.primary,
                                backgroundColor: gray[800],
                                borderColor: gray[700],
                            },
                            [`&.${tabClasses.selected}`]: {
                                color: '#fff',
                            },
                        }),
                    }),
                },
            },
            MuiStepConnector: {
                styleOverrides: {
                    line: ({theme}) => ({
                        borderTop: '1px solid',
                        borderColor: (theme.vars || theme).palette.divider,
                        flex: 1,
                        borderRadius: '99px',
                    }),
                },
            },
            MuiStepIcon: {
                styleOverrides: {
                    root: ({theme}) => ({
                        color: 'transparent',
                        border: `1px solid ${gray[400]}`,
                        width: 12,
                        height: 12,
                        borderRadius: '50%',
                        '& text': {
                            display: 'none',
                        },
                        '&.Mui-active': {
                            border: 'none',
                            color: (theme.vars || theme).palette.primary.main,
                        },
                        '&.Mui-completed': {
                            border: 'none',
                            color: (theme.vars || theme).palette.success.main,
                        },
                        ...theme.applyStyles('dark', {
                            border: `1px solid ${gray[700]}`,
                            '&.Mui-active': {
                                border: 'none',
                                color: (theme.vars || theme).palette.primary.light,
                            },
                            '&.Mui-completed': {
                                border: 'none',
                                color: (theme.vars || theme).palette.success.light,
                            },
                        }),
                        variants: [
                            {
                                props: {completed: true},
                                style: {
                                    width: 12,
                                    height: 12,
                                },
                            },
                        ],
                    }),
                },
            },
            MuiStepLabel: {
                styleOverrides: {
                    label: ({theme}) => ({
                        '&.Mui-completed': {
                            opacity: 0.6,
                            ...theme.applyStyles('dark', {opacity: 0.5}),
                        },
                    }),
                },
            },
            MuiAccordion: {
                defaultProps: {
                    elevation: 0,
                    disableGutters: true,
                },
                styleOverrides: {
                    root: ({theme}) => ({
                        padding: 4,
                        overflow: 'clip',
                        backgroundColor: (theme.vars || theme).palette.background.default,
                        border: '1px solid',
                        borderColor: (theme.vars || theme).palette.divider,
                        ':before': {
                            backgroundColor: 'transparent',
                        },
                        '&:not(:last-of-type)': {
                            borderBottom: 'none',
                        },
                        '&:first-of-type': {
                            borderTopLeftRadius: (theme.vars || theme).shape.borderRadius,
                            borderTopRightRadius: (theme.vars || theme).shape.borderRadius,
                        },
                        '&:last-of-type': {
                            borderBottomLeftRadius: (theme.vars || theme).shape.borderRadius,
                            borderBottomRightRadius: (theme.vars || theme).shape.borderRadius,
                        },
                    }),
                },
            },
            MuiAccordionSummary: {
                styleOverrides: {
                    root: ({theme}) => ({
                        border: 'none',
                        borderRadius: 8,
                        '&:hover': {backgroundColor: gray[50]},
                        '&:focus-visible': {backgroundColor: 'transparent'},
                        ...theme.applyStyles('dark', {
                            '&:hover': {backgroundColor: gray[800]},
                        }),
                    }),
                },
            },
            MuiAccordionDetails: {
                styleOverrides: {
                    root: {mb: 20, border: 'none'},
                },
            },
            MuiPaper: {
                defaultProps: {
                    elevation: 0,
                },
            },
            MuiCard: {
                styleOverrides: {
                    root: ({theme}) => {
                        return {
                            padding: 16,
                            gap: 16,
                            transition: 'all 100ms ease',
                            backgroundColor: gray[50],
                            borderRadius: (theme.vars || theme).shape.borderRadius,
                            border: `1px solid ${(theme.vars || theme).palette.divider}`,
                            boxShadow: 'none',
                            ...theme.applyStyles('dark', {
                                backgroundColor: gray[800],
                            }),
                            variants: [
                                {
                                    props: {
                                        variant: 'outlined',
                                    },
                                    style: {
                                        border: `1px solid ${(theme.vars || theme).palette.divider}`,
                                        boxShadow: 'none',
                                        background: 'hsl(0, 0%, 100%)',
                                        ...theme.applyStyles('dark', {
                                            background: alpha(gray[900], 0.4),
                                        }),
                                    },
                                },
                            ],
                        };
                    },
                },
            },
            MuiCardContent: {
                styleOverrides: {
                    root: {
                        padding: 0,
                        '&:last-child': {paddingBottom: 0},
                    },
                },
            },
            MuiCardHeader: {
                styleOverrides: {
                    root: {
                        padding: 0,
                    },
                },
            },
            MuiCardActions: {
                styleOverrides: {
                    root: {
                        padding: 0,
                    },
                },
            },
            MuiChartsAxis: {
                styleOverrides: {
                    root: ({theme}) => ({
                        [`& .${axisClasses.line}`]: {
                            stroke: gray[300],
                        },
                        [`& .${axisClasses.tick}`]: {stroke: gray[300]},
                        [`& .${axisClasses.tickLabel}`]: {
                            fill: gray[500],
                            fontWeight: 500,
                        },
                        ...theme.applyStyles('dark', {
                            [`& .${axisClasses.line}`]: {
                                stroke: gray[700],
                            },
                            [`& .${axisClasses.tick}`]: {stroke: gray[700]},
                            [`& .${axisClasses.tickLabel}`]: {
                                fill: gray[300],
                                fontWeight: 500,
                            },
                        }),
                    }),
                },
            },
            MuiChartsTooltip: {
                styleOverrides: {
                    mark: ({theme}) => ({
                        ry: 6,
                        boxShadow: 'none',
                        border: `1px solid ${(theme.vars || theme).palette.divider}`,
                    }),
                    table: ({theme}) => ({
                        border: `1px solid ${(theme.vars || theme).palette.divider}`,
                        borderRadius: theme.shape.borderRadius,
                        background: 'hsl(0, 0%, 100%)',
                        ...theme.applyStyles('dark', {
                            background: gray[900],
                        }),
                    }),
                },
            },
            MuiChartsLegend: {
                styleOverrides: {
                    root: {
                        [`& .${legendClasses.mark}`]: {
                            ry: 6,
                        },
                    },
                },
            },
            MuiChartsGrid: {
                styleOverrides: {
                    root: ({theme}) => ({
                        [`& .${chartsGridClasses.line}`]: {
                            stroke: gray[200],
                            strokeDasharray: '4 2',
                            strokeWidth: 0.8,
                        },
                        ...theme.applyStyles('dark', {
                            [`& .${chartsGridClasses.line}`]: {
                                stroke: gray[700],
                                strokeDasharray: '4 2',
                                strokeWidth: 0.8,
                            },
                        }),
                    }),
                },
            },
            MuiDataGrid: {
                styleOverrides: {
                    root: ({theme}) => ({
                        '--DataGrid-overlayHeight': '300px',
                        overflow: 'clip',
                        borderColor: (theme.vars || theme).palette.divider,
                        backgroundColor: (theme.vars || theme).palette.background.default,
                        [`& .${gridClasses.columnHeader}`]: {
                            backgroundColor: (theme.vars || theme).palette.background.paper,
                        },
                        [`& .${gridClasses.footerContainer}`]: {
                            backgroundColor: (theme.vars || theme).palette.background.paper,
                        },
                        [`& .${checkboxClasses.root}`]: {
                            padding: theme.spacing(0.5),
                            '& > svg': {
                                fontSize: '1rem',
                            },
                        },
                        [`& .${tablePaginationClasses.root}`]: {
                            marginRight: theme.spacing(1),
                            '& .MuiIconButton-root': {
                                maxHeight: 32,
                                maxWidth: 32,
                                '& > svg': {
                                    fontSize: '1rem',
                                },
                            },
                        },
                    }),
                    cell: ({theme}) => ({borderTopColor: (theme.vars || theme).palette.divider}),
                    menu: ({theme}) => ({
                        borderRadius: theme.shape.borderRadius,
                        backgroundImage: 'none',
                        [`& .${paperClasses.root}`]: {
                            border: `1px solid ${(theme.vars || theme).palette.divider}`,
                        },

                        [`& .${menuItemClasses.root}`]: {
                            margin: '0 4px',
                        },
                        [`& .${listItemIconClasses.root}`]: {
                            marginRight: 0,
                        },
                        [`& .${listClasses.root}`]: {
                            paddingLeft: 0,
                            paddingRight: 0,
                        },
                    }),

                    row: ({theme}) => ({
                        '&:last-of-type': {borderBottom: `1px solid ${(theme.vars || theme).palette.divider}`},
                        '&:hover': {
                            backgroundColor: (theme.vars || theme).palette.action.hover,
                        },
                        '&.Mui-selected': {
                            background: (theme.vars || theme).palette.action.selected,
                            '&:hover': {
                                backgroundColor: (theme.vars || theme).palette.action.hover,
                            },
                        },
                    }),
                    iconButtonContainer: ({theme}) => ({
                        [`& .${iconButtonClasses.root}`]: {
                            border: 'none',
                            backgroundColor: 'transparent',
                            '&:hover': {
                                backgroundColor: alpha(theme.palette.action.selected, 0.3),
                            },
                            '&:active': {
                                backgroundColor: gray[200],
                            },
                            ...theme.applyStyles('dark', {
                                color: gray[50],
                                '&:hover': {
                                    backgroundColor: gray[800],
                                },
                                '&:active': {
                                    backgroundColor: gray[900],
                                },
                            }),
                        },
                    }),
                    menuIconButton: ({theme}) => ({
                        border: 'none',
                        backgroundColor: 'transparent',
                        '&:hover': {
                            backgroundColor: gray[100],
                        },
                        '&:active': {
                            backgroundColor: gray[200],
                        },
                        ...theme.applyStyles('dark', {
                            color: gray[50],
                            '&:hover': {
                                backgroundColor: gray[800],
                            },
                            '&:active': {
                                backgroundColor: gray[900],
                            },
                        }),
                    }),
                    filterForm: ({theme}) => ({
                        gap: theme.spacing(1),
                        alignItems: 'flex-end',
                    }),
                    columnsManagementHeader: ({theme}) => ({
                        paddingRight: theme.spacing(3),
                        paddingLeft: theme.spacing(3),
                    }),
                    columnHeaderTitleContainer: {
                        flexGrow: 1,
                        justifyContent: 'space-between',
                    },
                    columnHeaderDraggableContainer: {paddingRight: 2},
                },
            },
            MuiPickersPopper: {
                styleOverrides: {
                    paper: ({theme}) => ({
                        marginTop: 4,
                        borderRadius: theme.shape.borderRadius,
                        border: `1px solid ${(theme.vars || theme).palette.divider}`,
                        backgroundImage: 'none',
                        background: 'hsl(0, 0%, 100%)',
                        boxShadow:
                            'hsla(220, 30%, 5%, 0.07) 0px 4px 16px 0px, hsla(220, 25%, 10%, 0.07) 0px 8px 16px -5px',
                        [`& .${menuItemClasses.root}`]: {
                            borderRadius: 6,
                            margin: '0 6px',
                        },
                        ...theme.applyStyles('dark', {
                            background: gray[900],
                            boxShadow:
                                'hsla(220, 30%, 5%, 0.7) 0px 4px 16px 0px, hsla(220, 25%, 10%, 0.8) 0px 8px 16px -5px',
                        }),
                    }),
                },
            },
            MuiPickersArrowSwitcher: {
                styleOverrides: {
                    spacer: {width: 16},
                    button: ({theme}) => ({
                        backgroundColor: 'transparent',
                        color: (theme.vars || theme).palette.grey[500],
                        ...theme.applyStyles('dark', {
                            color: (theme.vars || theme).palette.grey[400],
                        }),
                    }),
                },
            },
            MuiPickersCalendarHeader: {
                styleOverrides: {
                    switchViewButton: {
                        padding: 0,
                        border: 'none',
                    },
                },
            },
            MuiPickersMonth: {
                styleOverrides: {
                    monthButton: ({theme}) => ({
                        fontSize: theme.typography.body1.fontSize,
                        color: (theme.vars || theme).palette.grey[600],
                        padding: theme.spacing(0.5),
                        borderRadius: theme.shape.borderRadius,
                        '&:hover': {
                            backgroundColor: (theme.vars || theme).palette.action.hover,
                        },
                        [`&.${pickersMonthClasses.selected}`]: {
                            backgroundColor: gray[700],
                            fontWeight: theme.typography.fontWeightMedium,
                        },
                        '&:focus': {
                            outline: `3px solid ${alpha(brand[500], 0.5)}`,
                            outlineOffset: '2px',
                            backgroundColor: 'transparent',
                            [`&.${pickersMonthClasses.selected}`]: {backgroundColor: gray[700]},
                        },
                        ...theme.applyStyles('dark', {
                            color: (theme.vars || theme).palette.grey[300],
                            '&:hover': {
                                backgroundColor: (theme.vars || theme).palette.action.hover,
                            },
                            [`&.${pickersMonthClasses.selected}`]: {
                                color: (theme.vars || theme).palette.common.black,
                                fontWeight: theme.typography.fontWeightMedium,
                                backgroundColor: gray[300],
                            },
                            '&:focus': {
                                outline: `3px solid ${alpha(brand[500], 0.5)}`,
                                outlineOffset: '2px',
                                backgroundColor: 'transparent',
                                [`&.${pickersMonthClasses.selected}`]: {backgroundColor: gray[300]},
                            },
                        }),
                    }),
                },
            },
            MuiPickersYear: {
                styleOverrides: {
                    yearButton: ({theme}) => ({
                        fontSize: theme.typography.body1.fontSize,
                        color: (theme.vars || theme).palette.grey[600],
                        padding: theme.spacing(0.5),
                        borderRadius: theme.shape.borderRadius,
                        height: 'fit-content',
                        '&:hover': {
                            backgroundColor: (theme.vars || theme).palette.action.hover,
                        },
                        [`&.${pickersYearClasses.selected}`]: {
                            backgroundColor: gray[700],
                            fontWeight: theme.typography.fontWeightMedium,
                        },
                        '&:focus': {
                            outline: `3px solid ${alpha(brand[500], 0.5)}`,
                            outlineOffset: '2px',
                            backgroundColor: 'transparent',
                            [`&.${pickersYearClasses.selected}`]: {backgroundColor: gray[700]},
                        },
                        ...theme.applyStyles('dark', {
                            color: (theme.vars || theme).palette.grey[300],
                            '&:hover': {
                                backgroundColor: (theme.vars || theme).palette.action.hover,
                            },
                            [`&.${pickersYearClasses.selected}`]: {
                                color: (theme.vars || theme).palette.common.black,
                                fontWeight: theme.typography.fontWeightMedium,
                                backgroundColor: gray[300],
                            },
                            '&:focus': {
                                outline: `3px solid ${alpha(brand[500], 0.5)}`,
                                outlineOffset: '2px',
                                backgroundColor: 'transparent',
                                [`&.${pickersYearClasses.selected}`]: {backgroundColor: gray[300]},
                            },
                        }),
                    }),
                },
            },
            MuiPickersDay: {
                styleOverrides: {
                    root: ({theme}) => ({
                        fontSize: theme.typography.body1.fontSize,
                        color: (theme.vars || theme).palette.grey[600],
                        padding: theme.spacing(0.5),
                        borderRadius: theme.shape.borderRadius,
                        '&:hover': {
                            backgroundColor: (theme.vars || theme).palette.action.hover,
                        },
                        [`&.${pickersDayClasses.selected}`]: {
                            backgroundColor: gray[700],
                            fontWeight: theme.typography.fontWeightMedium,
                        },
                        '&:focus': {
                            outline: `3px solid ${alpha(brand[500], 0.5)}`,
                            outlineOffset: '2px',
                            backgroundColor: 'transparent',
                            [`&.${pickersDayClasses.selected}`]: {backgroundColor: gray[700]},
                        },
                        ...theme.applyStyles('dark', {
                            color: (theme.vars || theme).palette.grey[300],
                            '&:hover': {
                                backgroundColor: (theme.vars || theme).palette.action.hover,
                            },
                            [`&.${pickersDayClasses.selected}`]: {
                                color: (theme.vars || theme).palette.common.black,
                                fontWeight: theme.typography.fontWeightMedium,
                                backgroundColor: gray[300],
                            },
                            '&:focus': {
                                outline: `3px solid ${alpha(brand[500], 0.5)}`,
                                outlineOffset: '2px',
                                backgroundColor: 'transparent',
                                [`&.${pickersDayClasses.selected}`]: {backgroundColor: gray[300]},
                            },
                        }),
                    }),
                },
            },
            MuiTreeItem2: {
                styleOverrides: {
                    root: ({theme}) => ({
                        position: 'relative',
                        boxSizing: 'border-box',
                        padding: theme.spacing(0, 1),
                        '& .groupTransition': {
                            marginLeft: theme.spacing(2),
                            padding: theme.spacing(0),
                            borderLeft: '1px solid',
                            borderColor: (theme.vars || theme).palette.divider,
                        },
                        '&:focus-visible .focused': {
                            outline: `3px solid ${alpha(brand[500], 0.5)}`,
                            outlineOffset: '2px',
                            '&:hover': {
                                backgroundColor: alpha(gray[300], 0.2),
                                outline: `3px solid ${alpha(brand[500], 0.5)}`,
                                outlineOffset: '2px',
                            },
                        },
                    }),
                    content: ({theme}) => ({
                        marginTop: theme.spacing(1),
                        padding: theme.spacing(0.5, 1),
                        overflow: 'clip',
                        '&:hover': {
                            backgroundColor: alpha(gray[300], 0.2),
                        },

                        '&.selected': {
                            backgroundColor: alpha(gray[300], 0.4),
                            '&:hover': {
                                backgroundColor: alpha(gray[300], 0.6),
                            },
                        },
                        ...theme.applyStyles('dark', {
                            '&:hover': {
                                backgroundColor: alpha(gray[500], 0.2),
                            },
                            '&:focus-visible': {
                                '&:hover': {
                                    backgroundColor: alpha(gray[500], 0.2),
                                },
                            },
                            '&.selected': {
                                backgroundColor: alpha(gray[500], 0.4),
                                '&:hover': {
                                    backgroundColor: alpha(gray[500], 0.6),
                                },
                            },
                        }),
                    }),
                },
            },
        },
    });

    return (
        <ThemeProvider theme={theme} disableTransitionOnChange>
            {props.children}
        </ThemeProvider>
    );
}
