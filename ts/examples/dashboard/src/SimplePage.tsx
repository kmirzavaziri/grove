import {Container, Box, Typography} from '@mui/material';
import React from 'react';

export interface SimplePageProps {
    code: string;
    message: string;
}

export const SimplePage: React.FC<SimplePageProps> = (props) => (
    <Container>
        <Box
            sx={{
                height: '100vh',
                display: 'flex',
                flexDirection: 'column',
                justifyContent: 'center',
                alignItems: 'center',
                textAlign: 'center',
            }}
        >
            <Typography variant="h1" component="h1" gutterBottom>
                {props.code}
            </Typography>
            <Typography variant="body2" color="text.secondary">
                {props.message}
            </Typography>
        </Box>
    </Container>
);
