<template>
    <div class="challenge" @click="openModal()">
        <div>
            <span class="title">{{ props.chall.title }}</span>
            <ul class="categories">
                <li v-for="category in props.chall.categories" :key="category">
                    {{ category }}
                </li>
            </ul>
            <span class="solves" :title="`This challenge has been sloved by ${props.chall.solves} teams`">{{
                props.chall.solves
            }}</span>
            <span class="score" :title="`This challenge gives ${props.chall.score} points to each team that solves it`">{{
                props.chall.score }}</span>
        </div>
    </div>
</template>
  
<script setup lang="ts">

const props = defineProps(['chall'])

const router = useRouter()

function openModal() {
    if (props.chall.id) {
        router.push()
    }
}

</script>
  
<style lang="scss" scoped>
$font-size-base: 1rem;
$font-family-sans-serif: Roboto, sans-serif;

$dark-transparent-background: #0004;

$font-challenge-color: #FFF;
$standard-border-radius: 0.5rem;

$challenge-color: #008eae;
$challenge-height: 8rem;
$challenge-min-width: 15rem;
$challenge-spacing: 1rem;
$challenge-chain-arrow-depth: 2rem;
$challenge-chain-spacing: 0.1rem;



.challenge {
    height: $challenge-height;
    min-width: $challenge-min-width;
    background-color: $challenge-color;
    border-color: $challenge-color;
    margin: $challenge-spacing/2 $challenge-spacing $challenge-spacing/2 0;
    padding: 0.7rem;
    flex-wrap: wrap;
    color: $font-challenge-color;
    border-radius: $standard-border-radius;
    font-weight: 600;
    position: relative;

    div {
        position: relative;
        height: 100%;
        width: 100%;
    }
}

.title {
    font-size: 1.4rem;
}

.solves {
    position: absolute;
    bottom: 0;
    left: 0;
    color: $dark-transparent-background;
}

.score {
    position: absolute;
    bottom: 0;
    right: 0;
    font-size: 3rem;
    line-height: 1;
}

.categories {
    list-style: none;
    margin: 0;
    padding: 0;
    display: flex;

    li {
        background-color: $dark-transparent-background;
        padding: 0rem 0.2rem;
        border-radius: 0.3rem;
        font-size: 1rem;
        margin-right: 0.2rem;
    }
}

.solved {
    background-color: #2cdd2c;
    border-color: #2cdd2c;
}

.locked {
    background-color: #888;
    border-color: #888;
    cursor: default;
}

.outsidelist {
    background-color: $challenge-color;
    border-color: $challenge-color;
    cursor: default;
}
</style>
  