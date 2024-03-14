<template>
    <AlertMsg :msg="alertMsg.msg" :type="alertMsg.type" :hidden="!alertMsg.isVisable"></AlertMsg>
    <div>
      <h1 class="pb-3">Register</h1>
        <v-sheet class="mx-auto" max-width="300">
            <v-form validate-on="submit lazy" @submit.prevent="sendData">
                <v-text-field
                    v-model="user.account"
                    label="Account"
                ></v-text-field>
                <v-text-field
                    v-model="user.password"
                    label="Password"
                ></v-text-field>
                <v-text-field
                    v-model="user.confirm"
                    label="Confirm Password"
                ></v-text-field>
                <v-text-field
                    v-model="user.name"
                    label="Name"
                ></v-text-field>
                <v-text-field
                    v-model="user.email"
                    label="Email"
                ></v-text-field>
                <p>or <router-link to="/login">login</router-link></p>

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
    name: 'RegisterPage',
    components: {
        AlertMsg
    },
    data(){
      return {
        loading: false,
        timeout: null,
        user: {
            account: "",
            password: "",
            confirm: "",
            name: "",
            email: ""
        },
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
        if (this.user.password !== this.user.confirm) {
            this.showAlert('Password and confirm password are not the same', 'error', true)
            return
        }

        this.loading = true
        axios.post('http://localhost:8080/register', {
            account: this.user.account,
            password: this.user.password,
            name: this.user.name,
            email: this.user.email
        })
        .then((res) => {
            this.loading = false
            if(res.data.code === 200)
                this.$router.push('/login')
            else
                this.showAlert(res.data.msg, 'error', true)
        })
      }
    }
}

</script>