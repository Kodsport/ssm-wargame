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
                <label for="imageUrl">Image URL</label>
                <input class="form-control" type="text" v-model="form.image_url" id="imageUrl">
            </div>
            <div class="form-group">
                <input class="form-check-input" type="checkbox" v-model="form.sponsor" id="isSponsor">
                <label class="ps-2" for="isSponsor">Sponsor</label>
            </div>
            <div class="form-group">
                <input class="form-check-input" type="checkbox" v-model="form.publish" id="publish">
                <label class="ps-2" for="publish">Publish</label>
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
                            <nuxt-link class="btn btn-secondary me-2" :to="`/admin/authors/${a.slug}/edit`">Edit</nuxt-link>
                            <button class="btn btn-danger" @click="deleteAuthor(a.id)">Delete</button>
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
    sponsor: false,
    publish: false,
})


onMounted(() => {
    store.getAuthors()
})

async function create() {
    await http('/admin/authors', {
        method: 'POST',
        body: form.value
    })
    store.getAuthors()
}

async function deleteAuthor(id) {
    await http('/admin/authors/' + id, {
        method: 'DELETE',
    })
    store.getAuthors()

}

</script>