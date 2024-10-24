<template>
  <div class="container">
    <div>
      <div v-if="monthly.status.value == 'success'">
        <h1 class="text-primary">
          Månadens utmaning - {{ monthly.data.value.display_month }}
        </h1>
        <div class="col pt-4 pt-md-0 mb-5">
          <MonthlyChallenge :chall="monthly.data.value.challenge" />
        </div>
      </div>

      <h1 class="text-primary">Förra månaders utmaningar</h1>
      <div
        v-for="prev_monthly in prev_monthlies.data.value.filter(
          (e) => monthly?.data?.value?.challenge_id !== e.challenge_id
        )"
        v-if="prev_monthlies.status.value == 'success'"
        class="mb-2"
      >
        <div class="bg-dark rounded">
          <div
            class="text-primary rounded p-4 position-relative bg-body-hover pointer-cursor hover-thing"
            v-if="show_prev_monthly != prev_monthly.challenge_id"
            @click="show_prev_monthly = prev_monthly.challenge_id"
          >
            {{ prev_monthly.challenge.title }} -

            {{ prev_monthly.display_month }}
            {{ new Date(prev_monthly.start_date * 1000).getFullYear() }}

            <div
              class="badge bg-primary position-absolute"
              style="top: 50%; right: 20px; transform: translateY(-50%)"
            >
              {{ prev_monthly.challenge.category }}
            </div>
          </div>

          <MonthlyChallenge
            v-if="show_prev_monthly == prev_monthly.challenge_id"
            :chall="prev_monthly.challenge"
            :displayMonth="
              prev_monthly.display_month +
              ' ' +
              new Date(prev_monthly.start_date * 1000).getFullYear()
            "
          />
        </div>
      </div>
      <p class="mt-4" style="font-size: 8px">
        Om du vill hjälpa till och skapa nästa månadens utmaning, kontakta
        Movitz (@mvtz på discord) och skriv "skibidi utmaning" som kodord!
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useAuthStore } from "../store/auth";

const http = useHttp();
const auth = useAuthStore();

const monthly = await useAsyncData("monthly", () =>
  http("/current_monthly_challenge")
);

const show_prev_monthly = ref("");

const prev_monthlies = await useAsyncData("prev_monthly", () =>
  http("/monthly_challenges")
);

watch(
  () => auth.user,
  () => {
    if (auth.user.id) {
      monthly.refresh();
      prev_monthlies.refresh();
    }
  }
);
</script>
<style scoped>
.hover-thing:hover {
  transition: background-color 0.2s;
  background-color: #003642;
  cursor: pointer;
}
</style>
