<template>
    <div>
        hixasd
    </div>
</template>

<script setup lang="ts">
    import { useAuthStore } from '../../store/auth';
    const auth = useAuthStore();
    if (!auth.isAuthed) {
        location.href = "/"
    }
    if (!(auth.user.role == "admin" || auth.user.role == "author")) {
        location.href = "/" 
    }
</script>