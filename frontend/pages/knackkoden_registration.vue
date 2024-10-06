<template>

    <div>
        <h1>Registrera din klass för Knäck Koden!</h1>

        <p>
            Vid problem, mejla movitz@kodsport.se
        </p>
        <div class="col-4">

            <div class="form-group">
                <label for="class_name">Klassens namn (t.ex. 4C)</label>
                <input class="form-control" type="text" v-model="form.class_name" name="class_name" id="class_name">
            </div>
            <div class="form-group">
                <label for="school_name">Skolans namn</label>
                <input class="form-control" type="text" v-model="form.school_name" name="school_name" id="school_name">
            </div>
            <div class="form-group">
                <label for="postal_code">Skolans postnummer</label>
                <input class="form-control" type="text" v-model="form.postal_code" name="postal_code" id="postal_code">
            </div>
            <div class="form-group">
                <label for="teacher_full_name">Lärarens namn</label>
                <input class="form-control" type="text" v-model="form.teacher_full_name" name="teacher_full_name"
                    id="teacher_full_name">
            </div>
            <div class="form-group">
                <label for="teacher_email">Lärarens mejl</label>
                <input class="form-control" type="text" v-model="form.teacher_email" name="teacher_email"
                    id="teacher_email">
            </div>
            <div class="form-group">
                <label for="teacher_phonenr">Lärarens telefonnr</label>
                <input class="form-control" type="text" v-model="form.teacher_phonenr" name="teacher_phonenr"
                    id="teacher_phonenr">
            </div>


            <button class="btn btn-primary mt-2" @click="register">Registrera klass</button>
        </div>

        <div class="alert alert-success mt-3" v-if="password != ''">
            Klassens lösenord är {{ password }}
            <br>
            Spara det. Lösenordet visas inte igen.
        </div>
    </div>
</template>

<script setup lang="ts">
const http = useHttp()

const form = ref({
    class_name: "",
    school_name: "",
    postal_code: "",
    teacher_full_name: "",
    teacher_email: "",
    teacher_phonenr: "",
})

const password = ref("")

async function register() {
    const resp = await http(`/knack_koden_register_class`, {
        method: 'POST',
        body: form.value,
    });

    password.value = resp.password

}

</script>