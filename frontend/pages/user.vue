<template>
    <div class="container">
        <div v-if="auth.user.id">
            <h1>{{ auth.user.full_name }}</h1>

            <div>
                <h3>Skola</h3>
                <p>
                    Pluggar du? Representera din skola på poängtavlan!
                </p>
                <div v-if="auth.user.school_name">
                    Du representerar <b class="pe-3"> {{ auth.user.school_name }}</b>
                    <button class="btn btn-danger" @click="leaveSchool">Lämna</button>
                </div>
                <div v-else>
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

</script>