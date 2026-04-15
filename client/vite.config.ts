import path from 'path';
import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import tailwindcss from '@tailwindcss/vite'

// https://vite.dev/config/
export default defineConfig({
  resolve: {
    alias: {
      '@': path.resolve(__dirname, 'src'),
      '@/app': path.resolve(__dirname, 'src/app'),
      '@/shared': path.resolve(__dirname, 'src/shared'),
      '@/pages': path.resolve(__dirname, 'src/pages'),
      '@/mudules': path.resolve(__dirname, 'src/modules'),
    },
  },
  plugins: [react(), tailwindcss(),],
})
