<template>
    <div class="container">
        <table class="table">
            <thead>
                <th>Name</th>
                <th>Email</th>
                <th>Role</th>
            </thead>
            <tbody>
                <tr v-for="user in users.users" :key="user.id">
                    <td>{{ user.first_name }}</td>
                    <td>{{ user.email }}</td>
                    <td>{{ user.role }}</td>
                </tr>
            </tbody>
        </table>
    </div>
</template>
<script setup lang="ts">

import { onMounted } from 'vue';
import { useUserStore } from '../../../store/admin/users'
const users = useUserStore()


onMounted(() => {
    users.getUsers()
})





</script>
