<template>
    <ChallengeModal v-if="chall" :chall="chall" @back="router.push(`/challenges`)" />
</template>

<script setup lang="ts">
import { useChallengeStore } from '../../store/challenges';

const route = useRoute()
const router = useRouter()
const store = useChallengeStore()

const chall = computed(() => store.getBySlug(route.params.slug))

const title = `SSM - ${chall.value.title}`

useHead({
    title,
})


useServerSeoMeta({
    ogTitle: title,
    title: title,
    description: chall.value.description,
    ogDescription: chall.value.description,
})

</script>
