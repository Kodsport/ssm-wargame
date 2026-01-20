<template>
  <div class="container">
    <h1>Poängtavla - {{ slug }}</h1>
    
    <table class="table table-hover align-middle scoreboard-table mt-4">
      <thead class="table-active">
        <tr>
          <th>#</th>
          <th>{{ isTeamBased ? "Lagnamn" : "Namn" }}</th>
          <th>Poäng</th>
          <th>Lösningar</th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="(row, i) in scoreboard"
          :key="row.username"
          :class="{
            'scoreboard-gold': i === 0,
            'scoreboard-silver': i === 1,
            'scoreboard-bronze': i === 2,
          }"
        >
          <td class="fw-bold">{{ i + 1 }}</td>
          <td>{{ row.username }}</td>
          <td class="fw-bold">{{ row.score }}</td>
          <td>{{ row.solves ? row.solves.length : 0 }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useRoute } from "vue-router";
import useHttp from "@/composables/use-http";

const route = useRoute();
const http = useHttp();
const slug = route.params.slug as string;
const scoreboard = ref<any[]>([]);
const isTeamBased = ref(false);

onMounted(async () => {
  try {
    scoreboard.value = await http(`/admin/ctfs/${slug}/scoreboard`);
    
    const ctfs = await http("/admin/ctfs");
    const ctf = ctfs.find((c: any) => c.slug === slug);
    if (ctf) {
      isTeamBased.value = ctf.team_based;
    }
  } catch (e) {
    console.error("Failed to load scoreboard:", e);
  }
});
</script>

<style scoped>
.scoreboard-table {
  font-size: 1.1rem;
  border-radius: 12px 12px 0 0;
  overflow: hidden;
}

.table > :not(caption) > * > * {
  background-color: #00000000;
}

.scoreboard-table th,
.scoreboard-table td {
  vertical-align: middle;
}

.scoreboard-gold td {
  background: #a1862a !important;
  color: #fff6d3 !important;
  font-weight: bold;
}

.scoreboard-silver td {
  background: #52575c !important;
  color: #e7e7e7 !important;
  font-weight: bold;
}

.scoreboard-bronze td {
  background: #7a4e21 !important;
  color: #ffe3c6 !important;
  font-weight: bold;
}
</style>
