<template>
    <ChallengeModal v-if="chall" :chall="chall" @back="router.push(`/challenges`)" />
</template>

<script setup lang="ts">
import { useChallengeStore } from '../../store/challenges';

const route = useRoute()
const router = useRouter()
const store = useChallengeStore()

const chall = computed(() => {
    const c = store.getBySlug(route.params.slug)

    if (!c) {
        throw createError({
            statusCode: 404,
            statusMessage: 'Page Not Found'
        })
    }

    return c;
})



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
