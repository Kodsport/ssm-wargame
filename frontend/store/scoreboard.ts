import { defineStore } from 'pinia'
import useHttp from '../composables/use-http'
const http = useHttp()

export const useScoreboardStore = defineStore('scoreboard', {
    state: () => ({
        userScores: [],
        schoolScores: [],
    }),
    actions: {
        async getUserScoreboards() {
            const scores = await http('/user_scoreboard')
            this.userScores = scores.scores.sort((a, b) => a.score < b.score)
        },
        async getSchoolScoreboards() {
            const scores = await http('/scoreboard')
            this.schoolScores = scores.scores.sort((a, b) => a.score < b.score)
        },
    },
    getters: {
        uniScores: (state) => state.schoolScores.filter(s => s.is_university),
        nonUniScores: (state) => state.schoolScores.filter(s => !s.is_university),
    }
})