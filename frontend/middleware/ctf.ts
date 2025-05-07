import { useCTFStore } from "~/store/ctf";
import { useAuthStore } from "~/store/auth";

export default defineNuxtRouteMiddleware(async (to) => {
  const ctfStore = useCTFStore();
  const authStore = useAuthStore();
  const password =
    (process.client && localStorage.getItem("ctf_password")) || undefined;
  const slug = to.params.slug as string;
  if (slug) {
    await ctfStore.getAll(slug, password);
  }
  if (password) {
    await authStore.getCTFUserOnce(slug);
  }
});
