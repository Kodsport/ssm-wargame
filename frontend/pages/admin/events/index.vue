<template>
    <div class="container">

        <h1>CTF Events</h1>


        <div class="row">
            <div class="col">
                <div>
                    <h3>Create new event</h3>
                    <div class="form-group">
                        <label>Name</label>
                        <input class="form-control" type="text" v-model="form.name">
                    </div>
                    <div class="form-group pt-2">

                        <button class="btn btn-primary" @click="createEvent">
                            Create
                        </button>
                    </div>
                </div>
            </div>
            <div class="col">
                <div>
                    <h3>Create import token</h3>
                    <div class="form-group">
                        <label>Name</label>
                        <input type="text" v-model="tokenForm.name">
                    </div>
                    <div class="form-group">
                        <label>Expires</label>
                        <select class="form-control" v-model="tokenForm.expires_in">
                            <option v-for="period in ['hour', 'week', 'year']" :value="period">{{ period }}</option>
                        </select>
                    </div>
                    <div class="form-group pt-2">

                        <button class="btn btn-primary" @click="createToken">
                            Create
                        </button>
                    </div>
                </div>
                <div v-if="newToken" class="mt-3 alert alert-success">
                    <p>{{ newToken }}</p>
                </div>
            </div>

        </div>

        <table class="table mt-4">
            <thead>
                <th>ID</th>
                <th>Name</th>
                <th>No. Challs</th>
                <th></th>
            </thead>
            <tbody>
                <tr v-for="event in challs.events">
                    <td>{{ event.id }}</td>
                    <td>{{ event.name }}</td>
                    <td>{{ challs.challenges.filter(c => c.ctf_event_id == event.id).length }}</td>
                    <td>
                        <button class="btn btn-danger" @click="deleteEvent(event.id)">Delete</button>
                    </td>
                </tr>
            </tbody>
        </table>


    </div>
</template>

<script lang="ts" setup>
import { useChallengeStore } from '../../../store/admin/challenges'

const challs = useChallengeStore()
const http = useHttp()

const newToken = ref(null)
const form = ref({
    name: '',
    event_id: null,
})

const tokenForm = ref({
    name: '',
    expires_in: 'hour'
})

onMounted(() => {
    challs.getEvents()
    challs.getChallenges()
})

async function createEvent() {
    await http('/admin/events', {
        method: 'POST',
        body: {
            name: form.value.name
        }
    })
    challs.getEvents()
}

async function createToken() {
    const resp = await http('/admin/import_token', {
        method: 'POST',
        body: tokenForm.value
    })

    newToken.value = resp.token
}

async function deleteEvent(eventId: string) {
    await http('/admin/events/' + eventId, {
        method: 'DELETE',
    })
    challs.getEvents()
    challs.getChallenges()
}

</script>