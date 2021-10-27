<template>
  <v-app>
    <!-- サイドバー部分  -->
    <v-navigation-drawer
      v-model="drawer"
      :mini-variant="miniVariant"
      :clipped="clipped"
      fixed
      app
    >
      <v-list>
        <v-list-item
          v-for="(item, i) in items"
          :key="i"
          :to="item.to"
          router
          exact
        >
          <v-list-item-action>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title v-text="item.title" />
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
    <!-- ヘッダー部分 -->
    <v-app-bar :clipped-left="clipped" fixed app>
      <v-app-bar-nav-icon @click.stop="drawer = !drawer" />
      <v-toolbar-title v-text="title"/>巨人を肩の上に乗せる
      <v-spacer />
      <v-card-actions>
        <v-btn color="white darken-1" class="black--text" @click="loginUser">
          ログイン
        </v-btn>
      </v-card-actions>
      <v-card-actions>
        <v-btn color="white darken-1" class="black--text" @click="registerUser">
          新規登録
        </v-btn>
      </v-card-actions>
    </v-app-bar>
    <!-- コンテンツ部分 -->
    <v-content>
      <v-container fluid pa-10>
        <nuxt />
      </v-container>
    </v-content>
    <!-- フッター部分 -->
    <v-footer :fixed="fixed" app>
      <span>© {{ new Date().getFullYear() }}</span>
    </v-footer>
  </v-app>
</template>

<script>
export default {
  data() {
    return {
      clipped: false,
      drawer: true,
      fixed: false,
      items: [
        {
          icon: 'mdi-apps',
          title: 'HOME',
          to: '/',
        },
        {
          icon: '',
          title: 'プロフィール',
          to: '/profile',
        },
        {
          icon: '',
          title: '論文投稿',
          to: '/upload',
        },
        {
          icon: 'mdi-account',
          title: 'ユーザー一覧',
          to: '/users',
        },
      ],
      miniVariant: false,
      title: 'Zaiya Scholar　',
    }
  },
  methods: {
    loginUser() {
      this.$router.push('/login')
    },
    registerUser() {
      this.$router.push('/register')
    },
  },
}
</script>