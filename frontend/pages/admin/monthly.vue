<template>
    <div <h1>Monthly challenge</h1>

        <table class="table">
            <thead>
                <tr>
                    <th>Title</th>
                    <th>Solvers</th>
                    <th>Interval</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="[monthly, chall] in monthlies" :key="monthly.id">
                    <td>{{ chall.title }}</td>
                    <td>{{ chall.solvers }}</td>
                    <td>
                        {{ monthly.start_date + " -> " + monthly.end_date }} ({{ monthly.display_month }})
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
</template>
  
<script lang="ts" setup>
import { useChallengeStore } from '../../store/challenges'

const chall = useChallengeStore()

const monthlies = ref([])


</script>