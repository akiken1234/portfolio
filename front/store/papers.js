export const state = () => ({
  papers: [],
  user_papers: [],
})

export const getters = {
  list(state) {
    return state.papers
  },
  listUser(state) {
    return state.user_papers
  }
}

export const mutations = {
  setList(state, papers) {
    state.papers = papers
  },
  setListUser(state, papers) {
    state.user_papers = papers
  }
}

export const actions = {
  async fetchList({ commit }) {
    await this.$axios.$get('/papers')
      .then(response => {
        commit('setList', response)
      })
      .catch(err => {
        console.log(err)
      })
  },
  async fetchListUser({ commit }) {
    const params = new URLSearchParams()
    params.append('id', this.$auth.user.ID)
    await this.$axios.$post('/papers/get', params)
      // await this.$axios.$get('/papers/get/', { params: { id: this.$auth.user.ID } })
      .then(response => {
        commit('setListUser', response)
      })
      .catch(err => {
        console.log(err)
      })
  }
}
