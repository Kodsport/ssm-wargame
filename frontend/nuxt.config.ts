// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },
  pages: true,
  modules: [
    "@pinia/nuxt",
    // 'nuxt-security' // TODO fix later
  ],
  css: ["~/assets/scss/main.scss"],
  app: {
    head: {
      link: [
        {
          rel: "stylesheet",
          href: "https://fonts.googleapis.com/css2?family=Material+Symbols+Outlined:opsz,wght,FILL,GRAD@20..48,100..700,0..1,-50..200",
        },
        {
          rel: "stylesheet",
          href: "https://cdn.jsdelivr.net/npm/bootstrap-icons@1.11.0/font/bootstrap-icons.css",
        },
      ],
      title: "Säkerhets-SM",
      meta: [
        {
          name: "description",
          content:
            "Säkerhets-SM är en tävling inom cybersäkerhet för grundskolan och gymnasiet",
        },
      ],
      htmlAttrs: {
        lang: "se",
      },
      script: [
        {
          // <script defer src="https://analytics.sakerhetssm.se/script.js" data-website-id="ec04f82f-8090-4b5a-b404-5a07099e1d78"></script>
          defer: true,
          src: "https://analytics.sakerhetssm.se/script.js",
          "data-website-id": "ec04f82f-8090-4b5a-b404-5a07099e1d78",
        },
      ],
    },
  },
  runtimeConfig: {
    public: {
      apiBase: process.env.NUXT_PUBLIC_API_BASE || "http://localhost:8000",
    },
  },
  routeRules: {
    "/admin/**": { ssr: false },
  },
  vite: {
    vue: {
      script: {
        defineModel: true,
      },
    },
  },
});
