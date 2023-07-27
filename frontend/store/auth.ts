import { defineStore } from 'pinia'
import useHttp from '../composables/use-http'


export const useAuthStore = defineStore('auth', {
    state: () => ({
        token: "",
        user: {}
    }),
    actions: {
        setToken(token: string) {
            this.token = token
            this.getUser()
        },
        async getUser() {
            try {
                const http = useHttp()

                const user = await http('/user/self')
                this.user = user
            } catch {
                this.user = {}
                // this.token = ''
            }
        }
    }
})