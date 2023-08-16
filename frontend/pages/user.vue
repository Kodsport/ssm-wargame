<template>
    <div class="container">
        <div v-if="auth.user.id">
            <h1>{{ auth.user.full_name }}</h1>

            <div class="row">
                <div class="col">
                    <h3>Namn</h3>
                    <p>Välj hur du syns på poängtavlan</p>
                    <form @submit.prevent="updateUsername">

                        <div class="form-group">
                            <input class="form-control" type="text" v-model="newUsername">
                            <button class="btn btn-primary mt-2">Uppdatera</button>
                        </div>
                    </form>
                </div>
                <div class="col">
                    <h3>Skola</h3>
                    <div v-if="auth.user.school_name">
                        Du representerar <b class="pe-3"> {{ auth.user.school_name }}</b>
                        <button class="btn btn-danger" @click="leaveSchool">Lämna</button>
                    </div>
                    <div v-else>
                        Pluggar du? Representera din skola på poängtavlan!
                        <input class="form-control" type="text" v-model="schoolQuery" @input="searchSchools"
                            placeholder="Sök skola">
                        <ul>
                            <li v-for="school in schools" :key="school.id">{{ school.name }} <span
                                    v-if="!school.is_university">({{ school.municipality_name
                                    }})</span>
                                <button class="btn btn-primary btn-xs mt-1 ms-2"
                                    @click="joinSchool(school.id)">Representera</button>
                            </li>
                        </ul>
                    </div>
                </div>

            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { useChallengeStore } from '../store/challenges'
import { useAuthStore } from '../store/auth'

const http = useHttp()
const challs = useChallengeStore()
const auth = useAuthStore()

onMounted(() => {
    challs.getChallenges()
})

var schoolQuery = ref("")
var schools = ref([])
const newUsername = ref(auth.user.full_name)

async function leaveSchool() {
    await http('/user/leave_school', {
        method: 'POST'
    })
    auth.getUser()
}

async function joinSchool(id) {
    await http('/user/join_school', {
        method: 'POST',
        body: {
            school_id: id
        }
    })
    auth.getUser()
}

async function searchSchools() {
    if (schoolQuery.value == "") {
        schools.value = []
        return
    }
    const resp = await http('/user/schools?q=' + encodeURIComponent(schoolQuery.value))
    schools.value = resp
}

async function updateUsername() {
    await http('/user/self', {
        method: 'POST',
        body: {
            full_name: newUsername.value
        }
    })
    auth.getUser()
}

</script>