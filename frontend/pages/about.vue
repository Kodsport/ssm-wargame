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
            <div v-for="person in people" class="
                    d-flex
                    flex-column
                    bg-dark
                    m-2
                    p-3
                    rounded
                " style="max-width: 22.5rem">
                <img class="mw-100 rounded" style="filter: grayscale(1)" :src="person.url" alt="">
                <h5 class="pt-3">{{ person.name }}</h5>
                <i>{{ person.role }}</i>
                <a :href="`mailto:${person.mail}`">{{ person.mail }}</a>
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
                <h3 class="text-primary">Kontakt</h3>
                <p>
                    Säkerhets-SM kan nås via mail på <a
                        href="mailto:sakerhetssm@kodsport.se">sakerhetssm@kodsport.se</a>
                </p>
                <h3 class="text-primary">Stöd</h3>
                <p>
                    Hosting av sakerhetssm.se sponsras av GleSYS,
                    vill eran organisation och stödja oss? Kontakta oss ovan!
                </p>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">

const people = [
    {
        name: "Movitz Sunar",
        role: "Projektledare",
        url: "/movitz.jpg",
        mail: "movitz.sunar@kodsport.se"
    },
    {
        name: "Martin Wennberg",
        role: "Vice Projektledare",
        url: "/martin.png",
        mail: "martin.wennberg@kodsport.se"
    },
    {
        name: "Mateusz Drwal",
        role: "Tekniskt ansvarig",
        url: "/mateusz.png",
        mail: "mateusz.drwal@kodsport.se"
    }
];

const url = useRequestURL()
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
