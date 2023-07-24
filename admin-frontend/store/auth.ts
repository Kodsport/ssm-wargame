import { ActionTree, GetterTree, MutationTree } from "vuex";
import { RootState } from ".";

export const state = () => ({
  bearerToken: localStorage.getItem("ssm_token") || "",
});

export type AuthState = ReturnType<typeof state>;

export const mutations: MutationTree<AuthState> = {
  SET_TOKEN(state, token: string) {
    state.bearerToken = token;
  },
};

export const actions: ActionTree<AuthState, RootState> = {
  setToken({ commit }, token) {
    commit("SET_TOKEN", token);
    localStorage.setItem("ssm_token", token);
    this.$axios.setToken(token, "Bearer");
  },
  logout({ commit }) {
    commit("SET_TOKEN", "");
    localStorage.setItem("ssm_token", "");
    this.$axios.setToken(false);
  },
};

export const getters: GetterTree<AuthState, RootState> = {
  jwtClaims(state) {
    if (state.bearerToken !== "") {
      const claims = state.bearerToken.split(".")[1];
      return JSON.parse(atob(claims));
    }
  },
  isAuthenticated: (state) => state.bearerToken !== "",
};

export const namespaced = false;
