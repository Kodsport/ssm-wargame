<template>
    <div class="container-fluid">
        <h1 class="text-primary">Utmaningar</h1>
        <p>
            Här kan du öva på utmaningar från tidigare års upplagor av Säkerhets-SM.
        </p>

        <ChallengeModal v-if="modalChall" :chall="modalChall" />
        <div class="row pt-4">

            <div class="col-3">
                <h3>Filtrera</h3>
                <div>
                    <div class="form-group">
                        <label>Kategori</label>
                        <select class="form-control" v-model="categoryFilter">
                            <option value="">Alla</option>
                            <option v-for="cat in categories" :value="cat">{{ cat }}</option>
                        </select>
                    </div>
                    <div v-if="challs.events.length" class="form-group pt-3">
                        <label>Tävling</label>

                        <div class="form-check" v-for="event in challs.events">
                            <input type="checkbox" class="form-check-input" v-model="eventFilter[event.id]">
                            <label class="form-check-label">{{ event.name }}</label>
                        </div>
                    </div>
                </div>
            </div>

            <div class="col pt-5">
                <div v-if="challenges.length != 0" class="row row-cols-5 g-4">
                    <div class="col" v-for="chall in challenges">
                        <Challenge3 @click="nav(chall.slug)" :chall="chall" />
                    </div>
                </div>
                <p v-else>
                    Inga utmaningar matchade filtreringen
                </p>
            </div>
        </div>
    </div>
</template>



<script setup lang="ts">
import { useChallengeStore } from '../../store/challenges'
import { useAuthStore } from '../../store/auth'

const http = useHttp()
const challs = useChallengeStore()
const auth = useAuthStore()
const route = useRoute()
const router = useRouter()

const eventFilter = ref({})
const categoryFilter = ref("")


const modalChall = computed(() => challs.getBySlug(route.params.slug[0]))

onMounted(() => {
    challs.getChallenges()
    challs.getEvents()

})

const categories = computed(() => challs.challenges.map(c => c.category).filter((v, i, a) => a.indexOf(v) == i))

const challenges = computed(() => {
    const allowedEvents: string[] = [];
    for (const [key, value] of Object.entries(eventFilter.value)) {
        if (!value) continue
        allowedEvents.push(key)
    }
    var res = challs.challenges
    if (allowedEvents.length != 0) {
        res = res.filter(v => allowedEvents.includes(v.ctf_event_id))
    }
    if (categoryFilter.value != "") {
        res = res.filter(v => v.category == categoryFilter.value)
    }

    return res
})


function showModal(name: string) {
    console.log(name)

}

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
