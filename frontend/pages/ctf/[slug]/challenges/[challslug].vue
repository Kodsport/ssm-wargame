<template>
  <div>
    <client-only>
      <CTFChallengeModal v-if="chall" :chall="chall" @back="router.push(`/ctf/${slug}/challenges`)" />
    </client-only>
  </div>
</template>

<script setup lang="ts">
import { useCTFStore } from "../../../../store/ctf";

definePageMeta({
  layout: "ctf",
  middleware: "ctf",
});

const route = useRoute();
const router = useRouter();
const store = useCTFStore();
const slug = route.params.slug as string;
const challSlug = route.params.challslug as string;
const ctfStore = useCTFStore();

const chall = computed(() => {
  return store.getBySlug(challSlug);
});

onMounted(async () => {
  await ctfStore.getCTF(slug);
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
