<template>
  <div class="container">
    <h1>Teams for {{ ctf?.name }}</h1>
    <button class="btn btn-primary mb-3" @click="showCreate = !showCreate">
      {{ showCreate ? "Cancel" : "Create New Team" }}
    </button>
    <form v-if="showCreate" @submit.prevent="createTeam" class="mb-4">
      <div class="form-group mb-2">
        <label>Team Name</label>
        <input class="form-control" v-model="teamForm.teamname" required />
      </div>
      <button class="btn btn-success" type="submit">Create Team</button>
    </form>
    <table class="table mt-4">
      <thead>
        <tr>
          <th>Team Name</th>
          <th>Team Code</th>
          <th></th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="team in teams" :key="team.id">
          <td>{{ team.teamname }}</td>
          <td>{{ team.password }}</td>
          <td>
            <button class="btn btn-danger btn-sm" @click="deleteTeam(team.id)">
              Delete
            </button>
          </td>
        </tr>
      </tbody>
    </table>
    <p v-if="error" class="text-danger mt-2">{{ error }}</p>
  </div>
</template>
<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useRoute } from "vue-router";
import useHttp from "@/composables/use-http";

const route = useRoute();
const http = useHttp();
const ctfId = route.params.id as string;
const teams = ref<any[]>([]);
const ctf = ref<any>(null);
const showCreate = ref(false);
const error = ref("");
const teamForm = ref({ teamname: "", password: "" });

onMounted(async () => {
  await fetchTeams();
  const ctfs = await http("/admin/ctfs");
  ctf.value = ctfs.find((c: any) => c.id === ctfId);
});

async function fetchTeams() {
  teams.value = await http(`/admin/ctfs/${ctfId}/teams`);
}

async function createTeam() {
  error.value = "";
  try {
    await http(`/admin/ctfs/${ctfId}/teams`, {
      method: "POST",
      body: { ...teamForm.value },
    });
    await fetchTeams();
    teamForm.value = { teamname: "", password: "" };
    showCreate.value = false;
  } catch (e: any) {
    error.value =
      "Failed to create team: " + (e.response?._data?.message || e.message);
  }
}

async function deleteTeam(id: string) {
  if (!confirm("Are you sure you want to delete this team?")) return;
  try {
    await http(`/admin/ctfs/${ctfId}/teams/${id}`, { method: "DELETE" });
    await fetchTeams();
  } catch (e: any) {
    error.value =
      "Failed to delete team: " + (e.response?._data?.message || e.message);
  }
}
</script>
