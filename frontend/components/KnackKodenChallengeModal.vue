<template>
    <div class="modal-backdrop fade show"></div>
    <div class="modal modal-xl d-block" tabindex="-1" @click="goBack" ref="modal">
        <div class="modal-dialog modal-dialog-centered">
            <div class="modal-content">
                <div class="pt-3 d-flex justify-content-center">
                    <h2 class="modal-title text-primary fw-bold">{{ props.chall.title }}</h2>
                </div>
                <div class="d-flex justify-content-center">
                    <div class="badge bg-primary me-1">{{ props.chall.category }}</div>
                    <span v-if="event" class="badge bg-primary">{{ event.name }}</span>
                </div>

                <h5>
                    <div class="d-flex justify-content-center pt-3">
                        <span class="mx-3 d-flex">
                            <span class="me-1">
                                {{ props.chall.score }}
                            </span>
                            <span class="material-symbols-outlined">control_point_duplicate</span>
                        </span>
                    </div>
                </h5>

                <div class="modal-body">
                    <div class="row">
                        <div class="col-12 col-lg-6 order-2 order-lg-1"
                            v-html="renderMarkdown(props.chall.description)">
                        </div>
                        <div class="col-12 col-lg-6 order-1 order-lg-2 mb-3 mb-lg-0">
                            <div v-if="props.chall?.authors?.length">
                                <span class="material-symbols-outlined pe-1">group</span>
                                <span class="material-symbols-outlined pe-1">edit</span>
                                <span>
                                    <template v-for="a in props.chall.authors">
                                        <nuxt-link v-if="a.publish" class="author" :to="`/author/${a.slug}`">
                                            {{ a.full_name }}
                                        </nuxt-link>
                                        <span v-else class="author">
                                            {{ a.full_name }}
                                        </span>
                                    </template>
                                </span>
                            </div>
                            <div class="d-flex" v-for="service in props.chall.services">
                                <span class="material-symbols-outlined pe-2">router</span>
                                <a v-if="service.hyperlink" :href="service.user_display" target="_blank"
                                    rel="nofollow">{{ service.user_display
                                    }}</a>
                                <span v-else>{{ service.user_display }}</span>
                            </div>
                            <div v-for="file in props.chall.files">
                                <FileDownload :file="file" />
                            </div>

                        </div>
                    </div>
                    <div class="row mt-3" :class="{ 'alert alert-danger': warn }">
                        <template v-if="!props.chall?.solved">

                            <input v-if="auth.knackKodenPassword" type="text" class="form-control" placeholder="SSM{..."
                                v-model="flagInput" @keypress.enter="submitFlag">
                            <InputReplacer v-else text="Logga in för att skicka in flaggor!" />
                        </template>
                        <InputReplacer v-else text="Löst!" />
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import renderMarkdown from '../utils/markdown';
import { useAuthStore } from '../store/auth';
import { useChallengeStore } from '../store/challenges';


const route = useRoute()
const auth = useAuthStore()
const http = useHttp()
const store = useChallengeStore()

const props = defineProps(['chall'])
const emit = defineEmits(['back'])

const modal = ref(null)
const event = computed(() => store.events.find(e => e.id == chall.ctf_event_id))

const chall = computed(() => store.getBySlug(route.params.slug))

const warn = ref(false)

const flagInput = ref("")

async function submitFlag() {
    try {
        await http(`/challenges/${chall.value.id}/knack_koden_attempt`, {
            method: 'POST',
            body: {
                flag: flagInput.value,
                password: auth.knackKodenPassword
            }
        })

        auth.getKnackKodenData(null)

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
        emit('back')
    }
}

</script>

<style scoped>
.author:after {
    content: ",";
    padding-right: 0.5em;
    text-decoration: none;
    display: inline-block;
    color: #ffffff;
    /* todo, use variable */
}

.author:last-child:after {
    content: none;
}
</style>
