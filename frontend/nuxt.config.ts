import { dev } from "process";

// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },
  pages: true,
  modules: [
    '@pinia/nuxt',
    // 'nuxt-security' // TODO fix later
  ],
  css: [
    '~/assets/scss/main.scss'
  ],
  app: {
    head: {
      link: [
        {
          rel: 'stylesheet', href: 'https://fonts.googleapis.com/icon?family=Material+Icons'
        }
      ]
    }
  },
  runtimeConfig: {
    public: {
      apiBase: 'http://localhost:8000/'
    }
  }
})
