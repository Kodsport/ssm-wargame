<template>
  <div class="form-group mt-2">
    <div class="challenge-picker-outer box card shadow-sm mb-2 p-2">
      <p class="mb-2">
        Challenge groups
        <span class="small text-muted"
          >Only appends not already selected challenges.
          <nuxt-link to="/admin/ctfs/challenge_groups"
            >View all groups</nuxt-link
          >.</span
        >
      </p>
      <div
        v-for="group in groups"
        :key="group.id"
        class="col-12 col-sm-6 col-md-4 col-lg-3 g-2"
      >
        <div
          class="card h-100 challenge-picker-item border-2 border-secondary"
          @click="apply(group.challenges)"
          style="cursor: pointer"
        >
          <div class="card-body py-1 px-2">
            <div class="d-flex justify-content-between align-items-center">
              <span class="fw-bold">{{ group.name }}</span>
            </div>
            <div class="d-flex justify-content-between align-items-center">
              <span class="text-muted small">{{ group.description }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import { ref, onMounted, defineModel } from "vue";
import { useChallengeStore } from "@/store/admin/challenges";
import useHttp from "@/composables/use-http";

const selectedChallenges = defineModel<any[]>();
const groups = ref<any[]>([]);
const http = useHttp();

onMounted(async () => {
  groups.value = await http("/admin/challenge_groups");
});

function apply(challenges: Record<string, unknown>[]) {
  for (const chall of challenges) {
    if (!selectedChallenges.value?.find((e) => e.id === chall.id)) {
      selectedChallenges.value.push({
        id: chall.id,
        custom_score: chall.custom_score,
        display_order: chall.display_order,
      });
    }
  }
}
</script>
<style scoped>
.challenge-picker-item {
  transition: border 0.2s, background 0.2s;
  user-select: none;
}
.challenge-picker-item:hover {
  background: #0000001f;
}
.box {
  overflow-y: auto;
  height: 200px;
  scrollbar-width: thin;
  scrollbar-color: #b0b0b0 #0000;
}
.box::-webkit-scrollbar {
  width: 6px;
  background: transparent;
}
.box::-webkit-scrollbar-thumb {
  background: #b0b0b0;
  border-radius: 4px;
}
.challenge-picker-outer {
  background: #0000001c;
}
</style>
