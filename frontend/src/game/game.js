import Phaser from 'phaser'
import BootScene from './scenes/BootScene'
import PlayScene from './scenes/PlayScene'
import MenuScene from './scenes/MenuScene'


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
    scene: [BootScene, PlayScene, MenuScene]
  })
}

export default launch
export { launch }
