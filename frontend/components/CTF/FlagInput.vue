<template>
  <div class="form-group">
    <template v-if="props.solved">
      <InputReplacer text="Löst!" />
    </template>
    <template v-else>
      <input
        v-if="!!auth.ctfUser.id"
        type="text"
        class="form-control"
        placeholder="SSM{..."
        :value="modelValue"
        @input="$emit('update:modelValue', $event.target.value)"
      />
      <InputReplacer v-else text="Logga in för att skicka in flaggor!" />
    </template>
  </div>
</template>

<script setup lang="ts">
import { useAuthStore } from "~/store/auth";
import { useCTFStore } from "~/store/ctf";

const auth = useAuthStore();
const ctfStore = useCTFStore();
const props = defineProps(["modelValue", "solved"]);
defineEmits(["update:modelValue"]);

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
