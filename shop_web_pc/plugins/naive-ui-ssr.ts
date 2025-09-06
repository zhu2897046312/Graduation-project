import { setup } from '@css-render/vue3-ssr'
import { defineNuxtPlugin } from '#app'

export default defineNuxtPlugin((nuxtApp) => {
  const { collect } = setup(nuxtApp.vueApp)
  useServerHead({
      style: () => {
          const stylesString = collect()
          const stylesArray = stylesString.split(/<\/style>/g).filter((style: any) => style)
          return stylesArray.map((styleString: string) => {
              const match = styleString.match(/<style cssr-id="([^"]*)">([\s\S]*)/)
              if (match) {
                  const id = match[1]
                  return { 'cssr-id': id, children: match[2] }
              }
              return {}
          })
      }
  })
})