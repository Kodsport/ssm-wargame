<template>
    <div class="container-fluid">
        <table class="table">
            <thead>
                <th style="width: 30%">Name</th>
                <th style="width: 40%">Email</th>
                <th style="width: 30%">Role</th>
            </thead>
            <tbody>
                <tr v-for="user in users.users" :key="(user as any).id">
                    <td style="width: 30%">
                        <NuxtLink 
                            :to="`/admin/users/${(user as any).id}`" 
                            class="user-link text-decoration-none"
                        >
                            {{ (user as any).full_name }}
                        </NuxtLink>
                    </td>
                    <td style="width: 40%">{{ (user as any).email }}</td>
                    <td style="width: 30%">{{ (user as any).role }}</td>
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