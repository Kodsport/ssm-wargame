<template>
    <div v-if="chall" class="container">
        <h1>{{ chall.title }}</h1>
        <div>
            <form @submit.prevent>
                <div class="form-group">
                    <label>Title</label>
                    <input class="form-control" placeholder="Enter title" v-model="form.title" />
                </div>
                <div class="form-group">
                    <label>Slug (shown in url)</label>
                    <input class="form-control" type="text" placeholder="Enter slug" v-model="form.slug" />
                </div>
                <div class="form-group">
                    <label>Score</label>
                    <input class="form-control" type="number" placeholder="Enter score" v-model.number="form.score"
                        step="50" min="0" />
                </div>
                <div class="form-group">
                    <label>Description</label>
                    <textarea class="form-control" rows="7" placeholder="Enter description" v-model="form.description" />
                </div>
                <div class="form-group">
                    <label>Category</label>
                    <select class="form-control" name="" v-model="form.categoryId">
                        <option v-for="cat in challs.categories" :value="cat.id">{{ cat.name }}</option>
                    </select>
                </div>
                <div class="form-group pt-2">
                    <label>Publish immediately</label>
                    <input class="" type="checkbox" v-model="publishImm">
                </div>
                <div class="form-group" v-if="!publishImm">
                    <label>Publish at</label>
                    <input class="form-control" type="datetime-local" placeholder="Enter publish date"
                        v-model="form.publishAt" />
                </div>
                <div class="form-group pt-2">
                    <label>Authors</label>
                    <ul>
                        <li v-for="(author, i) in form.authors">{{ users.byId(author).full_name }} ({{
                            users.byId(author).email }}) <span @click="form.authors = form.authors.splice(i, i)">[x]</span>
                        </li>
                    </ul>
                    <select ref="authSelect" class="form-control" @change="pushAuthor">
                        <option value="" disabled selected>Choose Authors</option>
                        <option v-for="user in users.users.filter(u => form.authors.indexOf(u.id) == -1)" :value="user.id">
                            {{ user.email }}</option>
                    </select>
                </div>
                <div class="form-group pt-2">
                    <button class="btn btn-primary" @click="updateChall">Update</button>
                </div>
            </form>

            <div class="pt-3">
                <h5 class="border-top">Flags</h5>

                <table v-if="chall.flags?.length" class="table">
                    <thead>
                        <tr>
                            <th>Flag</th>
                            <th>Solves</th>
                            <th></th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="flag in chall.flags" :key="flag.id">
                            <td>{{ flag.flag }}</td>
                            <td>TODO</td>
                            <td>
                                <button class="btn btn-danger" @click="deleteFlag(flag.id)">
                                    Delete
                                </button>
                            </td>
                        </tr>
                    </tbody>
                </table>

                <div class="pb-4">
                    <div class="row">
                        <div class="col">
                            <input class="form-control" placeholder="SSM{..." type="text" v-model="newFlag">

                        </div>
                        <div class="col">
                            <button class="btn btn-primary" @click="addFlag">
                                Add
                            </button>
                        </div>

                    </div>
                </div>

            </div>

            <div v-if="chall.files?.length">
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
                <input ref="fileInput" class="form-control-file" type="file" @change="checkFile" />
            </div>
            <button class="btn btn-primary" v-bind:class="{
                disabled: !hasFile,
            }" @click="uploadFile">
                Upload
            </button>
        </div>
    </div>
</template>

<script setup lang="ts">
import md5 from "js-md5";
import { useChallengeStore } from "../../../../store/admin/challenges";
import { useUserStore } from "../../../../store/admin/users";

const http = useHttp()
const route = useRoute()
const router = useRouter()
const challs = useChallengeStore()
const users = useUserStore()

const authSelect = ref(null)

var hasFile = ref(false);
var form = ref({
    title: "",
    description: "",
    score: "",
    slug: "",
    categoryId: "",
    publishAt: null,
    authors: []
})

const fileInput = ref(null)

var chall = computed(() => challs.getBySlug(route.params.slug))
var newFlag = ref("")
const publishImm = ref(false)

const meta = ref(null)


onMounted(async () => {

    challs.getCategories()
    users.getUsers()
    if (!chall.value) {
        await challs.getChallenges()
    }

    form.value = {
        title: chall.value.title,
        description: chall.value.description,
        score: chall.value.score,
        slug: chall.value.slug,
        categoryId: chall.value.category_id,
        publishAt: chall.value.publish_at == null ? '' : new Date(chall.value.publish_at * 1000).toISOString().slice(0, 16), // hacky af,pls fix
        authors: chall.value.authors || []

    };
    publishImm.value = chall.value.publish_at == null

    meta.value = await http(`/admin/challenges/${chall.value.id}`)

})

async function uploadFile() {
    // @ts-ignore
    const file = fileInput.value.files[0] as File;

    const hash = md5.base64(await file.arrayBuffer());
    const res = await http(`/admin/challenges/${chall.value.id}/file_url`,
        {
            method: 'POST',
            body: {
                md5: hash,
                filename: file.name,
                size: file.size,
            }
        }
    );

    await fetch(res.url, {
        body: await file.arrayBuffer(),
        method: "PUT",
        headers: {
            "content-md5": hash,
        },
    });

    challs.getChallenges()
}

function checkFile() {
    // @ts-ignore
    hasFile.value = fileInput.value.file?.files.length !== 0;
}

async function updateChall() {
    await http(`/admin/challenges/${chall.value.id}`, {
        method: 'PUT',
        body: {
            title: form.value.title,
            description: form.value.description,
            score: form.value.score,
            slug: form.value.slug,
            category_id: form.value.categoryId,
            publish_at: publishImm.value ? null : new Date(form.value.publishAt).valueOf() / 1000,
            authors: form.value.authors
        }
    });

    challs.getChallenges()
    router.replace(`/admin/challenges/${form.value.slug}/edit`)
}

function fileSize(size: number): String {
    if (size === 0) {
        return "0 Bytes";
    }

    const sizes = ["Bytes", "KB", "MB", "GB", "TB"];

    const i = Math.floor(Math.log(size) / Math.log(1000));

    return parseFloat((size / Math.pow(1000, i)).toFixed(2)) + " " + sizes[i];
}

async function deleteFile(fileId: string) {
    await http(`/admin/files/${fileId}`, {
        method: 'DELETE'
    });
    challs.getChallenges()
}

async function addFlag() {
    await http(`/admin/challenges/${chall.value.id}/flags`, {
        method: 'POST',
        body: {
            flag: newFlag.value
        }
    })
    challs.getChallenges()
    newFlag.value = ""
}

async function deleteFlag(flagId: string) {
    await http(`/admin/challenges/${chall.value.id}/flags/${flagId}`, {
        method: 'DELETE'
    })
    challs.getChallenges()

    if (newFlag.value == "") {
        newFlag.value = chall.value.flags.find((f: any) => f.id == flagId).flag
    }
}


function pushAuthor(event) {
    form.value.authors.push(event.target.value)
    authSelect.value.value = ''
}
</script>
  