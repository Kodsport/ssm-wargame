<template>
  <div class="container">
    <h1>Users</h1>

    <div>
      <table class="table table-striped">
        <thead>
          <tr>
            <th>Email</th>
            <th>Name</th>
            <th>Role</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="user in users.users" :key="user.id">
            <td>{{ user.email }}</td>
            <td>{{ user.first_name + " " + user.last_name }}</td>
            <td>{{ user.role }}</td>
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
  name: "UsersPage",
  computed: {
    ...mapState(["users"]),
  },
  mounted() {
    this.$store.dispatch("users/fetchUsers");
  },
});
</script>
