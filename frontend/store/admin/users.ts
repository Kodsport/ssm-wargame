import { defineStore } from 'pinia'
import useHttp from '../../composables/use-http'
const http = useHttp()

export const useUserStore = defineStore('users', {
    state: () => ({
        users: [],
        userScoreboard: []
    }),
    actions: {
        async getUsers() {
            const challs = await http('/admin/users')
            this.users = challs
        },
        async getUserScoreboard() {
            const scoreboard = await http('/user_scoreboard')
            this.userScoreboard = scoreboard
            return scoreboard
        }
    },
    getters: {
        byId: (state) => (id: string) => state.users.find((u: any) => u.id === id),
        getUserSubmissions: (state) => (userId: string) =>
            state.userScoreboard.filter((entry: any) => entry.user_id === userId)
    }
})