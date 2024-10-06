import { defineStore } from 'pinia'
import useHttp from '../composables/use-http'


export const useAuthStore = defineStore('auth', {
    state: () => ({
        token: "",
        user: {},
        knackKodenPassword: ""
    }),
    actions: {
        setKnackKodenPassword(pw: string) {
            this.knackKodenPassword = pw;
        },
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
                this.token = ''
            }
        }
    },
    getters: {
        isAuthed(state) {
            return !!state.user.id
        }
    }
})