<template>
    <div>
        <div>
            <div v-if="montly">
                <h1>Montly challenge - {{ montly.display_month }}</h1>
                <div>
                    <challenge :chall="montly.challenge"></challenge>
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

const montly = computed(() => challs.getCurrentMontly)

onMounted(() => {
    challs.getChallenges()
    challs.getMonthlies()
})


</script>