<template>
  <v-sheet
    class="leaderboard my-9 px-5 pt-1"
    rounded
    color="secondary"
    elevation="4"
    align="center"
  >
    <h1>Leaderboard</h1>
    <v-expansion-panels>
      <v-expansion-panel>
        <v-expansion-panel-header>
          <v-img
          alt="goldencup1"
          class="shrink mr-4"
          width="25"
          height="25"
          src="../../assets/top1.png"
        />
          FatChicken277
        </v-expansion-panel-header>
        <v-expansion-panel-content>
          Maximum score: 204
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel>
        <v-expansion-panel-header>
          <v-img
          alt="goldencup1"
          class="shrink mr-4"
          width="25"
          height="25"
          src="../../assets/top2.png"
        />
          Trickshot
        </v-expansion-panel-header>
        <v-expansion-panel-content>
          Maximum score: 124
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel>
        <v-expansion-panel-header>
          <v-img
          alt="goldencup1"
          class="shrink mr-4"
          width="25"
          height="25"
          src="../../assets/top3.png"
        />
          Lexxion
        </v-expansion-panel-header>
        <v-expansion-panel-content>
          Maximum score: 102
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel>
        <v-expansion-panel-header>
          A$ASP Rocky
        </v-expansion-panel-header>
        <v-expansion-panel-content>
          Maximum score: 96
        </v-expansion-panel-content>
      </v-expansion-panel>
    </v-expansion-panels>
  </v-sheet>
</template>

<script>
</script>

<style>
  .leaderboard h1 {
    text-shadow: 4px 4px black;
  }

  .player {
    justify-content: space-between;
  }

  .leaderboard h1 {
    margin: 2rem 2rem;
  }

  .leaderboard {
    min-height: 30rem;
  }

  .v-expansion-panel-header, .v-expansion-panel-content {
    background-color: #525252;
  }
</style>
