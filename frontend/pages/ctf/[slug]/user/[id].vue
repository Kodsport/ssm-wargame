<template>
  <div class="container mb-4">
    <h2>Lösta Utmaningar - {{ username }}</h2>
    <div>
      <div v-if="solves.length">
        <table class="table">
          <thead>
            <tr>
              <th>Utmaning</th>
              <th>Löst</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="solve in solves" :key="solve.challenge_id">
              <td>
                {{
                  ctfStore.challenges.find(
                    (c: any) => c.id === solve.challenge_id
                  )?.title || solve.challenge_id
                }}
              </td>
              <td>{{ new Date(solve.solved_at).toLocaleString() }}</td>
            </tr>
          </tbody>
        </table>
      </div>
      <div v-else>
        <p>Inga utmaningar lösta ännu.</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from "vue";
import { useRoute } from "vue-router";
import useHttp from "@/composables/use-http";
import { useCTFStore } from "~/store/ctf";

definePageMeta({
  layout: "ctf",
  middleware: "ctf",
});

const route = useRoute();
const http = useHttp();
const ctfStore = useCTFStore();
const slug = route.params.slug as string;
const userId = route.params.id as string;
const store = useCTFStore();
const challSlug = route.params.challslug as string;

const data = await http(`/ctfs/${slug}/user/${userId}/solves`);
const username = ref(data.username);
const solves = ref<Array<{ challenge_id: string; solved_at: number }>>(
  data.solves
);

const chall = computed(() => {
  return store.getBySlug(challSlug);
});

onMounted(async () => {
  await ctfStore.getCTF(slug);
  loadTheme();
});

watch(
  () => ctfStore.ctf.theme,
  () => {
    loadTheme();
  }
);

function loadTheme() {
  const existingTheme = document.querySelector("link[data-ctf-theme]");
  if (existingTheme) {
    existingTheme.remove();
  }

  const themeName = ctfStore.ctf.theme || "ctf-theme";

  const link = document.createElement("link");
  link.rel = "stylesheet";
  link.href = `/themes/${themeName}.css`;
  link.setAttribute("data-ctf-theme", "true");
  link.onerror = () => {
    const fallbackLink = document.createElement("link");
    fallbackLink.rel = "stylesheet";
    fallbackLink.href = `/assets/themes/${themeName}.css`;
    fallbackLink.setAttribute("data-ctf-theme", "true");
    fallbackLink.onload = () => {};
    document.head.appendChild(fallbackLink);
  };

  document.head.appendChild(link);
}
</script>
<style scoped>
.table > :not(caption) > * > * {
  background-color: #001125 !important;
  background-color: #00000000 !important;
}
</style>
