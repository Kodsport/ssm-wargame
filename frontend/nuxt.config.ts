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
          rel: 'stylesheet', href: "https://fonts.googleapis.com/css2?family=Material+Symbols+Outlined:opsz,wght,FILL,GRAD@20..48,100..700,0..1,-50..200"
        }
      ],
      title: 'Säkerhets-SM',
      meta: [
        {
          name: 'description', content: 'Säkerhets-SM är en tävling inom cybersäkerhet för grundskolan och gymnasiet'
        }
      ],
      htmlAttrs: {
        lang: 'se'
      }
    }
  },
  runtimeConfig: {
    public: {
      apiBase: 'http://localhost:8000/'
    }
  },
  routeRules: {
    '/admin/**': { ssr: false },
  }
})
