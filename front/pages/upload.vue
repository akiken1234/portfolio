<template>
  <div class="mt-3">
    <v-card class="mt-n5 mx-auto" max-width="600">
      <v-container>
        <v-row justify="center">
          <p cols="12" class="mt-3 display-1 grey--text">論文投稿</p>
        </v-row>
        <v-form ref="form" v-model="valid" lazy-validation>
          <v-row justify="center">
            <v-col cols="12" md="10" sm="10">
              <v-text-field
                v-model="title"
                label="タイトル"
                counter="50"
                outlined
                :rules="[rules.required, rules.titleMin]"
              />
            </v-col>
          </v-row>
          <v-row justify="center">
            <v-col cols="12" md="10" sm="10">
              <v-textarea
                v-model="abstract"
                label="概要"
                counter="200"
                outlined
                :rules="[rules.required, rules.abstractMin]"
              />
            </v-col>
          </v-row>
          <v-row justify="center">
            <v-col cols="12" md="10" sm="10">
              <v-file-input
                v-model="file"
                accept=".pdf"
                label="ファイル名"
                outlined
                :rules="[rules.required]"
              ></v-file-input>
            </v-col>
          </v-row>
          <v-row justify="center">
            <v-col cols="12" md="5" sm="5" class="mt-n5">
              <v-btn block class="mr-4 blue white--text" @click="upload_file">
                アップロード
              </v-btn>
            </v-col>
          </v-row>
        </v-form>
      </v-container>
    </v-card>
  </div>
</template>

<script>
export default {
  data() {
    return {
      title: '',
      abstract: '',
      file: null,
      valid: true,
      rules: {
        required: (value) => !!value || '',
        titleMin: (value) => (value && value.length <= 50) || '',
        abstractMin: (value) => (value && value.length <= 200) || '',
        file: (value) => (value && value.size < 1000000) || '1MB以下',
      },
      show: false,
    }
  },
  methods: {
    upload_file() {
      if (this.$refs.form.validate()) {
        const params = new FormData()
        params.append('title', this.title)
        params.append('abstract', this.abstract)
        params.append('file_name', this.file.name)
        params.append('file', this.file)
        params.append('user_id', this.$auth.user.ID)
        this.$axios
          .$post('/papers/create', params)
          .then((res) => {
            // バリデーション入力値を初期化
            this.$refs.form.reset()
            alert('アップロードしました')
          })
          .catch((e) => {
            console.log(e)
            alert('アップロードできませんでした')
          })
      }
    },
  },
}
</script>