import { useAuthStore } from "../store/auth";

export default defineNuxtRouteMiddleware((to, from) => {
  const auth = useAuthStore();

  if (
    to.path.startsWith("/admin") &&
    !(auth.user.role = "admin" || auth.user.role == "author")
  ) {
    return abortNavigation();
  }
});
