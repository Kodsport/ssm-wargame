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
.ctf-gradient-bg {
  min-height: 100vh;
  background: linear-gradient(135deg, #004353 0%, #005053 100%);
}
</style>
