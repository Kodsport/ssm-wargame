<template>
  <header class="d-flex d-md-none p-5 justify-content-center">
    <h1>
      <nuxt-link class="text-light text-decoration-none" to="/">Säkerhets-SM</nuxt-link>
    </h1>
  </header>
  <nav class="navbar navbar-expand-md bg-dark">
    <div class="container-fluid">
      <button class="navbar-toggler color-primary" type="button" @click="toggleCollapse = !toggleCollapse">
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" :class="{ 'show': toggleCollapse }" id="navbarSupportedContent">
        <nuxt-link class="navbar-brand nav-link d-none d-md-inline" to="/">SSM</nuxt-link>
        <ul class="navbar-nav me-auto">
          <li class="nav-item d-none d-md-inline">
            <nuxt-link active-class="active" class="nav-link" to="/challenges">
              Utmaningar
            </nuxt-link>
          </li>
          <li class="nav-item">
            <nuxt-link active-class="active" class="nav-link" to="/scoreboard">
              Poängtavla
            </nuxt-link>
          </li>
          <li class="nav-item">
            <nuxt-link active-class="active" class="nav-link" to="/learn">
              Läroresurser
            </nuxt-link>
          </li>
          <li class="nav-item">
            <nuxt-link active-class="active" class="nav-link" to="/about">
              Om oss
            </nuxt-link>
          </li>
        </ul>
        <client-only>
          <ul v-if="auth.user" class="navbar-nav pe-3">
            <template v-if="auth.user.role == 'admin'">

              <li class="nav-item">
                <nuxt-link active-class="active" class="nav-link" to="/admin/users">
                  Users
                </nuxt-link>
              </li>
              <li class="nav-item">
                <nuxt-link active-class="active" class="nav-link" to="/admin/challenges">
                  Challenges
                </nuxt-link>
              </li>
              <li class="nav-item">
                <nuxt-link active-class="active" class="nav-link" to="/admin/monthly">
                  Monthly
                </nuxt-link>
              </li>
              <li class="nav-item">
                <nuxt-link active-class="active" class="nav-link" to="/admin/events">
                  Events
                </nuxt-link>
              </li>
            </template>
            <li class="nav-item">
              <nuxt-link v-if="!!auth.user.id" active-class="active" class="nav-link btn border border-primary"
                to="/user">
                <span class="d-flex">
                  {{ auth.user.full_name }}
                  <span class="material-icons text-primary">person</span>
                </span>

              </nuxt-link>
            </li>
          </ul>
          <template class="d-none d-md-inline">
            <a v-if="!isAuthenticated" class="btn btn-primary" @click="login">Logga in</a>
            <a v-else class="btn btn-primary" @click="logout">Logga ut</a>
          </template>
          <template #fallback class="d-none d-md-inline">
            <a class="btn btn-primary">Logga in</a>
          </template>
        </client-only>
      </div>
    </div>
  </nav>
</template>

<script setup lang="ts">
import { useAuthStore } from "../store/auth"
const http = useHttp()
const auth = useAuthStore()
const router = useRouter()

var isAuthenticated = computed(() => auth.token != "");

var toggleCollapse = ref(false)

// temporary hacky solution...
setInterval(() => {

  const part = auth.token.split('.')[1]
  if (!part) {
    return
  }
  const jwtData = JSON.parse(atob(part))
  if (jwtData * 1000 < Date.now()) {
    auth.setToken('')
    localStorage.removeItem('ssm-token')
  }

}, 1000)

onMounted(() => {
  const jwt = localStorage.getItem('ssm-token')
  if (jwt) {
    auth.setToken(jwt)
  }
})

function logout() {
  auth.setToken('')
  localStorage.removeItem('ssm-token')
  router.push('/')
}

async function login() {
  const x = await http("/auth/discord/url");

  const w = window.open(x.url);

  if (!w) {
    // TODO error handling
    return;
  }

  const intv = setInterval(async () => {
    try {
      if (w.location.host !== window.location.host) {
        return;
      }

      const params = new URLSearchParams(w.location.search);

      clearInterval(intv);
      w.close();

      const resp = await http("/auth/discord/exchange", {
        method: 'POST',
        body: {
          code: params.get("code"),
          state: params.get("state"),
        }
      });

      auth.setToken(resp.jwt)
      auth.getUser()
      localStorage.setItem('ssm-token', resp.jwt)
    } catch (error) {
      // DOMException is thrown when w is on other domain
    }
  }, 30);
}
</script>

