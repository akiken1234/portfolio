<template>
  <div class="mt-3">
    <v-card class="mt-5 mx-auto" max-width="600">
      <v-container>
        <v-row justify="center">
          <p cols="12" class="mt-3 display-1 grey--text">ログイン</p>
        </v-row>
        <v-row justify="center">
          <v-col cols="12" md="10" sm="10">
            <v-text-field v-model="email" label="Eメールアドレス" />
            <p class="caption mb-0" />
          </v-col>
        </v-row>
        <v-row justify="center">
          <v-col cols="12" md="10" sm="10">
            <v-text-field
              v-model="password"
              type="password"
              label="パスワード"
            />
          </v-col>
        </v-row>
        <v-row justify="center">
          <v-col cols="12" md="5" sm="5">
            <v-btn block class="mr-4 blue white--text" @click="login">
              ログイン
            </v-btn>
          </v-col>
        </v-row>
      </v-container>
    </v-card>
  </div>
</template>

<script>
export default {
  auth: false,
  data() {
    return {
      email: '',
      password: '',
    }
  },
  methods: {
    login() {
      this.$auth
        .loginWith('local', {
          data: {
            email: this.email,
            password: this.password,
          },
        })
        .then(
          (response) => {
            this.$router.push('/')
            alert('ログインしました！')
            return response
          },
          (error) => {
            return error
          }
        )
    },
  },
}
</script>