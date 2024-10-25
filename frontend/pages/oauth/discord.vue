<template>
  <div class="d-flex flex-column align-items-center mt-5">
    <h1 class="text-primary">Var god vänta...</h1>
    <div class="text-primary">
      Om du är fast här, stäng denna sida och försök logga in igen.
    </div>
    <p class="mt-2 red-error" v-if="errorMsg">
      {{ errorMsg }}
    </p>
    <img
      class="ks-logo mt-5"
      src="~/assets/kodsport.png"
      alt="Kodsport Logga"
      width="400"
    />
  </div>
</template>

<script setup lang="ts">
import { useAuthStore } from "../../store/auth";
const auth = useAuthStore();
const router = useRouter();

const errorMsg = ref("");

async function getToken() {
  if (process.server) return;

  const params = new URLSearchParams(window.location.search);

  if (auth.isAuthed || !params.has("code")) {
    router.push("/");
  }

  window.opener.postMessage({
    type: "auth",
    code: params.get("code"),
    state: params.get("state"),
  });

  window.close();
}

getToken();
</script>

<style scoped>
.red-error {
  color: var(--bs-red);
  font-weight: 700;
}
</style>
