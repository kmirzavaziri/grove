export type Align = 'inherit'
    | 'center'
    | 'justify'
    | 'start'
    | 'end';

export function align2MUI(align?: Align): "inherit" | "left" | "center" | "right" | "justify" | undefined {
    switch (align) {
        case 'start':
            return 'left';
        case 'end':
            return 'right';
        case undefined:
            return undefined;
        default:
            return align;
    }
}

export type Color = ''
    | 'primary'
    | 'secondary'
    | 'success'
    | 'error'
    | 'info'
    | 'warning'
    | string;

export type TypographyVariant = 'inherit'
    | 'body1'
    | 'body2'
    | 'button'
    | 'caption'
    | 'h1'
    | 'h2'
    | 'h3'
    | 'h4'
    | 'h5'
    | 'h6'
    | 'overline'
    | 'subtitle1'
    | 'subtitle2';

export type ButtonVariant = 'item'
    | 'contained'
    | 'outlined'
    | 'text';

export type Size = ''
    | 'small'
    | 'medium'
    | 'large'
    | string;
