import { VitePWA } from "vite-plugin-pwa";

export default defineConfig({
  plugins: [
    VitePWA({
      injectRegister: "auto",
    }),
  ],
});
