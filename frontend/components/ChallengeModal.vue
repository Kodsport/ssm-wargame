<template>
    <div class="modal-backdrop fade show"></div>
    <div class="modal modal-xl d-block" tabindex="-1" @click="goBack" ref="modal">
        <div class="modal-dialog">
            <div class="modal-content">


                <div class="pt-3 d-flex justify-content-center">
                    <h2 class="modal-title text-primary fw-bold">{{ props.chall.title }}</h2>
                </div>
                <div class="d-flex justify-content-center">
                    <div class="badge bg-secondary">{{ props.chall.category }}</div>
                </div>

                <h5>
                    <div class="d-flex justify-content-center pt-3">
                        <span class="mx-3 d-flex">
                            <span class="me-1">
                                {{ props.chall.score }}
                            </span>
                            <span class="material-icons text-primary">control_point_duplicate</span>
                        </span>
                        <span class="mx-3 d-flex">
                            <span class="me-1">
                                {{ props.chall.solves }}
                            </span>
                            <span class="material-icons text-primary">flag</span>
                        </span>
                    </div>
                </h5>

                <div class="modal-body">
                    <div class="row">
                        <div class="col">
                            {{ props.chall.description }}
                        </div>
                        <div class="col">
                            <div class="d-flex">
                                <span class="material-icons text-primary pe-1">group</span>
                                <span class="material-icons text-primary pe-1">edit</span>
                                <span>Movitz Sunar</span>
                            </div>
                            <div v-for="file in props.chall.files">

                                <span class="material-icons text-primary">description</span>
                                <a :href="file.url">{{ file.filename }}</a>
                            </div>
                            <div class="pt-3">
                                <b>Första lösarna</b>
                                <ol>
                                    <li>Bumbodosan</li>
                                    <li>SNHT</li>
                                    <li>Royalroppers</li>
                                </ol>
                            </div>
                        </div>
                    </div>
                    <form @submit.prevent>

                        <div class="row mt-3">
                            <div v-if="!props.chall.solved" class="form-group alert"
                                v-bind:class="{ 'alert-danger': warn }">
                                <input v-if="!!auth.user.id" type="text" class="form-control" placeholder="SSM{..."
                                    @keypress.enter="submitFlag" v-model="flag">
                                <p class="alert alert-secondary" v-else>Logga in för att skicka in flaggor</p>
                            </div>
                            <div v-else>
                                Solved!

                            </div>
                        </div>
                    </form>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { useAuthStore } from '../store/auth';
import { useChallengeStore } from '../store/challenges';

const props = defineProps(['chall'])

const router = useRouter();
const auth = useAuthStore()
const http = useHttp()
const store = useChallengeStore()

const modal = ref(null)


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

        store.getChallenges()
        store.getMonthlies()

    } catch (error) {
        console.log(error)
        warn.value = true
        setTimeout(() => {
            warn.value = false
        }, 1000)
    }

}

function goBack(event) {
    if (event.target == modal.value) {
        router.replace('/challenges')
    }
}


</script>
