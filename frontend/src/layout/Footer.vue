<template>
  <v-footer
      color="secondary"
      dark
      absolute
      height="70"
      class="px-9"
  >
    <v-col cols="10">
      <div class="about">
        Created by FatChicken277 - 2020
      </div>
    </v-col>
    <v-col cols="2">
      <div class="social-media">
        <v-btn
          color="primary"
          fab
          x-small
          dark
          href="https://github.com/FatChicken277"
        >
          <v-icon>mdi-github</v-icon>
        </v-btn>
        <v-btn
          color="primary"
          fab
          x-small
          dark
          href="https://twitter.com/alejorc277"
        >
          <v-icon>mdi-twitter</v-icon>
        </v-btn>
        <v-btn
          color="primary"
          fab
          x-small
          dark
          href="https://www.linkedin.com/in/alejandro-ramirez-ciceros/"
        >
          <v-icon>mdi-linkedin</v-icon>
        </v-btn>
      </div>
    </v-col>
  </v-footer>
</template>

<script>
  export default {
    name: "Footer"
  }
</script>

<style>
  .social-media {
    display: flex;
    padding-left: 3rem;
    justify-content: space-between;
  }
  .social-media .v-btn .v-icon {
    color: #1E1E1E;
  }
</style>
