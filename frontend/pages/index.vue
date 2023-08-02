<template>
    <div class="container">
        <div class="row">
            <div class="col">
                <h1 class="text-primary">Vad är Säkerhets-SM?</h1>
                <p>
                    Har du ett intresse för datorrelaterad problemlösning, eller vill du lära dig vad det egentligen innebär
                    att hacka? Säkerhets-SM är en tävling med fokus på att lära ut koncept inom bland annat cybersäkerhet,
                    men även att hitta de bästa eleverna i Sverige! Tävlingen är skapad för att passa alla kunskapsnivåer,
                    så det spelar ingen roll om du är total nybörjare eller erfaren. Går det riktigt bra får du chansen att
                    bli uttagen till landslaget och åka på EM!
                </p>
                <p>
                    Tävlingen är riktad främst mot gymnasie- och högstadieelever, och består av ett kval och en final. Man
                    tävlar i lag med upp till 3 personer, så du kan tävla själv eller samla två kompisar för en ännu
                    roligare upplevelse! Kvalet hålls online den 10-12 mars, med en träningstävling veckan innan för att
                    hjälpa till att värma upp. Finalen hålls den 26-28 maj på Chalmers tekniska högskola i Göteborg för de
                    ~30 bästa eleverna, med betald resa, mat och boende.
                </p>
                <p>
                    Joina gärna
                    <a href="https://discord.gg/edKFKKU">Kodsports Discordserver</a>
                    där du kan nå oss organisatörer, prata med andra deltagare, visa dina egna projekt eller ställa allmänna
                    datorrelaterade frågor.
                    Vi har också en
                    <a href="https://mail.sakerhetssm.se/subscription/form">mailinglista</a>

                    där vi skickar ut uppdateringar om tävlingen.
                </p>
            </div>

            <div class="col">
                <div v-if="monthly">
                    <h1 class="text-primary">Månadens utmaning - {{ monthly.display_month }}</h1>
                    <challenge :chall="monthly.challenge"></challenge>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { useChallengeStore } from '../store/challenges'
import { useAuthStore } from '../store/auth'

const http = useHttp()
const challs = useChallengeStore()
const auth = useAuthStore()

const monthly = computed(() => challs.getCurrentMonthly)

onMounted(() => {
    challs.getChallenges()
    challs.getMonthlies()
})


</script>