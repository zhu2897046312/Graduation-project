// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2025-07-15',
  devtools: { enabled: true },

  build: {
    transpile: [
      'naive-ui',
      'vueuc',
      '@css-render/vue3-ssr',
    ],
    
  },
  css: [
    '~/assets/css/main.css',
    // '~/assets/css/naive-ui.css'
  ],
  imports: {
    autoImport: true // 确保自动导入开启
  },
  modules: [ '@nuxtjs/tailwindcss'],
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
      link: [
        { 
          rel: 'stylesheet', 
          href: 'https://cdn.jsdelivr.net/npm/tailwindcss@2.2.19/dist/tailwind.min.css' 
        }
      ]
    },
    
  },
  $development: {
    runtimeConfig: {
      public: {
        apiUrl: 'http://localhost:8080/api/client'
      }
    },
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
  vite: {
    server: {
      watch: {
        usePolling: true, // 在某些环境下更好的文件监听
        interval: 1000    // 轮询间隔
      },
      allowedHosts: ['true']
    }
  },
  routeRules: {
    '/blogs': { redirect: '/' },
  },
})

