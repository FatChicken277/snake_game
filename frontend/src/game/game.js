import Phaser from 'phaser'
import BootScene from './scenes/BootScene'
import PlayScene from './scenes/PlayScene'


function launch(containerId) {
  return new Phaser.Game({
    type: Phaser.AUTO,
    width: 740,
    height: 440,
    backgroundColor: '#96bb58',
    parent: containerId,
    autoCenter: Phaser.Scale.CENTER_BOTH,
    physics: {
      default: 'arcade',
      arcade: {
        gravity: { y: 0 },
        debug: false
      }
    },
    scene: [BootScene, PlayScene]
  })
}

export default launch
export { launch }
