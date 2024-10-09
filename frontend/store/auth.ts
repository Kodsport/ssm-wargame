import { defineStore } from 'pinia'
import useHttp from '../composables/use-http'


export const useAuthStore = defineStore('auth', {
    state: () => ({
        token: "",
        user: {},
        knackKodenPassword: "",
        knackKodenData: null
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
                this.token = ''
            }
        },
        async getKnackKodenData(pw: string | null) {
            const http = useHttp()

            const resp = await http("/knack_koden_get_class", {
                method: "POST",
                body: {
                    password: pw || this.knackKodenPassword
                }
            })
            this.knackKodenData = resp;
            if (pw) {
                this.knackKodenPassword = pw
            }

        }
    },
    getters: {
        isAuthed(state) {
            return !!state.user.id
        }
    }
})