<template>
    <div v-bind:class="{'alert': props.chall.solved, 'alert-success': props.chall.solved}">
        <h5>{{ props.chall.title }}</h5>
        <p>{{ props.chall.description }}</p>
        <p>{{ props.chall.solves }} Solves</p>

        <div v-if="!props.chall.solved" class="alert" v-bind:class="{ 'alert-danger': warn }">
            <input class="form-control" type="text" placeholder="SSM{" v-model="flag" @keypress.enter="submitFlag">
        </div>
    </div>
</template>

<script setup lang="ts">
import { useChallengeStore } from '../store/challenges';

const challs = useChallengeStore()
const props = defineProps(['chall'])
const http = useHttp()

const flag = ref("")
const warn = ref(false)

async function submitFlag() {
    try {
        const resp = await http(`/challenges/${props.chall.id}/attempt`, {
            method: 'POST',
            body: {
                flag: flag.value
            }
        })

        challs.getChallenges()

    } catch (error) {
        console.log(error)
        warn.value = true
        setTimeout(() => {
            warn.value = false
        }, 1000)
    }

}

</script>