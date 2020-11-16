<template>
  <v-app>
    <v-main>
      <router-view name="header" />
      <fade-transition>
        <router-view/>
      </fade-transition>
      <router-view name="footer" />
    </v-main>
  </v-app>
</template>

<script>
import { FadeTransition } from 'vue2-transitions'

export default {
  name: 'App',

  components: {
    FadeTransition
  },

  data: () => ({
    //
  }),
};
</script>

<style>
  .v-main {
    font-family: 'VT323', monospace;
  }
</style>
