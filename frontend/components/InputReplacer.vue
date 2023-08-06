<template>
    <input type="text" disabled class="bg-secondary form-control" :placeholder="props.text">
</template>

<script setup lang="ts">
const props = defineProps(['text'])
</script>