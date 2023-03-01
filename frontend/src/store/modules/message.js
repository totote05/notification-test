const api = "http://localhost:5000/message";

const state = {
  messages: [],
  sendErrorMessage: "",
  fetchErrorMessage: "",
  loading: false,
  sending: false,
};

const mutations = {
  setMessage(state, value) {
    state.messages = value;
  },
  setSendErrorMessage(state, value) {
    state.sendErrorMessage = value;
  },
  setFetchErrorMessage(state, value) {
    state.fetchErrorMessage = value;
  },
  setLoading(state, value) {
    state.loading = value;
  },
  setSending(state, value) {
    state.sending = value;
  },
};

const actions = {
  async fetchMessages({ commit }) {
    commit("setLoading", true);
    commit("setFetchErrorMessage", "");
    try {
      const data = await fetch(api).then((result) => result.json());
      commit("setMessage", data);
    } catch (e) {
      commit("setFetchErrorMessage", "The server is not responding");
      console.error(e);
    } finally {
      commit("setLoading", false);
    }
  },
  async sendMessage({ commit, dispatch }, payload) {
    commit("setSending", true);
    try {
      const result = await fetch(api, {
        method: "post",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(payload),
      });

      if (result.status === 201) {
        dispatch("fetchMessages");
      } else if (result.status === 400) {
        commit("setSendErrorMessage", await result.text());
      } else {
        throw new Error();
      }
    } catch (e) {
      commit("setSendErrorMessage", "Unknown error");
      console.error(e);
    } finally {
      commit("setSending", false);
    }
  },
};

const getters = {
  messages: (state) => state.messages,
  sendErrorMessage: (state) => state.sendErrorMessage,
  fetchErrorMessage: (state) => state.fetchErrorMessage,
  loading: (state) => state.loading,
  sending: (state) => state.sending,
};

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters,
};
