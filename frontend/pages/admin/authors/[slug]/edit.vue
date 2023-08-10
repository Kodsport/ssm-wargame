<template>
    <div class="container">
        <div v-if="a">
            <h1>{{ a.full_name }}</h1>
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
            <button class="btn btn-primary mt-3" @click="update">
                Update
            </button>
        </div>
    </div>
</template>


<script setup lang="ts">
import { useChallengeStore } from '../../../../store/admin/challenges'

const store = useChallengeStore()
const route = useRoute()
const http = useHttp()

const a = computed(() => store.getAuthorBySlug(route.params.slug))

const form = ref({
    full_name: '',
    slug: '',
    description: '',
    image_url: '',
    sponsor: '',
    publish: '',
})

onMounted(async () => {
    await store.getAuthors()
    form.value = store.getAuthorBySlug(route.params.slug)
})

async function update() {
    await http('/admin/authors/' + a.value.id, {
        method: 'PUT',
        body: form.value
    })
    await store.getAuthors()
}

</script>