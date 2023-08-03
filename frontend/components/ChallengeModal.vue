<template>
    <div class="modal-backdrop fade show"></div>
    <div class="modal modal-xl d-block" tabindex="-1" @keydown.escape="goBack">
        <div class="modal-dialog">
            <div class="modal-content">


                <div class="pt-3 d-flex justify-content-center">
                    <h2 class="modal-title text-primary fw-bold">{{ props.chall.title }}</h2>
                </div>
                <div class="d-flex justify-content-center">
                    <div class="badge bg-secondary">{{ props.chall.category }}</div>
                </div>

                <h5>
                    <div class="d-flex justify-content-center pt-3">
                        <span class="mx-3 d-flex">
                            <span class="me-1">
                                {{ props.chall.score }}
                            </span>
                            <span class="material-icons text-primary">control_point_duplicate</span>
                        </span>
                        <span class="mx-3 d-flex">
                            <span class="me-1">
                                {{ props.chall.solves }}
                            </span>
                            <span class="material-icons text-primary">flag</span>
                        </span>
                    </div>
                </h5>

                <div class="modal-body">
                    <div class="row">
                        <div class="col">
                            {{ props.chall.description }}
                        </div>
                        <div class="col">
                            <div class="d-flex">
                                <span class="material-icons text-primary pe-1">group</span>
                                <span class="material-icons text-primary pe-1">edit</span>
                                <span>Movitz Sunar</span>
                            </div>
                            <div v-for="file in props.chall.files">

                                <span class="material-icons text-primary">description</span>
                                <a :href="file.url">{{ file.filename }}</a>
                            </div>
                            <div class="pt-3">
                                <b>Första lösarna</b>
                                <ol>
                                    <li>Bumbodosan</li>
                                    <li>SNHT</li>
                                    <li>Royalroppers</li>
                                </ol>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">

const props = defineProps(['chall'])

const router = useRouter();

function goBack() {
    router.replace('/challenges')
}


</script>
