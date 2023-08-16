import { defineStore } from 'pinia'
import useHttp from '../../composables/use-http'
const http = useHttp()

export const useChallengeStore = defineStore('admin-challenges', {
    state: () => ({
        challenges: [],
        categories: [],
        monthlies: [],
        events: [],
        authors: [],
        courses: [],
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
        },
        async getAuthors() {
            const authors = await http('/admin/authors')
            this.authors = authors
        },
        async getCourses() {
            const courses = await http('/admin/courses')
            this.courses = courses
        }
    },
    getters: {
        getBySlug: (state) => (slug: string) => state.challenges.find(c => c.slug == slug),
        getAuthorBySlug: (state) => (slug: string) => state.authors.find(c => c.slug == slug),
        getById: (state) => (id: string) => state.challenges.find(c => c.id == id),
        getCategory: (state) => (id: string) => state.categories.find(c => c.id == id),
        getCurrentMonthly: () => null
    }
})