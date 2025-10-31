<template>
  <div class="container">
    <h1>CTFs</h1>
    <p>
      If you want to edit challenge groups, go
      <nuxt-link to="/admin/ctfs/challenge_groups">here</nuxt-link>.
    </p>
    <h2>{{ edit ? "Edit CTF" : "Create New CTF" }}</h2>
    <form @submit.prevent="createCTF">
      <div class="d-flex">
        <div class="form-group col-6 pe-2">
          <label>Name</label>
          <input
            class="form-control"
            v-model="form.name"
            placeholder="Name"
            required
          />
        </div>
        <div class="form-group col-6">
          <label>Slug</label>
          <input
            class="form-control"
            v-model="form.slug"
            placeholder="Slug"
            required
          />
        </div>
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
      <div class="d-flex">
        <div class="form-group col-6 pe-2">
          <label>Start time</label>
          <input
            class="form-control"
            type="datetime-local"
            placeholder="Enter start date"
            v-model="form.start_time"
          />
        </div>
        <div class="form-group col-6">
          <label>End time</label>
          <input
            class="form-control"
            type="datetime-local"
            placeholder="Enter end date"
            v-model="form.end_time"
          />
        </div>
      </div>
      <div class="form-check mt-2">
        <input
          class="form-check-input"
          type="checkbox"
          id="team_based"
          v-model="form.team_based"
        />
        <label class="form-check-label" for="team_based"> Team based </label>
      </div>
      <div class="form-check mt-2">
        <input
          class="form-check-input"
          type="checkbox"
          id="freeze_enabled"
          v-model="freezeEnabled"
        />
        <label class="form-check-label" for="freeze_enabled">
          Scoreboard Freeze
        </label>
      </div>
      <div v-if="freezeEnabled" class="d-flex mt-2">
        <div class="form-group col-6 pe-2">
          <label>Freeze start time</label>
          <input
            class="form-control"
            type="datetime-local"
            placeholder="Enter freeze start date"
            v-model="form.scoreboard_freeze_start"
          />
        </div>
        <div class="form-group col-6">
          <label>Freeze end time</label>
          <input
            class="form-control"
            type="datetime-local"
            placeholder="Enter freeze end date"
            v-model="form.scoreboard_freeze_end"
          />
        </div>
      </div>
      <div class="form-group mt-2">
        <label>Theme</label>
        <select class="form-control" v-model="form.theme">
          <option v-for="theme in themes" :key="theme.filename" :value="theme.filename">
            {{ theme.name }}
          </option>
        </select>
      </div>
      <ChallengeGroupPicker v-model="form.challenges"></ChallengeGroupPicker>
      <ChallengePicker v-model="form.challenges"></ChallengePicker>
      <button class="btn btn-primary mt-2" type="submit">
        {{ edit ? "Save" : "Create CTF" }}
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
          <th>Slug</th>
          <th>Start</th>
          <th>End</th>
          <th></th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="ctf in ctfs" :key="ctf.id">
          <td class="align-middle">
            <nuxt-link :to="'/ctf/' + ctf.slug" target="_blank">{{
              ctf.name
            }}</nuxt-link>
          </td>
          <td class="align-middle">{{ ctf.slug }}</td>
          <td class="align-middle">
            {{ new Date(ctf.start_time).toLocaleString() }}
          </td>
          <td class="align-middle">
            {{ new Date(ctf.end_time).toLocaleString() }}
          </td>
          <td class="text-end">
            <button class="btn btn-success me-2" @click="viewScoreboard(ctf.slug)">
              Poängtavla
            </button>
            <button class="btn btn-info me-2" @click="viewUsers(ctf.id)">
              Manage users
            </button>
            <button class="btn btn-info me-2" @click="viewTeams(ctf.id)">
              Manage teams
            </button>
            <button class="btn btn-info me-2" @click="editCTF(ctf.id)">
              Edit
            </button>
            <button class="btn btn-danger" @click="deleteCTF(ctf.id)">
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
import { getAvailableThemes } from "@/utils/themes";

const router = useRouter();
const challStore = useChallengeStore();
const http = useHttp();
const ctfs = ref<any[]>([]);
const themes = ref<Array<{name: string, filename: string}>>([]);
const form = ref({
  name: "",
  description: "",
  start_time: "",
  end_time: "",
  slug: "",
  team_based: false,
  theme: "",
  scoreboard_freeze_start: "",
  scoreboard_freeze_end: "",
  challenges: [] as Array<{
    id: string;
    custom_score: number;
    display_order: number;
  }>,
});
const error = ref("");
const edit = ref(false);
const editCTFId = ref("");
const challenges = ref<any[]>([]);
const freezeEnabled = ref(false);

