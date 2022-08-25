import { defineConfig } from "vite";
import { svelte } from "@sveltejs/vite-plugin-svelte";
import removeConsole from "vite-plugin-remove-console";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [svelte(), removeConsole()],
  server: {
    proxy: {
      "/api": {
        target: "http://pc:50831",
        changeOrigin: true,
        secure: false,
        ws: true,
      },
    },
  },
});
