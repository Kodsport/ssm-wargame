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
                    <label class="me-2">Hide</label>
                    <input class="" type="checkbox" v-model="form.hide">
                </div>
                <div class="form-group">
                    <label class="me-2">Dynamic scoring</label>
                    <input class="" type="checkbox" v-model="dynamicScoring">
                </div>
                <div v-if="!dynamicScoring" class="form-group">
                    <label>Static score</label>
                    <input class="form-control" type="number" placeholder="Enter score" v-model.number="form.static_score"
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
                <div class="form-group pt-2">
                    <label>Authors</label>
                    <ul>
                        <li v-for="(author, i) in store.authors.filter(a => form.authors.indexOf(a.id) != -1)">
                            {{ author.full_name
                            }}
                            <span @click="form.authors.splice(i, 1)">[x]</span>
                        </li>
                    </ul>
                    <select ref="authSelect" class="form-control" @change="pushAuthor">
                        <option value="" disabled selected>Choose Authors</option>
                        <option v-for="author in store.authors.filter(u => form.authors.indexOf(u.id) == -1)"
                            :value="author.id">
                            {{ author.full_name }}</option>
                    </select>
                </div>
                <div class="form-group">
                    <label class="me-2">Publish immediately</label>
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
                        <th>Score</th>
                        <th />
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="chall in challenges" :key="chall.id">
                        <td>
                            <NuxtLink :to="`/admin/challenges/${chall.slug}`" class="pe-2">
                                {{ chall.title }}
                            </NuxtLink>
                            <span v-if="!chall.flags" title="No flags">⚠️⚠️⚠️</span>
                            <span v-if="chall.hide" class="material-symbols-outlined text-primary">
                                visibility_off
                            </span>
                        </td>
                        <td>{{ chall.solves }}</td>
                        <td>{{ getCategoryName(chall.category_id) }}</td>
                        <td>{{ chall.static_score || 'Dynamic' }}</td>
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

const authSelect = ref(null)

onMounted(() => {
    store.getChallenges()
    store.getCategories()
    users.getUsers()
})

const challenges = computed(() => store.challenges)
const dynamicScoring = ref(true)

const publishImm = ref(false)
const form = ref({
    title: "",
    slug: "",
    static_score: 250,
    description: "",
    publishAt: "",
    hide: true,
    category_id: "",
    authors: []
})

async function createChall() {
    await http("/admin/challenges", {
        method: 'POST',
        body: {
            title: form.value.title,
            slug: form.value.slug,
            static_score: dynamicScoring.value ? null : form.value.static_score,
            description: form.value.description,
            category_id: form.value.category_id,
            publish_at: publishImm.value ? null : new Date(form.value.publishAt).valueOf() / 1000,
            authors: form.value.authors,
            hide: form.value.hide,
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