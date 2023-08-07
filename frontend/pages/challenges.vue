<template>
    <div class="container-fluid">
        <h1 class="text-primary">Utmaningar</h1>
        <p>
            Här kan du öva på utmaningar från tidigare års upplagor av Säkerhets-SM.
        </p>

        <div class="row pt-4">

            <div class="col-3">
                <h3>Filtrera</h3>
                <div>
                    <div class="form-group">
                        <label>Kategori</label>
                        <select class="form-control" v-model="challFilter.categoryFilter">
                            <option value="">alla</option>
                            <option class="text-lowercase" v-for="cat in categories" :value="cat">{{ cat.toLowerCase() }}
                            </option>
                        </select>
                    </div>
                    <div v-if="challs.events.length" class="form-group pt-3">
                        <label>Tävling</label>

                        <div class="form-check" v-for="event in challs.events">
                            <input type="checkbox" class="form-check-input" v-model="challFilter.eventFilter[event.id]">
                            <label class="form-check-label">{{ event.name }}</label>
                        </div>
                    </div>
                </div>
            </div>

            <div class="col">
                <div class="row" v-for="category in categories">
                    <div
                        v-if="(challFilter.categoryFilter == '' || category == challFilter.categoryFilter) && challenges.filter(c => c.category == category).length">

                        <div class="text-primary border-primary border-bottom border-3 pb-2 me-3">
                            <h4 class="text-lowercase">{{ category }}</h4>
                        </div>

                        <div class="row row-cols-3 row-cols-xl-5 g-4 mt-0 mb-2">
                            <div class="col" v-for="chall in challenges.filter(c => c.category == category)">
                                <ChallengePreview @click="nav(chall.slug)" :chall="chall" />
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <NuxtPage />
    </div>
</template>



<script setup lang="ts">
import { useChallengeStore } from '../store/challenges'
import { storeToRefs } from "pinia";

useHead({
    title: 'SSM - Utmaningar'
})


const challs = useChallengeStore()
const route = useRoute()
const router = useRouter()

const { challFilter } = storeToRefs(useChallengeStore())

onMounted(() => {
    challs.getChallenges()
    challs.getEvents()
})

const categories = computed(() => challs.challenges.map(c => c.category).filter((v, i, a) => a.indexOf(v) == i))

const challenges = computed(() => {
    const allowedEvents: string[] = [];
    for (const [key, value] of Object.entries(challFilter.value.eventFilter)) {
        if (!value) continue
        allowedEvents.push(key)
    }
    var res = challs.challenges
    if (allowedEvents.length != 0) {
        res = res.filter(v => allowedEvents.includes(v.ctf_event_id))
    }
    if (challFilter.value.categoryFilter != "") {
        res = res.filter(v => v.category == challFilter.value.categoryFilter)
    }

    return res
})

function nav(slug: string) {
    router.push(`/challenges/${slug}`)
}

</script>


<style scoped>
.challenge-container {
    display: flex;
    flex-wrap: wrap;
}

h5 {
    margin-bottom: 0;
}
</style>
