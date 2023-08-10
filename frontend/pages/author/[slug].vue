<template>
    <div class="container">
        <h1>{{ author.full_name }}</h1>
        <div v-html="renderMarkdown(author.description)"></div>
    </div>
</template>

<script setup lang="ts">
import renderMarkdown from '../../utils/markdown';
import { useChallengeStore } from '../../store/challenges'

const route = useRoute()
const http = useHttp()
const challs = useChallengeStore()

const authors = await useAsyncData('authors', () => http('/authors'))
const author = computed(() => authors.data.value.find(a => a.slug == route.params.slug))

onMounted(() => {
    challs.getChallenges()
})

</script>