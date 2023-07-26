<template>
  <div class="container">
    <h1>Monthly challenge</h1>

    <table class="table">
      <thead>
        <tr>
          <th>Title</th>
          <th>Solvers</th>
          <th>Interval</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="[monthly, chall] in montlies" :key="monthly.id">
          <td>{{ chall.title }}</td>
          <td>{{ chall.solvers }}</td>
          <td>
            {{ monthly.start_date + " -> " + monthly.end_date }} ({{monthly.display_month}})
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import { mapState } from "vuex";

export default Vue.extend({
  name: "UsersPage",
  computed: {
    ...mapState(["challenges"]),
    montlies() {
      return this.challenges.monthlyChallenges.map((chall: any) => {
        return [
          chall,
          this.challenges.challenges.find(
            (c: any) => (c.id = chall.challenge_id)
          ),
        ];
      });
    },
  },
  mounted() {
    this.$store.dispatch("challenges/fetchChallenges");
    this.$store.dispatch("challenges/fetchMonthlyChallenges");
  },
});
</script>