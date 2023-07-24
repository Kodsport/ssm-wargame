<template>
  <div>
    <header class="masthead">
      <div class="jumbotron container-fluid mb-0 title-img top-colors">
        <h1 class="text-color text-center">
          <span class="custom-title pr-1 pl-1">Säkerhets-SM - Månadens problem</span>
        </h1>
      </div>
    </header>


    <nav class="navbar navbar-expand-sm sticky-top top-colors">
      <div class="container">
        <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarcontent" aria-controls="navbarcontent" aria-expanded="false">
          <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="navbarcontent">
          <ul class="navbar-nav mx-auto">
            <li class="nav-item active">
              <a class="nav-link" href="/">Problem</a>
            </li>
            <li class="nav-item">
              <a class="nav-link" href="/results">Resultat</a>
            </li>
          </ul>
        </div>
      </div>
    </nav>

    <div class="container-fluid custom-body pt-2 pb-2">
      
    <div class="row">
      <div class="mx-auto col-lg-7">
        <h2 class="text-accent-color text-center">Månadens CTF-problem</h2>
        <p class="text-color">
          Säkerhets-SM är en årlig gymnasietävling i IT-säkerhet på det så kallade <em><a href="https://ctftime.org/ctf-wtf/">jeopordy-style Capture the Flag</a></em>-formatet.
          Tävlingen går ut på att lösa någon slags uppgift med ursprung inom IT-säkerhet, så som kryptografi, websäkerhet, forensik, programmering, o.s.v.
          Dold i uppgiften finns en "flagga", en textsträng på formen "SSM{...}". Denna sträng är uppgiftens lösning, och ska skickas in som svar.
        </p>
        <p class="text-color">
          Mellan tävlingarna publicerar vi ett CTF-problem i månaden, där vi lottar ut priser bland lösarna.
          Du hittar uppgifterna och kan skicka in en lösning här.
        </p>
        <p class="text-color">
          Tävlingen arrangeras av föreningen <a href="https://kodsport.se">Kodsport Sverige</a>.
          <strong>OBS:</strong> för att vara del av prisutlottningen måste du vara medlem i föreningen.
          Det blir du gratis <a href="https://ebas.ungaforskare.se/signups/index/63">på följande länk</a>.
        </p>
        <p class="text-color">
          Du kan hitta lösningar på tidigare månadens problem <a href="https://github.com/Kodsport/sakerhetssm-monthly-solutions">här</a>.
        </p>

        <div v-for="chall in challenges" :key="chall.id">
          <h2 class="text-accent-color text-center">{{chall.display_month}}</h2>
          <p class="text-color">Du hittar uppgiften och eventuella tillhörande filer här: <a href="https://sakerhetssm.se/monthly/aprmaj-2022/">se uppgiften här.</a></p>
          <p class="text-color">Du kan skicka in svar på uppgiften här: <a href="/svar/aprmaj-2022">skicka in svar.</a></p>
        </div>

      </div>
    </div>

  </div>
    <div class="container-fluid top-colors pt-2 pb-2"> 
      <div class="text-center">
        <a href="https://www.kodsport.se/"><img src="~assets/img/kodsport.png" class="img-fluid footer-img"></a>
      </div>
    </div>

  </div>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  name: 'IndexPage',

  data() {
    return {
      challenges: []
    };
  },
  async mounted() {
    
    const challs = await this.$axios.get('http://localhost:8080/monthly_challenges');
    console.log(challs)
    this.challenges = challs.data;
  }

})
</script>

<style>
.nav-item.active .nav-link,
.nav-item:hover .nav-link,
.text-color {
    color: #fff;
}

.jumbotron {
	border-radius: 0;
}

.navbar-toggler-icon {
  background-image: url("~assets/img/hamburger.svg");
}

.top-colors {
	background-color: #002c36;
}

.custom-body {
	min-height: 65vh;
	background-color: #004454;
}

.navbar-nav .nav-link,
.text-accent-color {
	color: #ffab1c;
}

.custom-title {
	background-color: rgba(0,44,54, 0.4);
}

* {
	font-family: "Consolas", Monospace; 
}

.logo-heights {
	max-height: 150px;
}

.footer-img {
	max-height: 100px;
}

.title-img {
  background-image: url("~assets/img/logo.svg");
	background-repeat: no-repeat;
	background-position: center;
	background-size: cover;
}

.almost-text-color {
	color: rgba(255,255,255, 0.6);
}

</style>