<template>
    <div class="pt-16">
    <v-container>
        <v-row mx-10 px-10>
            <v-col cols="4">
                <v-card
                    class="mx-auto"
                    href="/login"
                    target="_blank"
                    max-width="244"
                    min-height="200"
                    hover
                    >
                    <v-card-item>
                        <v-card-title>
                            Ticket Management
                        </v-card-title>
                    </v-card-item>
                </v-card>
            </v-col>
            <v-col cols="4">
                <v-card
                    class="mx-auto"
                    href="/ticket/list"
                    max-width="244"
                    min-height="200"
                    hover
                >
                <v-card-title>
                    Buy Ticket
                </v-card-title>

                </v-card>
            </v-col>
        </v-row>
        
    </v-container></div>
</template>

<script>
export default {
    name: 'HomePage',
    data() {
        return {
        }
    },
    beforeCreate() {
        if (localStorage.getItem('token') === null) {
            this.$router.push('/login');
        }
    }
}
</script>