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
                <div class="pt-3">
                    <button class="btn btn-primary" @click="createCourse">Create</button>
                </div>
            </form>
        </div>
        <div>
            <table class="table">
                <thead>
                    <tr>
                        <th>Name</th>
                        <th>Category</th>
                        <th>Difficulty</th>
                        <th />
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="course in store.courses" :key="course.id">
                        <td>
                            {{ course.title }}
                        </td>
                        <td>{{ course.category }}</td>
                        <td>{{ course.difficulty }}</td>
                        <td class="text-right">
                            <NuxtLink class="btn btn-primary" :to="`/admin/courses/${course.slug}/edit`">
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
import { useChallengeStore } from '../../../store/admin/challenges'
import { useUserStore } from '../../../store/admin/users';

const http = useHttp()

const store = useChallengeStore()
const users = useUserStore()

const form = ref({
    title: "",
    slug: "",
    difficulty: "",
    category: "",
    description: "",
    publish: false,
    author_ids: []
})

onMounted(() => {
    store.getCourses()
    store.getAuthors()
})

async function createCourse() {
    await http('/admin/courses', {
        method: 'POST',
        body: form.value
    })
    await store.getCourses()
}

</script>