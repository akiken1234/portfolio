<template>
  <v-row justify="center" align="center">
    <v-col cols="12" sm="8" md="6">
      <v-card class="logo py-4 d-flex justify-center">
        <NuxtLogo />
        <VuetifyLogo />
      </v-card>
      <v-card>
        <v-card-title class="headline">
          Welcome to the Vuetify + Nuxt.js template
        </v-card-title>
        <div>
          <ul>
            <li v-for="(post, index) in posts" :key="index">
              <a :href="'post.url'" target="_blank" rel="noopener noreferrer">{{
                post.title
              }}</a>
            </li>
          </ul>
        </div>
        <v-card-actions>
          <v-spacer />
          <v-btn color="primary" nuxt to="/inspire"> Continue </v-btn>
        </v-card-actions>
      </v-card>
    </v-col>
  </v-row>
</template>

<script>
export default {
  async asyncData(context) {
    const posts = await context.$axios.$get('/papers')
    console.log(posts)
    return { posts }
  },
}
</script>
