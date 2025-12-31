import { defineConfig } from 'vite'

export default defineConfig({
  build: {
    manifest: true,
    rollupOptions: {
      input: 'src/main.js',
    },
  },
  server: {
    origin: 'http://localhost:5173',
  },
})
