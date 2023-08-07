<template>
    <div class="card h-100" :class="{ 'bg-info': !props.chall.solved, 'bg-success': props.chall.solved }">
        <div class="card-body">
            <h4>{{ props.chall.title }}</h4>
            <h5>
                <span v-if="event" class="badge bg-primary">{{ event.name }}</span>
            </h5>
        </div>
        <div class="card-body d-flex justify-content-between">
            <div class="align-items-end d-flex">
                <span v-if="props.chall.solves != 0">{{ props.chall.solves }} lösare</span>
                <span v-else>Olöst</span>
            </div>
            <h3 class="align-items-end d-flex mb-0">{{ props.chall.score }}</h3>
        </div>
    </div>
</template>
  
<script setup lang="ts">

import { useChallengeStore } from '../store/challenges';

const store = useChallengeStore()
const props = defineProps(['chall'])

const event = computed(() => store.events.find(e => e.id == props.chall.ctf_event_id))


</script>
