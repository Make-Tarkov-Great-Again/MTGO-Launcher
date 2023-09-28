import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import path from 'path';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, 'src'),
      '@Event': path.resolve(__dirname, 'src/main.js') //Vue event emit thingie ma-bob
    },
  },
  template: {
    compilerOptions: {
      isCustomElement: (tag) => ['cus-'].includes(tag),

    }
  }
});