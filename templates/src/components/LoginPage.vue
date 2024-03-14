<template>
    <AlertMsg :msg="alertMsg.msg" :type="alertMsg.type" :hidden="!alertMsg.isVisable"></AlertMsg>
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
                <p>or <router-link to="/register">register</router-link></p>

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
import AlertMsg from './AlertMsg.vue'

export default {
    name: 'LoginPage',
    components: {
        AlertMsg
    },
    data(){
      return {
        loading: false,
        timeout: null,
        account: '',
        password: '',
        alertMsg: {
            msg: '',
            type: 'info',
            isVisable: false
        }
      }
    },
    methods: {
      async showAlert(msg, type, isVisable) {
        this.alertMsg.msg = msg
        this.alertMsg.type = type
        this.alertMsg.isVisable = isVisable
        this.timeout = setTimeout(() => {
            this.alertMsg.isVisable = false
        }, 2000)

      },
      async sendData () {
        this.loading = true
        axios.post('http://localhost:8080/login', {
            account: this.account,
            password: this.password
        })
        .then(response => {
            if (response.data.code === 200) {
                // Redirect to home page if login is successful
                localStorage.setItem('token', response.data.token)
                localStorage.setItem('uid', response.data.uid)
                this.$router.push('/');
            } else {
                // Show error message
                this.showAlert(response.data.msg, 'error', true)
            }
            this.loading = false
        })
      },
    },
  }
</script>