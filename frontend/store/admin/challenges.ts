import { defineStore } from 'pinia'
import useHttp from '../../composables/use-http'
const http = useHttp()

export const useChallengeStore = defineStore('admin-challenges', {
    state: () => ({
        challenges: [],
        categories: [],
        monthlies: [],
        events: [],
    }),
    actions: {
        async getChallenges() {
            const challs = await http('/admin/challenges')
            this.challenges = challs
        },
        async getMonthlies() {
            const monthlies = await http('/admin/monthly_challenges')
            this.monthlies = monthlies
        },
        async getCategories() {
            const cats = await http('/admin/categories')
            this.categories = cats
        },
        async getEvents() {
            const events = await http('/admin/events')
            this.events = events
        }
    },
    getters: {
        getBySlug: (state) => (slug: string) => state.challenges.find(c => c.slug == slug),
        getById: (state) => (id: string) => state.challenges.find(c => c.id == id),
        getCategory: (state) => (id: string) => state.categories.find(c => c.id == id),
        getCurrentMonthly: () => null
    }
})