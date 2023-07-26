<template>
  <nav class="navbar navbar-expand-lg navbar-light bg-light">
    <div class="container-fluid">
      <nuxt-link class="navbar-brand" to="/">ssm?</nuxt-link>
      <div class="collapse navbar-collapse" id="navbarSupportedContent">
        <ul class="navbar-nav me-auto mb-2 mb-lg-0">
          <li class="nav-item">
            <nuxt-link active-class="active" class="nav-link" to="/users">
              Users
            </nuxt-link>
          </li>
          <li class="nav-item">
            <nuxt-link active-class="active" class="nav-link" to="/challenges">
              Challenges
            </nuxt-link>
          </li>
          <li class="nav-item">
            <nuxt-link active-class="active" class="nav-link" to="/monthly">
              Monthly
            </nuxt-link>
          </li>
          <li>
            <a v-if="!isAuthenticated" class="btn btn-primary" @click="login">
              Login
            </a>
            <a v-else class="btn btn-primary" @click="logout">Logout</a>
          </li>
        </ul>
      </div>
    </div>
  </nav>
</template>

<script lang="ts">
import Vue from "vue";
import { mapGetters } from "vuex";

export default Vue.extend({
  computed: {
    ...mapGetters(["isAuthenticated"]),
  },
  methods: {
    logout() {
      this.$store.dispatch("logout");
    },
    async login() {
      const x = await this.$axios.get("/auth/discord/url");

      const w = window.open(x.data.url);

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

          const resp = await this.$axios.post("/auth/discord/exchange", {
            code: params.get("code"),
            state: params.get("state"),
          });

          this.$store.dispatch("setToken", resp.data.jwt);
        } catch (error) {
          // DOMException is thrown when w is on other domain
        }
      }, 30);
    },
  },
});
</script>
