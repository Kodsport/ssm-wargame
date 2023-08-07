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
      ],
      title: 'Säkerhets-SM',
      meta: [
        {
          name: 'description', content: 'Säkerhets-SM är en tävling inom cybersäkerhet för grundskolan och gymnasiet'
        }
      ]
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
