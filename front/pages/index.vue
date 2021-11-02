<template>
  <v-layout column justify-center align-center>
    <v-card v-if="papers">
      <v-card-title>
        <v-spacer />
        <v-text-field
          v-model="searchText"
          append-icon="mdi-magnify"
          label="検索"
          sigle-line
        />
        <v-spacer />
      </v-card-title>
      <v-data-table
        :headers="headers"
        :items="papers"
        :items-per-page="5"
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
      const params = new URLSearchParams()
      params.append('file_name', item.file_name)
      this.$axios
        .$post('/papers/get', params, { responseType: 'blob' })
        .then((res) => {
          const alink = document.createElement('a')
          alink.download = item.file_name // [download] のファイル名
          alink.href = URL.createObjectURL(res) // ファイルのURL
          alink.click() // クリック実行
        })
        .catch((e) => {
          console.log(e)
        })
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