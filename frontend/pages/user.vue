<template>
    <div>
        <div v-if="auth.user.id">
            <h1>User</h1>

            <div>
                <table>
                    <tbody>
                        <tr v-for="[k, v] in Object.entries(auth.user)" :key="k">
                            <th>{{ k }}</th>
                            <td>{{ v }}</td>
                        </tr>
                    </tbody>
                </table>
            </div>
            <div>
                <div v-if="auth.user.school_id">
                    in school {{ auth.user.school_id }}
                    <button class="btn btn-danger" @click="leaveSchool">Leave</button>
                </div>
                <div v-else>
                    <input class="form-control" type="text" v-model="schoolQuery" @input="searchSchools">
                    <ul>
                        <li v-for="school in schools" :key="school.id">{{ school.name }} ({{ school.municipality_name
                        }})
                            <button class="btn btn-primary btn-xs" @click="joinSchool(school.id)">Join</button>
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