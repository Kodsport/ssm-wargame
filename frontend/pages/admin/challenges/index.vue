<template>
    <div class="container">
        <h1>Challs</h1>
        <div>
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
                    <input class="form-control" type="number" placeholder="Enter score" v-model.number="form.score"
                        step="50" min="0" max="1500" />
                </div>
                <div class="form-group">
                    <label>Description</label>
                    <textarea class="form-control" placeholder="Enter description" v-model="form.description" />
                </div>
                <div class="form-group">
                    <label>Category</label>
                    <select class="form-control" name="" v-model="form.category_id">
                        <option value="" disabled selected>Category</option>
                        <option v-for="cat in store.categories" :value="cat.id">{{ cat.name }}</option>
                    </select>
                </div>
                <div class="form-group">
                    <label>Authors</label>
                    <ul>
                        <li v-for="(author, i) in form.authors">{{ users.byId(author).full_name }} ({{
                            users.byId(author).email }}) <span @click="form.authors = form.authors.splice(i, i)">[x]</span>
                        </li>
                    </ul>
                    <select ref="authSelect" class="form-control" @change="pushAuthor">
                        <option value="" disabled selected>Choose Authors</option>
                        <option v-for="user in users.users.filter(u => form.authors.indexOf(u.id) == -1)" :value="user.id">
                            {{ user.email }}</option>
                    </select>
                </div>
                <div class="form-group">
                    <label>Publish immediately</label>
                    <input class="" type="checkbox" v-model="publishImm">
                </div>
                <div class="form-group" v-if="!publishImm">
                    <label>Publish at</label>
                    <input class="form-control" type="datetime-local" placeholder="Enter publish date"
                        v-model="form.publishAt" />
                </div>
                <div class="pt-3">
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
                        <th>Category</th>
                        <th />
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="chall in challenges" :key="chall.id">
                        <td>
                            <NuxtLink :to="`/admin/challenges/${chall.slug}`">
                                {{ chall.title }}
                                <span v-if="!chall.flags" title="No flags">⚠️⚠️⚠️</span>
                            </NuxtLink>
                        </td>
                        <td>{{ chall.solves }}</td>
                        <td>{{ getCategoryName(chall.category_id) }}</td>
                        <td class="text-right">
                            <NuxtLink class="btn btn-primary" :to="`/admin/challenges/${chall.slug}/edit`">
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
import { useUserStore } from '../../../store/admin/users';

const http = useHttp()

const store = useChallengeStore()
const users = useUserStore()

onMounted(() => {
    store.getChallenges()
    store.getCategories()
    users.getUsers()
})

const challenges = computed(() => store.challenges)

const publishImm = ref(false)
const form = ref({
    title: "",
    slug: "",
    score: 0,
    description: "",
    publishAt: "",
    category_id: "",
    authors: []
})

async function createChall() {
    await http("/admin/challenges", {
        method: 'POST',
        body: {
            title: form.value.title,
            slug: form.value.slug,
            score: form.value.score,
            description: form.value.description,
            category_id: form.value.category_id,
            publish_at: publishImm.value ? null : new Date(form.value.publishAt).valueOf() / 1000,
            authors: form.value.authors
        }
    });

    store.getChallenges()
}

function pushAuthor(event) {
    form.value.authors.push(event.target.value)
    authSelect.value.value = ''
}

function getCategoryName(id) {
    return store.getCategory(id)?.name
}

</script>