onMounted(async () => {
  challenges.value = challStore.challenges;
  ctfs.value = await http("/admin/ctfs");
  // laddar från frontend
  themes.value = getAvailableThemes();
  fixCTFs();
});

async function createCTF() {
  error.value = "";
  try {
    const body: any = {
      ...form.value,
      start_time: new Date(form.value.start_time).valueOf() / 1000,
      end_time: new Date(form.value.end_time).valueOf() / 1000,
    };
    
    if (freezeEnabled.value && form.value.scoreboard_freeze_start && form.value.scoreboard_freeze_end) {
      body.scoreboard_freeze_start = new Date(form.value.scoreboard_freeze_start).valueOf() / 1000;
      body.scoreboard_freeze_end = new Date(form.value.scoreboard_freeze_end).valueOf() / 1000;
    } else {
      delete body.scoreboard_freeze_start;
      delete body.scoreboard_freeze_end;
    }
    
    const ctf = await http(
      edit.value ? "/admin/ctfs/" + editCTFId.value : "/admin/ctfs",
      {
        method: edit.value ? "PUT" : "POST",
        body,
      }
    );
    if (edit.value) {
      ctfs.value = ctfs.value.map((c) => (c.id === ctf.id ? ctf : c));
      edit.value = false;
      editCTFId.value = "";
    } else {
      ctfs.value.push(ctf);
    }
    fixCTFs();
    clearForm();
  } catch (e: any) {
    error.value = "Failed to create/edit CTF: " + e.response._data.message;
  }
}

function clearForm() {
  form.value = {
    name: "",
    description: "",
    start_time: "",
    end_time: "",
    slug: "",
    team_based: false,
    theme: "",
    scoreboard_freeze_start: "",
    scoreboard_freeze_end: "",
    challenges: [],
  };
  freezeEnabled.value = false;
}

async function deleteCTF(id: string) {
  const confirmDelete = confirm("Are you sure you want to delete this CTF?");
  if (!confirmDelete) return;
  try {
    await http(`/admin/ctfs/${id}`, { method: "DELETE" });
    ctfs.value = ctfs.value.filter((c) => c.id !== id);
  } catch (e: any) {
    error.value = "Failed to delete CTF: " + e.response._data.message;
  }
}

function editCTF(id: string) {
  edit.value = true;
  editCTFId.value = id;
  const ctf = ctfs.value.find((c) => c.id === id);
  if (ctf) {
    form.value = {
      name: ctf.name,
      description: ctf.description,
      start_time: toLocalInput(ctf.start_time),
      end_time: toLocalInput(ctf.end_time),
      slug: ctf.slug,
      team_based: ctf.team_based,
      theme: ctf.theme || "",
      scoreboard_freeze_start: ctf.scoreboard_freeze_start ? toLocalInput(ctf.scoreboard_freeze_start) : "",
      scoreboard_freeze_end: ctf.scoreboard_freeze_end ? toLocalInput(ctf.scoreboard_freeze_end) : "",
      challenges: JSON.parse(JSON.stringify(ctf.challenges)),
    };
    freezeEnabled.value = !!(ctf.scoreboard_freeze_start && ctf.scoreboard_freeze_end);
  }
}

function toLocalInput(isoString: string): string {
  let d = new Date(isoString);
  const pad = (n: number) => String(n).padStart(2, "0");
  return `${d.getFullYear()}-${pad(d.getMonth() + 1)}-${pad(d.getDate())}T${pad(
    d.getHours()
  )}:${pad(d.getMinutes())}`;
}

function viewScoreboard(slug: string) {
  router.push(`/admin/ctfs/${slug}/scoreboard`);
}
function viewUsers(id: string) {
  router.push(`/admin/ctfs/${id}/users`);
}
function viewTeams(id: string) {
  router.push(`/admin/ctfs/${id}/teams`);
}
function fixCTFs() {
  ctfs.value = ctfs.value.map((ctf) => {
    ctf.challenges = ctf.challenges.map((chall: any) => {
      if (chall.custom_score === undefined) {
        chall.custom_score = null;
      }
      return chall;
    });
    return ctf;
  });
}
</script>
