import Box from '@mui/material/Box';
import Divider from '@mui/material/Divider';
import React from 'react';

export interface DDividerProps {}

export const DDivider: React.FC<DDividerProps> = (_props) => {
    return (
        <Box sx={{pb: 1}}>
            <Divider sx={{position: 'absolute', left: 0, right: 0}} />
        </Box>
    );
};
