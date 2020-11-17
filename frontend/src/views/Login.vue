<template>
  <div class="login">
    <v-container fluid class="pa-0">
			<v-img 
        class="home-img"
        src=""
        lazy-src="../assets/background.jpg"
			>
				<template v-slot:placeholder>
					<v-row
            class="fill-height ma-0 placeholder"
            align="center"
            justify="center"
					>
						<v-sheet
							color="secondary"
							dark
							elevation="24"
							rounded
							width="350"
							align="center"
						>
              <v-alert v-if="alert.status"
                dense
                :color="alert.type === 'error' ? '#cc4949' : '#5beb34'"
              >
                {{ alert.message }}
              </v-alert>
							<v-col cols="9">
								<div class="login py-7">
									<h1>Sign in</h1>
									<div class="login-inputs mt-4">
										<v-form
											ref="form"
											v-model="valid"
										>
											<v-text-field
												v-model="username"
												:rules="usernameRules"
												label="Username"
											></v-text-field>
											<v-text-field
												v-model="password"
												type="password"
												:rules="passwordRules"
												label="Password"
											></v-text-field>
										</v-form>
									</div>
									<v-btn
										block
										color="primary"
										outlined
										class="mt-7"
										:disabled="!valid"
										@click="login"
									>
										Login
									</v-btn>
								</div>
							</v-col>
						</v-sheet>
					</v-row>
				</template>
			</v-img>
    </v-container>
  </div>
</template>

<script>

export default {
  name: 'Login',
	data: () => ({
		valid: false,
		username: '',
		usernameRules: [
			v => !!v || 'Username is required',
			v => v.length < 30 || 'Username must be less than 30 characters',
		],
		password: '',
		passwordRules: [
			v => !!v || 'Password is required',
			v => v.length > 6 || 'Password must be great than 6 characters',
		],
    alert: {
      status: false,
      type: '#cc4949',
      message: '',
    }
	}),
	methods: {
		async login() {
			await this.$store.dispatch("signIn", {
        username: this.username,
        password: this.password,
      })
        .then(() => {
          this.$router.push("/game")
        })
        .catch(error => {
          this.alert = {
            status: true,
            type: 'error',
            message: error.response.data.message.toUpperCase(),
          }
          setTimeout(() => {
            this.alert.status = false
          }, 2500);
        })

      this.password = ''
		}
	}
}
</script>

<style>
	.login h1 {
		font-size: 3rem;
	}
</style>
