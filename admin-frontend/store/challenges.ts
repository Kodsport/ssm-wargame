import { ActionTree, GetterTree, MutationTree } from "vuex";
import { RootState } from ".";

export const state = (): { challenges: any[]; monthlyChallenges: any[] } => ({
  challenges: [],
  monthlyChallenges: [],
});

export type ChallengeState = ReturnType<typeof state>;

export const mutations: MutationTree<ChallengeState> = {
  SET_CHALLENGES(state, challenges) {
    state.challenges = challenges;
  },
  SET_MONTHLY_CHALLENGES(state, monthlyChallenges) {
    state.monthlyChallenges = monthlyChallenges;
  },
};

export const actions: ActionTree<ChallengeState, RootState> = {
  async fetchChallenges({ commit }) {
    const { data } = await this.$axios.get("/admin/challenges");
    commit("SET_CHALLENGES", data);
  },
  async fetchMonthlyChallenges({ commit }) {
    const { data } = await this.$axios.get("/monthly_challenges");
    commit("SET_MONTHLY_CHALLENGES", data);
  },
};

export const getters: GetterTree<ChallengeState, RootState> = {
  getChallengeById: (state: ChallengeState) => (id: string) =>
    state.challenges.find((u) => u.id === id),

  getChallengeBySlug: (state: ChallengeState) => (slug: string) =>
    state.challenges.find((c: any) => c.slug === slug),
};
