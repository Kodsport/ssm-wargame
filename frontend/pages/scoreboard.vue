<template>
    <div class="container-fluid">
        <div class="row">
            <div class="col-12 col-md-4">
                <h1>Personer</h1>
                <table class="table">
                    <thead>
                        <tr>
                            <th>Rang</th>
                            <th>Namn</th>
                            <th>Poäng</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="score, i in scoreboard.userScores.slice(0, 10)" :key="score.user_id">
                            <td>{{ i + 1 }}</td>
                            <td>{{ score.name }}</td>
                            <td>{{ score.score }}</td>
                        </tr>
                    </tbody>
                </table>

            </div>
            <div class="col-12 col-md-4">
                <h1>Skolor</h1>
                <table class="table">
                    <thead>
                        <tr>
                            <th>Rang</th>
                            <th>Skola</th>
                            <th>Poäng</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="score, i in scoreboard.nonUniScores.slice(0, 10)" :key="score.school_name">
                            <td>{{ i + 1 }}</td>
                            <td>{{ score.school_name }}</td>
                            <td>{{ score.score }}</td>
                        </tr>
                    </tbody>
                </table>
            </div>

            <div class="col-12 col-md-4">
                <h1>Universitet</h1>
                <table class="table">
                    <thead>
                        <tr>
                            <th>Rang</th>
                            <th>Skola</th>
                            <th>Poäng</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="score, i in scoreboard.uniScores.slice(0, 10)" :key="score.school_name">
                            <td>{{ i + 1 }}</td>
                            <td>{{ score.school_name }}</td>
                            <td>{{ score.score }}</td>
                        </tr>
                    </tbody>
                </table>
            </div>

        </div>
    </div>
</template>

<script setup lang="ts">
import { useScoreboardStore } from '../store/scoreboard'

const scoreboard = useScoreboardStore()

await useAsyncData('userScoreboard', scoreboard.getUserScoreboards)
await useAsyncData('scoreboard', scoreboard.getSchoolScoreboards)

useHead({
    title: 'SSM - Poängtavla'
})


</script>