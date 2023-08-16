<template>
    <div class="container">
        <div class="d-flex justify-content-center">
            <h1>Kurser</h1>
        </div>
        <div class="row pt-5">
            <div class="col-3 h-100" v-for="cat in categories">
                <div class="rounded"
                    :class="{ 'bg-info border border-1 border-primary': selectedCategory == cat, 'bg-dark': selectedCategory != cat }"
                    @click="selectedCategory = cat">
                    <div class="d-flex justify-content-center py-3">
                        <h2 class="text-primary">{{ cat }}</h2>
                    </div>

                    <div class="p-3 d-flex justify-content-center">
                        <h5>{{ store.courses.filter(a => a.category == cat).length }} kurser</h5>
                    </div>
                </div>
            </div>
        </div>
        <div v-if="selectedCategory" class="pt-5">

            <div v-for="[diff, diffName] in difficulties">
                <h3 class="border-bottom py-3">{{ diffName }}</h3>
                <div>
                    <div class="bg-info p-3 m-3 rounded"
                        v-for="course in store.courses.filter(c => c.category == selectedCategory && c.difficulty == diff)">
                        <div class="d-flex justify-content-between">
                            <span class="d-flex justify-content-center flex-column">{{ course.title }}</span>
                            <div>
                                <span v-if="course.enrolled && !course.completed"
                                    class="badge bg-primary text-dark me-3">Påbörjad</span>
                                <span v-else-if="course.enrolled && course.completed"
                                    class="badge bg-primary text-dark me-3">Avslutad</span>

                                <span class="badge bg-primary text-dark me-3">{{ course.course_items.length }} moment</span>
                                <button class="btn btn-primary" @click="router.push(`/courses/${course.slug}`)">
                                    {{ course.enrolled ? 'Visa' : 'Förhandsgranska' }}
                                </button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>


        </div>
    </div>
</template>

<script setup lang="ts">
import { useChallengeStore } from '../../store/challenges';

const store = useChallengeStore()
const router = useRouter()
const selectedCategory = ref("web")

const categories = ["web", "pwn", "crypto", "rev"];
const difficulties = [['easy', 'enkla'], ['difficult', 'svåra'], ['advanced', 'advancerade']];

await useAsyncData('courses', store.getCourses)

</script>