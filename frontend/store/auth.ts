import { defineStore } from "pinia";
import useHttp from "../composables/use-http";

// Unsure about the correct types here. Please replace string with the correct types of you find that a string type is not enough
interface User {
  onboarding_fone: boolean;
  id: string;
  email: string;
  full_name: string;
  role: string;
}

interface CTFUser {
  id: string;
  username: string;
  password: string;
  ctfSlug: string;
}

export const useAuthStore = defineStore("auth", {
  state: () => ({
    token: "",
    user: <User>{},
    knackKodenPassword: "",
    knackKodenData: null,
    ctfUser: <CTFUser>{},
  }),

  actions: {
    async setToken(token: string) {
      this.token = token;
      await this.getUser();
    },
    async getUser() {
      try {
        const http = useHttp();

        const user = await http("/user/self");
        this.user = user;
      } catch {
        this.user = {} as User;
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
    async loginCTFUser(slug: string, password: string) {
      const http = useHttp();
      const user = await http(`/ctfs/${slug}/user`, {
        method: "GET",
        params: {
          password: password,
        },
      });
      this.ctfUser = {
        id: user.id,
        username: user.username,
        password: password,
        ctfSlug: slug,
      };
      localStorage.setItem("ctf_password", password);
    },
    async getCTFUser(slug: string) {
      const password = localStorage.getItem("ctf_password");
      if (!password) return;

      try {
        await this.loginCTFUser(slug, password);
      } catch (e) {
        localStorage.removeItem("ctf_password");
        this.ctfUser = {} as CTFUser;
        return;
      }
    },
    async getCTFUserOnce(slug: string) {
      if (!this.ctfUser.id) {
        await this.getCTFUser(slug);
      }
    },
    async registerCTFUser(slug: string, username: string) {
      const http = useHttp();

      const user = await http(`/ctfs/${slug}/users`, {
        method: "POST",
        body: { username: username },
      });

      const password = user.password;
      if (!password) throw new Error("No password found");
      localStorage.setItem("ctf_password", password);

      this.ctfUser = {
        id: user.id,
        username: user.username,
        password: user.password,
        ctfSlug: slug,
      };

      return this.ctfUser;
    },
    logoutCTFUser() {
      localStorage.removeItem("ctf_password");
      this.ctfUser = {} as CTFUser;
    },
  },

  getters: {
    isAuthed(state) {
      return !!state.user.id;
    },
  },
});
