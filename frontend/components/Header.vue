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
            <nuxt-link v-if="!!auth.user.id" active-class="active" class="nav-link btn border border-primary" to="/user">
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

<style scoped>
header {
  background-image: url('~/assets/logo.svg');
  flex: 0 0 185px;
  background-repeat: no-repeat;
  background-position: 50%;
  background-size: cover;
}
</style>