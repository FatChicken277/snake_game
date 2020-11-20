import { Scene, Actions, Math as PhaserMath } from 'phaser'
import snake_head from '@/game/assets/head.svg'
import snake_body from '@/game/assets/body.svg'
import background from '@/game/assets/background.svg'
import green_apple from '@/game/assets/green_apple.svg'
import red_apple from '@/game/assets/red_apple.svg'

const SNAKE_HEAD = 'snake_head'
const SNAKE_BODY = 'snake_body'
const BACKGROUND = 'background'
const GREEN_APPLE = 'green_apple'
const RED_APPLE = 'red_apple'
const HEIGHT = 440
const WIDTH = 740

export default class PlayScene extends Scene {
  constructor () {
    super({ key: 'PlayScene' })
  }

  init () {
    this.score = 0
  }

  preload () {
    this.load.image(BACKGROUND, background)
    this.load.image(SNAKE_HEAD, snake_head)
    this.load.image(SNAKE_BODY, snake_body)
    this.load.image(GREEN_APPLE, green_apple)
    this.load.image(RED_APPLE, red_apple)
    this.time.addEvent({ delay: 80, callback: this.movePlayerManager, callbackScope: this, loop: true })
  }

  create () {
    this.background = this.add.tileSprite(0, 0, 1920, 1920, BACKGROUND)
    this.background.setOrigin(0)
    this.physics.world.setBounds(0, 0, this.background.width, this.background.height, true, true, true, true);

    this.body = this.physics.add.group({key: SNAKE_BODY, frameQuantity: 2})

    this.tail = this.body.getLast(true)

    this.head = this.physics.add.image(WIDTH/2, HEIGHT/2, SNAKE_HEAD)
    this.head.setCollideWorldBounds(true);
    this.head.onWorldBounds = true;
    this.head.setDepth(1);

    this.input.on('pointermove', pointer => {
      this.physics.moveTo(this.head, pointer.worldX, pointer.worldY, 150)
    }, this)

    this.cameras.main.startFollow(this.head, this.cameras.main.FOLLOW_LOCKON, 0.05, 0.05)

    this.food = this.physics.add.group()
    this.generateFood(GREEN_APPLE)
    this.time.addEvent({ delay: 2000, callback: this.generateFood, args: [GREEN_APPLE], callbackScope: this, loop: true })
    this.time.addEvent({ delay: Math.floor(Math.random() * 10000) + 5000, callback: this.generateFood, args: [RED_APPLE], callbackScope: this, loop: true })

    this.physics.add.overlap(this.head, this.food, this.eat, null, this)
  }

  generateFood (type) {
    let randomX = PhaserMath.Between(this.background.width * 0.1, this.background.width * 0.9)
    let randomY = PhaserMath.Between(this.background.height * 0.1, this.background.height * 0.9)
    this.food.create(randomX, randomY, type);
  }

  eat(player, food) {
    if (food.texture.key === GREEN_APPLE) {
      this.body.create(-10000, 0, SNAKE_BODY)
      this.score += 1
    } else {
      for (let i = 0; i < this.bodyparts.length/2; i++) {
        this.tail = this.body.getLast(true)
        if (this.tail != null) {
          this.tail.destroy(true)
          this.score -= 1
        }
      }
    }

    food.destroy(true);
  }

  update () {
    this.bodyparts = this.body.getChildren()

    if (this.bodyparts.length > 1) {
      Actions.ShiftPosition(this.bodyparts, this.head.x, this.head.y)
    } else {
      console.log("dead")
    }
  }
}
