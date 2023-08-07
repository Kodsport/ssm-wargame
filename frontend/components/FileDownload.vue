<template>
    <span class="material-icons text-primary">description</span>
    <a rel="nofollow" :href="props.file.url">{{ decodeURIComponent(props.file.filename) }}</a>
</template>

<script setup lang="ts">
const props = defineProps(['file'])
</script>