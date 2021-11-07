<template>
  <v-layout column justify-center align-center>
    <v-card v-if="papers" width="1050">
      <v-card-title>
        <v-spacer />
        <v-col cols="4">
          <v-text-field
            v-model="searchText"
            append-icon="mdi-magnify"
            label="検索"
            sigle-line
          />
        </v-col>
        <v-spacer />
      </v-card-title>
      <v-data-table
        :headers="headers"
        :items="papers"
        :items-per-page="5"
        :search="searchText"
        sort-by="CreatedAt"
        :sort-desc="true"
        class="elevation-0"
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
  async fetch() {
    await this.$store.dispatch('papers/fetchList')
  },
  data() {
    return {
      searchText: '',
      headers: [
        {
          text: 'タイトル',
          value: 'title',
          width: '24%',
          class: 'text-center',
          sortable: false,
        },
        {
          text: '著者',
          value: 'User.name',
          width: '14%',
          class: 'text-center',
          sortable: false,
        },
        {
          text: '概要',
          value: 'abstract',
          width: '50%',
          class: 'text-center',
          sortable: false,
        },
        {
          text: '投稿日',
          value: 'CreatedAt',
          align: 'center',
          width: '10%',
        },
        {
          text: 'PDF',
          value: 'downLoad-action',
          align: 'center',
          width: '2%',
          sortable: false,
        },
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
        .$post('/papers/download', params, { responseType: 'blob' })
        .then((res) => {
          const alink = document.createElement('a')
          alink.download = item.file_name // [download] のファイル名
          alink.href = URL.createObjectURL(res) // ファイルのURL
          alink.click() // クリック実行
        })
        .catch((e) => {
          console.log(e)
          alert('ダウンロードできませんでした')
        })
    },
  },
}
</script>