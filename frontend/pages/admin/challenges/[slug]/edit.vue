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
                    <label>Dynamic scoring</label>
                    <input class="" type="checkbox" v-model="dynamicScoring">
                </div>
                <div v-if="!dynamicScoring" class="form-group">
                    <label>Static score</label>
                    <input class="form-control" type="number" placeholder="Enter score"
                        v-model.number="form.static_score" step="50" min="0" max="1500" />
                </div>
                <div class="form-group">
                    <label>Description</label>
                    <textarea class="form-control" rows="7" placeholder="Enter description"
                        v-model="form.description" />
                </div>
                <div class="form-group">
                    <label>Category</label>
                    <select class="form-control" name="" v-model="form.categoryId">
                        <option v-for="cat in challs.categories" :value="cat.id">{{ cat.name }}</option>
                    </select>
                </div>
                <div class="form-group pt-2">
                    <label>Hide</label>
                    <input class="" type="checkbox" v-model="form.hide">
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
                    <button class="btn btn-primary" @click="updateChall">Update</button>
                </div>
            </form>

            <div class="pt-3">
                <h5 class="border-top">Flags</h5>

                <table v-if="chall.flags?.length" class="table">
                    <thead>
                        <tr>
                            <th>Flag</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="flag in chall.flags" :key="flag.id">
                            <td>{{ flag.flag }}</td>
                        </tr>
                    </tbody>
                </table>



            </div>

            <div v-if="chall.files?.length">
                <h5 class="border-top">Files</h5>
                <table class="table">
                    <thead>
                        <tr>
                            <th>Filename</th>
                            <th></th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="file in chall.files" :key="file.id">
                            <td>{{ file.filename }}</td>
                            <td>
                                <a class="btn btn-primary" :href="file.url">Download</a>

                            </td>
                        </tr>
                    </tbody>
                </table>
            </div>

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
    static_score: 0,
    slug: "",
    categoryId: "",
    publishAt: null,
    hide: true,
    authors: []
})

const fileInput = ref(null)

var chall = computed(() => challs.getBySlug(route.params.slug))
var newFlag = ref("")
const publishImm = ref(false)
const dynamicScoring = ref(true)

const meta = ref(null)


onMounted(async () => {

    challs.getCategories()
    users.getUsers()
    challs.getAuthors()
    if (!chall.value) {
        await challs.getChallenges()
    }

    form.value = {
        slug: chall.value.slug,
        publishAt: chall.value.publish_at == null ? '' : new Date(chall.value.publish_at * 1000).toISOString().slice(0, 16), // hacky af,pls fix
        hide: chall.value.hide,
        authors: chall.value.authors || [],
    };
    publishImm.value = chall.value.publish_at == null
    dynamicScoring.value = chall.value.static_score == null

    meta.value = await http(`/admin/challenges/${chall.value.id}`)

})

async function updateChall() {
    await http(`/admin/challenges/${chall.value.id}`, {
        method: 'PUT',
        body: {
            slug: form.value.slug,
            publish_at: publishImm.value ? null : new Date(form.value.publishAt).valueOf() / 1000,
            hide: form.value.hide,
        }
    });

    challs.getChallenges()
    router.replace(`/admin/challenges/${form.value.slug}/edit`)
}

</script>