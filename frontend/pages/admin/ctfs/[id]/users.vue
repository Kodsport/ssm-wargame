<template>
  <div class="container">
    <h1>Users - {{ ctf.name }}</h1>
    <h6 class="text-white-50">{{ ctf.slug }}</h6>
    <table class="table mt-4">
      <thead>
        <tr>
          <th>Username</th>
          <th>Password (hover)</th>
          <th></th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="user in users" :key="user.id">
          <td>{{ user.username }}</td>
          <td>
            <span class="hover"> {{ user.password }}</span>
          </td>
          <td>
            <button class="btn btn-info me-2" @click="editUsername(user.id)">
              Edit username
            </button>
            <button class="btn btn-danger" @click="deleteUser(user.id)">
              Delete
            </button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useRouter } from "vue-router";
import { useChallengeStore } from "@/store/admin/challenges";
import useHttp from "@/composables/use-http";

const router = useRouter();
const http = useHttp();
const users = ref<any[]>([]);
const ctfId = ref<string>("");
const ctf = ref<any>({
  name: "",
  slug: "",
});
const error = ref("");

onMounted(async () => {
  ctfId.value = router.currentRoute.value.params.id as string;
  await getUsers();
  const ctfs = await http("/admin/ctfs");
  ctf.value = ctfs.find((ctf) => ctf.id === ctfId.value);
});

async function getUsers() {
  users.value = await http(`/admin/ctfs/${ctfId.value}/users`);
}

async function deleteUser(userId: string) {
  const confirmDelete = confirm("Are you sure you want to delete this user?");
  if (!confirmDelete) return;
  await http(`/admin/ctfs/${ctfId.value}/users/${userId}`, {
    method: "DELETE",
  });
  await getUsers();
}

async function editUsername(userId: string) {
  error.value = "";
  const newUsername = prompt("Enter new username:");
  if (!newUsername) return;
  try {
    await http(`/admin/ctfs/${ctfId.value}/users/${userId}`, {
      method: "PATCH",
      body: { username: newUsername },
    });
    await getUsers();
  } catch (e: any) {
    error.value = "Failed to update username: " + e.response._data.message;
  }
}
</script>
<style scoped>
.hover {
  padding: 0 15px;
  color: black;
  background: black;
  border-radius: 5px;
}
.hover:hover {
  color: white;
  background: inherit;
  transition: all 0.2s ease-in-out;
}
</style>
