<template>
    <AlertMsg :msg="alertMsg.msg" :type="alertMsg.type" :hidden="!alertMsg.isVisable"></AlertMsg>
    <div class="pt-10">
        <v-dialog max-width="750">
            <template v-slot:activator="{ props: activatorProps }">
                <v-btn
                    v-bind="activatorProps"
                    color="surface-variant"
                    text="Create New"
                    variant="flat"
                ></v-btn>
            </template>

            <template v-slot:default="{ isActive }">
                <v-card title="Create Ticket">
                    <v-card-text>
                        <v-text-field
                        v-model="newTicket.name"
                        label="Name"
                        >
                        </v-text-field>
                        <v-text-field
                        v-model.number="newTicket.count"
                        label="Count"
                        >
                        </v-text-field>
                        <v-textarea
                            label="Introduce"
                            v-model="newTicket.introduce"
                            variant="outlined"
                            auto-grow
                            counter
                        >
                        </v-textarea>
                        <v-row>
                            <v-col cols="auto">
                                <v-date-picker
                                    title="Start Date"
                                    v-model="newTicket.startDate"
                                    time
                                ></v-date-picker>
                                <v-text-field
                                    label="Start Time"
                                    v-model="newTicket.startTime"
                                    suffix="NST"
                                    type="time"
                                ></v-text-field>
                            </v-col>
                            <v-col cols="auto">
                                <v-date-picker
                                    title="End Date"
                                    v-model="newTicket.endDate"
                                ></v-date-picker>
                                <v-text-field
                                    label="End Time"
                                    v-model="newTicket.endTime"
                                    suffix="NST"
                                    type="time"
                                ></v-text-field>
                            </v-col>
                        </v-row>
                    </v-card-text>

                    <v-card-actions>
                        <v-spacer></v-spacer>

                        <v-btn
                        text="Cancel"
                        @click="isActive.value = false"
                        >
                        </v-btn>
                        <v-btn
                        text="Create"
                        @click="sendCreateTicketData"
                        ></v-btn>
                    </v-card-actions>
                </v-card>
            </template>
        </v-dialog>
    </div>
    <div class="pt-10">
        <v-row class="px-3">
            <v-col cols="6">
                <v-card
                    class="mx-auto"
                >
                    <v-card-title>
                        Online
                    </v-card-title>
                    <div>
                        <v-card
                            class="mx-auto"
                            height="100"
                        >
                        </v-card>
                    </div>
                </v-card>
            </v-col>
            <v-col cols="6">
                <v-card
                    class="mx-auto"
                >
                    <v-card-title>
                        Wait On
                    </v-card-title>
                </v-card>
            </v-col>
        </v-row>
        
    </div>
</template>
<script>
import axios from 'axios';
import AlertMsg from './AlertMsg.vue'

export default {
    name: "OrderList",
    components: {
        AlertMsg
    },
    data() {
        return {
            online: [],
            waitOn: [],
            newTicket: {
                name: '',
                count: '',
                introduce: '',
                startDate: null,
                startTime: null,
                endDate: null,
                endTime: null
            },
            alertMsg: {
                msg: '',
                type: 'info',
                isVisable: false
            }
        }
    },
    created() {
        this.getListData();
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
        async getListData() {
            axios.post('http://localhost:8080/api/ticket/list',{},{
                headers: {
                    Authorization: localStorage.getItem('token')
                }
            }).then((res) => {
                if (res.data.code === 200) {
                    this.online = res.data.data.online;
                    this.waitOn = res.data.data.waitOn;
                } else if(res.data.code === 401) {
                    this.showAlert(res.data.msg, 'error', true)
                    this.$router.push('/login');
                } else {
                    this.showAlert(res.data.msg, 'error', true)
                }
            })
        },
        async sendCreateTicketData() {
            axios.post('http://localhost:8080/api/ticket/create', {
                name: this.newTicket.name,
                count: this.newTicket.count,
                introduce: this.newTicket.introduce,
                start_at: new Date(this.newTicket.startDate).toISOString().slice(0, 10)+ ' ' + this.newTicket.startTime + ':00',
                expires_at: new Date(this.newTicket.endDate).toISOString().slice(0, 10) + ' ' + this.newTicket.endTime + ':00'
            }, {
                headers: {
                    Authorization: localStorage.getItem('token')
                }
            }).then((res) => {
                if (res.data.code === 200) {
                    this.getListData();
                    this.showAlert(res.data.msg, 'success', true)
                    this.isActive.value = false
                } else if(res.data.code === 401) {
                    this.showAlert(res.data.msg, 'error', true)
                    this.$router.push('/login');
                } else {
                    this.showAlert(res.data.msg, 'error', true)
                }

            }).catch((err) => {
                console.log(err);
            });
        }
    }
}
</script>