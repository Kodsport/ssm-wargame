<template>
    <div class="container">
        <div class="row">
            <div class="col-12 col-md-6">
                <h1 class="text-primary">Vad är Säkerhets-SM?</h1>
                <p>
                    Har du ett intresse för datorrelaterad problemlösning, eller vill du lära dig vad det egentligen
                    innebär
                    att hacka? Säkerhets-SM är en tävling för gymnasie- och högstadieelever med fokus på att lära ut
                    koncept
                    inom cybersäkerhet och hackning. Tävlingen är skapad för att passa alla kunskapsnivåer,
                    så det spelar ingen roll om du är total nybörjare eller erfaren. Går det riktigt bra får du chansen
                    att
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
                            Kval online <b>17-19 januari 2025</b>
                        </li>
                        <li>
                            <span class="material-symbols-outlined">emoji_events</span>
                            Final i Stockholm <b>7-9 mars 2025</b>
                        </li>
                        <li>
                            <span class="material-symbols-outlined">azm</span>
                            Bästa lagen går vidare till <a href="https://snht.se/">landslagsuttagningen</a>
                        </li>
                    </ul>
                </div>
                <p>
                    Registrera dig på tävlingsplatformen för att delta!
                    <a class="btn btn-primary" href="https://ctf.sakerhetssm.se/signin">Till tävlingsplatformen</a>
                </p>
                <p>
                    Joina gärna
                    <a :href="discordUrl">Kodsports Discordserver</a>
                    där du kan nå oss organisatörer, prata med andra deltagare, visa dina egna projekt eller ställa
                    allmänna
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


                <div class="mt-5 text-center">
                    <h3 class="">I samarbete med</h3>
                    <div class="m-5">
                        <a href="https://www.fra.se/">
                            <img
                            class="img-fluid "
                            src="~/assets/fra.png"
                            alt="FRA - Försvarets Radioanstalt"
                            >
                        </a>
                    </div>
                    <div class="m-5">
                        <a href="https://www.msab.com" title="MSAB">
                            <svg width="170" height="50" viewBox="0 0 170 50" fill="none" xmlns="http://www.w3.org/2000/svg">
                            <path d="M87.4823 35.0472C87.2459 32.1513 85.709 29.2553 83.3446 27.364C80.9802 25.4137 78.3203 24.468 76.3696 23.8179L75.6603 23.5815C73.5915 22.9314 71.5226 22.1631 69.9858 20.9219C69.1582 20.2718 68.5671 19.5626 68.1534 18.7352C67.6805 17.7305 67.5032 16.6666 67.6214 15.3073C68.0943 11.4657 71.5226 8.5697 75.483 8.8652C76.9607 8.98341 78.3794 9.51532 79.5025 10.4018C80.6847 11.1702 81.6304 12.1158 82.8126 13.5342L82.9308 13.7115L85.4726 8.39239L85.4135 8.33329C83.1082 5.31911 79.6798 3.42785 75.8968 3.13234C75.6012 3.13234 75.2466 3.07324 74.951 3.07324C69.2173 3.07324 64.0748 6.85575 62.4197 12.3522V12.4113C62.2424 13.0614 62.065 13.8297 61.9468 14.6572C61.8286 15.4255 61.9468 17.0212 61.9468 17.0803C62.065 18.5579 62.4197 19.8581 63.0108 21.0992C64.0748 23.3451 65.9663 25.2364 68.508 26.6548C70.4586 27.7777 72.4684 28.4279 73.7688 28.8416L74.2417 29.0189C76.3105 29.7281 78.3794 30.4964 79.7389 31.6784C81.1575 32.9196 81.926 34.2198 81.8077 36.9385C81.2758 41.312 77.3745 44.5035 72.9413 44.208C70.6951 44.0307 68.6853 43.026 67.2076 41.312C66.1436 40.3664 64.7841 38.6524 64.2521 38.0023L64.1339 37.825L61.5922 43.2624L61.6513 43.3215C64.0157 47.104 68.0943 49.5862 72.5866 49.8818C72.9413 49.8818 73.2959 49.9409 73.5915 49.9409C80.6255 49.9409 86.5957 44.6808 87.5414 37.7068V37.4704C87.6005 36.9385 87.4823 35.1063 87.4823 35.0472Z" fill="white"></path>
                            <path d="M169.705 33.0969C168.936 29.5508 166.808 26.714 163.853 24.8818C165.862 23.1679 167.34 20.8038 167.813 18.0851C167.931 17.3759 167.99 16.6076 167.99 15.8392C167.99 14.9527 167.872 14.1253 167.695 13.2388C167.045 10.3428 165.39 7.86053 163.025 6.20568C160.72 4.55083 157.764 3.66431 154.572 3.66431H142.041V49.1135H155.459C158.947 49.0544 162.198 48.0497 164.68 46.2175C167.281 44.3263 169.113 41.5485 169.705 38.2979C169.823 37.4705 169.941 36.7021 169.941 35.8747C170 34.9291 169.882 33.9835 169.705 33.0969ZM164.503 37.3523C164.148 39.2435 163.143 40.7802 161.547 41.9622C160.011 43.0851 157.823 43.7943 155.341 43.7943H147.479V29.7872V27.896V22.5177V18.9125V9.04256H154.691C156.878 9.04256 158.71 9.63357 160.011 10.5792C161.311 11.5248 162.198 12.766 162.552 14.4208C162.67 14.8936 162.73 15.3664 162.73 15.8392C162.73 16.253 162.67 16.6667 162.611 17.0804C162.02 20.0355 159.597 22.2222 155.696 22.5768H152.208V27.9551H155.755C158.119 28.0142 160.129 28.7234 161.607 29.8463C163.084 30.9693 164.089 32.4468 164.503 34.279C164.621 34.8109 164.68 35.4019 164.68 35.9338C164.621 36.4066 164.562 36.8794 164.503 37.3523Z" fill="white"></path>
                            <path d="M101.492 49.1726L114.082 16.0757L126.673 49.1726H132.643L114.082 0L95.5216 49.1726H101.492Z" fill="white"></path>
                            <path d="M5.61544 49.1726L16.0188 16.9031L26.7768 47.8132L37.4757 16.9031L47.879 49.1726H53.5536L37.653 0L26.7768 31.4421L15.8414 0L0 49.1726H5.61544Z" fill="white"></path>
                            </svg>
                        </a>
                        <p class="mt-5">
                            Kodsport's infrastructure is sponsored by <a href="https://glesys.com/">GleSYS</a>
                        </p>
                    </div>
            </div>
            </div>

            <div class="col pt-4 pt-md-0" v-if="monthly.status.value == 'success'">
                <h1 class="text-primary">Månadens utmaning - {{ monthly.data.value.display_month }}</h1>
                <MonthlyChallenge :chall="monthly.data.value.challenge" />
            </div>
            <div v-else-if="scoreboard.nonUniScores" class="col pt-4 pt-md-0 text-primary">
                <h1>Poängtavla</h1>
                <table class="table">
                    <thead>
                        <tr>
                            <th>Rang</th>
                            <th>Namn</th>
                            <th>Poäng</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="score, i in scoreboard.nonUniScores.slice(0, 10)">
                            <td>{{ i + 1 }}</td>
                            <td>{{ score.school_name }}</td>
                            <td>{{ score.score }}</td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { useChallengeStore } from '../store/challenges'
import { useScoreboardStore } from '../store/scoreboard'
import { useAuthStore } from "../store/auth";

const http = useHttp()
const scoreboard = useScoreboardStore()
const challs = useChallengeStore()
const auth = useAuthStore();

const discordUrl = 'https://discord.gg/edKFKKU'
const mailFormId = 'f6b65d3d-5ba1-4da9-8c42-4e3be8b6277f';

const url = useRequestURL();
useHead({
    title: 'Säkerhets-SM',
    link: {
        rel: 'canonical',
        href: `${url.protocol}//${url.host}/`,
    },
})

const monthly = await useAsyncData('monthly', () => http('/current_monthly_challenge'))
if (monthly.error.value) {
    await useAsyncData('scoreboard', scoreboard.getSchoolScoreboards)
}

watch(
  () => auth.user,
  () => {
    if (auth.user.id) {
      monthly.refresh();
      prev_monthlies.refresh();
    }
  }
);

onMounted(() => {
    challs.getChallenges()
    monthly.refresh()
})


</script>
