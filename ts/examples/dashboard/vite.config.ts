import react from '@vitejs/plugin-react';
import {defineConfig} from 'vite';

export default defineConfig({
    root: './',
    server: {port: 3000},
    plugins: [react()],
});
