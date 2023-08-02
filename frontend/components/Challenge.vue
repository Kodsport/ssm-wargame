<template>
    <div v-bind:class="{ 'alert': props.chall.solved, 'alert-success': props.chall.solved }">
        <h5>{{ props.chall.title }}</h5>
        <p>{{ props.chall.description }}</p>
        <p>{{ props.chall.solves }} lösare</p>


        <div v-if="props.chall.files">
            Filer:
            <a class="px-2" v-for="file in props.chall.files" :href="file.url">{{ file.filename }}</a>
        </div>
        <div v-if="!props.chall.solved && !!auth.user.id" class="alert" v-bind:class="{ 'alert-danger': warn }">
            <input class="form-control" type="text" placeholder="SSM{..." v-model="flag" @keypress.enter="submitFlag">
        </div>
    </div>
</template>

<script setup lang="ts">
import { useChallengeStore } from '../store/challenges';
import { useAuthStore } from '../store/auth';


const challs = useChallengeStore()
const auth = useAuthStore()
const props = defineProps(['chall'])
const http = useHttp()

const flag = ref("")
const warn = ref(false)

async function submitFlag() {
    try {
        await http(`/challenges/${props.chall.id}/attempt`, {
            method: 'POST',
            body: {
                flag: flag.value
            }
        })

        challs.getChallenges()
        challs.getMonthlies()

    } catch (error) {
        console.log(error)
        warn.value = true
        setTimeout(() => {
            warn.value = false
        }, 1000)
    }

}

</script>