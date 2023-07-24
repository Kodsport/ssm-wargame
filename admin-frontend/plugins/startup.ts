import { Plugin } from '@nuxt/types'

const plugin: Plugin = (context) => {
  // The JWT-token is persisted in localStorage
  const token = localStorage.getItem('ssm_token')
  if (token) {
    context.$axios.setToken(token, 'Bearer')
  }
}

export default plugin
