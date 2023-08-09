<template>
    <div class="container">
        <div class="row">
            <div class="col-12 col-md-6">
                <h1 class="text-primary">Vad är Säkerhets-SM?</h1>
                <p>
                    Har du ett intresse för datorrelaterad problemlösning, eller vill du lära dig vad det egentligen innebär
                    att hacka? Säkerhets-SM är en tävling för gymnasie- och högstadieelever med fokus på att lära ut koncept
                    inom cybersäkerhet och hackning. Tävlingen är skapad för att passa alla kunskapsnivåer,
                    så det spelar ingen roll om du är total nybörjare eller erfaren. Går det riktigt bra får du chansen att
                    bli uttagen till landslaget och åka på EM!
                </p>
                <div class="">
                    <ul style="list-style-type: none;">
                        <li>
                            <span class="material-symbols-outlined">groups_2</span>
                            1-3 personer per lag
                        </li>
                        <li>
                            <span class="material-symbols-outlined">swords</span>
                            Kval online (datum saknas)
                        </li>
                        <li>
                            <span class="material-symbols-outlined">emoji_events</span>
                            Final i Stockholm (datum saknas)
                        </li>
                        <li>
                            <span class="material-symbols-outlined">emoji_events</span>
                            Bästa lagen går vidare till <a href="https://snht.se/">landslagsuttagningen</a>
                        </li>
                    </ul>
                </div>
                <p>
                    Joina gärna
                    <a :href="discordUrl">Kodsports Discordserver</a>
                    där du kan nå oss organisatörer, prata med andra deltagare, visa dina egna projekt eller ställa allmänna
                    datorrelaterade frågor. Vi har också en <a
                        href="https://mail.sakerhetssm.se/subscription/form">mailinglista</a> där vi skickar ut
                    uppdateringar om tävlingen.
                </p>

                <form method="post" action="https://mail.sakerhetssm.se/subscription/form" class="listmonk-form">
                    <input type="hidden" name="nonce" class="form-control" />
                    <input id="f6b65" type="hidden" name="l" checked :value="mailFormId" />

                    <div class="input-group">
                        <input class="form-control" type="email" name="email" required placeholder="E-mail" />
                        <div class="input-group-append">
                            <button class="btn btn-outline-primary" type="submit">Prenumerera</button>
                        </div>
                    </div>
                </form>

                <a class="btn me-md-2 mt-2" style="background-color: #5865f2;" :href="discordUrl">
                    Gå med i Discord!
                </a>

            </div>

            <div class="col pt-4 pt-md-0" v-if="monthly.status.value == 'success'">
                <h1 class="text-primary">Månadens utmaning - {{ monthly.data.value.display_month }}</h1>
                <MonthlyChallenge :chall="monthly.data.value.challenge" />
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { useChallengeStore } from '../store/challenges'

const http = useHttp()
const discordUrl = 'https://discord.gg/edKFKKU'
const mailFormId = 'f6b65d3d-5ba1-4da9-8c42-4e3be8b6277f';

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