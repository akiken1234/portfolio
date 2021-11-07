<template>
  <div class="mt-3">
    <v-card class="mt-5 mx-auto" max-width="600">
      <v-container>
        <v-row justify="center">
          <p cols="12" class="mt-3 display-1 grey--text">新規登録</p>
        </v-row>
        <v-form ref="form" v-model="valid" lazy-validation>
          <v-row justify="center">
            <v-col cols="12" md="10" sm="10">
              <v-text-field
                v-model="name"
                label="名前"
                outlined
                :rules="[rules.required]"
              />
            </v-col>
          </v-row>
          <v-row justify="center">
            <v-col cols="12" md="10" sm="10">
              <v-text-field
                v-model="email"
                label="メールアドレス"
                outlined
                :rules="[rules.required, rules.email]"
              />
            </v-col>
          </v-row>
          <v-row justify="center">
            <v-col cols="12" md="10" sm="10">
              <v-text-field
                v-model="password"
                label="パスワード"
                :append-icon="toggle.icon"
                :type="toggle.type"
                outlined
                autocomplete="on"
                @click:append="show = !show"
                :rules="[rules.required, rules.password]"
              />
            </v-col>
          </v-row>
          <v-row justify="center">
            <v-col cols="12" md="5" sm="5">
              <v-btn block class="blue white--text" @click="register">
                新規登録
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
  auth: false,
  data() {
    return {
      name: '',
      email: '',
      password: '',
      valid: true,
      rules: {
        required: (value) => !!value || '',
        email: (value) => /.+@.+\..+/.test(value) || '',
        password: (value) =>
          /^[0-9a-zA-Z\\d]{8,32}$/.test(value) || '半角英数字8文字以上',
      },
      show: false,
    }
  },
  computed: {
    toggle() {
      const icon = this.show ? 'mdi-eye' : 'mdi-eye-off'
      const type = this.show ? 'text' : 'password'
      return { icon, type }
    },
  },
  methods: {
    register() {
      if (this.$refs.form.validate()) {
        this.$axios
          .$post('/users', {
            name: this.name,
            email: this.email,
            password: this.password,
          })
          .then((res) => {
            this.login()
            // バリデーション入力値を初期化
            this.$refs.form.reset()
            alert('登録しました')
            this.$router.push('/')
          })
          .catch((e) => {
            console.log(e)
            alert('登録できませんでした')
          })
      }
    },
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