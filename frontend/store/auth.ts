import { defineStore } from 'pinia'

export const useAuthStore = defineStore('auth', {
    state: () => ({
        token: "",
    }),
    actions: {
        async setToken(token: string) {
            this.token = token
        }
    }
})