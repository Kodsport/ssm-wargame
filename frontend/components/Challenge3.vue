<template>
    <div class="card bg-info ">
        <div class="card-body border-dark">
            <h4 class="card-title">{{ props.chall.title }}</h4>
            <h5>
                <span class="badge bg-primary">{{ props.chall.category }}</span>
                <span v-if="event" class="badge bg-primary">{{ event.name }}</span>
            </h5>
            <!-- hack -->
            <br v-if="!event">
            <div class="d-flex justify-content-between">
                <div class="align-items-end d-flex">
                    <span v-if="props.chall.solves != 0">{{ props.chall.solves }} lösare</span>
                    <span v-else>Olöst</span>
                </div>
                <h3 class="align-items-end d-flex mb-0">{{ props.chall.score }}</h3>
            </div>
        </div>
    </div>
</template>
  
<script setup lang="ts">

import { useChallengeStore } from '../store/challenges';

const store = useChallengeStore()
const props = defineProps(['chall'])

const event = computed(() => store.events.find(e => e.id == props.chall.ctf_event_id))


</script>
