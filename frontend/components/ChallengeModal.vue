<template>
    <div class="modal-backdrop fade show"></div>
    <div class="modal modal-xl d-block" tabindex="-1" @click="goBack" ref="modal">
        <div class="modal-dialog modal-dialog-centered">
            <div class="modal-content">
                <div class="pt-3 d-flex justify-content-center">
                    <h2 class="modal-title text-primary fw-bold">{{ props.chall.title }}</h2>
                </div>
                <div class="d-flex justify-content-center">
                    <div class="badge me-1" :class="getCategoryClass()">{{ props.chall.category }}</div>
                    <span v-if="event" class="badge bg-primary me-1">{{ (event as any).name }}</span>
<!--                     <span class="difficulty-badge me-1" :class="getDifficultyClass()">
                        {{ getDifficulty() }}
                    </span> -->
                </div>

                <h5>
                    <div class="d-flex justify-content-center pt-3">
                        <span class="mx-3 d-flex">
                            <span class="me-1">
                                {{ props.chall.score }}
                            </span>
                            <span class="material-symbols-outlined">control_point_duplicate</span>
                        </span>
                        <span class="mx-3 d-flex">
                            <span class="me-1">
                                {{ props.chall.solves }}
                            </span>
                            <span class="material-symbols-outlined">flag</span>
                        </span>
                    </div>
                </h5>

                <div class="modal-body">
                    <div class="row">
                        <div class="col-12 col-lg-6 order-2 order-lg-1 text-break"
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
                            <div class="pt-3" v-if="props.chall.solvers">
                                <b>Första lösarna</b>
                                <ol>
                                    <li v-for="solver in props.chall.solvers">
                                        {{ solver.full_name }} <span class="badge bg-info">{{
        timeAgo(solver.solved_at)
    }}</span>
                                    </li>
                                </ol>
                            </div>
                        </div>
                    </div>
                    <div class="row mt-3">
                        <FlagInput :class="{ 'alert alert-danger': warn }" v-model="flag" @keypress.enter="submitFlag"
                            :solved="props.chall.solved" />
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
import * as moment from 'moment'


const route = useRoute()
const auth = useAuthStore()
const http = useHttp()
const store = useChallengeStore()

const props = defineProps(['chall'])
const emit = defineEmits(['back'])

const modal = ref(null)
const event = computed(() => store.events.find((e: any) => e.id == props.chall.ctf_event_id))

const chall = computed(() => store.getBySlug(route.params.slug as string))

const flag = ref("")
const warn = ref(false)

// fixar voting senare, detta kan dock användas ifall man vill.
/* const getDifficulty = () => {
    const score = props.chall.score || props.chall.static_score || 0
    if (score <= 150) return 'enkel'
    if (score <= 300) return 'medel'
    return 'svår'
}

const getDifficultyClass = () => {
    const difficulty = getDifficulty()
    return {
        'difficulty-easy': difficulty === 'enkel',
        'difficulty-medium': difficulty === 'medel',
        'difficulty-hard': difficulty === 'svår',
    }
} */

const getCategoryClass = () => {
    const category = props.chall.category?.toLowerCase()
    switch (category) {
        case 'pwn':
            return 'category-pwn'
        case 'crypto':
            return 'category-crypto'
        case 'web':
            return 'category-web'
        case 'forensics':
            return 'category-forensics'
        case 'misc':
            return 'category-misc'
        case 'osint':
            return 'category-osint'
        case 'rev':
        case 'reverse':
        case 'reversing':
            return 'category-rev'
        default:
            return 'bg-secondary'
    }
}

async function submitFlag() {
    try {
        await http(`/challenges/${chall.value.id}/attempt`, {
            method: 'POST',
            body: {
                flag: flag.value.trim()
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

function timeAgo(unixTime) {
    return moment.default(new Date(unixTime * 1000)).fromNow()
}

function goBack(event) {
    if (event.target == modal.value) {
        emit('back')
    }
}

</script>

<style scoped>
.difficulty-badge {
    padding: 0.15rem 0.35rem;
    border-radius: 0.2rem;
    font-size: 0.65rem;
    font-weight: 600;
    border: 1px solid #00000077 !important;
}

.difficulty-easy {
    background-color: #1b9a41;
    color: #ffffff;
}

.difficulty-medium {
    background-color: #ddb531;
    color: #ffffff;
}

.difficulty-hard {
    background-color: #8e3039;
    color: #ffffff;
}

.category-pwn {
    background-color: #b01d2b;
    color: white;
}

.category-crypto {
    background-color: #ffbf00;
    color: #ffffff;
}

.category-web {
    background-color: #2281a7;
    color: #ffffff;
}

.category-forensics {
    background-color: #144690;
    color: white;
}

.category-misc {
    background-color: #6f42c1;
    color: white;
}

.category-osint {
    background-color: #6c757d;
    color: white;
}

.category-rev {
    background-color: #b47523;
    color: white;
}

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
