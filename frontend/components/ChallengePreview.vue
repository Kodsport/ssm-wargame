<template>
    <div class="card h-100 challenge-card" :class="{ 'bg-info': !props.chall.solved, 'bg-success': props.chall.solved }">
        <div class="card-body pb-2">
            <div class="d-flex justify-content-between align-items-start mb-2">
                <h4 class="card-title flex-grow-1 me-2">{{ props.chall.title }}</h4>
<!--                 <span class="difficulty-badge" :class="getDifficultyClass()">
                    {{ getDifficulty() }}
                </span> -->
            </div>
            <div class="mb-2">
                <span v-if="event" class="badge bg-primary me-1">{{ (event as any).name }}</span>
                <span class="badge me-1" :class="getCategoryClass()">{{ props.chall.category }}</span>
            </div>
        </div>
        <div class="card-body pt-0 d-flex justify-content-between align-items-end">
            <div v-if="!props.hideSolves" class="text small">
                <i class="bi bi-people me-1"></i>
                <span v-if="props.chall.solves != 0">{{ props.chall.solves }} lösare</span>
                <span v-else>Olöst</span>
            </div>
            <div class="text-end">
                <h3 class="mb-0 text-white">{{ props.chall.score || props.chall.static_score }}</h3>
                <small class="text">poäng</small>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">

import { useChallengeStore } from '../store/challenges';

const store = useChallengeStore()
const props = defineProps(['chall', 'hideSolves'])

const event = computed(() => store.events.find((e: any) => e.id == props.chall.ctf_event_id))

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

</script>

<style scoped>
.challenge-card {
    transition: transform 0.2s ease-in-out, box-shadow 0.2s ease-in-out;
    min-height: 200px;
}

.challenge-card:hover {
    transform: translateY(-2px);
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.card-title {
    font-size: 1.1rem;
    font-weight: 600;
    margin-bottom: 0.5rem;
}

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
</style>
