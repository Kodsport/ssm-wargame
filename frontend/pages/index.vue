<template>
    <div class="container">
        <div class="row">
            <div class="col-12 col-md-6">
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
                    datorrelaterade frågor. Vi har också en <a
                        href="https://mail.sakerhetssm.se/subscription/form">mailinglista</a> där vi skickar ut
                    uppdateringar om tävlingen.
                </p>
                <a class="btn me-md-2 mb-1" style="background-color: #5865f2;" href="https://discord.gg/edKFKKU">Joina
                    Discord!</a>
                <a class="btn btn-secondary" href="https://mail.sakerhetssm.se/subscription/form">Gå med i
                    mailinglistan!</a>
            </div>

            <div v-if="monthly.status.value == 'success'" class="col pt-3 pt-md-0">
                <h1 class="text-primary">Månadens utmaning - {{ monthly.data.value.display_month }}</h1>
                <MonthlyChallenge :chall="monthly.data.value.challenge" />
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { useChallengeStore } from '../store/challenges'

const http = useHttp()

useHead({
    title: 'Säkerhets-SM'
})

const challs = useChallengeStore()

const monthly = await useAsyncData('monthly', () => http('/current_monthly_challenge'))

onMounted(() => {
    challs.getChallenges()
    monthly.refresh()
})


</script>