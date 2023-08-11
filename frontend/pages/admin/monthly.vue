<template>
    <div class="container">
        <h1>Monthly challenge</h1>

        <div>
            <form @submit.prevent>
                <div class="form-group">
                    <label>Challenge</label>
                    <select class="form-control" v-model="form.challenge_id">
                        <option disabled selected value="">Challenge</option>
                        <option v-for="c in chall.challenges" :key="c.id" :value="c.id">{{ c.title }}</option>
                    </select>
                </div>
                <div class="form-group">
                    <label>Display month</label>
                    <input class="form-control" placeholder="April" v-model="form.display_month" />
                </div>
                <div class="form-group">
                    <label>Start date</label>
                    <input class="form-control" type="datetime-local" placeholder="Enter start date"
                        v-model="form.start_date" />
                </div>
                <div class="form-group">
                    <label>End date</label>
                    <input class="form-control" type="datetime-local" placeholder="Enter end date"
                        v-model="form.end_date" />
                </div>
                <div class="pt-3">
                    <button class="btn btn-primary" @click="createMonthly">Create</button>
                </div>
            </form>

        </div>

        <table class="table">
            <thead>
                <tr>
                    <th>Title</th>
                    <th>Solvers</th>
                    <th>Interval</th>
                    <th></th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="[monthly, c] in chall.monthlies.map(m => [m, chall.getById(m.challenge_id)])" :key="monthly.id">
                    <td><nuxt-link :to="`/admin/challenges/${c.slug}`">{{ c.title }}</nuxt-link></td>
                    <td>{{ c.solves }}</td>
                    <td>
                        {{ moment.default(monthly.start_date * 1000).format('ll') + " -> " +
                            moment.default(monthly.end_date * 1000).format('ll') }} ({{
        monthly.display_month }})
                    </td>
                    <td>
                        <button class="btn btn-danger" @click="deleteMonthly(c.id)">Delete</button>
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
</template>
  
<script lang="ts" setup>
import * as moment from 'moment'
import { useChallengeStore } from '../../store/admin/challenges'
const http = useHttp()

const form = ref({
    challenge_id: "",
    start_date: 0,
    end_date: 0,
    display_month: ""
})


const chall = useChallengeStore()

onMounted(() => {
    chall.getMonthlies()
    chall.getChallenges()
})

async function createMonthly() {
    await http('/admin/monthly_challenges', {
        method: 'POST',
        body: {
            ...form.value,
            start_date: new Date(form.value.start_date).valueOf() / 1000,
            end_date: new Date(form.value.end_date).valueOf() / 1000,
        }
    })

    chall.getMonthlies()
}

async function deleteMonthly(id: string) {
    await http('/admin/monthly_challenges/' + id, {
        method: 'DELETE'
    })

    chall.getMonthlies()
}


</script>