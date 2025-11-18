// https://nuxt.com/docs/api/configuration/nuxt-config
import tailwindcss from "@tailwindcss/vite"

export default defineNuxtConfig({
  compatibilityDate: '2025-07-15',
  devtools: { enabled: true },

  app: {
    head: {
      title: 'Thunderstruck Festival',
      htmlAttrs: {
        lang: 'nl-NL',
      },
      link: [
        { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' },
      ],
    },
  },

  modules: [
    '@nuxt/eslint',
    '@nuxt/icon',
    '@nuxt/fonts',
  ],

  css: [ '~/assets/css/main.css' ],

  runtimeConfig: {
    public: {
      serverBaseURL: process.env.SERVER_BASE_URL,
      clientBaseURL: process.env.CLIENT_BASE_URL,
    },
  },

  vite: {
    plugins: [
      tailwindcss(),
    ],
  },
})
