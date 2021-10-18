<template>
  <v-layout column justify-center>
    <v-card v-if="users">
      <v-card-title>
        ユーザー一覧
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
        :items="users"
        :items-per-page="10"
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
  async fetch({ store }) {
    const users = await store.dispatch('users/fetchList')
    store.commit('users/setList', users)
  },
  data() {
    return {
      searchText: '',
    }
  },
  computed: {
    users() {
      return this.$store.getters['users/list']
    },
    headers() {
      return [
        { text: 'ID', value: 'ID' },
        { text: '名前', value: 'name' },
        { text: 'メールアドレス', value: 'email' },
        { text: '', value: 'edit-action' },
        { text: '', value: 'delete-action' },
      ]
    },
  },
}
</script>