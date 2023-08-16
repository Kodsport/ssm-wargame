<template>
    <div class="container">

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
                <label>Description</label>
                <textarea class="form-control" placeholder="Enter description" v-model="form.description" />
            </div>
            <div class="form-group">
                <label>Category</label>
                <select class="form-control" name="" v-model="form.category">
                    <option selected disabled value="">choose category</option>
                    <option v-for="cat in ['pwn', 'web', 'rev', 'crypto']" :value="cat">{{ cat }}</option>
                </select>
            </div>
            <div class="form-group">
                <label>Difficulty</label>
                <select class="form-control" name="" v-model="form.difficulty">
                    <option selected disabled value="">choose difficulty</option>
                    <option v-for="diff in ['easy', 'difficult', 'advanced']" :value="diff">{{ diff }}</option>
                </select>
            </div>
            <div class="form-group pt-2">
                <label>Publish</label>
                <input class="ms-2" type="checkbox" v-model="form.publish">
            </div>
            <div class="pt-3">
                <button class="btn btn-primary" @click="updateCourse">Update</button>
            </div>
        </form>
    </div>
</template>

<script setup lang="ts">
import { useChallengeStore } from '~/store/admin/challenges'

const http = useHttp()
const route = useRoute()
const router = useRouter()
const store = useChallengeStore()

const form = ref({
    title: "",
    slug: "",
    difficulty: "",
    category: "",
    description: "",
    publish: false,
    author_ids: []
})

const course = computed(() => store.courses.find(c => c.slug == route.params.slug))

onMounted(async () => {
    store.getAuthors()
    if (!course.value) {
        await store.getCourses()
    }

    form.value = {
        title: course.value.title,
        slug: course.value.slug,
        difficulty: course.value.difficulty,
        category: course.value.category,
        description: course.value.description,
        publish: course.value.publish,
        author_ids: course.value.author_ids
    }
})


async function updateCourse() {
    console.log(form.value)
    await http(`/admin/courses/${course.value.id}/edit`, {
        method: 'PUT',
        body: form.value
    })
    await store.getCourses()
    router.replace(`/admin/course/${form.value.slug}`)
}
</script>