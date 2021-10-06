export const state = () => ({
  papers: []
})

export const getters = {
  list (state) {
    return state.papers
  }
}

export const mutations = {
  setList (state, data) {
    state.papers = data
  }
}

export const actions = {
  async fetchList () {
    return await this.$axios.$get('/papers')
      .catch(err => {
        console.log(err)
      })
  }
}
