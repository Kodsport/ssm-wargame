<template>
    <div>
        <div>
            <div v-if="monthly">
                <h1 class="text-primary">Monthly challenge - {{ monthly.display_month }}</h1>
                <div>
                    <challenge :chall="monthly.challenge"></challenge>
                </div>
            </div>


        </div>

    </div>
</template>

<script setup lang="ts">
import { useChallengeStore } from '../store/challenges'
import { useAuthStore } from '../store/auth'

const http = useHttp()
const challs = useChallengeStore()
const auth = useAuthStore()

const monthly = computed(() => challs.getCurrentMonthly)

onMounted(() => {
    challs.getChallenges()
    challs.getMonthlies()
})


</script>