<style scoped lang="scss">
header {
  background-image: url('data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHZpZXdCb3g9IjAgMCAxMjcxLjk0IDY1OC40NyI+PGRlZnM+PHN0eWxlPi5jbHMtMXtmaWxsOiMwMDJjMzY7fS5jbHMtMntmaWxsOm5vbmU7c3Ryb2tlLXdpZHRoOjRweDt9LmNscy0yLC5jbHMtNHtzdHJva2U6I2Q0ODYwMDtzdHJva2UtbWl0ZXJsaW1pdDoxMDt9LmNscy0zLC5jbHMtNHtmaWxsOiNkNDg2MDA7fTwvc3R5bGU+PC9kZWZzPjx0aXRsZT5TU00yPC90aXRsZT48ZyBpZD0iYmdfZmlsbCIgZGF0YS1uYW1lPSJiZyBmaWxsIj48cmVjdCBjbGFzcz0iY2xzLTEiIHdpZHRoPSIxMjcxLjk0IiBoZWlnaHQ9IjY1OC40NyIvPjxwYXRoIGNsYXNzPSJjbHMtMiIgZD0iTTEwMTAsNDU5LjA4bDkxLjc2LTk0LjczWiIgdHJhbnNmb3JtPSJ0cmFuc2xhdGUoLTMxNS42NCAtMTkyLjM2KSIvPjwvZz48ZyBpZD0ibGluZXMiPjxwYXRoIGNsYXNzPSJjbHMtMiIgZD0iTTYxMS40MiwzNzAuNjciIHRyYW5zZm9ybT0idHJhbnNsYXRlKC0zMTUuNjQgLTE5Mi4zNikiLz48cGF0aCBjbGFzcz0iY2xzLTIiIGQ9Ik02ODkuMjUsMjYyLjY3IiB0cmFuc2Zvcm09InRyYW5zbGF0ZSgtMzE1LjY0IC0xOTIuMzYpIi8+PGxpbmUgY2xhc3M9ImNscy0yIiB4MT0iOTYuNCIgeTE9IjI0MC4yMyIgeDI9Ijc5Ljg2IiB5Mj0iNDM0LjY0Ii8+PGxpbmUgY2xhc3M9ImNscy0yIiB4MT0iODIuMDIiIHkxPSI0MzQuNjQiIHgyPSIxMjAuMDIiIHkyPSIzNDkuNDgiLz48bGluZSBjbGFzcz0iY2xzLTIiIHgxPSI5NC4wMiIgeTE9IjQ1NC40NCIgeDI9IjI5OS4zNiIgeTI9IjQ3OS42NCIvPjxsaW5lIGNsYXNzPSJjbHMtMiIgeDE9IjEwNS4yMyIgeTE9IjI0MC4zOSIgeDI9IjMyMy42MSIgeTI9IjU4Mi42NCIvPjxsaW5lIGNsYXNzPSJjbHMtMiIgeDE9IjMyOC44NiIgeTE9IjU3OS42NCIgeDI9IjMxNi44NSIgeTI9IjUwMi4zMyIvPjxsaW5lIGNsYXNzPSJjbHMtMiIgeDE9IjM0MS4xOCIgeTE9IjU5MC4zMiIgeDI9IjQ1OS42NyIgeTI9IjU2OS42NCIvPjxsaW5lIGNsYXNzPSJjbHMtMiIgeDE9IjU5Mi4xOSIgeTE9IjUwMy4xNCIgeDI9IjYwNy4wMiIgeTI9IjU4MS40OCIvPjxsaW5lIGNsYXNzPSJjbHMtMiIgeDE9IjYxOC45OCIgeTE9IjU4NS44OSIgeDI9Ijc3Ni4zNiIgeTI9IjQyOC45OCIvPjxsaW5lIGNsYXNzPSJjbHMtMiIgeDE9IjM5Ni42NSIgeTE9IjcyLjc3IiB4Mj0iNzgyLjAyIiB5Mj0iMTY1LjI3Ii8+PGxpbmUgY2xhc3M9ImNscy0yIiB4MT0iNTYxLjczIiB5MT0iNzYuMjciIHgyPSI2NjYuOTMiIHkyPSI2Ny45NSIvPjxsaW5lIGNsYXNzPSJjbHMtMiIgeDE9Ijc5Ny43MyIgeTE9IjE2My43NyIgeDI9Ijg4Mi44NiIgeTI9IjExOC4zOSIvPjxsaW5lIGNsYXNzPSJjbHMtMiIgeDE9Ijg5OS44MSIgeTE9IjEyMi44NCIgeDI9Ijk2My44NiIgeTI9IjIyMS4xNSIvPjxsaW5lIGNsYXNzPSJjbHMtMiIgeDE9IjkwMy45MiIgeTE9IjExNy43OCIgeDI9IjEwMjAuMzYiIHkyPSIxODAuMzkiLz48bGluZSBjbGFzcz0iY2xzLTIiIHgxPSIxMDUzLjE3IiB5MT0iMTk5LjI3IiB4Mj0iMTE0Mi4yNCIgeTI9IjI1Ni4zNiIvPjxsaW5lIGNsYXNzPSJjbHMtMiIgeDE9IjExNTUuMDYiIHkxPSIyNzMuNiIgeDI9IjExODUuMjMiIHkyPSI0MzguMzciLz48bGluZSBjbGFzcz0iY2xzLTIiIHgxPSIxMTc3LjU3IiB5MT0iNDQ0LjgxIiB4Mj0iOTc1Ljg2IiB5Mj0iMzA1LjY0Ii8+PGxpbmUgY2xhc3M9ImNscy0yIiB4MT0iMTE0NC4zOCIgeTE9IjI3MS4zNSIgeDI9IjEwNDUuNiIgeTI9IjQyNi43NiIvPjxsaW5lIGNsYXNzPSJjbHMtMiIgeDE9IjkxOC4wMiIgeTE9IjQwMy4zMSIgeDI9Ijk2My41MiIgeTI9IjUxMi45OCIvPjxsaW5lIGNsYXNzPSJjbHMtMiIgeDE9IjEwMzQuNzciIHkxPSIyMDguODkiIHgyPSI5OTAuOTgiIHkyPSI1MDYuMzkiLz48bGluZSBjbGFzcz0iY2xzLTIiIHgxPSI5NjEuNTIiIHkxPSI1MTcuNTYiIHgyPSI3OTIuNjkiIHkyPSI0MjcuNjQiLz48bGluZSBjbGFzcz0iY2xzLTIiIHgxPSI5NjIuMzYiIHkxPSI1MTUuMzEiIHgyPSI1NTcuNzMiIHkyPSI4MC41OCIvPjxsaW5lIGNsYXNzPSJjbHMtMiIgeDE9IjkyNC43NyIgeTE9IjMwNy4zMSIgeDI9Ijk1NC40MiIgeTI9IjMwMC4xNCIvPjxsaW5lIGNsYXNzPSJjbHMtMiIgeDE9Ijk3MC45MiIgeTE9IjI0MS41OCIgeDI9Ijk2Ny42NSIgeTI9IjI4Ny4yOSIvPjxsaW5lIGNsYXNzPSJjbHMtMiIgeDE9IjkxMy45NyIgeTE9IjI5Ni40NyIgeDI9Ijk2Mi45OCIgeTI9IjIzOC4zNiIvPjxsaW5lIGNsYXNzPSJjbHMtMiIgeDE9IjEwMjEuNDgiIHkxPSIyMDAuMTQiIHgyPSI5ODEuNTQiIHkyPSIyMjUuODkiLz48bGluZSBjbGFzcz0iY2xzLTIiIHgxPSIxMDE3LjA0IiB5MT0iMTkyLjMyIiB4Mj0iNjk0LjM0IiB5Mj0iMjY5LjM5Ii8+PGxpbmUgY2xhc3M9ImNscy0yIiB4MT0iNjk1Ljg2IiB5MT0iMjcxLjk4IiB4Mj0iOTU5LjQ4IiB5Mj0iMjMwLjE0Ii8+PGxpbmUgY2xhc3M9ImNscy0yIiB4MT0iNzg4LjIzIiB5MT0iNDMxLjUyIiB4Mj0iODQ3Ljk0IiB5Mj0iNTU3LjY2Ii8+PGxpbmUgY2xhc3M9ImNscy0yIiB4MT0iNjA3LjAyIiB5MT0iNDg4LjY0IiB4Mj0iODQxLjIzIiB5Mj0iNTY1Ljk4Ii8+PGxpbmUgY2xhc3M9ImNscy0yIiB4MT0iODk1LjIzIiB5MT0iMzM3LjUxIiB4Mj0iODU2LjEyIiB5Mj0iNTU2Ljc2Ii8+PGxpbmUgY2xhc3M9ImNscy0yIiB4MT0iOTA2LjI4IiB5MT0iMzg1LjMxIiB4Mj0iNjMwLjc2IiB5Mj0iMTEyLjAyIi8+PGxpbmUgY2xhc3M9ImNscy0yIiB4MT0iNjMxLjUxIiB5MT0iOTguMjUiIHgyPSI2NjguMDciIHkyPSI3NS43NyIvPjxsaW5lIGNsYXNzPSJjbHMtMiIgeDE9IjY4Ny43NiIgeTE9Ijc2LjI3IiB4Mj0iNzg0Ljg2IiB5Mj0iMTYwLjAyIi8+PGxpbmUgY2xhc3M9ImNscy0yIiB4MT0iNjI2LjYiIHkxPSIxMTUuMTIiIHgyPSI2NjIuMDQiIHkyPSIyNTkuOTIiLz48bGluZSBjbGFzcz0iY2xzLTIiIHgxPSI1NjkuMzYiIHkxPSIyMDAuNjQiIHgyPSI2NDguODYiIHkyPSIyNjYuNzMiLz48bGluZSBjbGFzcz0iY2xzLTIiIHgxPSI0NTUuNzciIHkxPSIyMjcuODQiIHgyPSI2NDIuMzciIHkyPSIyNzcuOTEiLz48bGluZSBjbGFzcz0iY2xzLTIiIHgxPSI2NjEuMzkiIHkxPSIzMDMuNCIgeDI9IjY0MS41MyIgeTI9IjM2My40NiIvPjxsaW5lIGNsYXNzPSJjbHMtMiIgeDE9IjY4Mi4zOCIgeTE9IjI5OS4zNiIgeDI9Ijc3MC41MyIgeTI9IjM5MS4zMyIvPjxsaW5lIGNsYXNzPSJjbHMtMiIgeDE9Ijc4NC43MSIgeTE9IjM5Ny4zMSIgeDI9IjkwMiIgeTI9IjM5NC4yMyIvPjxsaW5lIGNsYXNzPSJjbHMtMiIgeDE9IjkxOC4zMiIgeTE9IjMzMS4zNCIgeDI9IjEwMjcuMzgiIHkyPSI0MzAuMTYiLz48bGluZSBjbGFzcz0iY2xzLTIiIHgxPSI2NDcuMzgiIHkxPSIzNzQuODkiIHgyPSI3NzQuMDciIHkyPSI0MTguODkiLz48bGluZSBjbGFzcz0iY2xzLTIiIHgxPSI2MDUuOTgiIHkxPSI0MTEuNTIiIHgyPSI1OTQuODYiIHkyPSI0NjMuMzEiLz48bGluZSBjbGFzcz0iY2xzLTIiIHgxPSI2MTMuOTUiIHkxPSI0MDMuNjQiIHgyPSI3NjgiIHkyPSIzOTcuMzEiLz48bGluZSBjbGFzcz0iY2xzLTIiIHgxPSI1NDkuODQiIHkxPSI4MS42NCIgeDI9IjQ3OS4zNiIgeTI9IjMyNi42MSIvPjxsaW5lIGNsYXNzPSJjbHMtMiIgeDE9IjU0Ni43NyIgeTE9Ijc4LjciIHgyPSI0NTMuMTEiIHkyPSIyMTcuODkiLz48bGluZSBjbGFzcz0iY2xzLTIiIHgxPSIzOTAuMTEiIHkxPSI3Ni43NyIgeDI9IjQzMi4xMSIgeTI9IjE0OC44OSIvPjxsaW5lIGNsYXNzPSJjbHMtMiIgeDE9IjQ0MCIgeTE9IjE2My4yMSIgeDI9IjQ0NS44NiIgeTI9IjIxNS42NCIvPjxsaW5lIGNsYXNzPSJjbHMtMiIgeDE9IjI5Mi40OCIgeTE9IjE5OC4wMiIgeDI9IjQyOC4wMyIgeTI9IjE1Ny4zNCIvPjxsaW5lIGNsYXNzPSJjbHMtMiIgeDE9IjI4OS43MyIgeTE9IjIxMi43NyIgeDI9IjU2Ni44NyIgeTI9IjQ2Ny4yMiIvPjxsaW5lIGNsYXNzPSJjbHMtMiIgeDE9IjQ4Mi42OSIgeTE9IjM0Ny42NCIgeDI9IjYwMC45OCIgeTI9IjQwMS43NyIvPjxsaW5lIGNsYXNzPSJjbHMtMiIgeDE9IjQ4Ny4wMiIgeTE9IjM0My4xNCIgeDI9IjYzMC43NiIgeTI9IjM2OS4yNyIvPjxsaW5lIGNsYXNzPSJjbHMtMiIgeDE9IjU2NC45OCIgeTE9IjIwMC45NSIgeDI9IjU4NC45OSIgeTI9IjQ2My4zMSIvPjxsaW5lIGNsYXNzPSJjbHMtMiIgeDE9IjQ3My4yOSIgeTE9IjU1Ny44NyIgeDI9IjU2MC4xNyIgeTI9IjIwMC42NCIvPjxsaW5lIGNsYXNzPSJjbHMtMiIgeDE9IjU1My42OSIgeTE9IjgxLjk4IiB4Mj0iNTYwLjg2IiB5Mj0iMTg5LjgxIi8+PGxpbmUgY2xhc3M9ImNscy0yIiB4MT0iMTQ0LjIzIiB5MT0iMzI0LjE0IiB4Mj0iMjY1LjIzIiB5Mj0iMjE0Ljc3Ii8+PGxpbmUgY2xhc3M9ImNscy0yIiB4MT0iMTQzLjYxIiB5MT0iMzI2LjMxIiB4Mj0iNDQwLjI4IiB5Mj0iMjI2LjE5Ii8+PGxpbmUgY2xhc3M9ImNscy0yIiB4MT0iMTQ1Ljk4IiB5MT0iMzI2LjMxIiB4Mj0iMzg5Ljc2IiB5Mj0iNDAxLjk3Ii8+PGxpbmUgY2xhc3M9ImNscy0yIiB4MT0iNDQ2Ljk4IiB5MT0iMjMyLjM1IiB4Mj0iNDA2LjYxIiB5Mj0iMzkzLjQ0Ii8+PGxpbmUgY2xhc3M9ImNscy0yIiB4MT0iNDEwLjc5IiB5MT0iNDEyLjA5IiB4Mj0iNDM5LjM2IiB5Mj0iNDM0LjY0Ii8+PGxpbmUgY2xhc3M9ImNscy0yIiB4MT0iNDY5LjM2IiB5MT0iMzQ5LjY4IiB4Mj0iNDQ4LjM2IiB5Mj0iNDMyLjk4Ii8+PGxpbmUgY2xhc3M9ImNscy0yIiB4MT0iNDA4LjUyIiB5MT0iMzk0LjczIiB4Mj0iNDYzLjE5IiB5Mj0iMzQ5LjY4Ii8+PGxpbmUgY2xhc3M9ImNscy0yIiB4MT0iMjc4LjU4IiB5MT0iMjE5LjYiIHgyPSIzMTEuMTkiIHkyPSI0NjkuNjQiLz48bGluZSBjbGFzcz0iY2xzLTIiIHgxPSI0MzIuODIiIHkxPSIxNjEuOCIgeDI9IjMyMS4zNiIgeTI9IjQ3MC42NCIvPjxsaW5lIGNsYXNzPSJjbHMtMiIgeDE9IjQzOS4zNiIgeTE9IjQ0NC4xNCIgeDI9IjMzMS4zNiIgeTI9IjQ4MS42NCIvPjxsaW5lIGNsYXNzPSJjbHMtMiIgeDE9IjE5NC41NCIgeTE9IjUxMy4xNCIgeDI9IjM3Ni42MSIgeTI9Ijc2LjI3Ii8+PHBhdGggY2xhc3M9ImNscy0yIiBkPSJNNjkyLjI1LDI2Ni44OCIgdHJhbnNmb3JtPSJ0cmFuc2xhdGUoLTMxNS42NCAtMTkyLjM2KSIvPjxsaW5lIGNsYXNzPSJjbHMtMiIgeDE9IjM3NC40NCIgeTE9Ijc0LjUyIiB4Mj0iMjg1LjkzIiB5Mj0iMTkwLjU4Ii8+PGxpbmUgY2xhc3M9ImNscy0yIiB4MT0iNDQ3LjQ1IiB5MT0iNDQ3LjY2IiB4Mj0iNDY3LjY0IiB5Mj0iNTU3Ljg3Ii8+PGxpbmUgY2xhc3M9ImNscy0yIiB4MT0iNTYwLjQ4IiB5MT0iNzguNyIgeDI9IjYxMS40NyIgeTI9IjEwMC4yMyIvPjxsaW5lIGNsYXNzPSJjbHMtMiIgeDE9IjYzMy40NiIgeTE9IjEwNi40NSIgeDI9Ijc4Mi4wMiIgeTI9IjE2MS4zOSIvPjxsaW5lIGNsYXNzPSJjbHMtMiIgeDE9IjEwNy44NCIgeTE9IjIyOC40IiB4Mj0iMzczLjYxIiB5Mj0iNzIuMjciLz48bGluZSBjbGFzcz0iY2xzLTIiIHgxPSIxMDguODQiIHkxPSIyMzIuMzUiIHgyPSIyNTguNDgiIHkyPSIyMDQuMzgiLz48bGluZSBjbGFzcz0iY2xzLTIiIHgxPSI3OTcuNzMiIHkxPSIxNjUuMjciIHgyPSI5NjAuNjkiIHkyPSIyMjUuMDUiLz48L2c+PGcgaWQ9ImJpZ2NpcmNsZXMiPjxjaXJjbGUgY2xhc3M9ImNscy0zIiBjeD0iMzE0LjczIiBjeT0iNDg0Ljg0IiByPSIxNy42MyIvPjxjaXJjbGUgY2xhc3M9ImNscy0zIiBjeD0iMTI1LjU0IiBjeT0iMzMxLjgxIiByPSIyNS4wNiIvPjxjaXJjbGUgY2xhc3M9ImNscy0zIiBjeD0iNjcwLjkxIiBjeT0iMjgzLjEzIiByPSIzMC4yNSIvPjxjaXJjbGUgY2xhc3M9ImNscy0zIiBjeD0iNTg3LjExIiBjeT0iNDg0Ljg0IiByPSIyNi44NCIvPjxjaXJjbGUgY2xhc3M9ImNscy0zIiBjeD0iNDc0LjM2IiBjeT0iMzQwLjMxIiByPSIxNC41OCIvPjxjaXJjbGUgY2xhc3M9ImNscy0yIiBjeD0iNDAxLjQiIGN5PSI0MDQuNjkiIHI9IjExLjk1Ii8+PGNpcmNsZSBjbGFzcz0iY2xzLTMiIGN4PSI0NDUuODYiIGN5PSI0NDAuNjQiIHI9IjExLjk1Ii8+PGNpcmNsZSBjbGFzcz0iY2xzLTMiIGN4PSI5MDAuMTEiIGN5PSIzMTUuMzEiIHI9IjI3LjYiLz48Y2lyY2xlIGNsYXNzPSJjbHMtNCIgY3g9Ijc4OS44OSIgY3k9IjE2Ni4xNCIgcj0iOS43MiIvPjxjaXJjbGUgY2xhc3M9ImNscy0zIiBjeD0iNTYyLjkiIGN5PSIxOTUuODIiIHI9IjkuNzIiLz48Y2lyY2xlIGNsYXNzPSJjbHMtMyIgY3g9IjU1My4xOCIgY3k9Ijc0LjUyIiByPSI5LjcyIi8+PGNpcmNsZSBjbGFzcz0iY2xzLTMiIGN4PSI2MDcuMDQiIGN5PSI0MDMuNjQiIHI9IjkuNzIiLz48Y2lyY2xlIGNsYXNzPSJjbHMtMyIgY3g9IjEwMzYuNzkiIGN5PSIxODkuODEiIHI9IjIzIi8+PGNpcmNsZSBjbGFzcz0iY2xzLTMiIGN4PSI5ODUuNTUiIGN5PSI1MzUuMTEiIHI9IjM0LjUzIi8+PGNpcmNsZSBjbGFzcz0iY2xzLTMiIGN4PSI3Ni45MSIgY3k9IjQ1MS4wMyIgcj0iMTkuNjEiLz48Y2lyY2xlIGNsYXNzPSJjbHMtMyIgY3g9IjI3Ni4xNiIgY3k9IjIwNC4zOCIgcj0iMTkuNjEiLz48Y2lyY2xlIGNsYXNzPSJjbHMtMyIgY3g9IjM4NC41NiIgY3k9IjY3LjQ0IiByPSIxNi43NCIvPjwvZz48ZyBpZD0idGlueWNpcmNsZXMiPjxjaXJjbGUgY2xhc3M9ImNscy0yIiBjeD0iNjM5LjAyIiBjeT0iMzcxLjQzIiByPSI4LjM1Ii8+PGNpcmNsZSBjbGFzcz0iY2xzLTIiIGN4PSI0NDguMzYiIGN5PSIyMjQiIHI9IjguMzUiLz48Y2lyY2xlIGNsYXNzPSJjbHMtMiIgY3g9IjQzNy44NiIgY3k9IjE1NS4xNCIgcj0iOC4zNSIvPjxjaXJjbGUgY2xhc3M9ImNscy0yIiBjeD0iMTAwLjQ4IiBjeT0iMjMyLjM1IiByPSI4LjM1Ii8+PGNpcmNsZSBjbGFzcz0iY2xzLTIiIGN4PSIxOTEuMjciIGN5PSI1MjAuMDQiIHI9IjguMzUiLz48Y2lyY2xlIGNsYXNzPSJjbHMtMiIgY3g9IjQ2OS44NyIgY3k9IjU2Ny44MiIgcj0iMTAuMiIvPjxjaXJjbGUgY2xhc3M9ImNscy0yIiBjeD0iNjEwLjU5IiBjeT0iNTkxLjY3IiByPSIxMC4yIi8+PGNpcmNsZSBjbGFzcz0iY2xzLTIiIGN4PSIzMzAuODUiIGN5PSI1ODkuNzkiIHI9IjEwLjM1Ii8+PGNpcmNsZSBjbGFzcz0iY2xzLTIiIGN4PSI5MTMuNDQiIGN5PSIzOTQuMjMiIHI9IjExLjQ0Ii8+PGNpcmNsZSBjbGFzcz0iY2xzLTIiIGN4PSIxMTg3Ljg2IiBjeT0iNDQ5LjgxIiByPSIxMS40NCIvPjxjaXJjbGUgY2xhc3M9ImNscy0yIiBjeD0iMTE1MS43OSIgY3k9IjI2Mi42NCIgcj0iMTEuNDQiLz48Y2lyY2xlIGNsYXNzPSJjbHMtMiIgY3g9Ijk3MC45MiIgY3k9IjIzMC4xNCIgcj0iMTEuNDQiLz48Y2lyY2xlIGNsYXNzPSJjbHMtMiIgY3g9Ijg5My4zNiIgY3k9IjExMy4zOSIgcj0iMTEuNDQiLz48Y2lyY2xlIGNsYXNzPSJjbHMtMiIgY3g9IjYyMi4wMiIgY3k9IjEwNC42NCIgcj0iMTEuNDQiLz48Y2lyY2xlIGNsYXNzPSJjbHMtMiIgY3g9IjY3OC4xMyIgY3k9IjcwLjMxIiByPSIxMS40NCIvPjxjaXJjbGUgY2xhc3M9ImNscy0yIiBjeD0iOTY1Ljg2IiBjeT0iMjk4LjMzIiByPSIxMS40NCIvPjxjaXJjbGUgY2xhc3M9ImNscy0yIiBjeD0iMTAzOC4xIiBjeT0iNDM3LjA2IiByPSIxMi43NSIvPjxjaXJjbGUgY2xhc3M9ImNscy0yIiBjeD0iNzc2LjM2IiBjeT0iMzk3LjMxIiByPSI4LjM1Ii8+PGNpcmNsZSBjbGFzcz0iY2xzLTIiIGN4PSI3ODMuNzEiIGN5PSI0MjMiIHI9IjkuNjQiLz48Y2lyY2xlIGNsYXNzPSJjbHMtMiIgY3g9Ijg1My4xNCIgY3k9IjU2Ny4xOCIgcj0iMTAuODQiLz48L2c+PC9zdmc+');
  flex: 0 0 185px;
  background-repeat: no-repeat;
  background-position: 50%;
  background-size: cover;
}
</style>