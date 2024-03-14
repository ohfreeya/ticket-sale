<template>
    <div class="pt-10">
        
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

export default {
    name: "OrderList",
    data() {
        return {
            online: [],
            waitOn: [],
        }
    },
    created() {
        this.getListData();
    },
    methods: {
        async getListData() {
            axios.post('/api/order/list').then((res) => {
                this.online = res.data.online;
                this.waitOn = res.data.waitOn;
            }).catch((err) => {
                console.log(err);
            });
        }
    }
}
</script>