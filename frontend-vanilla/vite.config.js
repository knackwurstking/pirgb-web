import { defineConfig } from "vite";
import removeConsole from "vite-plugin-remove-console";

export default defineConfig({
  plugins: [removeConsole()],
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
