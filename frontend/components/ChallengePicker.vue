<template>
  <div class="form-group mt-2">
    <div class="challenge-picker-outer box card shadow-sm mb-2 p-2">
      <div class="form-group">
        <label
          class="form-label d-flex justify-content-between align-items-center"
        >
          <span>Challenges - Right click challenge for more info</span>

          <span>
            <span v-if="selectedChallenges.length === 0" class="text-muted">
              No challenges selected
            </span>
            <span v-else class="text-primary">
              {{ selectedChallenges.length }} challenges selected
            </span>
          </span>
        </label>
        <input
          class="form-control"
          v-model="searchChall"
          placeholder="Search for challenge name..."
        />
      </div>

      <div v-for="cat in categories" :key="cat.id" class="g-0 row">
        <h6
          class="text-primary pt-2 mb-0"
          v-if="
            challenges.find(
              (c) =>
                c.category_id === cat.id &&
                c.title.toLowerCase().includes(searchChall.toLowerCase())
            )
          "
        >
          {{ cat.name }}
        </h6>
        <div
          v-for="chall in challenges.filter(
            (c) =>
              c.category_id === cat.id &&
              c.title.toLowerCase().includes(searchChall.toLowerCase())
          )"
          :key="chall.id"
          class="col-12 col-sm-6 col-md-4 col-lg-3 g-2"
        >
          <div
            class="card h-100 challenge-picker-item border-2"
            :class="{
              'border-primary': !!selectedChallenges?.find(
                (e) => e.id === chall.id
              ),
              'border-secondary': !selectedChallenges?.find(
                (e) => e.id === chall.id
              ),
            }"
            @click="toggle(chall.id)"
            @contextmenu="showFlagModal(chall, $event)"
            style="cursor: pointer"
          >
            <div class="card-body py-1 px-2">
              <div class="d-flex justify-content-between align-items-center">
                <span class="fw-bold">{{ chall.title }}</span>
                <span class="badge bg-secondary">
                  {{ chall.category }}
                </span>
              </div>
              <div class="d-flex justify-content-between align-items-center">
                <span
                  class="small"
                  :class="{
                    'text-primary': !!selectedChallenges?.find(
                      (e) => e.id == chall.id
                    )?.custom_score,
                    'text-muted': !selectedChallenges?.find(
                      (e) => e.id == chall.id
                    )?.custom_score,
                  }"
                  >{{
                    selectedChallenges?.find((e) => e.id == chall.id)
                      ?.custom_score ?? chall.static_score
                  }}</span
                >
                <span
                  v-if="chall.hide"
                  class="material-symbols-outlined text-secondary"
                >
                  visibility_off
                </span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="mt-2"></div>
    <div
      v-if="flagModal"
      class="flag-modal-backdrop model-content"
      @click="closeFlagModal"
    >
      <div class="flag-modal" @click.stop>
        <div class="d-flex justify-content-between align-items-center mb-2">
          <strong>{{ flagModal.title }}</strong>
        </div>
        <div class="mt-2">
          <span class="text-monospace">{{ flagModal.flag }}</span>
        </div>
        <div class="mt-2">
          <div class="form-group">
            <label>Custom score (Leave blank for default)</label>
            <input
              class="form-control"
              v-model="flagModal.custom_score"
              @input="applyCustomScore(flagModal.custom_score)"
              placeholder="69"
            />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import { ref, onMounted, defineModel } from "vue";
import { useChallengeStore } from "@/store/admin/challenges";

const challStore = useChallengeStore();
const selectedChallenges = defineModel<any[]>();
const challenges = ref<any[]>([]);
const categories = ref<any[]>([]);
const searchChall = ref<string>("");
const flagModal = ref<{
  id: string;
  flag: string;
  title: string;
  custom_score: number | null;
} | null>(null);

onMounted(async () => {
  await challStore.getChallenges();
  await challStore.getCategories();
  challenges.value = challStore.challenges;
  categories.value = challStore.categories;
  challenges.value = challenges.value.map((chall) => {
    const category = challStore.getCategory(chall.category_id);
    return {
      ...chall,
      category: category ? category.name : "???",
    };
  });
});

function toggle(id: string) {
  const idx = selectedChallenges.value.findIndex((e) => e.id === id);
  if (idx === -1)
    selectedChallenges.value?.push({
      id: id,
      custom_score: null,
    });
  else selectedChallenges.value.splice(idx, 1);
}

function showFlagModal(chall: any, event: MouseEvent) {
  event.preventDefault();
  const flags = chall.flags
    .map((flag: any) => "SSM{" + flag.flag + "}")
    .join(", ");
  const custom_score = selectedChallenges.value.find(
    (e) => e.id === chall.id
  )?.custom_score;
  flagModal.value = {
    id: chall.id,
    flag: flags,
    title: chall.title,
    custom_score: custom_score ?? null,
  };
}

function applyCustomScore(score: string | null) {
  if (flagModal.value) {
    const idx = selectedChallenges.value.findIndex(
      (e) => e.id === flagModal.value?.id
    );
    if (idx !== -1 && score !== null && score !== "") {
      selectedChallenges.value[idx].custom_score = Number(score) ?? null;
    } else {
      selectedChallenges.value[idx].custom_score = null;
    }
  }
}

function closeFlagModal() {
  flagModal.value = null;
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
  height: 300px;
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
.flag-modal-backdrop {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.3);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}
.flag-modal {
  background: var(--bs-body-bg);
  border-radius: 8px;
  padding: 20px 24px;
  margin: 0 10px;
  min-width: 260px;
  box-shadow: 0 2px 16px rgba(0, 0, 0, 0.2);
}
.text-monospace {
  font-family: "Fira Mono", "Consolas", "Menlo", monospace;
  font-size: 1.1em;
  color: #c2c2c2;
  word-break: break-all;
}
</style>
