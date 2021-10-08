export const state = () => ({
  users: []
})

export const getters = {
  list(state) {
    return state.users
  }
}

export const mutations = {
  setList(state, data) {
    state.users = data
  }
}

export const actions = {
  async fetchList() {
    return await this.$axios.$get('/users')
      .catch(err => {
        console.log(err)
      })
  }
}
