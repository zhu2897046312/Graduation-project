import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import { resolve } from 'path';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      "/@": resolve(__dirname, "src")
    }
  },
  server: {
    host: "0.0.0.0",
    port: 3000,
    // 是否开启 https
    proxy: {
      "/oss": {
        target: "http://172.25.13.23:9077",
        changeOrigin: true,
      }
    }
  },
})
