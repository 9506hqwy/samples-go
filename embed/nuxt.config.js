import { defineNuxtConfig } from "nuxt/config";

export default defineNuxtConfig({
  nitro: {
    output: {
      publicDir: "../dist"
    }
  },
  experimental: {
    payloadExtraction: false
  }
});
