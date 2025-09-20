<template>
  <div class="ctf-gradient-bg">
    <CTFHeader />
    <div class="px-3 py-4">
      <slot />
    </div>
  </div>
</template>
<script setup lang="ts">
import { useCTFStore } from "~/store/ctf";
import { useRoute } from "vue-router";
import { onMounted, onUnmounted } from "vue";

const ctfStore = useCTFStore();
const route = useRoute();
let intervalId: any;

const time = 15 * 1000; // 15 seconds

onMounted(async () => {
  await new Promise((resolve) => setTimeout(resolve, time));

  // Refresh the CTF data every 15 seconds
  intervalId = setInterval(() => {
    const slug = route.params.slug as string;
    const password =
      (process.client && localStorage.getItem("ctf_password")) || undefined;
    if (slug) ctfStore.getAll(slug, password, true);
  }, time);
});

onUnmounted(() => {
  if (intervalId) clearInterval(intervalId);
});
</script>
<style scoped>
.ctf.gradient-bg {
  border-bottom: 3px solid;
  border-image: linear-gradient(to right, var(--bs-primary), #004353) 1;
}
</style>
