<template>
  <div v-if="chall" class="container">
    <h1>{{ chall.title }}</h1>
    <div>
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
          <p class="btn btn-primary" @click="updateChall">Update</p>
        </div>
      </form>


      <div v-if="chall.files.length">
        <h5 class="border-top">Files</h5>
        <table class="table">
          <thead>
            <tr>
              <th>Filename</th>
              <th>Bucket location</th>
              <th>Size</th>
              <th>MD5</th>
              <th></th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="file in chall.files" :key="file.id">
              <td>{{ file.filename }}</td>
              <td>{{ file.bucket + "/" + file.key }}</td>
              <td>{{ fileSize(file.size) }}</td>
              <td>{{ file.md5 }}</td>
              <td>
                <a class="btn btn-primary" :href="file.url">Download</a>
                <button class="btn btn-danger" @click="deleteFile(file.id)">
                  Delete
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <h1>Upload file</h1>
      <div class="form-group">
        <input
          ref="file"
          class="form-control-file"
          type="file"
          @change="checkFile"
        />
      </div>
      <p
        class="btn btn-primary"
        v-bind:class="{
          disabled: !hasFile,
        }"
        @click="uploadFile"
      >
        Upload
      </p>
    </div>
  </div>
</template>



<script lang="ts">
import md5 from "js-md5";
import Vue from "vue";
import { mapGetters, mapState } from "vuex";

export default Vue.extend({
  name: "ChallengeViewPage",
  data() {
    return {
      hasFile: false,
      form: {},
    };
  },
  computed: {
    ...mapState(["challenges"]),
    ...mapGetters({
      getChallengeBySlug: "challenges/getChallengeBySlug",
    }),
    chall(): any {
      return this.getChallengeBySlug(this.$route.params.slug);
    },
  },
  async mounted() {
    if (!this.chall) {
      await this.$store.dispatch("challenges/fetchChallenges");
    }

    this.form = {
      title: this.chall.title,
      description: this.chall.description,
      score: this.chall.score,
      slug: this.chall.slug,
    };
  },
  methods: {
    async uploadFile() {
      // @ts-ignore
      const file = this.$refs.file.files[0] as File;

      const hash = md5.base64(await file.arrayBuffer());
      const res = await this.$axios.post(
        "/admin/challenges/" + this.chall.id + "/file_url",
        {
          md5: hash,
          filename: file.name,
          size: file.size,
        }
      );

      await fetch(res.data.url, {
        body: await file.arrayBuffer(),
        method: "PUT",
        headers: {
          "content-md5": hash,
        },
      });

      this.$store.dispatch("challenges/fetchChallenges");
    },
    checkFile() {
      // @ts-ignore
      this.hasFile = this.$refs.file?.files.length !== 0;
    },
    updateChall() {},
    fileSize(size: number): String {
      if (size === 0) {
        return "0 Bytes";
      }

      const sizes = ["Bytes", "KB", "MB", "GB", "TB"];

      const i = Math.floor(Math.log(size) / Math.log(1000));

      return parseFloat((size / Math.pow(1000, i)).toFixed(2)) + " " + sizes[i];
    },
    async deleteFile(fileId: string) {
      await this.$axios.delete(`/admin/files/${fileId}`);
      this.$store.dispatch("challenges/fetchChallenges");
    },
  },
});
</script>
