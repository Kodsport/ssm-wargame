<template>
    <h1 class="text-primary">Knäck Koden</h1>
    <p>
        Knäck koden är en tävling för mellanstadieelever som vill lära sig mer om hacking och cybersäkerhet.
        Vill du ha Knäck Koden-lektioner i din skola? Spana in <a
            href="https://www.cyberarena.se/knackkoden">informationen hos Kodcentrum!</a>
    </p>
    <h1 class="text-primary">Föreläsningar för gymnasiet!</h1>
    <p>
        Vi kommer gärna och föreläser för era studenter! Vi förklarar hur man hackar och hur man kommer in i svenska hackinglandslaget! Kontakta sakerhetssm@kodsport.se för att boka! 
    </p>

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
