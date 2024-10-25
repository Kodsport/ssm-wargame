import { defineStore } from "pinia";
import useHttp from "../composables/use-http";

// Unsure about the correct types here. Please replace string with the correct types of you find that a string type is not enough
interface User {
    onboarding_fone: boolean,
    id: string,
    email: string,
    full_name: string,
    role: string
}

export const useAuthStore = defineStore('auth', {
    state: () => ({
        token: "",
        user: <User>{},
        knackKodenPassword: "",
        knackKodenData: null
    }),
    actions: {
        async setToken(token: string) {
            this.token = token
            await this.getUser()
        },
        async getUser() {
            try {
                const http = useHttp()

                const user = await http('/user/self')
                this.user = user
            } catch {
                this.user = {} as User;
                this.token = ''
            }
        },
        async getKnackKodenData(pw: string | null) {
            const http = useHttp()

            const resp = await http("/knack_koden_get_class", {
                method: "POST",
                body: {
                    password: pw || this.knackKodenPassword
                }
            })
            this.knackKodenData = resp;
            if (pw) {
                this.knackKodenPassword = pw
            }

        }
    },
    async getUser() {
      try {
        const http = useHttp();

        const user = await http("/user/self");
        this.user = user;
      } catch {
        this.user = {};
        this.token = "";
      }
    },
    async getKnackKodenData(pw: string | null) {
      const http = useHttp();

      const resp = await http("/knack_koden_get_class", {
        method: "POST",
        body: {
          password: pw || this.knackKodenPassword,
        },
      });
      this.knackKodenData = resp;
      if (pw) {
        this.knackKodenPassword = pw;
      }
    },
  },
  getters: {
    isAuthed(state) {
      return !!state.user.id;
    },
  },
});
