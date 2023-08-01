<template>
    <div>
        <h1>Challenges</h1>

        <div v-for="cat in challs.challenges.map(c => c.category).filter((v, i, a) => a.indexOf(v) == i)" :key="cat">
            <h3>{{ cat }}</h3>
            <Challenge class="border-top pt-3" v-for="chall in challs.challenges.filter(v => v.category === cat)"
                v-bind:chall="chall">
            </Challenge>
        </div>
    </div>
</template>

<script setup lang="ts">
import { useChallengeStore } from '../store/challenges'
import { useAuthStore } from '../store/auth'

const http = useHttp()
const challs = useChallengeStore()
const auth = useAuthStore()

onMounted(() => {
    challs.getChallenges()
})

</script>