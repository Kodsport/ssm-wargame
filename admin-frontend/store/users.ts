import { ActionTree, GetterTree, MutationTree } from "vuex";
import { RootState } from ".";

export const state = (): { users: any[] } => ({
  users: [],
});

export type UserState = ReturnType<typeof state>;

export const mutations: MutationTree<UserState> = {
  SET_USERS(state, users) {
    state.users = users;
  },
};

export const actions: ActionTree<UserState, RootState> = {
  async fetchUsers({ commit }) {
    const { data } = await this.$axios.get("/admin/users");
    commit("SET_USERS", data);
  },
};

export const getters: GetterTree<UserState, RootState> = {
  getUserById: (state: UserState) => (id: string) =>
    state.users.filter((u) => u.id === id)[0],
};
