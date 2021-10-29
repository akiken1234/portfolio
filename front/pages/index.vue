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
      // const params = new FormData()
      // params.append('file_name', item.file_name)
      // this.axios('/papers/download', {
      //   data: params,
      //   method: 'POST',
      //   responseType: 'blob',
      // })
      //   .then((response) => {
      //     console.log(response.file_1)
      //     // Create a Blob from the PDF Stream
      //     const file = new Blob([response.file_1], { type: 'application/pdf' })
      //     // Build a URL from the file
      //     const fileURL = URL.createObjectURL(file)
      //     // Open the URL on new Window
      //     window.open(fileURL)
      //   })
      //   .catch((error) => {
      //     console.log('エラー')
      //     console.log(error)
      //   })
      const params = new FormData()
      params.append('file_name', item.file_name)
      this.$axios
        .$post('/papers/download', params, { responseType: 'blob' })
        .then((res) => {
          console.log(res.file_1)
          console.log(res.file_2)
          const blob = new Blob([res.file_1], {
            type: 'application/octet-stream',
          })
          const alink = document.createElement('a')
          alink.download = item.file_name // [download] のファイル名
          alink.href = window.URL.createObjectURL(blob) // ファイルのURL
          alink.click() // クリック実行
        })
        .catch((e) => {
          console.log(e)
        })
    },
    upload_file() {
      const params = new FormData()
      params.append('title', this.title)
      params.append('abstract', this.abstract)
      params.append('file_name', this.file.name)
      params.append('file', this.file)
      this.$axios
        .$post('/papers', params)
        .then((res) => {
          alert('アップロードしました！')
          this.title = ''
          this.abstract = ''
          this.file = null
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