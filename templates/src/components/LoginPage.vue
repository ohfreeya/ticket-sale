<template>
    <div class="pt-5">
        <h1 class="pb-3">Login</h1>
        <v-sheet class="mx-auto" max-width="300">
            <v-form validate-on="submit lazy" @submit.prevent="sendData">
                <v-text-field
                    v-model="account"
                    label="Account"
                ></v-text-field>
                <v-text-field
                    v-model="password"
                    label="Password"
                ></v-text-field>

                <v-btn
                    :loading="loading"
                    class="mt-2"
                    text="Submit"
                    type="button"
                    @click="sendData"
                    block
                ></v-btn>
            </v-form>
        </v-sheet>
    </div>
</template>

<script>
import axios from 'axios'

export default {
    name: 'LoginPage',
    data(){
      return {
        loading: false,
        timeout: null,
        account: '',
        password: '',
      }
    },
    methods: {
      async sendData () {
        console.log(this.account, this.password)
        this.loading = true

        try {
          axios.post('http://localhost:8080/login', {
            headers: {
              'Content-Type': 'application/json',
            },
            body: {
              account: this.account,
              password: this.password,
            },
          }).then(
            (response) => {
              if (response.status === 200) {
                this.$router.push('/')
              } else {
                alert('Login failed '+ response.msg)
                this.loading = false
              }
            }
          )
        } catch (error) {
          console.error(error)
          this.loading = false
        }
      },
    },
  }
</script>