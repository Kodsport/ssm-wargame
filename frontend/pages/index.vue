<template>
    <div>
        <div>
            <h1>challs</h1>

            <Challenge v-for="chall in challs.challenges" v-bind:chall="chall"></Challenge>
        </div>
    </div>
</template>

<script setup lang="ts">
import { useChallengeStore } from '../store/challenges'

const challs = useChallengeStore()

onMounted(() => {
    challs.getChallenges()
})

</script>