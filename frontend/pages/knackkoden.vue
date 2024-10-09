<template>
    <h1 class="text-primary">Knäck Koden</h1>
    <p>
        Knäck Koden är en tävling inom hacking för mellanstadieelever. Fråga din lärare om lösenordet till klassen!
    </p>

    <div v-if="!auth.knackKodenPassword" class="form-group col-3">
        <label for="password">Klasslösenord</label>
        <input class="form-control" type="text" v-model="password">
        <button class="btn btn-primary mt-2" @click="login(password)">Logga in</button>
    </div>
    <div v-else>
        Du representerar {{ auth.knackKodenData.class_name }}

        <button class="btn btn-primary mt-2" @click="logout()">Byt klass</button>

    </div>

    <div class="d-flex flex-column flex-lg-row pt-4">

        <div class="flex-grow-1 d-flex flex-column col-6">
            <template v-for="category in categories">
                <div v-if="challenges.filter(c => c.category == category).length">
                    <div class="text-primary border-primary border-bottom border-3 pt-2 pb-2">
                        <h4 class="text-lowercase">{{ category }}</h4>
                    </div>

                    <div class="ssm-grid pt-3 pb-2">
                        <div v-for="chall in challenges.filter(c => c.category == category)">
                            <ChallengePreview :hideSolves="true" class="pointer" @click="nav(chall.slug)"
                                :chall="chall" />
                            <!-- for seo -->
                            <a class="d-none" :href="`/knackkoden/${chall.slug}`">{{ chall.title }}</a>
                        </div>
                    </div>
                </div>
            </template>
        </div>
        <div class="col-6 p-2">
            <table class="table">
                <thead>
                    <tr>
                        <th>Klass</th>
                        <th>Poäng</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="row in scoreboard.scores">
                        <td>{{ row.school_name }}</td>
                        <td>{{ row.score }}</td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
    <NuxtPage />
</template>

<script setup lang="ts">
import { useChallengeStore } from '../store/challenges'
import { useAuthStore } from '../store/auth'

const password = ref("")

useHead({
    title: 'SSM - Knäck Koden'
})

useServerSeoMeta({
    title: 'SSM - Knäck Koden'
})

const challs = useChallengeStore()
const router = useRouter()
const http = useHttp()
const auth = useAuthStore()

await useAsyncData('challenges', challs.getChallenges)
await useAsyncData('events', challs.getEvents)

onMounted(() => {
    challs.getChallenges()

    const pw = localStorage.getItem('ssm-knackkoden-password')
    if (pw) {

        login(pw)
    }
})

const catOrder = ["kryptografi", "generellt", "web"]

const categories = computed(() => challs.challenges
    .map(c => c.category)
    .filter((v, i, a) => a.indexOf(v) == i)
    .sort((a, b) => catOrder.indexOf(a) - catOrder.indexOf(b))
)

const challenges = computed(() => {

    var res = challs.challenges
        .filter(c => c.chall_namespace == "knackkoden")
        .map(c => ({ ...c, solved: auth.knackKodenData?.solves?.includes(c.id) }))
        .sort((a, b) => a.score - b.score)

    return res
})

function nav(slug: string) {
    router.push(`/knackkoden/${slug}`)
}

const scoreboard = await http("/knack_koden_scoreboard")

function login(pw: string) {
    try {
        auth.getKnackKodenData(pw)
        localStorage.setItem("ssm-knackkoden-password", pw)

    } catch (error) {
        localStorage.removeItem("ssm-knackkoden-password")
        alert("Fel klasslösenord")

    }
}

function logout() {
    localStorage.removeItem("ssm-knackkoden-password")
    auth.knackKodenData = null;
    auth.knackKodenPassword = null;

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
