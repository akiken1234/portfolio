<template>
  <v-layout column>
    <v-row justify="center" class="my-2">
      <div>名前：{{ this.$auth.user.name }}</div>
    </v-row>
    <v-row justify="center" class="my-2">
      <div>メールアドレス：{{ this.$auth.user.email }}</div>
    </v-row>
    <v-row justify="center" class="my-1">
      <v-btn class="my-2" @click="$router.push('/edit')">編集</v-btn>
    </v-row>
    <v-row justify="center" class="my-1">
      <v-card v-if="papers" width="1050" class="my-2">
        <v-data-table
          :headers="headers"
          :items="papers"
          :items-per-page="5"
          sort-by="CreatedAt"
          :sort-desc="true"
          class="elevation-0"
        >
          <template v-slot:[`item.delete-action`]="{ item }">
            <v-icon @click="onClickDeleteIcon(item)">mdi-delete</v-icon>
          </template>
        </v-data-table>
      </v-card>
    </v-row>
  </v-layout>
</template>

<script>
export default {
  async fetch() {
    await this.$store.dispatch('papers/fetchListUser')
  },
  data() {
    return {
      headers: [
        {
          text: 'タイトル',
          value: 'title',
          width: '28%',
          class: 'text-center',
          sortable: false,
        },
        {
          text: '概要',
          value: 'abstract',
          width: '54%',
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
          text: '削除',
          value: 'delete-action',
          align: 'center',
          width: '8%',
          sortable: false,
        },
      ],
    }
  },
  computed: {
    papers() {
      return this.$store.getters['papers/listUser']
    },
  },
  methods: {
    onClickDeleteIcon(item) {
      // IDが０なので、file_nameを送信
      this.$axios
        .$delete('/papers/', { params: { file_name: item.file_name } })
        .then((res) => {
          this.$fetch()
          alert('削除しました')
        })
        .catch((e) => {
          console.log(e)
          alert('削除できませんでした')
        })
    },
  },
}
</script>