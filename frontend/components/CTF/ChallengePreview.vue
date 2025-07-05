<template>
  <div
    class="card h-100 challenge-preview-bg"
    :class="getBackgroundClass()"
  >
    <div class="card-body pb-2">
      <h5 class="title-text">{{ props.chall.title }}</h5>
    </div>
    <div class="card-body d-flex justify-content-between">
      <div v-if="!props.hideSolves" class="align-items-end d-flex solve-text">
        <template v-if="props.chall.numTeamsSolved !== undefined">
          <span v-if="props.chall.numTeamsSolved > 0"
            >{{ props.chall.numTeamsSolved }} lag</span
          >
          <span v-else>Olöst</span>
        </template>
        <template v-else>
          <span v-if="props.chall.numUsersSolved || props.chall.solves"
            >{{ props.chall.numUsersSolved || props.chall.solves }} lösare</span
          >
          <span v-else>Olöst</span>
        </template>
      </div>
      <h3 class="align-items-end d-flex mb-0 score-text">
        {{ props.chall.score }}
      </h3>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, watch } from 'vue';
import { useCTFStore } from '~/store/ctf';

const ctfStore = useCTFStore();
const props = defineProps(["chall", "hideSolves"]);

function getBackgroundClass() { // färger beroende på hur en chall e löst
  if (props.chall.solved) {
    return 'bg-success'; // ifall solvern är den som är inloggad
  }
  
  if (props.chall.solved_in_team) {
    return 'bg-warning'; // ifall nån i temeat har löst den men inte den som är inloggad
  }
  
  return 'bg-info'; // ingen i laget har löst den/ du har inte löst den heller
}

onMounted(async () => {
  const slug = ctfStore.ctf.slug;
  if (slug) {
    await ctfStore.getCTF(slug);
  }
  loadTheme();
});

watch(() => ctfStore.ctf.theme, () => {
  loadTheme();
});

function loadTheme() {
  const existingTheme = document.querySelector('link[data-ctf-theme]');
  if (existingTheme) {
    existingTheme.remove();
  }

  const themeName = ctfStore.ctf.theme || 'ctf-theme';

  const link = document.createElement('link');
  link.rel = 'stylesheet';
  link.href = `/themes/${themeName}.css`;
  link.setAttribute('data-ctf-theme', 'true');
  link.onerror = () => {
    const fallbackLink = document.createElement('link');
    fallbackLink.rel = 'stylesheet';
    fallbackLink.href = `/assets/themes/${themeName}.css`;
    fallbackLink.setAttribute('data-ctf-theme', 'true');
    fallbackLink.onload = () => {
    };
    document.head.appendChild(fallbackLink);
  };

  document.head.appendChild(link);
}
</script>
<style>
.title-text {
  font-size: 1.3rem;
  min-height: 50px;
}
.score-text {
  font-size: 1.5rem;
}
.solve-text span {
  font-size: 0.9rem;
}
.challenge-preview-bg {
  position: relative;
  overflow: hidden;
}
.challenge-preview-bg::before {
  content: "";
  position: absolute;
  inset: 0;
  z-index: 0;
  background-image: url('data:image/svg+xml;utf8,<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 900 675"><defs><filter id="blur1" x="-10%" y="-10%" width="120%" height="120%"><feFlood flood-opacity="0" result="BackgroundImageFix"/><feBlend mode="normal" in="SourceGraphic" in2="BackgroundImageFix" result="shape"/><feGaussianBlur stdDeviation="167" result="effect1_foregroundBlur"/></filter></defs><rect width="900" height="675" fill="%23008eae"/><g filter="url(%23blur1)"><circle cx="633" cy="560" fill="%230cc2b9" r="371"/><circle cx="522" cy="283" fill="%23008eae" r="371"/><circle cx="523" cy="46" fill="%230cc2b9" r="371"/><circle cx="303" cy="84" fill="%230cc2b9" r="371"/><circle cx="250" cy="448" fill="%23008eae" r="371"/><circle cx="775" cy="58" fill="%230cc2b9" r="371"/></g></svg>');
  background-size: cover;
  background-position: center;
  opacity: 0.5;
  pointer-events: none;
}
.card-body,
.title-text,
.score-text,
.solve-text {
  position: relative;
  z-index: 1;
}
</style>
