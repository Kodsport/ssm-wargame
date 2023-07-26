<template>
  <div class="container">
    <h1>Challs</h1>
    <div class="col-6">
      <form @submit.prevent>
        <div class="form-group">
          <label>Title</label>
          <input
            class="form-control"
            placeholder="Enter title"
            v-model="form.title"
          />
        </div>
        <div class="form-group">
          <label>Slug (shown in url)</label>
          <input
            class="form-control"
            type="text"
            placeholder="Enter slug"
            v-model="form.slug"
          />
        </div>
        <div class="form-group">
          <label>Score</label>
          <input
            class="form-control"
            type="number"
            placeholder="Enter score"
            v-model.number="form.score"
          />
        </div>
        <div class="form-group">
          <label>Description</label>
          <textarea
            class="form-control"
            placeholder="Enter description"
            v-bind="form.description"
          />
        </div>
        <div>
          <p class="btn btn-primary" @click="createChall">Create</p>
        </div>
      </form>
    </div>
    <div>
      <table class="table">
        <thead>
          <tr>
            <th>Title</th>
            <th>Solvers</th>
            <th />
          </tr>
        </thead>
        <tbody>
          <tr v-for="chall in challenges.challenges" :key="chall.id">
            <td>{{ chall.title }}</td>
            <td>{{ chall.solves }}</td>
            <td class="text-right">
              <nuxt-link
                class="btn btn-primary"
                :to="`/challenges/${chall.slug}`"
              >
                Edit
              </nuxt-link>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import { mapState } from "vuex";

export default Vue.extend({
  name: "ChallengePage",
  data() {
    return {
      form: {
        title: "",
        slug: "",
        score: 0,
        description: "",
      },
    };
  },
  mounted() {
    this.$store.dispatch("challenges/fetchChallenges");
  },
  computed: {
    ...mapState(["challenges"]),
  },
  methods: {
    async createChall() {
      await this.$axios.post("/admin/challenges", {
        ...this.form,
      });

      this.$store.dispatch("challenges/fetchChallenges");
    },
  },
});
</script>
