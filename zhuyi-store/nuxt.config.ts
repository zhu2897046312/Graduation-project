// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  modules: [
    '@nuxt/eslint',
    '@nuxt/ui'
  ],

  devtools: {
    enabled: true
  },

  css: ['~/assets/css/main.css'],

  // 运行时配置
  // - public.apiUrl：浏览器用，由 NUXT_PUBLIC_API_URL 覆盖（如 http://localhost:8080/api/client）
  // - apiUrl（仅服务端）：SSR 时容器内请求 backend，由 NUXT_API_URL 覆盖（如 http://backend:8080/api/client）
  runtimeConfig: {
    apiUrl: 'http://localhost:8080/api/client',
    public: {
      apiUrl: 'http://localhost:8080/api/client'
    }
  },

  routeRules: {
    // 关闭首页预渲染，改为每次访问都 SSR，这样容器内会请求 backend，否则预渲染在构建时执行拿不到数据且之后不会请求后端
    '/': { ssr: true }
  },

  // 开发服务器配置（自定义端口）
  devServer: {
    port: 3009
  },

  compatibilityDate: '2025-01-15',

  eslint: {
    config: {
      stylistic: {
        commaDangle: 'never',
        braceStyle: '1tbs'
      }
    }
  }
})
