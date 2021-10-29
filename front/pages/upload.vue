<template>
  <div class="mt-3">
    <v-card class="mt-5 mx-auto" max-width="600">
      <v-form ref="form" v-model="valid" lazy-validation>
        <v-container>
          <v-row justify="center">
            <p cols="12" class="mt-3 display-1 grey--text">アップロード</p>
          </v-row>
          <v-row justify="center">
            <v-col cols="12" md="10" sm="10">
              <v-text-field v-model="title" label="タイトル" counter="50" />
              <p class="caption mb-0" />
            </v-col>
          </v-row>
          <v-row justify="center">
            <v-col cols="12" md="10" sm="10">
              <v-textarea v-model="abstract" label="概要" counter="140" />
            </v-col>
          </v-row>
          <v-row justify="center">
            <v-col cols="12" md="10" sm="10">
              <v-file-input
                v-model="file"
                accept="pdf/*"
                label="ファイル名"
                show-size
              ></v-file-input>
            </v-col>
          </v-row>
          <v-row justify="center">
            <v-col cols="12" md="10" sm="10">
              <v-btn block class="mr-4 blue white--text" @click="upload_file">
                アップロード
              </v-btn>
            </v-col>
          </v-row>
        </v-container>
      </v-form>
    </v-card>
  </div>
</template>

<script>
import FormData from 'form-data'
export default {
  auth: false,
  data() {
    return {
      title: '',
      abstract: '',
      file: null,
    }
  },
  methods: {
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