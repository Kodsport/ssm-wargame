import { defineStore } from 'pinia'
import useHttp from '../composables/use-http'
const http = useHttp()

export const useChallengeStore = defineStore('challenges', {
    state: () => ({
        challenges: [],
        monthlies: [],
    }),
    actions: {
        async getChallenges() {
            const challs = await http('/challenges')
            this.challenges = challs
        },
        async getMonthlies() {
            const challs = await http('/monthly_challenges')
            this.monthlies = challs
        }
    },
    getters: {
        getBySlug(state) {
            return (slug: string) => state.challenges.find(c => c.slug == slug)
        },
        getCurrentMontly: (state) => {
            const now = new Date().valueOf() / 1000;
            return state.monthlies.find(m => m.start_date < now && m.end_date > now)
        }
    }
})