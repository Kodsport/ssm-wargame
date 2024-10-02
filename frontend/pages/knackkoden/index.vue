<template>
    <h1 class="text-primary">Knäck Koden</h1>
    <p>
        Knäck Koden är en tävling inom hacking för mellanstadieelever. Man kan inte ännu samla poäng för klassen, den
        funktionaliteten kommer senare!
    </p>

    <div class="d-flex flex-column flex-lg-row pt-4">

        <div class="flex-grow-1 d-flex flex-column">
            <template v-for="category in categories">
                <div v-if="challenges.filter(c => c.category == category).length">
                    <div class="text-primary border-primary border-bottom border-3 pt-2 pb-2">
                        <h4 class="text-lowercase">{{ category }}</h4>
                    </div>

                    <div class="ssm-grid pt-3 pb-2">
                        <div v-for="chall in challenges.filter(c => c.category == category)">
                            <ChallengePreview class="pointer" @click="nav(chall.slug)" :chall="chall" />
                            <!-- for seo -->
                            <a class="d-none" :href="`/challenges/${chall.slug}`">{{ chall.title }}</a>
                        </div>
                    </div>
                </div>
            </template>
        </div>
    </div>
    <NuxtPage />
</template>

<script setup lang="ts">
import { useChallengeStore } from '../../store/challenges'

useHead({
    title: 'SSM - Knäck Koden'
})

useServerSeoMeta({
    title: 'SSM - Knäck Koden'
})

const challs = useChallengeStore()
const router = useRouter()


await useAsyncData('challenges', challs.getChallenges)
await useAsyncData('events', challs.getEvents)

onMounted(() => {
    challs.getChallenges()
})

const categories = computed(() => challs.challenges.map(c => c.category).filter((v, i, a) => a.indexOf(v) == i))

const challenges = computed(() => {

    var res = challs.challenges.filter(c => c.chall_namespace == "knackkoden")

    return res
})

function nav(slug: string) {
    router.push(`/challenges/${slug}`)
}
</script>

<style scoped>
.pointer {
    cursor: pointer;
}

.ssm-grid {
    display: grid;
    /* Magic auto-sizing grid: https://css-tricks.com/auto-sizing-columns-css-grid-auto-fill-vs-auto-fit/ */
    grid-template-columns: repeat(auto-fill, minmax(15rem, 1fr));
    gap: 1rem;
}

/* Bootstrap large */
@media (min-width: 992px) {
    .ssm-filter {
        width: 21rem;
    }
}
</style>
