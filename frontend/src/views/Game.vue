<template>
  <div class="game">
    <v-container>
      <v-row>
        <v-col cols="8">
          <v-sheet
            class="game-sheet my-9"
            rounded
            color="secondary"
            elevation="4"
            padding="1"
          >
          </v-sheet>
        </v-col>
        <v-spacer />
        <v-col cols="3">
          <Leaderboard />
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>
import Leaderboard from '@/components/Game/Leaderboard.vue'

export default {
   name: 'Home',
   components: {
    Leaderboard
  },
}
</script>

<style>
  .game-sheet {
    min-height: 30rem;
  }
</style>