// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devServer: {
    port: 3001, // 设置启动端口为 3000
  },
  compatibilityDate: '2024-11-01',
  devtools: { enabled: true},
  modules: ['@nuxtjs/tailwindcss', 'nuxtjs-naive-ui'],
  css: [
    '@/assets/css/main.css'
  ],
  nitro: {
    preset: 'node-server', 
  },
  app: {
    head: {
      charset: 'utf-8',
      viewport: 'width=device-width, initial-scale=1',
      htmlAttrs: {
        style: '--nuxt-devtools-safe-area-top: 0px; --nuxt-devtools-safe-area-right: 0px; --nuxt-devtools-safe-area-bottom: 0px; --nuxt-devtools-safe-area-left: 0px;'
      },
    }
  },
  $development: {
    runtimeConfig: {
      public: {
        apiUrl: 'http://localhost:8080/api/client'
      }
    }
  },
  $production: {
    routeRules: {
      '/**': { isr: true }
    },
    runtimeConfig: {
      public: {
        apiUrl: 'https://www.earring18.com/api/client'
      }
    }
  },
})