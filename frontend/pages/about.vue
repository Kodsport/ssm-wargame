<template>
    <div class="container">
        <div>
            <h2 class="text-primary d-flex justify-content-center">SSM Projektledning</h2>

            <div class="row justify-content-center">
                <div class="col-3 bg-dark rounded p-3 m-2" v-for="x in [0, 0, 0, 0, 0, 0]">
                    <div class="d-flex justify-content-center flex-column">

                        <img class="rounded"
                            src="https://upload.wikimedia.org/wikipedia/commons/3/34/Terry_A._Davis_2017.jpg" alt="">
                        <h5 class="pt-3">Terry Davis</h5>
                        <i>Eventkoordinator</i>
                        <nuxt-link to="/challenges/terry">Se Terry's utmaningar</nuxt-link>


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
                <img src="~/assets/kodsport.png" alt="Kodsport Logga">
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

<style scoped>
img {
    height: 10vw;
}
</style>

<script setup lang="ts">
useHead({
    title: 'SSM - Om oss'
})
</script>