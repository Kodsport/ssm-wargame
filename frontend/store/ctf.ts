import { defineStore } from "pinia";
import useHttp from "../composables/use-http";
const http = useHttp();

export const useCTFStore = defineStore("ctf", {
  state: () => ({
    ctf: {} as Record<string, string>,
    challenges: [] as Array<Record<string, string>>,
    scoreboard: [] as Array<Record<string, string>>,
    hasFetchedCTF: false,
    hasFetchedChallenges: false,
    hasFetchedScoreboard: false,
  }),
  actions: {
    async getCTF(slug: string) {
      try {
        const ctf = await http("/ctfs/" + slug);
        this.ctf = ctf;
      } catch (e) {
        this.ctf = {};
        console.error(e);
      }
    },
    async getChallenges(slug: string, password?: string, tryAgain = true) {
      try {
        const challs = await http("/ctfs/" + slug + "/challenges", {
          method: "GET",
          params: password
            ? {
                password,
              }
            : {},
        });
        this.challenges = challs;
      } catch (e) {
        console.error(e);
        // If the password stored in local storage is invalid, we need to refetch the challenges
        if (tryAgain)
          setTimeout(() => this.getChallenges(slug, undefined, false), 1000);
      }
    },
    async getScoreboard(slug: string) {
      try {
        const scoreboard = await http("/ctfs/" + slug + "/scoreboard");
        this.scoreboard = scoreboard;
      } catch (e) {
        console.error(e);
      }
    },
    async getAll(slug: string, password?: string, force = false) {
      let promises = [];
      if (!this.hasFetchedCTF || force) {
        this.hasFetchedCTF = true;
        promises.push(this.getCTF(slug));
      }
      if ((!this.hasFetchedChallenges && process.client) || force) {
        this.hasFetchedChallenges = true;
        promises.push(this.getChallenges(slug, password));
      }
      if (!this.hasFetchedScoreboard || force) {
        this.hasFetchedScoreboard = true;
        promises.push(this.getScoreboard(slug));
      }
      try {
        await Promise.all(promises);
      } catch (e) {
        console.error(e);
      }
    },
  },
  getters: {
    getBySlug(state) {
      return (slug: string) =>
        state.challenges.find((c: Record<string, string>) => c.slug == slug);
    },
  },
});
