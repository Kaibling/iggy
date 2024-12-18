import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
  base: "/",
  plugins: [react()],
  preview: {
   port: 7080,
   strictPort: true,
  },
  server: {
   port: 7080,
   strictPort: true,
   host: true,
   origin: "http://0.0.0.0:7080",
  },
 });