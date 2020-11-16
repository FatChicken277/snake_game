<template>
  <v-app-bar
      app
      color="secondary"
      dark
      class="px-9"
  >
    <div class="d-flex align-center">
      <v-img
          alt="Snake Logo"
          class="shrink mr-2"
          src="https://lh3.googleusercontent.com/proxy/-ys5Oia_X1N78Z3nOWwkv06B-nPwec7_OYwcqVO5Ndm_2aCuP981sBXkjgcELs4w7La1YYTCW57OQ5U6ig1hv_nOko-lszFO"
          transition="scale-transition"
          width="50"
      />
      <h1>Snakee</h1>
    </div>

    <v-spacer></v-spacer>

    <router-link to="/register">
      <v-btn
        color="primary"
        outlined
        class="mr-4"
      >
        Register
      </v-btn>
    </router-link>
    <router-link to="/login">
      <v-btn
        color="primary"
        style="color: #1E1E1E;"
      >
        Login
      </v-btn>
    </router-link>
  </v-app-bar>
</template>

<script>
  export default {
    name: "Header"
  }
</script>

<style>
</style>
