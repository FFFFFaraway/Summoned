export default {
  state: {
    signed: false,
  },
  mutations: {
    set(state, value) {
      state.signed = value
    }
  },
  namespaced: true,
};
