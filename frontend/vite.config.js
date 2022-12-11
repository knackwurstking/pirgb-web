import { defineConfig } from "vite";
import { svelte } from "@sveltejs/vite-plugin-svelte";
import { VitePWA } from "vite-plugin-pwa";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [svelte(), VitePWA({ registerType: "autoUpdate" })],
  server: {
    proxy: {
      "/api/v1/devices": {
        target: "https://localhost:50831",
        changeOrigin: true,
        secure: false,
      },
    },
  },
});
