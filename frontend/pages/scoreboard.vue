<template>
    <div class="container">
        <div class="row">
            <div class="col-12 col-md-6">
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
                        <tr v-for="score, i in userScoreboard.data.value.scores" :key="score.user_id">
                            <td>{{ i + 1 }}</td>
                            <td>{{ score.name }}</td>
                            <td>{{ score.score }}</td>
                        </tr>
                    </tbody>
                </table>

            </div>
            <div class="col-12 col-md-6">
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
                        <tr v-for="score, i in schoolScoreboard.data.value.scores" :key="score.school_name">
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

const http = useHttp()

const userScoreboard = await useAsyncData('userScoreboard', () => http('/user_scoreboard'))
const schoolScoreboard = await useAsyncData('schoolScoreboard', () => http('/scoreboard'))

useHead({
    title: 'SSM - Poängtavla'
})


</script>