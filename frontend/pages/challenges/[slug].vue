<template>
    <div class="modal-backdrop fade show"></div>
    <div class="modal modal-xl d-block" tabindex="-1" @click="goBack" ref="modal" v-if="chall">
        <div class="modal-dialog">
            <div class="modal-content">


                <div class="pt-3 d-flex justify-content-center">
                    <h2 class="modal-title text-primary fw-bold">{{ chall.title }}</h2>
                </div>
                <div class="d-flex justify-content-center">
                    <div class="badge bg-primary me-1">{{ chall.category }}</div>
                    <span v-if="event" class="badge bg-primary">{{ event.name }}</span>
                </div>

                <h5>
                    <div class="d-flex justify-content-center pt-3">
                        <span class="mx-3 d-flex">
                            <span class="me-1">
                                {{ chall.score }}
                            </span>
                            <span class="material-icons text-primary">control_point_duplicate</span>
                        </span>
                        <span class="mx-3 d-flex">
                            <span class="me-1">
                                {{ chall.solves }}
                            </span>
                            <span class="material-icons text-primary">flag</span>
                        </span>
                    </div>
                </h5>

                <div class="modal-body">
                    <div class="row">
                        <div class="col-12 col-lg-6 order-2 order-lg-1" v-html="renderMarkdown(chall.description)"></div>
                        <div class="col-12 col-lg-6 order-1 order-lg-2 mb-3 mb-lg-0">
                            <div class="d-flex" v-if="chall?.authors?.length || chall.other_authors?.length">
                                <span class="material-icons text-primary pe-1">group</span>
                                <span class="material-icons text-primary pe-1">edit</span>
                                <span>
                                    <nuxt-link class="author" v-for="a in chall.authors" :to="`/author/${a.id}`">
                                        {{ a.full_name }}
                                    </nuxt-link>
                                    <span class="author" v-for="a in chall.other_authors">
                                        {{ a }}
                                    </span>
                                </span>
                            </div>
                            <div class="d-flex" v-for="service in chall.services">
                                <span class="material-icons text-primary">router</span>
                                <a v-if="service.hyperlink" :href="service.user_display">{{ service.user_display }}</a>
                                <span v-else>{{ service.user_display }}</span>
                            </div>
                            <div class="d-flex" v-for="file in chall.files">
                                <FileDownload :file="file" />
                            </div>
                            <div class="pt-3" v-if="chall.solvers">
                                <b>Första lösarna</b>
                                <ol>
                                    <li v-for="solver in chall.solvers">
                                        {{ solver.full_name }} <span class="badge bg-info">{{
                                            timeAgo(solver.solved_at)
                                        }}</span>
                                    </li>
                                </ol>
                            </div>
                        </div>
                    </div>
                    <form @submit.prevent>
                        <div class="row mt-3">
                            <div v-if="!chall.solved" class="form-group alert" v-bind:class="{ 'alert-danger': warn }">
                                <input v-if="!!auth.user.id" type="text" class="form-control" placeholder="SSM{..."
                                    @keypress.enter="submitFlag" v-model="flag">
                                <InputReplacer v-else text="Logga in för att skicka in flaggor!" />
                            </div>
                            <div v-else>
                                <InputReplacer text="Löst!" />
                            </div>
                        </div>
                    </form>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import renderMarkdown from '../../utils/markdown';
import { useAuthStore } from '../../store/auth';
import { useChallengeStore } from '../../store/challenges';
import * as moment from 'moment'


const router = useRouter();
const route = useRoute()
const auth = useAuthStore()
const http = useHttp()
const store = useChallengeStore()


const modal = ref(null)
const event = computed(() => store.events.find(e => e.id == chall.ctf_event_id))

const chall = computed(() => store.getBySlug(route.params.slug))

const flag = ref("")
const warn = ref(false)

async function submitFlag() {
    try {
        await http(`/challenges/${chall.id}/attempt`, {
            method: 'POST',
            body: {
                flag: flag.value
            }
        })

        store.getChallenges()
        store.getMonthlies()

    } catch (error) {
        console.log(error)
        warn.value = true
        setTimeout(() => {
            warn.value = false
        }, 1000)
    }

}

function goBack(event) {
    if (event.target == modal.value) {
        router.replace('/challenges')
    }
}

function timeAgo(unixTime) {
    return moment.default(new Date(unixTime * 1000)).fromNow()
}
</script>

<style scoped>
.author:after {
    content: ",";
    padding-right: 0.5em;
    text-decoration: none;
    display: inline-block;
}

.author:last-child:after {
    content: none;
}
</style>