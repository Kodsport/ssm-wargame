<template>
    <div class="modal-backdrop fade show"></div>
    <div class="modal modal-xl d-block" tabindex="-1">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header pt-3">
                    <h2 class="modal-title text-primary fw-bold">Välkommen {{ auth.user.full_name }}!</h2>
                    <button type="button" class="btn-close" aria-label="Close" @click="finish"></button>
                </div>
                <div class="modal-body">
                    <div>
                        <p>
                            Vi hoppas att du lär dig något här, vi välkomnar både nybörjare och proffs! Be gärna
                            om hjälp i <a href="">Kodsports Discordserver</a> om fastnar.
                        </p>
                        <p>
                            Vill du representera din skola på poängtavlan?
                        </p>
                        <div v-if="!grund && !uni">
                            <button class="btn btn-primary me-2 mt-1" @click="grund = true">Jag studerar i
                                grundskola/gymnasie</button>
                            <button class="btn btn-primary me-2 mt-1" @click="uni = true">Jag studerar på
                                universitet/högskola</button>
                            <button class="btn btn-primary mt-1" @click="finish">Jag studerar inte</button>
                        </div>

                        <div v-if="grund || uni">

                            <input class="form-control" type="text" v-model="schoolQuery" @input="searchSchools"
                                placeholder="Sök skola">
                            <ul>
                                <li v-for="school in schools" :key="school.id">
                                    {{ school.name }} <span v-if="grund">({{ school.municipality_name }})</span>
                                    <button class="btn btn-primary btn-xs mt-1"
                                        @click="joinSchool(school.id)">Representera</button>
                                </li>
                            </ul>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { useAuthStore } from '../store/auth';


const auth = useAuthStore()
const http = useHttp()

const grund = ref(false)
const uni = ref(false)

var schoolQuery = ref("")
var schools = ref([])

async function joinSchool(id) {
    await http('/user/join_school', {
        method: 'POST',
        body: {
            school_id: id
        }
    })
    await finish()
}

async function searchSchools() {
    if (schoolQuery.value == "") {
        schools.value = []
        return
    }
    const resp = await http(`/user/schools?university=${uni.value}&q=${encodeURIComponent(schoolQuery.value)}`)

    schools.value = resp
}


async function finish() {
    await http('/user/complete_onboarding', { method: 'POST' })
    await auth.getUser()
}

</script>
