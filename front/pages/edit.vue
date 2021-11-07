<template>
  <div class="mt-3">
    <v-card class="mt-5 mx-auto" max-width="600">
      <v-container>
        <v-row justify="center">
          <p cols="12" class="mt-3 display-1 grey--text">プロフィール編集</p>
        </v-row>
        <v-form class="mt-3" ref="form" v-model="valid" lazy-validation>
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
            <v-col cols="12" md="5" sm="5">
              <v-btn block class="blue white--text" @click="register">
                変更
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
  async asyncData(context) {
    const name = await context.$auth.user.name
    const email = await context.$auth.user.email
    return { name, email }
  },
  data() {
    return {
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
          .$put('/users', {
            id: this.$auth.user.ID,
            name: this.name,
            email: this.email,
          })
          .then((res) => {
            alert('変更しました')
            this.$auth.fetchUser()
            this.$router.push('/profile')
          })
          .catch((e) => {
            console.log(e)
            alert('変更できませんでした')
          })
      }
    },
  },
}
</script>