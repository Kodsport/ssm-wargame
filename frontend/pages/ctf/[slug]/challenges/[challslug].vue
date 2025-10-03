<template>
  <div>
    <client-only>
      <CTFChallengeModal
        v-if="chall"
        :chall="chall"
        @back="router.push(`/ctf/${slug}/challenges`)"
      />
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
</script>
