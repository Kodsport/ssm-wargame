<template>
    <div class="border border-dark bg-dark rounded p-4">
        <h4>{{ props.chall.title }}</h4>

        <div class="d-flex pb-2" v-if="props.chall?.authors?.length">
            <span class="material-icons text-primary pe-1">group</span>
            <span class="material-icons text-primary pe-1">edit</span>
            <span>{{ props.chall.authors.map(a => a.full_name).join(', ') }}</span>
        </div>

        <div class="d-flex pb-2" v-for="file in props.chall.files">
            <span class="material-icons text-primary">description</span>
            <a :href="file.url">{{ file.filename }}</a>
        </div>

        <p v-html="renderMarkdown(props.chall.description)"></p>

        <div class="pt-4 pb-3">
            <div v-if="!!auth.user.id">
                <div v-if="!props.chall.solved" class="alert" v-bind:class="{ 'alert-danger': warn }">
                    <input class="form-control" type="text" placeholder="SSM{..." v-model="flag"
                        @keypress.enter="submitFlag">
                </div>
                <InputReplacer v-else text="Löst!" />

            </div>
            <div v-else>
                <InputReplacer text="Logga in för att lösa skicka in flaggor" />
            </div>
        </div>
        <p>{{ props.chall.solves }} lösare</p>
    </div>
</template>

<script setup lang="ts">
import { useChallengeStore } from '../store/challenges';
import { useAuthStore } from '../store/auth';
import renderMarkdown from '../utils/markdown';


const challs = useChallengeStore()
const auth = useAuthStore()
const props = defineProps(['chall'])
const http = useHttp()

const flag = ref("")
const warn = ref(false)

async function submitFlag() {
    let f = flag.value.trim()
    if (f == "") {
        return;
    }
    try {
        await http(`/challenges/${props.chall.id}/attempt`, {
            method: 'POST',
            body: {
                flag: f
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