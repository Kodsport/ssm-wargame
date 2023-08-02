<template>
    <div>
        <h1>Poängtavla</h1>
        <table class="table">
            <thead>
                <th>Rang</th>
                <th>Skola</th>
                <th>Poäng</th>
            </thead>
            <tbody>
                <tr v-for="score, i in scoreboard.sort((a, b) => a.score < b.score)" :key="score.school_name">
                    <td>{{ i + 1 }}</td>
                    <td>{{ score.school_name }}</td>
                    <td>{{ score.score }}</td>
                </tr>
            </tbody>
        </table>
    </div>
</template>

<script setup lang="ts">

const http = useHttp()

var scoreboard = ref([])

onMounted(async () => {
    const resp = await http('/scoreboard')
    scoreboard.value = resp.scores
})

</script>