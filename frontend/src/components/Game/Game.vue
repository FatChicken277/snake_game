<template>
  <v-row>
    <v-col class="12">
      <div :id="containerId" v-if="loaded" />
      <div class="loader" v-else>
        <v-progress-linear
          :size="70"
          :width="7"
          color="primary"
          indeterminate
        ></v-progress-linear>
      </div>
    </v-col>
  </v-row>
</template>

<script>
export default {
  name: 'Game',
  data() {
    return {
      loaded: false,
      gameInstance: null,
      containerId: 'game-container'
    }
  },
  async mounted() {
    const game = await import('@/game/game')
    this.loaded = true
    this.$nextTick(() => {
      this.gameInstance = game.launch(this.containerId)
    })
  },
  destroyed() {
    this.gameInstance.destroy(false)
  }
}
</script>

<style>
  .loader {
    display: flex;
    justify-content: center;
    align-items: center;
  }
</style>
