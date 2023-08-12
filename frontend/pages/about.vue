<template>
    <div class="container">
        <div>
            <h2 class="text-primary d-flex justify-content-center">SSM Projektledning</h2>

            <div class="row">
                <div class="col-3" v-for="person in people">
                    <div class="h-100 bg-dark d-flex justify-content-center flex-column p-3 rounded">

                        <img class="rounded img-fluidx" :src="person.url" alt="">
                        <h5 class="pt-3">{{ person.name }}</h5>
                        <i>{{ person.role }}</i>
                        <nuxt-link v-if="false" to="/challenges/terry">Se Terry's utmaningar</nuxt-link>


                        <span class="text-primary pt-2">
                            <span class="material-symbols-outlined text-decoration-none pe-2">email</span>
                            <a href="mailto:terry@kodsport.se">terry@kodsport.se</a>
                        </span>
                    </div>
                </div>
            </div>
        </div>

        <div class="row pt-3">
            <div class="col d-flex flex-row justify-content-center">
                <img class="img-fluid" src="~/assets/kodsport.png" alt="Kodsport Logga">
            </div>
            <div class="col">
                <h1 class="text-primary">Föreningen</h1>
                <p>
                    <a href="https://www.kodsport.se/">Kodsport</a> är en ideell förening som arbetar för att lyfta Sverige
                    som
                    IT-nation. Föreningen arrangerar tävlingar och träningsläger i tävlingsprogrammering och Capture The
                    Flag (datorsäkerhet).
                </p>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">

const people = [
    {
        name: "Aron Bergman",
        role: "Projektledning",
        url: "/aron.png"
    },
    {
        name: "Xenia Lööv",
        role: "Projektledning",
        url: "/xenia.png"
    },
    {
        name: "Oskar Tornevall",
        role: "Projektledning",
        url: "/oskar.png"
    },
    {
        name: "Mateusz Drwal",
        role: "Infra",
        url: "/mateusz.png"
    }
];

useHead({
    title: 'SSM - Om oss'
})
</script>