<template>
  <div class="container">
    <h1>CTF Challenge Groups</h1>
    <p>Just a collection of CTF Challenges.</p>
    <h2>{{ edit ? "Edit Group" : "Create New Group" }}</h2>
    <form @submit.prevent="createCG">
      <div class="form-group">
        <label>Name</label>
        <input
          class="form-control"
          v-model="form.name"
          placeholder="Name"
          required
        />
      </div>
      <div class="form-group">
        <label>Description</label>
        <input
          class="form-control"
          v-model="form.description"
          placeholder="Description"
          required
        />
      </div>
      <ChallengePicker v-model="form.challenge_ids"></ChallengePicker>
      <button class="btn btn-primary mt-2" type="submit">
        {{ edit ? "Save" : "Create Group" }}
      </button>
      <button
        v-if="edit"
        class="btn btn-danger mt-2 ms-2"
        type="button"
        @click="(edit = false), clearForm()"
      >
        Cancel edit
      </button>
    </form>
    <p v-if="error" class="text-danger mt-2">{{ error }}</p>

    <table class="table mt-4">
      <thead>
        <tr>
          <th>Name</th>
          <th>Description</th>
          <th></th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="group in groups" :key="group.id">
          <td class="align-middle">{{ group.name }}</td>
          <td class="align-middle">{{ group.description }}</td>
          <td class="align-middle text-end">
            <button class="btn btn-info me-2" @click="editGroup(group.id)">
              Edit
            </button>
            <button class="btn btn-danger" @click="deleteGroup(group.id)">
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

const challStore = useChallengeStore();
const http = useHttp();
const groups = ref<any[]>([]);
const form = ref({
  name: "",
  description: "",
  challenge_ids: [] as string[],
});
const error = ref("");
const edit = ref(false);
const editCTFId = ref("");

const challenges = ref<any[]>([]);

onMounted(async () => {
  await challStore.getChallenges();
  challenges.value = challStore.challenges;
  groups.value = await http("/admin/challenge_groups");
});

async function createCG() {
  error.value = "";
  try {
    const ctf = await http(
      edit.value
        ? "/admin/challenge_groups/" + editCTFId.value
        : "/admin/challenge_groups",
      {
        method: edit.value ? "PUT" : "POST",
        body: form.value,
      }
    );
    if (edit.value) {
      groups.value = groups.value.map((c) => (c.id === ctf.id ? ctf : c));
      edit.value = false;
      editCTFId.value = "";
    } else {
      groups.value.push(ctf);
    }
    clearForm();
  } catch (e: any) {
    error.value = "Failed to create/edit Group: " + e.response._data.message;
  }
}

function clearForm() {
  form.value = {
    name: "",
    description: "",
    challenge_ids: [],
  };
}

async function deleteGroup(id: string) {
  const confirmDelete = confirm(
    "Are you sure you want to delete this Challenge Group?"
  );
  if (!confirmDelete) return;
  try {
    await http(`/admin/challenge_groups/${id}`, { method: "DELETE" });
    groups.value = groups.value.filter((c) => c.id !== id);
  } catch (e: any) {
    error.value = "Failed to delete Group: " + e.response._data.message;
  }
}

function editGroup(id: string) {
  edit.value = true;
  editCTFId.value = id;
  const ctf = groups.value.find((c) => c.id === id);
  if (ctf) {
    form.value = {
      name: ctf.name,
      description: ctf.description,
      challenge_ids: JSON.parse(JSON.stringify(ctf.challenge_ids)),
    };
  }
}
</script>
