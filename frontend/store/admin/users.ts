import { defineStore } from 'pinia'
import useHttp from '../../composables/use-http'
const http = useHttp()

export const useUserStore = defineStore('users', {
    state: () => ({
        users: [],
    }),
    actions: {
        async getUsers() {
            const challs = await http('/admin/users')
            this.users = challs
        }
    },
    getters: {
        byId: (state) => (id: string) => state.users.find(u => u.id === id)
    }
})