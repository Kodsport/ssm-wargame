<template>
    <div class="container">
        <h1>Challs</h1>
        <div class="col-6">
            <form @submit.prevent>
                <div class="form-group">
                    <label>Title</label>
                    <input class="form-control" placeholder="Enter title" v-model="form.title" />
                </div>
                <div class="form-group">
                    <label>Slug (shown in url)</label>
                    <input class="form-control" type="text" placeholder="Enter slug" v-model="form.slug" />
                </div>
                <div class="form-group">
                    <label>Score</label>
                    <input class="form-control" type="number" placeholder="Enter score" v-model.number="form.score" />
                </div>
                <div class="form-group">
                    <label>Description</label>
                    <textarea class="form-control" placeholder="Enter description" v-model="form.description" />
                </div>
                <div>
                    <button class="btn btn-primary" @click="createChall">Create</button>
                </div>
            </form>
        </div>
        <div>
            <table class="table">
                <thead>
                    <tr>
                        <th>Title</th>
                        <th>Solvers</th>
                        <th />
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="chall in challenges" :key="chall.id">
                        <td>{{ chall.title }}</td>
                        <td>{{ chall.solves }}</td>
                        <td class="text-right">
                            <NuxtLink class="btn btn-primary" :to="`/admin/challenges/${chall.slug}`">
                                Edit
                            </NuxtLink>
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
</template>
  
<script setup lang="ts">
import { computed, ref } from 'vue'
import { useChallengeStore } from '../../../store/admin/challenges'

const http = useHttp()



const store = useChallengeStore()

onMounted(() => {
    store.getChallenges()
})

const challenges = computed(() => store.challenges)

const form = ref({
    title: "",
    slug: "",
    score: 0,
    description: ""
})

async function createChall() {
    await http("/admin/challenges", {
        method: 'POST',
        body: form.value
    });

    store.getChallenges()
}



</script>