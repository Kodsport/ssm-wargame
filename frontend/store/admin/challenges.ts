import { defineStore } from 'pinia'
import useHttp from '../../composables/use-http'
const http = useHttp()

export const useChallengeStore = defineStore('admin-challenges', {
    state: () => ({
        challenges: [],
        categories: [],
    }),
    actions: {
        async getChallenges() {
            const challs = await http('/admin/challenges')
            this.challenges = challs
        },
        async getCategories() {
            const cats = await http('/admin/categories')
            this.categories = cats
        }
    },
    getters: {
        getBySlug: (state) => (slug: string) => state.challenges.find(c => c.slug == slug),
        getCategory: (state) => (id: string) => state.categories.find(c => c.id == id)
    }
})