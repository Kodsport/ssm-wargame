<template>
    <div class="container">
        <h1>{{ courseMeta.title }}</h1>
        <div v-html="renderMarkdown(courseMeta.description)"></div>
        <div class="d-flex justify-content-center pt-5">
            <button v-if="!courseMeta.enrolled" class="btn btn-primary" @click="enroll()" :disabled="!auth.isAuthed">
                Påbörja kurs
            </button>
            <button v-else class="btn btn-primary"
                @click="router.push(`/courses/${courseMeta.slug}/${challenges[0].slug}`)">
                Visa kurs
            </button>
        </div>
        <div class="my-5 border-top">
        </div>
        <div>
            <h1>Kursinnehåll</h1>

            <table class="table table-striped table-dark">
                <thead>
                    <tr>
                        <th>Moment</th>
                        <th>Utmaning</th>
                        <th>Poäng</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="item in challenges">
                        <td>{{ item.title }}
                            <span v-if="item.solved" class="material-symbols-outlined text-success">
                                done
                            </span>
                        </td>
                        <td>{{ item.title }}</td>
                        <td>
                            {{ item.score }}
                        </td>
                    </tr>
                </tbody>
            </table>

        </div>
    </div>
</template>

<script setup lang="ts">
import { useChallengeStore } from '~/store/challenges';
import { useAuthStore } from '~/store/auth';

import renderMarkdown from '~/utils/markdown'


const store = useChallengeStore()
const auth = useAuthStore()
const route = useRoute()
const router = useRouter()
const http = useHttp()

const courses = await useAsyncData('courses', store.getCourses)
const courseMeta = computed(() => store.courses.find(c => c.slug == route.params.slug))
const challenges = computed(() => store.courseChalls[courseMeta.value.id])

onMounted(() => {
    courses.refresh()
    store.getCourseChalls(courseMeta.value.id)
})


async function enroll() {
    await http(`/courses/${courseMeta.value.id}/enroll`, {
        method: 'POST'
    })
}

</script>