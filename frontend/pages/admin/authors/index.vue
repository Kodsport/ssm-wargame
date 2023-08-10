<template>
    <div class="container">
        <h1>Authors</h1>

        <div>
            <div class="form-group">
                <label>Name</label>
                <input class="form-control" type="text" v-model="form.full_name">
            </div>
            <div class="form-group">
                <label>Slug</label>
                <input class="form-control" type="text" v-model="form.slug">
            </div>
            <div class="form-group">
                <label>Description</label>
                <textarea class="form-control" type="text" v-model="form.description"></textarea>
            </div>
            <div class="form-group">
                <label>Image URL</label>
                <input class="form-control" type="text" v-model="form.image_url">
            </div>
            <div class="form-group">
                <input class="form-check-input" type="checkbox" v-model="form.sponsor">
                <label class="ps-2">Sponsor</label>
            </div>
            <div class="form-group">
                <input class="form-check-input" type="checkbox" v-model="form.publish">
                <label class="ps-2">Publish</label>
            </div>
            <button class="btn btn-primary mt-3" @click="create">
                Create
            </button>
        </div>


        <div class="pt-3">
            <table class="table">
                <thead>
                    <tr>
                        <th>Name</th>
                        <th></th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="a in store.authors">
                        <td>
                            {{ a.full_name }}
                            <span v-if="!a.publish" class="material-symbols-outlined">visibility_off</span>
                            <span v-if="a.sponsor" class="material-symbols-outlined">monetization_on</span>
                        </td>
                        <td>
                            <nuxt-link class="btn btn-secondary" :to="`/admin/authors/${a.slug}/edit`">Edit</nuxt-link>
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>

    </div>
</template>

<script setup lang="ts">
import { useChallengeStore } from '../../../store/admin/challenges'

const store = useChallengeStore()
const http = useHttp()



const form = ref({
    full_name: '',
    slug: '',
    description: '',
    image_url: '',
    sponsor: '',
    publish: '',
})


onMounted(() => {
    store.getAuthors()
})

async function create() {
    http('/admin/author', {
        method: 'POST',
        body: form.value
    })
}

</script>