<template>
  <v-layout
    column
    justify-center
  >
    <v-card v-if="papers">
      <v-card-title>
        論文一覧
        <v-spacer />
        <v-text-field
          v-model="searchText"
          append-icon="mdi-magnify"
          label="検索"
          sigle-line
        />
      </v-card-title>
      <v-data-table
        :headers="headers"
        :items="papers"
        :items-per-page="5"
        :search="searchText"
        sort-by="id"
        :sort-desc="true"
        class="elevation-1"
      >
      </v-data-table>
    </v-card>
  </v-layout>
</template>

<script>
export default {
  layout: 'index',
  async fetch ({ store }) {
    const papers = await store.dispatch('papers/fetchList')
    store.commit('papers/setList', papers)
  },
  data () {
    return {
      searchText: '',
    }
  },
  computed: {
    papers () {
      return this.$store.getters['papers/list']
    },
    headers () {
     return [
        { text: 'タイトル', value: 'title' },
        { text: '著者名', value: 'User.name' },
        { text: '', value: 'edit-action' },
        { text: '', value: 'delete-action' },
      ]
    }
  }
}
</script>