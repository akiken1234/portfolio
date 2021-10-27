<template>
  <v-layout column justify-center align-center>
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
        :items-per-page="10"
        :search="searchText"
        sort-by="CreatedAt"
        :sort-desc="true"
        class="elevation-1"
      >
        <template v-slot:[`item.downLoad-action`]="{ item }">
          <v-icon @click="onClickDownLoadIcon(item)">mdi-download</v-icon>
        </template>
      </v-data-table>
    </v-card>
  </v-layout>
</template>

<script>
export default {
  auth: false,
  async fetch({ store }) {
    const papers = await store.dispatch('papers/fetchList')
    store.commit('papers/setList', papers)
  },
  data() {
    return {
      searchText: '',
      headers: [
        { text: 'タイトル', value: 'title' },
        { text: '著者', value: 'User.name' },
        { text: '概要', value: 'abstract' },
        { text: '投稿日', value: 'CreatedAt' },
        { text: 'PDF', value: 'downLoad-action' },
      ],
    }
  },
  computed: {
    papers() {
      return this.$store.getters['papers/list']
    },
  },
  methods: {
    onClickDownLoadIcon(item) {
      const alink = document.createElement('a')
      alink.download = item.file_name + '.pdf' // [download] のファイル名
      alink.href = item.file_name + '.pdf' // サーバのファイルのURL
      alink.click() // クリック実行
    },
  },
}
</script>

<style>
.text-start {
  min-width: 80px;
  max-width: 2000px;
}
</style>