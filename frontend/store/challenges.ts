import { defineStore } from 'pinia'
import useHttp from '../composables/use-http'
const http = useHttp()

export const useChallengeStore = defineStore('challenges', {
    state: () => ({
        challenges: [],
    }),
    actions: {
        async getChallenges() {
            const challs = await http('/challenges')
            this.challenges = challs
        }
    },
    getters: {
        getBySlug(state) {
            return (slug: string) => state.challenges.find(c => c.slug == slug)
        }
    }
})