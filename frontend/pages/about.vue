<template>
    <div class="
        d-flex
        flex-column
        align-items-center
        px-2
        mw-50
    ">
        <h2 class="text-primary">SSM Projektledning</h2>
        <div class="
            d-flex
            flex-column
            flex-md-row
            ssm-row
        ">
            <div
                v-for="person in people"
                class="
                    d-flex
                    flex-column
                    bg-dark
                    m-2
                    p-3
                    rounded
                "
                style="max-width: 22.5rem"
            >
                <img class="mw-100 rounded" :src="person.url" alt="">
                <h5 class="pt-3">{{ person.name }}</h5>
                <i>{{ person.role }}</i>
            </div>
        </div>

        <div class="
            d-flex
            flex-column
            flex-md-row
            align-items-center
            ssm-row
        ">
            <img class="ks-logo" src="~/assets/kodsport.png" alt="Kodsport Logga">
            <div>
                <h1 class="text-primary">Föreningen</h1>
                <p>
                    <a href="https://www.kodsport.se/">Kodsport</a> är en
                    ideell förening som arbetar för att lyfta Sverige
                    som IT-nation. Föreningen arrangerar tävlingar och
                    träningsläger i tävlingsprogrammering och Capture
                    The Flag (datorsäkerhet).
                </p>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">

const people = [
    {
        name: "Martin Wennberg",
        role: "Projektledare",
        url: "/martin.png"
    },
    {
        name: "Antoni Kowalik",
        role: "Projektledare",
        url: "/antoni.png"
    },
    {
        name: "Mateusz Drwal",
        role: "Infra",
        url: "/mateusz.png"
    }
];

useHead({
    title: 'SSM - Om oss',
    link: [
        {
            rel: 'canonical',
            href: `${url.protocol}//${url.host}/about`,
        },
    ]
})
</script>

<style scoped>
  .ks-logo {
    max-width: 100%;
  }
  /* Bootstrap's md breakpoint */
  @media (min-width: 768px) {
    .ssm-row {
      max-width: 75rem;
    }
    .ks-logo {
      max-width: 50%;
    }
  }
</style>
