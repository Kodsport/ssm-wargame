<template>
    <ChallengeModal v-if="chall" :chall="chall" @back="router.push(`/knackkoden`)" />
</template>

<script setup lang="ts">
import { useChallengeStore } from '../../store/challenges';

const route = useRoute()
const router = useRouter()
const store = useChallengeStore()

const chall = computed(() => store.getBySlug(route.params.slug))

const title = `SSM - ${chall.value.title}`
const url = useRequestURL()

useHead({
    title,
    link: [
        {
            rel: 'canonical',
            href: `${url.protocol}//${url.host}/challenges/` + chall.value.slug,
        },
    ]
})


useServerSeoMeta({
    ogTitle: title,
    title: title,
    description: chall.value.description,
    ogDescription: chall.value.description,
})

</script>
