<template>
    <div v-if="chall" class="container">
        <h1>{{ chall.title }}</h1>
        <div>
            <h1>Solves</h1>
            <table v-if="meta.solvers" class="table">
                <thead>
                    <th>User</th>
                    <th>Time</th>
                </thead>
                <tbody>
                    <tr v-for="solve in meta.solvers">
                        <td v-if="users.users.length">{{ users.byId(solve.user_id).email }}</td>
                        <td v-else>{{ users.id }}</td>
                        <td>{{ new Date(solve.solved_at * 1000) }}</td>
                    </tr>
                </tbody>
            </table>
            <h1>Attempts</h1>
            <table v-if="meta.submissions" class="table">
                <thead>
                    <th>User</th>
                    <th>Input</th>
                    <th>Time</th>
                </thead>
                <tbody>
                    <tr v-for="solve in meta.submissions">
                        <td v-if="users.users.length">{{ users.byId(solve.user_id).email }}</td>
                        <td v-else>{{ users.id }}</td>
                        <td :class="{ 'bg-success': solve.successful, 'bg-danger': !solve.successful }">{{ solve.input }}
                        </td>
                        <td>{{ new Date(solve.submitted_at * 1000) }}</td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
</template>
  
<script setup lang="ts">
import { useChallengeStore } from "../../../../store/admin/challenges";
import { useUserStore } from "../../../../store/admin/users";


const http = useHttp()
const route = useRoute()
const challs = useChallengeStore()
const users = useUserStore()

var chall = computed(() => challs.getBySlug(route.params.slug))

const meta = ref({})


onMounted(async () => {
    users.getUsers()

    if (!chall.value) {
        await challs.getChallenges()
    }


    meta.value = await http(`/admin/challenges/${chall.value.id}`)

})

</script>
    