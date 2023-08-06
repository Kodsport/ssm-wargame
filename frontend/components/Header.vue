<template>
  <nav class="navbar navbar-expand-lg navbar-nav bg-dark">
    <div class="container-fluid">
      <nuxt-link class="navbar-brand nav-link" to="/">SSM</nuxt-link>
      <div class="collapse navbar-collapse" id="navbarSupportedContent">
        <ul class="navbar-nav me-auto mb-2 mb-lg-0">
          <li class="nav-item">
            <nuxt-link active-class="active" class="nav-link" to="/challenges">
              Utmaningar
            </nuxt-link>
          </li>
          <li class="nav-item">
            <nuxt-link active-class="active" class="nav-link" to="/scoreboard">
              Poängtavla
            </nuxt-link>
          </li>
        </ul>
        <ul class="navbar-nav pe-3 mb-2 mb-lg-0">

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
              <nuxt-link v-if="!!auth.user.id" active-class="active" class="nav-link btn border border-primary"
                to="/user">
                <span class="d-flex">
                  {{ auth.user.full_name }}
                  <span class="material-icons text-primary">person</span>
                </span>

              </nuxt-link>
            </li>
          </template>
        </ul>
        <a v-if="!isAuthenticated" class="btn btn-primary" @click="login">Logga in</a>
        <a v-else class="btn btn-primary" @click="logout">Logga ut</a>
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
